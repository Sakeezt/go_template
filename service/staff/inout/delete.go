package inout

type StaffDeleteInput struct {
	ID string `json:"-" validate:"required"`
}
