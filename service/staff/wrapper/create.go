package wrapper

import (
	"context"

	"go-template/service/staff/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create(ctx context.Context, input *inout.StaffCreateInput) (id string, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Create")
	defer sp.Finish()

	sp.LogKV("Name", input.Name)
	sp.LogKV("Tel", input.Tel)

	id, err = wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("Error", err)

	return id, err
}
