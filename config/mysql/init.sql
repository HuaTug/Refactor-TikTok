create database if not exists TikTok;
use TikTok;

-- Table structure of sys_setting --
drop table if exists `sys_settings`;
create table `sys_settings`(
    `id` bigint not null auto_increment,
    `audit_policy` longtext not null,
    `audit_open` tinyint not null default '0' comment '0:disable 1:enable',
    `hot_limit` varchar(255) not null default '100',
    `allow_ip` varchar(255) not null,
    `auth` tinyint not null default '0' comment '0:disable 1:enable',
    `value` varchar(255) not null,
    `created_at` varchar(255) not null,
    `updated_at` varchar(255) not null,
    primary key (id)
)engine = InnoDB  auto_increment=1 default  charset = utf8mb4;

-- Table structure of role --
drop table if exists `roles`;
create table `roles`(
    `role_id` bigint not null auto_increment,
    `role` varchar(255) not null,
    primary key (role_id)
) engine = InnoDB  auto_increment=1 default  charset = utf8mb4;
INSERT INTO `roles` (`role_id`,`role`) VALUES (1,'admin'),(2,'user'),(3,'guest');-- 完成了对角色的权限划分

-- Table structure of role_permission --
drop table if exists `role_permissions`;
create table `role_permissions`(
    `permission_id` bigint not null auto_increment,
    `role_id` bigint not null,
    primary key (permission_id)
)engine = InnoDB  auto_increment=1 default  charset = utf8mb4;

-- Table structure of user --
drop table if exists `users`;
create table   `users`(
    `user_id` bigint not null auto_increment ,
    `user_name` varchar(255) not null ,
    `password` varchar(255) not null ,
    `email` varchar(30) not null,
    `sex` tinyint(1) not null, -- 0:female 1:male
    `avatar_url` varchar(255) ,
    `created_at` varchar(255) not null,
    `updated_at` varchar(255) not null,
    `deleted_at` varchar(255) ,
    primary key (user_id) ,
    key `username_password_index` (user_name,password) using btree
) engine = InnoDB  auto_increment=1 default  charset = utf8mb4;

-- -- 创建其他分表 users_1, users_2, users_3
-- CREATE TABLE `users_1` LIKE `users_0`;
-- CREATE TABLE `users_2` LIKE `users_0`;
-- CREATE TABLE `users_3` LIKE `users_0`;

-- Table structure of user_role --
drop table if exists `user_roles`;
create table `user_roles`(
    `role_id` bigint not null,
    `user_id` bigint not null,
    `role` varchar(255) not null
);


drop table if exists `user_behaviors`;
create table `user_behaviors`(
    `user_behavior_id` bigint not null auto_increment,
    `user_id` bigint not null,
    `video_id` bigint not null,
    `behavior_type` varchar(50) not null, -- 'view' 'like' 'share' 'comment'
    `behavior_time` varchar(255) not null,
    unique key(user_id,video_id,behavior_type),
    primary key (user_behavior_id)
)engine InnoDB auto_increment=1  default  charset=utf8mb4;

