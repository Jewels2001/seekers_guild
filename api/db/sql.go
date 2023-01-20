package db

var create_users_table string = `
CREATE TABLE IF NOT EXISTS 'users' (
	'id' INTEGER PRIMARY KEY AUTOINCREMENT,
	'name' VARCHAR(64) NOT NULL,
	'discord_name' VARCHAR(64),
	'discord_id' VARCHAR(18),
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

var insert_user string = `
    INSERT INTO users (name, discord_name, discord_id, cur_rank, prestige, tokens)
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
