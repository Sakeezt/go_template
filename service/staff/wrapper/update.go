package wrapper

import (
	"context"

	"go-template/service/staff/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Update(ctx context.Context, input *inout.StaffUpdateInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Update")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)
	sp.LogKV("Name", input.Name)
	sp.LogKV("Tel", input.Tel)

	err = wrp.Service.Update(ctx, input)

	sp.LogKV("Error", err)

	return err
}
