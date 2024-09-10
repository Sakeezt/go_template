package wrapper

import (
	"context"

	"go-template/service/staff/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Read(ctx context.Context, input *inout.StaffReadInput) (view *inout.StaffView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Read")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	view, err = wrp.Service.Read(ctx, input)

	sp.LogKV("Data", view)
	sp.LogKV("Err", err)

	return view, err
}
