package wrapper

import (
	"go-template/service/staff"
)

type Wrapper struct {
	Service staff.Service
}

func _(service staff.Service) staff.Service {
	return &Wrapper{service}
}
