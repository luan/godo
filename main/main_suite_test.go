package main_test

import (
	"bytes"
	"database/sql"
	"github.com/codegangsta/martini"
	"github.com/coopernurse/gorp"
	"github.com/luan/godo/godo"
	. "github.com/luan/godo/main"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	response       *httptest.ResponseRecorder
	martiniClassic *martini.ClassicMartini
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	martiniClassic = martini.Classic()
	InitDb()
	SetRoutes(martiniClassic)
	RunSpecs(t, "Main Suite")
}

func MethodsFor(path string) []string {
	return martiniClassic.MethodsFor(path)
}

func Get(route string) {
	request, _ := http.NewRequest("GET", route, nil)
	response = httptest.NewRecorder()
	martiniClassic.ServeHTTP(response, request)
}

func Post(route string, params map[string]string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range params {
		writer.WriteField(k, v)
	}

	request, _ := http.NewRequest("POST", route, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()

	response = httptest.NewRecorder()
	martiniClassic.ServeHTTP(response, request)
}

func Patch(route string, params map[string]string) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	for k, v := range params {
		writer.WriteField(k, v)
	}
	request, _ := http.NewRequest("PATCH", route, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	response = httptest.NewRecorder()
	martiniClassic.ServeHTTP(response, request)
}

func SetUpDb() {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, _ := sql.Open("sqlite3", "/tmp/tasks_test.bin")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbmap.AddTableWithName(godo.Task{}, "tasks").SetKeys(true, "Id")
	godo.Dbmap = dbmap

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	_ = dbmap.CreateTablesIfNotExists()
}
