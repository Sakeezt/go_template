package logs_test

import (
	"testing"

	"go-template/service/util/logs/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
