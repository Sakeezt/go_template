package wrapper

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) List(ctx context.Context, opt *domain.PageOption) (total int, list []*inout.StaffView, err error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "service.staff.List")
	defer sp.Finish()

	sp.LogKV("Page", opt.Page)
	sp.LogKV("PerPage", opt.PerPage)
	sp.LogKV("Filters", opt.Filters)

	total, list, err = wrp.Service.List(ctx, opt)

	sp.LogKV("Total", total)
	sp.LogKV("List", list)
	sp.LogKV("Error", err)

	return total, list, err
}
