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
