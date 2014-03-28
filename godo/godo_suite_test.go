package godo_test

import (
	"github.com/luan/godo/godo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestGodo(t *testing.T) {
	RegisterFailHandler(Fail)
	dbmap := godo.InitDb("test_tasks.bin")
	defer dbmap.Db.Close()
	RunSpecs(t, "Godo Suite")
}
