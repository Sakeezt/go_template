package wrapper

import (
	"context"

	"go-template/service/staff/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Delete(ctx context.Context, input *inout.StaffDeleteInput) (err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.Delete")
	defer sp.Finish()

	sp.LogKV("ID", input.ID)

	err = wrp.Service.Delete(ctx, input)

	sp.LogKV("Error", err)

	return err
}