-- INSERT INTO `user_behaviors` (`user_id`, `video_id`, `behavior_type`, `behavior_time`) VALUES
-- (1, 25, 'view', '2024-11-01 08:30:00'),
-- (6, 3, 'like', '2024-11-02 09:15:00'),
-- (2, 17, 'share', '2024-11-02 11:45:00'),
-- (4, 9, 'comment', '2024-11-03 14:00:00'),
-- (5, 12, 'view', '2024-11-03 15:30:00'),
-- (3, 7, 'like', '2024-11-04 10:00:00'),
-- (1, 19, 'comment', '2024-11-04 11:15:00'),
-- (6, 21, 'share', '2024-11-05 09:45:00'),
-- (4, 6, 'view', '2024-11-05 16:00:00'),
-- (5, 15, 'like', '2024-11-06 13:30:00'),
-- (2, 11, 'share', '2024-11-06 14:45:00'),
-- (3, 8, 'comment', '2024-11-07 12:00:00'),
-- (1, 22, 'view', '2024-11-07 08:15:00'),
-- (6, 28, 'like', '2024-11-08 10:30:00'),
-- (4, 14, 'comment', '2024-11-08 11:45:00'),
-- (5, 26, 'share', '2024-11-09 09:00:00'),
-- (2, 20, 'view', '2024-11-09 10:15:00'),
-- (3, 30, 'like', '2024-11-10 11:30:00'),
-- (1, 2, 'comment', '2024-11-10 12:45:00'),
-- (6, 13, 'view', '2024-11-11 13:00:00'),
-- (4, 18, 'like', '2024-11-11 14:15:00'),
-- (5, 24, 'share', '2024-11-12 15:30:00'),
-- (2, 16, 'comment', '2024-11-12 16:00:00'),
-- (3, 4, 'view', '2024-11-13 10:00:00'),
-- (1, 31, 'like', '2024-11-13 11:30:00'),
-- (6, 5, 'share', '2024-11-14 12:45:00'),
-- (4, 10, 'comment', '2024-11-14 13:15:00'),
-- (5, 23, 'view', '2024-11-15 08:30:00'),
-- (2, 29, 'like', '2024-11-15 09:45:00'),
-- (3, 1, 'comment', '2024-11-16 14:00:00'),
-- (1, 27, 'view', '2024-11-16 15:30:00'),
-- (6, 32, 'like', '2024-11-17 10:00:00'),
-- (4, 33, 'share', '2024-11-17 11:15:00'),
-- (5, 34, 'comment', '2024-11-18 12:00:00'),
-- (2, 35, 'view', '2024-11-18 13:30:00'),
-- (3, 36, 'like', '2024-11-19 14:45:00'),
-- (1, 37, 'share', '2024-11-19 16:00:00'),
-- (6, 38, 'comment', '2024-11-20 10:15:00'),
-- (4, 39, 'view', '2024-11-20 11:45:00'),
-- (5, 40, 'like', '2024-11-21 12:00:00'),
-- (2, 41, 'share', '2024-11-21 13:30:00'),
-- (3, 42, 'comment', '2024-11-22 14:00:00'),
-- (1, 43, 'view', '2024-11-22 15:15:00'),
-- (6, 44, 'like', '2024-11-23 09:45:00'),
-- (4, 45, 'share', '2024-11-23 11:00:00'),
-- (5, 46, 'comment', '2024-11-24 12:30:00'),
-- (2, 47, 'view', '2024-11-24 14:00:00'),
-- (3, 48, 'like', '2024-11-25 15:45:00'),
-- (1, 49, 'share', '2024-11-25 16:30:00'),
-- (6, 50, 'comment', '2024-11-26 09:00:00');
-- Table structure of videos --
drop table if exists `videos`;
create table `videos`(
    `video_id` bigint not null auto_increment,
    `user_id` bigint not null ,
    `video_url` varchar(255) not null ,
    `cover_url` varchar(255) not null ,
    `title` varchar(255) not null ,
    `description` varchar(255) not null ,
    `visit_count` varchar(255) default '0' not null,
    `share_count` varchar(255) default '0' not null ,
    `likes_count` varchar(255) default '0' not null,
    `favorites_count` varchar(255) default '0' not null,
    `comment_count` varchar(255) default '0' not null,
    `history_count` varchar(255) default '0' not null,
    `open` tinyint not null default '0' comment '0:private 1:public',
    `audit_status` tinyint not null default '0' comment '0:unreviewed 1:reviewed',
    `label_names` varchar(255) default '' not null,
    `category` varchar(255) default '' not null,
    `created_at` varchar(255) not null ,
    `updated_at` varchar(255) not null ,
    `deleted_at` varchar(255) ,
    primary key (video_id),
    key `time` (created_at) using btree ,
    key `author` (user_id) using btree
)engine InnoDB auto_increment=1  default  charset=utf8mb4;

