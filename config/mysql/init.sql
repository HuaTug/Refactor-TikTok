create database if not exists Hertz;
use Hertz;
drop table if exists `users`;
create table   `users`(
    `user_id` bigint not null auto_increment ,
    `user_name` varchar(255) not null ,
    `password` varchar(255) not null ,
    `avatar_url` varchar(255) ,
    `created_at` varchar(255) not null,
    `updated_at` varchar(255) not null,
    `deleted_at` varchar(255) ,
    primary key (user_id) ,
    key `username_password_index` (user_name,password) using btree
) engine = InnoDB  auto_increment=1 default  charset = utf8mb4;

drop table if exists `videos`;
create table `videos`(
    `video_id` bigint not null auto_increment,
    `user_id` bigint not null ,
    `video_url` varchar(255) not null ,
    `cover_url` varchar(255) not null ,
    `title` varchar(255) not null ,
    `description` varchar(255) not null ,
    `visit_count` varchar(255) not null ,
    `created_at` varchar(255) not null ,
    `updated_at` varchar(255) not null ,
    `deleted_at` varchar(255) ,
    primary key (video_id),
    key `time` (created_at) using btree ,
    key `author` (user_id) using btree
)engine InnoDB auto_increment=1  default  charset=utf8mb4;

drop table if exists `commenst`;
create table `comments`(
    `comment_id` bigint not null auto_increment,
    `user_id` bigint not null ,
    `video_id` bigint not null ,
    `parent_id` bigint not null ,
    `content` varchar(255) not null ,
    `created_at` varchar(255) not null ,
    `updated_at` varchar(255) not null ,
    `deleted_at` varchar(255)  ,
    primary key (comment_id) ,
    key `vide_index` (video_id) using btree
)engine =InnoDB auto_increment=1 default charset = utf8mb4;

drop table if exists `video_likes`;
create table `video_likes`(
    `video_likes_id` bigint not null ,
    `user_id` bigint not null ,
    `video_id` bigint not null ,
    `created_at` varchar(255) not null ,
    `deleted_at` varchar(255)  ,
    primary key (video_likes_id),
    unique key `user_id_video_id_no_duplicate` (user_id,video_id),
    key `user_id_video_id_index`(user_id,video_id) using btree ,
    key `user_id_index` (user_id) using btree ,
    key `video_id_index` (video_id) using btree
)engine = InnoDB auto_increment=1 default charset = utf8mb4;

drop table if exists `comment_likes`;
create table `comment_likes`(
    `comment_likes_id` bigint not null ,
    `user_id` bigint not null ,
    `comment_id` bigint not null ,
    `created_at` varchar(255) not null ,
    `deleted_at` varchar(255) ,
    primary key (comment_likes_id) ,
    unique key `user_id_comment_id_no_duplicate` (user_id,comment_id) ,
    key `user_id_comment_id_index` (user_id,comment_id) using btree ,
    key `user_id_index` (user_id) using btree ,
    key `comment_id_index` (comment_id) using btree
)engine = InnoDB auto_increment=1  default charset = utf8mb4 ;

drop table  if exists `follows`;
create table `follows`(
    `follow_id` bigint not null auto_increment,
    `following_id` bigint not null ,
    `followers_id` bigint not null ,
    `created_at` varchar(255) not null ,
    `deleted_at` varchar(255) ,
    primary key (follow_id) ,
    unique key `followers_following_no_duplicate` (followers_id,following_id) ,
    key `following_id_followers_id_index` (following_id,followers_id) using btree ,
    key `followers_id_index` (followers_id) using btree ,
    key `following_id_index` (following_id) using btree 
)engine = InnoDB auto_increment=1  default charset = utf8mb4;