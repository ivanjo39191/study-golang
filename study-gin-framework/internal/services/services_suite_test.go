package services

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestVideoService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blog Service Test Suite")
}
