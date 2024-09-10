package implement

import (
	"context"

	"go-template/service/staff/inout"
	"go-template/service/util"
)

func (impl *implementation) Create(ctx context.Context, input *inout.StaffCreateInput) (id string, err error) {
	if err = impl.Validator.Validate(input); err != nil {
		return "", util.ValidationCreateErr(err)
	}

	input.ID = impl.UUID.Generate()
	staff := input.ToDomain(impl.DateTime)

	_, err = impl.RepoStaff.Create(ctx, staff)
	if err != nil {
		return "", util.RepoCreateErr(err)
	}

	return input.ID, nil
}
