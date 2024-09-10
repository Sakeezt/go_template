package inout

type StaffReadInput struct {
	ID string `json:"-" validate:"required"`
}
