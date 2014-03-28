package godo

import (
	"errors"
	"github.com/coopernurse/gorp"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

type Task struct {
	Id     int
	Name   string
	Status string
}

func (t *Task) NextStatus() string {
	if t.Status == "done" {
		return "pending"
	}
	return "done"
}

var Dbmap *gorp.DbMap

func InitDb(dbname string) *gorp.DbMap {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("sqlite3", "/tmp/"+dbname)
	CheckErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(Task{}, "tasks").SetKeys(true, "Id")
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


func AddTask(name string) Task {
	t := Task{Name: name, Status: "pending"}
	_ = Dbmap.Insert(&t)
	return t
}

func FindTask(id int) (t Task, err error) {
	obj, err := Dbmap.Get(Task{}, id)
	if err == nil && obj != nil {
		t = *(obj.(*Task))
	} else if err == nil {
		err = errors.New("Couldn't find task with Id: "+strconv.Itoa(id))
	}

	return
}

func Tasks() (tasks []Task) {
	_, _ = Dbmap.Select(&tasks, "select * from tasks order by id")
	return
}

func UpdateTask(task Task) (_, err error) {
	_, err = Dbmap.Update(&task)
	return
}

func ResetTasks() {
	Dbmap.TruncateTables()
}
