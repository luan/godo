package godo_test

import (
	"github.com/luan/godo/main"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGodo(t *testing.T) {
	RegisterFailHandler(Fail)
	main.InitDb()
	RunSpecs(t, "Godo Suite")
}
