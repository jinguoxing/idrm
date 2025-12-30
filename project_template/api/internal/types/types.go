package types

// HealthResp 健康检查响应
type HealthResp struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// PageInfo 分页信息
type PageInfo struct {
	Page     int64 `form:"page,default=1"`
	PageSize int64 `form:"pageSize,default=20"`
}

// PageResult 分页返回结果
type PageResult struct {
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	PageSize int64       `json:"pageSize"`
	List     interface{} `json:"list"`
}

// IdReq ID请求
type IdReq struct {
	Id int64 `path:"id"`
}

// EmptyResp 空响应
type EmptyResp struct{}
