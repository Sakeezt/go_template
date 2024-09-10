package domain

type PageOption struct {
	Page    int      `form:"page,default=1" validate:"min=1"`
	PerPage int      `form:"per_page,default=15" validate:"min=1"`
	Filters []string `form:"filters"`
	Search  []string `form:"search"`
	Sorts   []string `form:"sorts"`
}

type QueryOption struct {
	Filters []string `form:"filters"`
	Search  []string `form:"search"`
	Sorts   []string `form:"sorts"`
	Limit   int64    `form:"limit"`
}

type SetOpParam struct {
	Filters      []string
	SetFieldName string
	Item         interface{}
}
