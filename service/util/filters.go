package util

import (
	"fmt"
)

//go:generate mockery --name=FilterString
type FilterString interface {
	MakeIDFilters(id string) (filters []string)
	MakeID(id string) (filter string)
	MakeNotEqualID(id string) (filter string)
	MakeDeletedAtIsNull() (filter string)
	MakeGateGroupNameFiltersNotEqualString(gateGroupName string) (filters string)
	MakeFilterEventAtBetweenString(start, end int64) (filters string)
}

type filter struct{}

func NewFilterString() (filterString FilterString) {
	return &filter{}
}

func (f *filter) MakeIDFilters(id string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", id),
	}
}

func (f *filter) MakeID(id string) (filter string) {
	return fmt.Sprintf("_id:eq:%s", id)
}

func (f *filter) MakeNotEqualID(id string) (filter string) {
	return fmt.Sprintf("_id:ne:%s", id)
}

func (f *filter) MakeDeletedAtIsNull() (filter string) {
	return "deleted_at:isNull"
}

func (f *filter) MakeGateGroupNameFiltersNotEqualString(gateGroupName string) (filters string) {
	return fmt.Sprintf("gateGroupName:ne:%s", gateGroupName)
}

func (f *filter) MakeFilterEventAtBetweenString(start, end int64) (filters string) {
	return fmt.Sprintf("eventAt:between:%d|%d", start, end)
}
