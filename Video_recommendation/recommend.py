import sys
import json
import pika
import mysql.connector
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity


# 连接到 MySQL
def get_mysql_connection():
    return mysql.connector.connect(
        host="localhost",
        user="root",
        password="root",
        database="TikTok"
    )


# 查询视频数据
def get_videos_from_db():
    conn = get_mysql_connection()
    cursor = conn.cursor(dictionary=True)
    cursor.execute("SELECT video_id, title, description, label_names, category FROM videos")
    videos = cursor.fetchall()
    cursor.close()
    conn.close()
    return videos


# 查询用户行为数据
def get_user_behaviors_from_db(user_id):
    conn = get_mysql_connection()
    cursor = conn.cursor(dictionary=True)
    query = """
        SELECT video_id, behavior_type
        FROM user_behaviors
        WHERE user_id = %s AND behavior_type IN ('view', 'like', 'share', 'comment')
    """
    cursor.execute(query, (user_id,))
    behaviors = cursor.fetchall()
    cursor.close()
    conn.close()
    return behaviors


# 查询用户观看历史
def get_user_watch_history_from_db(user_id):
    conn = get_mysql_connection()
    cursor = conn.cursor(dictionary=True)
    cursor.execute("SELECT video_id FROM user_video_watch_histories WHERE user_id = %s", (user_id,))
    watch_history = cursor.fetchall()
    cursor.close()
    conn.close()
    return [record['video_id'] for record in watch_history]


# 推荐视频
def recommend_videos(user_id, top_n=10):
    # 获取视频数据
    videos = get_videos_from_db()

    # 获取用户行为数据
    user_behaviors = get_user_behaviors_from_db(user_id)
    user_watch_history = get_user_watch_history_from_db(user_id)

    # 提取文本特征
    texts = [f"{video['title']} {video['description']} {video['label_names']} {video['category']}" for video in videos]
    vectorizer = TfidfVectorizer()
    features = vectorizer.fit_transform(texts)
    similarity_matrix = cosine_similarity(features, features)

    # 计算用户兴趣向量
    user_interest_vector = [0] * len(videos)
    for behavior in user_behaviors:
        video_index = next((i for i, v in enumerate(videos) if v['video_id'] == behavior['video_id']), None)
        if video_index is not None:
            if behavior['behavior_type'] == 'view':
                user_interest_vector[video_index] += 1
            elif behavior['behavior_type'] == 'like':
                user_interest_vector[video_index] += 2
            elif behavior['behavior_type'] == 'share':
                user_interest_vector[video_index] += 3
            elif behavior['behavior_type'] == 'comment':
                user_interest_vector[video_index] += 4

    # 计算推荐分数
    scores = similarity_matrix.dot(user_interest_vector)

    # 推荐视频
    recommended_indices = scores.argsort()[::-1][:top_n]
    recommended_videos = [videos[i] for i in recommended_indices if videos[i]['video_id'] not in user_watch_history]

    return recommended_videos


# 主函数
def main():
    if len(sys.argv) != 2:
        print(json.dumps({"error": "Usage: python recommend.py <user_id>"}))
        sys.exit(1)

    try:
        user_id = int(sys.argv[1])
        recommended_videos = recommend_videos(user_id)

        # 发送结果到消息队列
        connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        channel = connection.channel()
        channel.queue_declare(queue='recommend_queue', durable=True)

        message = json.dumps(recommended_videos)
        channel.basic_publish(
            exchange='',
            routing_key='recommend_queue',
            body=message.encode(),
            properties=pika.BasicProperties(
                delivery_mode=2,  # make message persistent
            )
        )

        print("Recommendations sent to queue")
        connection.close()
    except Exception as e:
        print(json.dumps({"error": str(e)}))
        sys.exit(1)


if __name__ == "__main__":
    main()
