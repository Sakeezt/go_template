package staff

import (
	"context"

	"go-template/service/staff/inout"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateStaff) CreateStructLevelValidation(structLV validator.StructLevel) {
	ctx := context.Background()
	input := structLV.Current().Interface().(inout.StaffCreateInput)

	c.checkIDUnique(ctx, structLV, input.ID)
	c.checkNameUnique(ctx, structLV, input.Name)
}
