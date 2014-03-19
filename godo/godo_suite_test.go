package godo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGodo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Godo Suite")
}
