package godo

import(
	"github.com/jmoiron/modl"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var Dbmap *modl.DbMap

func InitDb(dbname string) *modl.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/"+dbname)
	CheckErr(err, "sql.Open failed")

	// construct a modl DbMap
	dbmap := modl.NewDbMap(db, modl.SqliteDialect{})

	// add a table, setting the table name to 'posts' and
	// specifying that the ID property is an auto incrementing PK
	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "ID")
	dbmap.AddTableWithName(Project{}, "projects").SetKeys(true, "ID")
	Dbmap = dbmap

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Create tables failed")

	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func ResetDatabase() {
	Dbmap.TruncateTables()
}
