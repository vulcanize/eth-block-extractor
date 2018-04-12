package rocks_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRocks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rocks Suite")
}
