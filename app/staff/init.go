package staff

import "go-template/service/staff"

type Controller struct {
	service staff.Service
}

func New(staffSvc staff.Service) (staff *Controller) {
	return &Controller{service: staffSvc}
}
