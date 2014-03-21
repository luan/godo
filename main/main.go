package main

import (
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/coopernurse/gorp"
	"github.com/luan/godo/godo"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func SetRoutes(m *martini.ClassicMartini) {
	m.Use(render.Renderer())
	tc := &TasksController{}

	m.Get("/tasks", tc.List)
	m.Post("/tasks", tc.Create)
	m.Patch("/tasks/:id", tc.Update)
	m.Post("/tasks/:id", tc.Update) // hack cos' HTML has no PATCH
}

func main() {
	dbmap := InitDb()
	defer dbmap.Db.Close()
	m := martini.Classic()
	SetRoutes(m)
	m.Run()
}

func InitDb() *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/tasks.bin")
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(godo.Task{}, "tasks").SetKeys(true, "Id")
	godo.Dbmap = dbmap

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
