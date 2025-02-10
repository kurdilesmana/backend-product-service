package helperModel

type BaseFilter struct {
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
	Search        string `json:"search"`
}
