package test

import (
	"go-template/service/staff/inout"
)

func (suite *PackageTestSuite) makeTestCreateInput() (input *inout.StaffCreateInput) {
	return &inout.StaffCreateInput{}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *inout.StaffReadInput) {
	return &inout.StaffReadInput{}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *inout.StaffUpdateInput) {
	return &inout.StaffUpdateInput{}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *inout.StaffDeleteInput) {
	return &inout.StaffDeleteInput{}
}
