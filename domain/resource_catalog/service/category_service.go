package service

import (
	"context"
	"errors"

	"idrm/domain/resource_catalog/entity"
	"idrm/domain/resource_catalog/repository"
)

var (
	ErrCategoryNotFound      = errors.New("类别不存在")
	ErrCategoryAlreadyExists = errors.New("类别已存在")
)

// CategoryService 类别领域服务
type CategoryService struct {
	repo repository.CategoryRepository
}

// NewCategoryService 创建类别服务
func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

// Create 创建类别
func (s *CategoryService) Create(ctx context.Context, category *entity.Category) error {
	// 业务验证
	if err := category.Validate(); err != nil {
		return err
	}

	// 检查编码是否已存在
	existing, err := s.repo.FindByCode(ctx, category.Code)
	if err != nil && err != ErrCategoryNotFound {
		return err
	}
	if existing != nil {
		return ErrCategoryAlreadyExists
	}

	// 创建类别
	return s.repo.Create(ctx, category)
}

// Update 更新类别
func (s *CategoryService) Update(ctx context.Context, category *entity.Category) error {
	// 业务验证
	if err := category.Validate(); err != nil {
		return err
	}

	// 检查是否存在
	existing, err := s.repo.FindByID(ctx, category.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrCategoryNotFound
	}

	// 更新类别
	return s.repo.Update(ctx, category)
}

// Delete 删除类别
func (s *CategoryService) Delete(ctx context.Context, id int64) error {
	// 检查是否存在
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrCategoryNotFound
	}

	// 检查是否有子类别
	children, err := s.repo.FindByParentID(ctx, id)
	if err != nil {
		return err
	}
	if len(children) > 0 {
		return errors.New("存在子类别，无法删除")
	}

	return s.repo.Delete(ctx, id)
}

// GetByID 根据ID获取类别
func (s *CategoryService) GetByID(ctx context.Context, id int64) (*entity.Category, error) {
	return s.repo.FindByID(ctx, id)
}

// GetByCode 根据编码获取类别
func (s *CategoryService) GetByCode(ctx context.Context, code string) (*entity.Category, error) {
	return s.repo.FindByCode(ctx, code)
}

// List 分页获取类别列表
func (s *CategoryService) List(ctx context.Context, page, pageSize int) ([]*entity.Category, int64, error) {
	return s.repo.List(ctx, page, pageSize)
}

// GetTree 获取类别树
func (s *CategoryService) GetTree(ctx context.Context) ([]*entity.Category, error) {
	// 获取所有类别
	categories, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: 构建树形结构
	// 这里可以实现树形结构的构建逻辑

	return categories, nil
}
