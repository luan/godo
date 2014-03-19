package main_test

import (
	"github.com/codegangsta/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var (
	response *httptest.ResponseRecorder
)

func Get(m *martini.ClassicMartini, route string) {
	request, _ := http.NewRequest("GET", route, nil)
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}

func Post(m *martini.ClassicMartini, route string, params map[string]string) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	for k, v := range params {
		writer.WriteField(k, v)
	}
	request, _ := http.NewRequest("POST", route, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}

func Patch(m *martini.ClassicMartini, route string, params map[string]string) {
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	for k, v := range params {
		writer.WriteField(k, v)
	}
	request, _ := http.NewRequest("PATCH", route, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	writer.Close()
	response = httptest.NewRecorder()
	m.ServeHTTP(response, request)
}
