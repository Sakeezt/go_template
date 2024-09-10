package staff

import (
	"context"

	"go-template/domain"
	"go-template/service/staff/inout"
)

//go:generate mockery --name=Service
type Service interface {
	List(ctx context.Context, opt *domain.PageOption) (total int, items []*inout.StaffView, err error)
	Create(ctx context.Context, input *inout.StaffCreateInput) (ID string, err error)
	Read(ctx context.Context, input *inout.StaffReadInput) (staff *inout.StaffView, err error)
	Update(ctx context.Context, input *inout.StaffUpdateInput) (err error)
	Delete(ctx context.Context, input *inout.StaffDeleteInput) (err error)
}
