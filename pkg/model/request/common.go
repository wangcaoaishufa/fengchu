package request

// PageParam 分页请求基础对象
type PageParam struct {
	PageNumber int    `json:"pageNumber" form:"pageNumber"` // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
	Sorter     string `json:"sorter" form:"sorter"`         // 排序字段
}

const (
	DefaultPageNo   = 1  // 页码最小值为 1
	DefaultPageSize = 15 // 每页条数, 默认15
)

// SortField 排序对象
type SortField struct {
	field string // 排序字段
	order string // 排序顺序
}

const (
	ORDER_ASC  = "ascend"  // 升序
	ORDER_DESC = "descend" // 降序
)
