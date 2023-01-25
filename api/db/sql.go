package db

// User SQL

var create_users_table string = `
CREATE TABLE IF NOT EXISTS 'users' (
	'id' INTEGER PRIMARY KEY AUTOINCREMENT,
	'name' VARCHAR(64) NOT NULL,
    'email' VARCHAR(255) UNIQUE NOT NULL,
    'password' VARCHAR(255) NOT NULL,
    -- 'discord_name' VARCHAR(64),
	-- 'discord_id' VARCHAR(18),
	'cur_rank' VARCHAR(255) NOT NULL,
	'prestige' FLOAT64 NOT NULL,
	'tokens' FLOAT64 NOT NULL
);
`

var get_all_users string = `
    SELECT * FROM users;
`

var get_user_by_id string = `
    SELECT * FROM users WHERE id = ?;
`

var get_user_by_email string = `
    SELECT * FROM users WHERE email = ?;
`

var insert_user string = `
    INSERT INTO users (name, email, password, cur_rank, prestige, tokens)
    VALUES (?, ?, ?, ?, ?, ?);
`

var delete_user string = `
    DELETE FROM users
    WHERE id = ?;
`

var update_prestige string = `
    UPDATE users SET prestige = round(prestige + ?, 2)
    WHERE id = ?;
`

var update_tokens string = `
    UPDATE users SET tokens = round(tokens + ?, 2)
    WHERE id = ?;
`

// AUTH SQL

var create_tokens_table string = `
CREATE TABLE IF NOT EXISTS tokens (
    'id' INTEGER PRIMARY KEY AUTOINCREMENT,
    'uid' varchar(16) NOT NULL,
    'aid' varchar(16) NOT NULL
);
`

var insert_token string = `
    INSERT INTO tokens(uid,aid)
    VALUES (?,?);
`

var delete_token string = `
    DELETE FROM tokens
    WHERE id = ?;
`

var check_token string = `
    SELECT * FROM tokens
    WHERE aid = ?
`
