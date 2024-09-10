package test

import (
	"go-template/app/view"
)

func (suite *PackageTestSuite) TestMsgGetDataSuccess() {
	msg := view.MsgGetDataSuccess(suite.ctx)
	suite.Equal(view.MsgData["get_data_success"].EN, msg)
}

func (suite *PackageTestSuite) TestMsgCreateDataSuccess() {
	msg := view.MsgCreateDataSuccess(suite.ctx)
	suite.Equal(view.MsgData["create_data_success"].EN, msg)
}

func (suite *PackageTestSuite) TestMsgUpdateDataSuccess() {
	msg := view.MsgUpdateDataSuccess(suite.ctx)
	suite.Equal(view.MsgData["update_data_success"].EN, msg)
}

func (suite *PackageTestSuite) TestMsgDeleteDataSuccess() {
	msg := view.MsgDeleteDataSuccess(suite.ctx)
	suite.Equal(view.MsgData["delete_data_success"].EN, msg)
}

func (suite *PackageTestSuite) TestMsgAttemptError() {
	msg := view.MsgAttemptError(suite.ctx)
	suite.Equal(view.MsgData["attempt_error"].EN, msg)
}

func (suite *PackageTestSuite) TestMsgNotFoundData() {
	msg := view.MsgNotFoundData(suite.ctx)
	suite.Equal(view.MsgData["not_found"].EN, msg)
}
