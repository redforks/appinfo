package appinfo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppinfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Appinfo Suite")
}
