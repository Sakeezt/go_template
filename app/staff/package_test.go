package staff_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"go-template/app/staff/test"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
