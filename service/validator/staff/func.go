package staff

import (
	"context"
	"fmt"

	"go-template/domain"

	"github.com/go-playground/validator/v10"
)

func (c *customValidateStaff) checkIDUnique(ctx context.Context, structLV validator.StructLevel, id string) {
	filters := []string{
		fmt.Sprintf("id:eq:%s", id),
	}
	if err := c.repo.Read(ctx, filters, &domain.Staff{}); err == nil {
		structLV.ReportError(id, "id", "ID", "unique", "")
	}
}

func (c *customValidateStaff) checkNameUnique(ctx context.Context, structLV validator.StructLevel, name string) {
	filters := []string{
		fmt.Sprintf("name:eq:%s", name),
	}
	if err := c.repo.Read(ctx, filters, &domain.Staff{}); err == nil {
		structLV.ReportError(name, "name", "name", "unique", "")
	}
}
