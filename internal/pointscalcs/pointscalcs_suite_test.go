package pointscalcs_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPointscalcs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pointscalcs Suite")
}
