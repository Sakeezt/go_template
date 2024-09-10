package test

import (
	"io/ioutil"
	"os"

	"github.com/acarl005/stripansi"
)

func (suite *PackageTestSuite) TestPrintln() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Println(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [PNT] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestPrintf() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Printf("%s", text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [PNT] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestDebug() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Debug(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [DEB] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestInfo() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Info(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [INF] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestWarning() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Warning(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [WAR] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestError() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.Error(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [ERR] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestFatal() {
	suite.T().Skip()
}

func (suite *PackageTestSuite) TestPanic() {
	suite.T().Skip()
}

func (suite *PackageTestSuite) TestSystem() {
	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	os.Stdout = w
	suite.NoError(err)

	text := "Test"
	suite.log.System(text)

	err = w.Close()
	suite.NoError(err)
	out, err := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	suite.NoError(err)

	expected := suite.datetime.GetNow().Format("2006-01-02 15:04:05") + " [SYS] " + text + "\n"
	actual := stripansi.Strip(string(out))
	suite.Equal(expected, actual)
}

func (suite *PackageTestSuite) TestStruct() {
	suite.T().Skip()
}

func (suite *PackageTestSuite) TestLine() {
	suite.T().Skip()
}
