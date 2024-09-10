package inout

import (
	"go-template/domain"
	"go-template/service/util"
)

type StaffView struct {
	ID        string `json:"id" example:"123456789012345678"`
	Name      string `json:"name" example:"John Smith"`
	Tel       string `json:"tel" example:"0900000000"`
	Status    string `json:"status" example:"active"`
	Detail    string `json:"detail" example:"detail"`
	CreatedAt string `json:"created_at" example:"2016-01-02 15:04:05"`
	UpdatedAt string `json:"updated_at" example:"2016-01-02 15:04:05"`
}

func StaffToView(staff *domain.Staff, datetime util.DateTime) (view *StaffView) {
	return &StaffView{
		ID:        staff.ID,
		Name:      staff.Name,
		Tel:       staff.Tel,
		Status:    staff.Status,
		Detail:    staff.Detail,
		CreatedAt: datetime.ConvertUnixToDateTime(staff.CreatedAt),
		UpdatedAt: datetime.ConvertUnixToDateTime(staff.UpdatedAt),
	}
}