-- INSERT INTO videos (user_id, video_url, cover_url, title, description, label_names, category, created_at, updated_at) VALUES
-- (1, 'https://example.com/video1.mp4', 'https://example.com/cover1.jpg', 'Amazing Nature', 'A breathtaking view of nature.', 'nature, scenery', 'Nature', '2024-11-01 10:00:00', '2024-11-01 10:00:00'),
-- (2, 'https://example.com/video2.mp4', 'https://example.com/cover2.jpg', 'Tech Innovations 2024', 'Latest technology innovations.', 'tech, innovation', 'Technology', '2024-11-02 11:00:00', '2024-11-02 11:00:00'),
-- (3, 'https://example.com/video3.mp4', 'https://example.com/cover3.jpg', 'Cooking Tips for Beginners', 'Easy cooking tips for everyone.', 'cooking, tips', 'Food', '2024-11-03 12:00:00', '2024-11-03 12:00:00'),
-- (4, 'https://example.com/video4.mp4', 'https://example.com/cover4.jpg', 'Top 10 Football Goals', 'The best football goals of the year.', 'football, sports', 'Sports', '2024-11-04 13:00:00', '2024-11-04 13:00:00'),
-- (5, 'https://example.com/video5.mp4', 'https://example.com/cover5.jpg', 'Yoga for Beginners', 'A beginner\'s guide to yoga.', 'yoga, fitness', 'Health & Fitness', '2024-11-05 14:00:00', '2024-11-05 14:00:00'),
-- (6, 'https://example.com/video6.mp4', 'https://example.com/cover6.jpg', 'Travel Vlog in Japan', 'Exploring the beauty of Japan.', 'travel, vlog', 'Travel', '2024-11-06 15:00:00', '2024-11-06 15:00:00'),
-- (7, 'https://example.com/video7.mp4', 'https://example.com/cover7.jpg', 'DIY Home Decor Ideas', 'Creative home decor projects.', 'DIY, home decor', 'Lifestyle', '2024-11-07 16:00:00', '2024-11-07 16:00:00'),
-- (8, 'https://example.com/video8.mp4', 'https://example.com/cover8.jpg', 'Best Moments in Basketball History', 'Highlights from basketball history.', 'basketball, sports highlights', 'Sports', '2024-11-08 17:00:00', '2024-11-08 17:00:00'),
-- (9, 'https://example.com/video9.mp4', 'https://example.com/cover9.jpg', 'Fitness Challenge Day 1!', 'Join me on my fitness journey!', 'fitness challenge, motivation', 'Health & Fitness', '2024-11-09 18:00:00', '2024-11-09 18:00:00'),
-- (10, 'https://example.com/video10.mp4', 'https://example.com/cover10.jpg', 'Exploring the Amazon Rainforest', 'A journey through the Amazon.', 'nature, adventure', 'Nature', '2024-11-10 19:00:00', '2024-11-10 19:00:00'),
-- (11, 1, 'https://example.com/video11.mp4', 'https://example.com/cover11.jpg', 'AI in Everyday Life', 'How AI is changing our daily routines.', 'AI, technology', 'Technology', '2024-11-01 10:30:00'),
-- (12, 2, 'https://example.com/video12.mp4', 'https://example.com/cover12.jpg', 'Healthy Smoothie Recipes', 'Delicious and healthy smoothie ideas.', 'smoothies, health', 'Food', '2024-11-02 11:30:00'),
-- (13, 3, 'https://example.com/video13.mp4', 'https://example.com/cover13.jpg', 'The Future of Space Exploration', 'What lies ahead in space travel?', 'space exploration, future', 'Science', '2024-11-03 12:30:00'),
-- (14, 1, 'https://example.com/video14.mp4', 'https://example.com/cover14.jpg', 'Mindfulness Meditation Techniques', 'Learn to meditate effectively.', 'meditation, mindfulness', 'Health & Fitness', '2024-11-04 13:30:00'),
-- (15, 2, 'https://example.com/video15.mp4', 'https://example.com/cover15.jpg', 'Top Travel Destinations in Europe', 'Must-see places in Europe.', 'travel, europe', 'Travel', '2024-11-05 14:30:00'),
-- (16, 3, 'https://example.com/video16.mp4', 'https://example.com/cover16.jpg', 'Best Coding Practices for Beginners', 'Tips for new programmers.', 'coding, programming tips', 'Technology', '2024-11-06 15:30:00'),
-- (17, 1, 'https://example.com/video17.mp4', 'https://example.com/cover17.jpg', 'Ultimate Workout Routine for Weight Loss', 'Effective workouts to lose weight.', 'workout, weight loss tips', 'Health & Fitness', '2024-07-06 16:30:00'),
-- (18, 2, 'https://example.com/video18.mp4', 'https://example.com/cover18.jpg', 'The Art of Photography', 'Tips for capturing stunning photos.', 'photography, art', 'Lifestyle', '2024-07-07 07:30:00'),
-- (19, 3, 'https://example.com/video19.mp4', 'https://example.com/cover19.jpg', 'How to Start a Podcast', 'A guide to launching your own podcast.', 'podcasting, media', 'Technology', '2024-07-08 08:30:00'),
-- (20, 1, 'https://example.com/video20.mp4', 'https://example.com/cover20.jpg', 'Fashion Trends of the Year', 'Latest trends in fashion.', 'fashion, trends', 'Lifestyle', '2024-07-09 09:30:00'),
-- (21, 2, 'https://example.com/video21.mp4', 'https://example.com/cover21.jpg', 'Exploring Ancient Civilizations', 'A look into ancient cultures.', 'history, exploration', 'Education', '2024-07-10 10:30:00'),
-- (22, 3, 'https://example.com/video22.mp4', 'https://example.com/cover22.jpg', 'The Science of Cooking', 'Understanding the chemistry behind cooking.', 'cooking, science', 'Food', '2024-07-11 11:30:00'),
-- (23, 1, 'https://example.com/video23.mp4', 'https://example.com/cover23.jpg', 'Digital Marketing Strategies for Success', 'Effective strategies to grow your business online.', 'marketing, business', 'Business', '2024-07-12 12:30:00'),
-- (24, 2, 'https://example.com/video24.mp4', 'https://example.com/cover24.jpg', 'The Best Hiking Trails in the World', 'Discover breathtaking hiking routes around the globe.', 'hiking, travel', 'Travel', '2024-07-13 13:30:00'),
-- (25, 3, 'https://example.com/video25.mp4', 'https://example.com/cover25.jpg', 'Understanding Cryptocurrency and Blockchain', 'A beginner\'s guide to cryptocurrency.', 'cryptocurrency, finance', 'Finance', '2024-07-14 14:30:00'),
-- (26, 1, 'https://example.com/video26.mp4', 'https://example.com/cover26.jpg', 'Gardening Tips for Beginners', 'Learn how to start your own garden.', 'gardening, tips', 'Lifestyle', '2024-07-15 15:30:00'),
-- (27, 2, 'https://example.com/video27.mp4', 'https://example.com/cover27.jpg', 'Exploring the World of Virtual Reality', 'An introduction to virtual reality experiences.', 'virtual reality, tech', 'Technology', '2024-07-16 16:30:00'),
-- (28, 3, 'https://example.com/video28.mp4', 'https://example.com/cover28.jpg', 'How to Build a Personal Brand', 'Tips for creating your own personal brand.', 'branding, entrepreneurship', 'Business', '2024-07-17 17:30:00'),
-- (29, 1, 'https://example.com/video29.mp4', 'https://example.com/cover29.jpg', 'The Future of Renewable Energy', 'Exploring new technologies in renewable energy.', 'energy, renewable, technology', 'Science', '2024-07-18 18:30:00'),
-- (30, 2, 'https://example.com/video30.mp4', 'https://example.com/cover30.jpg', 'How to Create Stunning Visual Effects', 'Learn how to add visual effects to your videos.', 'VFX, film, tutorials', 'Film & Animation', '2024-07-19 19:30:00'),
-- (31, 3, 'https://example.com/video31.mp4', 'https://example.com/cover31.jpg', 'Virtual Fitness Classes', 'Join our virtual fitness classes from home.', 'fitness, virtual, workout', 'Health & Fitness', '2024-07-20 20:30:00'),
-- (32, 1, 'https://example.com/video32.mp4', 'https://example.com/cover32.jpg', 'Mindful Eating Practices', 'Learn the art of mindful eating for better health.', 'mindful eating, health', 'Food', '2024-07-21 21:30:00'),
-- (33, 2, 'https://example.com/video33.mp4', 'https://example.com/cover33.jpg', 'Building a Startup from Scratch', 'Tips for entrepreneurs starting their own businesses.', 'startup, entrepreneurship', 'Business', '2024-07-22 22:30:00'),
-- (34, 3, 'https://example.com/video34.mp4', 'https://example.com/cover34.jpg', 'Essential Coding Skills for Web Developers', 'A guide to essential skills for web development.', 'coding, web development', 'Technology', '2024-07-23 23:30:00'),
-- (35, 1, 'https://example.com/video35.mp4', 'https://example.com/cover35.jpg', 'How to Stay Productive While Working from Home', 'Productivity tips for remote workers.', 'productivity, remote work', 'Lifestyle', '2024-07-24 00:30:00'),
-- (36, 2, 'https://example.com/video36.mp4', 'https://example.com/cover36.jpg', 'Exploring the Oceans: Marine Life', 'A dive into the wonders of the ocean.', 'ocean, marine life', 'Science', '2024-07-25 01:30:00'),
-- (37, 3, 'https://example.com/video37.mp4', 'https://example.com/cover37.jpg', 'How to Master Digital Art', 'Tips and techniques for creating stunning digital art.', 'art, digital art', 'Art & Design', '2024-07-26 02:30:00'),
-- (38, 1, 'https://example.com/video38.mp4', 'https://example.com/cover38.jpg', 'Fitness Myths Busted', 'Debunking common fitness myths and misconceptions.', 'fitness, health', 'Health & Fitness', '2024-07-27 03:30:00'),
-- (39, 2, 'https://example.com/video39.mp4', 'https://example.com/cover39.jpg', 'Introduction to Machine Learning', 'An introductory course on machine learning concepts.', 'machine learning, AI', 'Technology', '2024-07-28 04:30:00'),
-- (40, 3, 'https://example.com/video40.mp4', 'https://example.com/cover40.jpg', 'Top 5 Best Tech Gadgets of 2024', 'A roundup of the best tech gadgets released in 2024.', 'tech, gadgets', 'Technology', '2024-07-29 05:30:00'),
-- (41, 1, 'https://example.com/video41.mp4', 'https://example.com/cover41.jpg', 'How to Build a Website from Scratch', 'Learn the basics of web development and build your first site.', 'web development, coding', 'Technology', '2024-07-30 06:30:00'),
-- (42, 2, 'https://example.com/video42.mp4', 'https://example.com/cover42.jpg', 'How to Stay Fit While Traveling', 'Tips for staying fit on the go.', 'fitness, travel', 'Health & Fitness', '2024-07-31 07:30:00'),
-- (43, 3, 'https://example.com/video43.mp4', 'https://example.com/cover43.jpg', 'Building Your Personal Finance Plan', 'A guide to managing your personal finances effectively.', 'finance, budgeting', 'Finance', '2024-08-01 08:30:00'),
-- (44, 1, 'https://example.com/video44.mp4', 'https://example.com/cover44.jpg', 'The Science Behind Climate Change', 'Understanding the causes and effects of climate change.', 'climate change, science', 'Science', '2024-08-02 09:30:00'),
-- (45, 2, 'https://example.com/video45.mp4', 'https://example.com/cover45.jpg', 'The Art of Public Speaking', 'Master the skills of effective public speaking.', 'public speaking, communication', 'Education', '2024-08-03 10:30:00'),
-- (46, 3, 'https://example.com/video46.mp4', 'https://example.com/cover46.jpg', 'Cooking for a Crowd', 'How to prepare meals for large groups of people.', 'cooking, large groups', 'Food', '2024-08-04 11:30:00'),
-- (47, 1, 'https://example.com/video47.mp4', 'https://example.com/cover47.jpg', 'Top 10 Productivity Hacks', 'Effective ways to boost your productivity every day.', 'productivity, hacks', 'Lifestyle', '2024-08-05 12:30:00'),
-- (48, 2, 'https://example.com/video48.mp4', 'https://example.com/cover48.jpg', 'How to Stay Motivated in Tough Times', 'Motivation tips for getting through challenging times.', 'motivation, self-help', 'Lifestyle', '2024-08-06 13:30:00'),
-- (49, 3, 'https://example.com/video49.mp4', 'https://example.com/cover49.jpg', 'Building a Sustainable Business', 'Tips for entrepreneurs aiming to build eco-friendly businesses.', 'sustainability, entrepreneurship', 'Business', '2024-08-07 14:30:00'),
-- (50, 1, 'https://example.com/video50.mp4', 'https://example.com/cover50.jpg', 'The Best Coding Tools for Developers', 'A list of must-have tools for every developer.', 'coding, tools', 'Technology', '2024-08-08 15:30:00');

