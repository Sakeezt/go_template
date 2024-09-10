package implement

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"
	"go-template/service/util"
)

func (impl *implementation) Read(ctx context.Context, input *inout.StaffReadInput) (view *inout.StaffView, err error) {
	staff := &domain.Staff{}
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoStaff.Read(ctx, filters, staff); err != nil {
		return nil, util.RepoReadErr(err)
	}

	return inout.StaffToView(staff, impl.DateTime), nil
}
