package handler

import (
	"net/http"
	"strconv"

	"idrm/domain/resource_catalog/entity"
	"idrm/domain/resource_catalog/service"
	"idrm/pkg/errorx"
	"idrm/pkg/response"
)

// CategoryHandler 类别处理器
type CategoryHandler struct {
	categoryService *service.CategoryService
}

// NewCategoryHandler 创建类别处理器
func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// Create 创建类别
func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: 解析请求体
	var req entity.Category
	// 需要从r.Body解析JSON到req

	// 调用领域服务
	err := h.categoryService.Create(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, req)
}

// Update 更新类别
func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: 解析请求体
	var req entity.Category
	// 需要从r.Body解析JSON到req

	err := h.categoryService.Update(r.Context(), &req)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, req)
}

// Delete 删除类别
func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// 获取ID参数
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, errorx.NewWithCode(errorx.ErrCodeParamInvalid))
		return
	}

	err = h.categoryService.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, nil)
}

// GetByID 根据ID获取类别
func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.Error(w, errorx.NewWithCode(errorx.ErrCodeParamInvalid))
		return
	}

	category, err := h.categoryService.GetByID(r.Context(), id)
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, category)
}

// List 分页获取类别列表
func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	categories, total, err := h.categoryService.List(r.Context(), page, pageSize)
	if err != nil {
		response.Error(w, err)
		return
	}

	result := map[string]interface{}{
		"list":      categories,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}

	response.Success(w, result)
}

// GetTree 获取类别树
func (h *CategoryHandler) GetTree(w http.ResponseWriter, r *http.Request) {
	tree, err := h.categoryService.GetTree(r.Context())
	if err != nil {
		response.Error(w, err)
		return
	}

	response.Success(w, tree)
}
