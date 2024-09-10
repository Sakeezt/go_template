package test

import (
	"github.com/uniplaces/carbon"
)

func (suite *PackageTestSuite) TestGetNow() {
	now, _ := carbon.NowInLocation("Asia/Bangkok")
	suite.Equal(suite.datetime.GetNow().DateTimeString(), now.DateTimeString())
}

func (suite *PackageTestSuite) TestGetUnixNow() {
	suite.Equal(suite.datetime.GetUnixNow(), carbon.Now().Unix())
}

func (suite *PackageTestSuite) TestConvertUnixToDateTime() {
	suite.Equal(suite.datetime.ConvertUnixToDateTime(int64(1577854800)), "2020-01-01 12:00:00")
}

func (suite *PackageTestSuite) TestConvertUnixToDateTimeWithFormat() {
	suite.Equal(suite.datetime.ConvertUnixToDateTimeWithFormat(int64(1577854800), "2006-01-02 15:04"), "2020-01-01 12:00")
}
