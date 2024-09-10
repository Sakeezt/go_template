package staff

import (
	"go-template/service/util"
)

type customValidateStaff struct {
	repo util.Repository
}

func New(repo util.Repository) (customValidate *customValidateStaff) {
	return &customValidateStaff{
		repo: repo,
	}
}
