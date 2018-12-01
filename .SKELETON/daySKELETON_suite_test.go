package day{{day}}_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDay{{day}}(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "{{year}} Day{{day}} Suite")
}