-- Table structure of user_video_watch_histories --
drop table if exists `user_video_watch_histories`;
create table `user_video_watch_histories`(
    `user_video_watch_history_id` bigint not null auto_increment,
    `user_id` bigint not null,
    `video_id` bigint not null,
    `watch_time` varchar(255) not null,
    `deleted_at` varchar(255),
    primary key (user_video_watch_history_id),
    unique key( user_id,video_id)
)engine InnoDB auto_increment=1  default  charset=utf8mb4;
-- INSERT INTO `user_video_watch_histories` (`user_id`, `video_id`, `watch_time`, `deleted_at`) VALUES
-- (1, 25, '2024-11-01 08:30:00', NULL),
-- (6, 3, '2024-11-02 09:15:00', NULL),
-- (2, 17, '2024-11-02 11:45:00', NULL),
-- (4, 9, '2024-11-03 14:00:00', NULL),
-- (5, 12, '2024-11-03 15:30:00', NULL),
-- (3, 7, '2024-11-04 10:00:00', NULL),
-- (1, 19, '2024-11-04 11:15:00', NULL),
-- (6, 21, '2024-11-05 09:45:00', NULL),
-- (4, 6, '2024-11-05 16:00:00', NULL),
-- (5, 15, '2024-11-06 13:30:00', NULL),
-- (2, 11, '2024-11-06 14:45:00', NULL),
-- (1, 8, '2024-11-07 12:00:00', NULL),
-- (1, 22, '2024-11-07 08:15:00', NULL),
-- (6, 28, '2024-11-08 10:30:00', NULL),
-- (3, 14, '2024-11-08 11:45:00', NULL),
-- (5, 26, '2024-11-09 09:00:00', NULL),
-- (2, 20, '2024-11-09 10:15:00', NULL),
-- (1, 30, '2024-11-10 11:30:00', NULL),
-- (1, 2, '2024-11-10 12:45:00', NULL),
-- (2, 13, '2024-11-11 13:00:00', NULL),
-- (4, 18, '2024-11-11 14:15:00', NULL),
-- (5, 24, '2024-11-12 15:30:00', NULL),
-- (2, 16, '2024-11-12 16:00:00', NULL),
-- (3, 4, '2024-11-13 10:00:00', NULL),
-- (1, 31, '2024-11-13 11:30:00', NULL),
-- (2, 5, '2024-11-14 12:45:00', NULL),
-- (4, 10, '2024-11-14 13:15:00', NULL),
-- (4, 23, '2024-11-15 08:30:00', NULL),
-- (2, 29, '2024-11-15 09:45:00', NULL),
-- (3, 1, '2024-11-16 14:00:00', NULL),
-- (5, 27, '2024-11-16 15:30:00', NULL),
-- (6, 32, '2024-11-17 10:00:00', NULL),
-- (4, 33, '2024-11-17 11:15:00', NULL),
-- (4, 34, '2024-11-18 12:00:00', NULL),
-- (2, 35, '2024-11-18 13:30:00', NULL),
-- (3, 36, '2024-11-19 14:45:00', NULL),
-- (1, 37, '2024-11-19 16:00:00', NULL),
-- (3, 38, '2024-11-20 10:15:00', NULL),
-- (4, 39, '2024-11-20 11:45:00', NULL),
-- (5, 40, '2024-11-21 12:00:00', NULL),
-- (2, 41, '2024-11-21 13:30:00', NULL),
-- (3, 42, '2024-11-22 14:00:00', NULL),
-- (1, 43, '2024-11-22 15:15:00', NULL),
-- (6, 44, '2024-11-23 09:45:00', NULL),
-- (2, 45, '2024-11-23 11:00:00', NULL),
-- (1, 46, '2024-11-24 12:30:00', NULL),
-- (2, 47, '2024-11-24 14:00:00', NULL),
-- (2, 48, '2024-11-25 15:45:00', NULL),
-- (1, 49, '2024-11-25 16:30:00', NULL),
-- (6, 50, '2024-11-26 09:00:00', NULL);

