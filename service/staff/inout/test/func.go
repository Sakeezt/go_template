package test

import (
	"go-template/domain"
	"go-template/service/staff/inout"

	"github.com/uniplaces/carbon"
)

func (suite *PackageTestSuite) makeTestStaff() (staff *domain.Staff) {
	now := carbon.Now()
	return &domain.Staff{
		Name:      "test",
		Tel:       "test",
		Detail:    "test",
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
}

func (suite *PackageTestSuite) makeTestCreateInput() (input *inout.StaffCreateInput) {
	return &inout.StaffCreateInput{
		ID:     "test",
		Name:   "test",
		Tel:    "test",
		Detail: "test",
	}
}

func (suite *PackageTestSuite) makeTestUpdateInput() (input *inout.StaffUpdateInput) {
	return &inout.StaffUpdateInput{
		ID:     "test",
		Name:   "test",
		Tel:    "test",
		Detail: "test",
	}
}
