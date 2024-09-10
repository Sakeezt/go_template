package inout

import (
	"go-template/domain"
	"go-template/service/util"

	"github.com/modern-go/reflect2"
)

type StaffUpdateInput struct {
	ID     string `json:"id" validate:"required" swaggerignore:"true"`
	Name   string `json:"name" validate:"required,not-empty" example:"John Smith"`
	Tel    string `json:"tel" example:"0900000000"`
	Status string `json:"status" validate:"required,oneof=active inactive" example:"active"`
	Detail string `json:"detail" example:"detail"`
}

func (input *StaffUpdateInput) ToDomain(datetime util.DateTime) (staff *domain.Staff) {
	if reflect2.IsNil(input) {
		return &domain.Staff{}
	}
	return &domain.Staff{
		Name:      input.Name,
		Tel:       input.Tel,
		Status:    input.Status,
		Detail:    input.Detail,
		UpdatedAt: datetime.GetUnixNow(),
	}
}