-- Table structure of video_likes --
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

-- Table structure of video_share --
drop table if exists `video_shares`;
create table `video_shares`(
    `video_share_id` bigint not null auto_increment,
    `user_id` bigint not null, -- 分享者
    `video_id` bigint not null, -- 被分享的视频
    `to_user_id` bigint not null, -- 被分享的用户
    `created_at` varchar(255) not null,
    `deleted_at` varchar(255),
    primary key (video_share_id)
)engine = InnoDB auto_increment=1 default charset = utf8mb4;

-- Table structure of favorites --
drop table if exists `favorites`;
create table `favorites`(
    `favorite_id` bigint not null auto_increment,
    `user_id` bigint not null,
    `name` varchar(255) not null,
    `description` varchar(255) default ''  not null,
    `cover_url` varchar(255) default '' not null,
    `created_at` varchar(255) not null,
    `deleted_at` varchar(255),
    primary key (favorite_id)
)engine = InnoDB auto_increment=1 default charset = utf8mb4;

-- Table structure of favorites_videos --
drop table if exists `favorites_videos`;
create table `favorites_videos`(
    `favorite_video_id` bigint not null auto_increment,
    `favorite_id` bigint not null, -- 收藏夹id
    `video_id` bigint not null, -- 被收藏的视频
    `user_id` bigint not null,
    primary key (favorite_video_id),
    unique key `fav_vid_usr_index` (favorite_id,video_id,user_id) using btree,
    key  `fav_usr_index` (user_id,favorite_id) using btree
)engine = InnoDB auto_increment=1 default charset = utf8mb4;

