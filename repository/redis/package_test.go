// +build integration

package redis_test

import (
	"testing"

	"go-template/repository/redis/test"

	"github.com/stretchr/testify/suite"
)

func TestPackageTestSuite(t *testing.T) {
	suite.Run(t, new(test.PackageTestSuite))
}
