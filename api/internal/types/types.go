package types

//  资源目录模块类型定义

type CategoryReq struct {
	Id int64 `path:"id"`
}

type CategoryResp struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ParentId    int64  `json:"parent_id"`
	Level       int    `json:"level"`
	Sort        int    `json:"sort"`
	Description string `json:"description,omitempty"`
	Status      int    `json:"status"`
}

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	ParentId    int64  `json:"parent_id,optional"`
	Level       int    `json:"level,optional,default=1"`
	Sort        int    `json:"sort,optional,default=0"`
	Description string `json:"description,optional"`
}

type UpdateCategoryReq struct {
	Id          int64  `path:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ParentId    int64  `json:"parent_id,optional"`
	Level       int    `json:"level,optional"`
	Sort        int    `json:"sort,optional"`
	Description string `json:"description,optional"`
	Status      int    `json:"status,optional"`
}

type ListCategoryReq struct {
	Page     int `form:"page,optional,default=1"`
	PageSize int `form:"page_size,optional,default=10"`
}

type ListCategoryResp struct {
	List  []CategoryResp `json:"list"`
	Total int64          `json:"total"`
}

// 数据视图模块类型（待添加）
// 数据理解模块类型（待添加）
