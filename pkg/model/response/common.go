package response

// PageResult 通用分页返回对象
type PageResult struct {
	Content interface{} `json:"content"`
	Total   int64       `json:"total"`
}
