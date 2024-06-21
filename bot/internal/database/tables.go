package database

const CREATE_BOT_TABLE string = `
	create table if not exists bot (
		id             integer primary key autoincrement,
		name           text not null unique,
		status         int default 0,
		domain         text not null,
		interval_secs  int default 0 not null,
		created        datetime default current_timestamp,
		updated        datetime default current_timestamp
	);`

const CREATE_BOT_INDEX string = `create index if not exists bot_index on bot(id);`

const CREATE_BOT_ALLOWED_DOMAINS_TABLE string = `
	create table if not exists allowed_domain (
		id      integer primary key autoincrement,
		bot_id  int references bot,
		domain  text not null
	);`

const CREATE_BOT_LOG_TABLE string = `
	create table if not exists bot_log (
		id         integer primary key autoincrement,
		timestamp  datetime default current_timestamp not null,
		bot_id     int references bot,
		message    text not null
	);`

const CREATE_HTTP_LOG_TABLE string = `
	create table if not exists http_log (
		id         integer primary key autoincrement,
		timestamp  datetime default current_timestamp not null,
		client_ip  text not null,
		remote_ip  text not null,
		method     text not null,
		url        text not null,
		host       text not null,
		path       text not null,
		body       text
	);`
