package implement

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"
	"go-template/service/util"

	"github.com/imdario/mergo"
)

func (impl *implementation) Update(ctx context.Context, input *inout.StaffUpdateInput) (err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return util.ValidationUpdateErr(err)
	}

	staff := &domain.Staff{}
	filters := impl.FilterString.MakeIDFilters(input.ID)

	if err = impl.RepoStaff.Read(ctx, filters, staff); err != nil {
		return util.RepoReadErr(err)
	}

	update := input.ToDomain(impl.DateTime)

	if err = mergo.Merge(staff, update, mergo.WithOverride); err != nil {
		return util.UnknownErr(err)
	}

	if err = impl.RepoStaff.Update(ctx, filters, staff); err != nil {
		return util.RepoUpdateErr(err)
	}

	return nil
}
