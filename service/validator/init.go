package validator

import (
	ioStaff "go-template/service/staff/inout"
	"go-template/service/util"
	vStaff "go-template/service/validator/staff"

	"github.com/go-playground/validator/v10"
)

type GoPlayGroundValidator struct {
	validate *validator.Validate
	Repo     *ValidatorRepository
}

type ValidatorRepository struct {
	Staff util.Repository
}

func New(vRepo *ValidatorRepository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate: validator.New(),
		Repo:     vRepo,
	}

	_ = v.validate.RegisterValidation("not-empty", notEmptyStringFunc)
	_ = v.validate.RegisterValidation("name-th", nameThFunc)
	_ = v.validate.RegisterValidation("name-en", nameEnFunc)

	// custom validate staff
	cStaff := vStaff.New(vRepo.Staff)
	v.validate.RegisterStructValidation(cStaff.CreateStructLevelValidation, &ioStaff.StaffCreateInput{})

	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
