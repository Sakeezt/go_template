package test

import (
	"go-template/service/staff/inout"
)

func (suite *PackageTestSuite) makeTestCreateInput() (input *inout.StaffCreateInput) {
	return &inout.StaffCreateInput{
		ID:   "test",
		Name: "test",
		Tel:  "test",
	}
}

func (suite *PackageTestSuite) makeTestReadInput() (input *inout.StaffReadInput) {
	return &inout.StaffReadInput{
		ID: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *inout.StaffUpdateInput) {
	return &inout.StaffUpdateInput{
		ID:   "test",
		Name: "test",
		Tel:  "test",
	}
}

func (suite *PackageTestSuite) makeTestDeleteInput() (input *inout.StaffDeleteInput) {
	return &inout.StaffDeleteInput{
		ID: "test",
	}
}
