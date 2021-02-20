package schema

//arrays of DDLs
var (
	CreatingDDLs = []string {CreatingCoursesTable, CreatingNewsTable, CreatingSubscribedUsersTable, CreatingAdminsLevelTable, CreatingAdminsTable, CreatingUsersTable}
	DroppingDDLs = []string {DroppingUsersTable, DroppingNewsTable, DroppingAdminsTable, DroppingAdminsLevelTable, DroppingUsersRolesTable, DroppingSubscribedUsersTable, DroppingCoursesTable}
)

//Creating tables
const (
CreatingNewsTable = `CREATE TABLE IF NOT EXISTS news
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	short_desc TEXT NOT NULL,
	expire_at BIGSERIAL NOT NULL,
	img VARCHAR(255) NOT NULL,
	full_desc TEXT NOT NULL,
	status BOOLEAN DEFAULT(TRUE)
);`

CreatingCoursesTable = `CREATE TABLE IF NOT EXISTS courses
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	course_durance VARCHAR(255) NOT NULL,
	img VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	plans TEXT NOT NULL,
	status BOOLEAN DEFAULT(TRUE)
);`

CreatingSubscribedUsersTable = `CREATE TABLE IF NOT EXISTS subscribed_users
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	UNIQUE(email)
);`

CreatingAdminsLevelTable = `CREATE TABLE IF NOT EXISTS admins_level
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	level VARCHAR(255) NOT NULL 
);`

CreatingAdminsTable = `CREATE TABLE IF NOT EXISTS admins
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	name  VARCHAR(255) NOT NULL,
	username VARCHAR(255) NOT NULL UNIQUE,
	password_hash VARCHAR(255) NOT NULL,
	level SERIAL REFERENCES admins_level(id) NOT NULL,
	UNIQUE(username)
);`

CreatingUsersTable = `CREATE TABLE IF NOT EXISTS users
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	middle_name VARCHAR(255),
	email VARCHAR(255) NOT NULL,
	about TEXT NOT NULL,
	cv VARCHAR(255) NOT NULL,
	course_id SERIAL REFERENCES courses(id) ON DELETE CASCADE NOT NULL ,
	UNIQUE(email, course_id)
);`
)

//Dropping tables
const (
	DroppingNewsTable = `DROP TABLE IF EXISTS news;`
	DroppingCoursesTable = `DROP TABLE IF EXISTS courses;`
	DroppingSubscribedUsersTable = `DROP TABLE IF EXISTS subscribed_users;`
	DroppingAdminsLevelTable = `DROP TABLE IF EXISTS admins_level;`
	DroppingAdminsTable = `DROP TABLE IF EXISTS admins;`
	DroppingUsersRolesTable = `DROP TABLE IF EXISTS users_roles;`
	DroppingUsersTable = `DROP TABLE IF EXISTS users;`
)
