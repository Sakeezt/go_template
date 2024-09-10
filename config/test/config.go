package test

import (
	"errors"

	"go-template/config"
)

func (suite *PackageTestSuite) TestGetConfig() {
	conf := config.New()
	if conf == nil {
		suite.NoError(errors.New("Cannot get config value"))
	}
}
