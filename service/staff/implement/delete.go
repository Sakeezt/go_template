package implement

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"
	"go-template/service/util"
)

func (impl *implementation) Delete(ctx context.Context, input *inout.StaffDeleteInput) (err error) {
	filters := []string{
		impl.FilterString.MakeID(input.ID),
		impl.FilterString.MakeDeletedAtIsNull(),
	}

	if err = impl.RepoStaff.Read(ctx, filters, &domain.Staff{}); err != nil {
		return util.RepoReadErr(err)
	}

	if err = impl.RepoStaff.SoftDelete(ctx, filters); err != nil {
		return util.RepoDeleteErr(err)
	}

	return nil
}
