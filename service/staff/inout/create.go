package inout

import (
	"go-template/domain"
	"go-template/service/util"

	"github.com/modern-go/reflect2"
)

type StaffCreateInput struct {
	ID     string `json:"id" swaggerignore:"true"`
	Name   string `json:"name" validate:"required,not-empty" example:"John Smith"`
	Tel    string `json:"tel" example:"0900000000"`
	Status string `json:"status" validate:"required,oneof=active inactive" example:"active"`
	Detail string `json:"detail" example:"detail"`
}

func (input *StaffCreateInput) ToDomain(datetime util.DateTime) (staff *domain.Staff) {
	if reflect2.IsNil(input) {
		return &domain.Staff{}
	}
	return &domain.Staff{
		ID:        input.ID,
		Name:      input.Name,
		Tel:       input.Tel,
		Status:    input.Status,
		Detail:    input.Detail,
		CreatedAt: datetime.GetUnixNow(),
		UpdatedAt: datetime.GetUnixNow(),
	}
}
