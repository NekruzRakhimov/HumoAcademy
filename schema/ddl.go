package schema

//arrays of DDLs
var (
	CreatingDDLs = []string {CreatingCoursesTable, CreatingNewsTable, CreatingSubscribedUsersTable, CreatingAdminsLevelTable, CreatingAdminsTable, CreatingUsersRolesTable}
	DroppingDDLs = []string {DroppingCoursesTable, DroppingCourseStatusTable, DroppingNewsTable, DroppingAdminsTable, DroppingAdminsLevelTable, DroppingUsersRolesTable, DroppingSubscribedUsersTable}
)

//Creating tables
const (
CreatingNewsTable = `CREATE TABLE IF NOT EXISTS news
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	short_desc TEXT NOT NULL,
	expire_at VARCHAR(255) NOT NULL,
	img VARCHAR(255) NOT NULL,
	full_desc TEXT NOT NULL,
	status BOOLEAN DEFAULT(TRUE)
);`

CreatingCoursesTable = `CREATE TABLE IF NOT EXISTS courses
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	course_durance SERIAL NOT NULL,
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

CreatingUsersRolesTable = `CREATE TABLE IF NOT EXISTS users_roles
(
	id SERIAL NOT NULL UNIQUE PRIMARY KEY,
	role VARCHAR(255) NOT NULL
);`

)

//Dropping tables
const (
	DroppingNewsTable = `DROP TABLE IF EXISTS news;`
	DroppingCoursesTable = `DROP TABLE IF EXISTS courses;`
	DroppingSubscribedUsersTable = `DROP TABLE IF EXISTS subscribed_users;`
	DroppingCourseStatusTable = `DROP TABLE IF EXISTS course_status;`
	DroppingAdminsLevelTable = `DROP TABLE IF EXISTS admins_level;`
	DroppingAdminsTable = `DROP TABLE IF EXISTS admins;`
	DroppingUsersRolesTable = `DROP TABLE IF EXISTS users_roles;`

)
