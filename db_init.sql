LOAD DATA INFILE '/var/lib/mysql-files/economicus_users.csv'
    INTO TABLE users
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (id, @vcreated_at, @vupdated_at, @vdeleted_at, user_active, access_level, @vlast_login, name, email, password, auth_resource)
    SET
        created_at = NULLIF(@vcreated_at,''),
        updated_at = NULLIF(@vupdated_at,''),
        deleted_at = NULLIF(@vdeleted_at,''),
        last_login = NULLIF(@vlast_login, '');

LOAD DATA INFILE '/var/lib/mysql-files/economicus_profiles.csv'
    INTO TABLE profiles
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (id, @vcreated_at, @vupdated_at, @vdeleted_at, user_id, nickname, profile_image, birth, email, @vphone, @vuser_url, @vintro_message, location_country, location_city)
    SET
        created_at = NULLIF(@vcreated_at,''),
        updated_at = NULLIF(@vupdated_at,''),
        deleted_at = NULLIF(@vdeleted_at,''),
        phone = NULLIF(@vphone, ''),
        user_url = NULLIF(@vuser_url, ''),
        intro_message = NULLIF(@vintro_message, '');

LOAD DATA INFILE '/var/lib/mysql-files/economicus_quants.csv'
    INTO TABLE quants
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (id, @vcreated_at, @vupdated_at, @vdeleted_at, user_id, active, name, profit_rate, description, private)
    SET
        created_at = NULLIF(@vcreated_at,''),
        updated_at = NULLIF(@vupdated_at,''),
        deleted_at = NULLIF(@vdeleted_at,'');

LOAD DATA INFILE '/var/lib/mysql-files/economicus_followings.csv'
    INTO TABLE followings
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (user_id, following_id);

LOAD DATA INFILE '/var/lib/mysql-files/economicus_user_favorite_quants.csv'
    INTO TABLE user_favorite_quants
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (user_id, quant_id);

LOAD DATA INFILE '/var/lib/mysql-files/economicus_comments.csv'
    INTO TABLE comments
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (id, @vcreated_at, @vupdated_at, @vdeleted_at, user_id, quant_id, content)
    SET
        created_at = NULLIF(@vcreated_at,''),
        updated_at = NULLIF(@vupdated_at,''),
        deleted_at = NULLIF(@vdeleted_at,'');

LOAD DATA INFILE '/var/lib/mysql-files/economicus_replies.csv'
    INTO TABLE replies
    FIELDS TERMINATED BY ','
    LINES TERMINATED BY '\n'
    (id, @vcreated_at, @vupdated_at, @vdeleted_at, comment_id, user_id, content)
    SET
        created_at = NULLIF(@vcreated_at,''),
        updated_at = NULLIF(@vupdated_at,''),
        deleted_at = NULLIF(@vdeleted_at,'');