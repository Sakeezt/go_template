package implement

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"
	"go-template/service/util"
)

func (impl *implementation) List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.StaffView, err error) {
	if err = impl.Validator.Validate(opt); err != nil {
		return 0, nil, util.ValidationParamOptionErr(err)
	}

	total, records, err := impl.RepoStaff.List(ctx, opt, &domain.Staff{})
	if err != nil {
		return 0, nil, util.RepoListErr(err)
	}

	items = make([]*inout.StaffView, len(records))
	for i, record := range records {
		items[i] = inout.StaffToView(record.(*domain.Staff), impl.DateTime)
	}

	return total, items, nil
}
