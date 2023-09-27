USE tiktok;

CREATE TABLE AUTHENTICATION_INFO(
	user_id BIGINT UNSIGNED,
    token VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id)
);


CREATE TABLE USER_INFO(
	user_id BIGINT UNSIGNED,
    user_name VARCHAR(255) NOT NULL,
    follow_count BIGINT UNSIGNED,
    follower_count BIGINT UNSIGNED,
    avatar TEXT,
    background_image TEXT,
    signature TEXT,
    total_favorited BIGINT UNSIGNED,
    work_count BIGINT UNSIGNED,
    favorite_count BIGINT UNSIGNED,
    PRIMARY KEY (user_id)
);

CREATE TABLE VIDEO_INFO(
	video_id BIGINT UNSIGNED,
    author_id BIGINT UNSIGNED,
    play_url TEXT NOT NULL,
    cover_url TEXT NOT NULL,
    title TEXT NOT NULL,
    favorite_count BIGINT UNSIGNED,
    comment_count BIGINT UNSIGNED,
    upload_time DATETIME NOT NULL,
    PRIMARY KEY (video_id, author_id),
    INDEX(upload_time)
);

CREATE TABLE COMMENT_INFO(
	comment_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    video_id BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL, 
    create_date DATE NOT NULL,
    INDEX(user_id, video_id)
);

CREATE TABLE FOLLOW_INFO(
	user_id BIGINT UNSIGNED NOT NULL,
    target_user_id BIGINT UNSIGNED NOT NULL,
    INDEX (user_id, target_user_id)
);

CREATE TABLE LIKE_INFO(
	user_id BIGINT UNSIGNED NOT NULL,
    video_id BIGINT UNSIGNED NOT NULL,
    like_time DATETIME NOT NULL,
    INDEX (user_id)
);

CREATE TABLE COLLECT_INFO(
	user_id BIGINT UNSIGNED NOT NULL,
    video_id BIGINT UNSIGNED NOT NULL,
    collect_time DATETIME NOT NULL,
    INDEX (user_id)
);

CREATE TABLE FRIEND_INFO(
	user_id BIGINT UNSIGNED NOT NULL,
    friend_id BIGINT UNSIGNED NOT NULL,
    msg_type INT NOT NULL,
    msg TEXT,
    last_chat_time DATETIME NOT NULL,
    INDEX(user_id)
);

CREATE TABLE CHAT_INFO(
	msg_id BIGINT UNSIGNED AUTO_INCREMENT,
    user_id BIGINT UNSIGNED NOT NULL,
    target_id BIGINT UNSIGNED NOT NULL,
    content TEXT NOT NULL,
    create_time DATETIME NOT NULL,
    INDEX(user_id),
    PRIMARY KEY (msg_id)
);