-- Table structure of user_perferences --
drop table if exists `user_perferences`;
create table `user_perferences`(
    `user_id` bigint not null,
    `label_names` varchar(255) not null   -- 以逗号分隔的用户偏好标签字符串
);

-- Table structure of comments --
drop table if exists `comments`;
create table `comments`(
    `comment_id` bigint not null auto_increment,
    `user_id` bigint not null ,
    `video_id` bigint not null ,
    `parent_id` bigint not null ,
    `like_count` bigint not null default '0',
    `child_count` bigint not null default '0',
    `content` varchar(255) not null ,
    `created_at` varchar(255) not null ,
    `updated_at` varchar(255) not null ,
    `deleted_at` varchar(255)  ,
    primary key (comment_id) ,
    key `vide_index` (video_id) using btree
)engine =InnoDB auto_increment=1 default charset = utf8mb4;

-- Table structure of comment_likes --
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

-- Table structure of follows --
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


/*
drop table if exists `messages`;
create table `messages`(
    `id`           bigint       not null auto_increment comment '自增记录序号',
    `from_user_id` bigint       not null comment '发送者ID',
    `to_user_id`   bigint       not null comment '接受者ID',
    `content`      varchar(255) not null comment '内容',
    `created_at`   bigint    not null comment '创建时间',
    `deleted_at`   bigint    not null comment '删除时间',
    primary key (`id`),
    foreign key (from_user_id) references users(uid) on delete cascade on update cascade,
    foreign key (to_user_id) references users(uid) on delete cascade on update cascade,
    key `from_user_id_to_user_id_index` (`from_user_id`,`to_user_id`) using btree comment '发送者与接受者索引',
    key `from_user_id_to_user_id_created_at_index` (`from_user_id`,`to_user_id`,`created_at`) using btree comment '发送者与接受者的时间段索引',
    key `from_user_id_created_at_index` (`from_user_id`,`created_at`) using btree comment '发送者与发送时间索引', 一般不会用到 
    key `created_at_index` (`created_at`) using btree comment '创建时间索引'  一般不会用到 
) engine =InnoDB auto_increment =10000 default charset =utf8mb4 comment '消息表';

*/
drop table if exists `messages_0`;
drop table if exists `messages_1`;
drop table if exists `messages_2`;
drop table if exists `messages_3`;
create table `messages_0`(
    `id`           bigint       not null auto_increment comment '自增记录序号',
    `from_user_id` bigint       not null comment '发送者ID',
    `to_user_id`   bigint       not null comment '接受者ID',
    `content`      varchar(255) not null comment '内容',
    `created_at`   bigint    not null comment '创建时间',
    `deleted_at`   bigint    not null comment '删除时间',
    primary key (`id`),
    key `from_user_id_to_user_id_index` (`from_user_id`,`to_user_id`) using btree comment '发送者与接受者索引',
    key `from_user_id_to_user_id_created_at_index` (`from_user_id`,`to_user_id`,`created_at`) using btree comment '发送者与接受者的时间段索引',
    key `from_user_id_created_at_index` (`from_user_id`,`created_at`) using btree comment '发送者与发送时间索引', /* 一般不会用到 */
    key `created_at_index` (`created_at`) using btree comment '创建时间索引' /* 一般不会用到 */
) engine =InnoDB auto_increment =10000 default charset =utf8mb4 comment '消息表';
create table `messages_1` like `messages_0`;
create table `messages_2` like `messages_0`;
create table `messages_3` like `messages_0`;