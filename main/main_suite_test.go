package main_test

import (
	"bytes"
	"github.com/codegangsta/martini"
	. "github.com/luan/godo/main"
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
