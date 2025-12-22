package gorm

import (
	"context"

	"idrm/model/resource_catalog"

	"gorm.io/gorm"
)

var _ resource_catalog.CategoryModel = (*CategoryModelGorm)(nil)

type CategoryModelGorm struct {
	db *gorm.DB
}

func NewCategoryModel(db *gorm.DB) resource_catalog.CategoryModel {
	return &CategoryModelGorm{db: db}
}

func (m *CategoryModelGorm) Insert(ctx context.Context, data *resource_catalog.Category) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *CategoryModelGorm) FindOne(ctx context.Context, id int64) (*resource_catalog.Category, error) {
	var category resource_catalog.Category
	err := m.db.WithContext(ctx).First(&category, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, resource_catalog.ErrNotFound
	}
	return &category, err
}

func (m *CategoryModelGorm) FindByCode(ctx context.Context, code string) (*resource_catalog.Category, error) {
	var category resource_catalog.Category
	err := m.db.WithContext(ctx).Where("code = ?", code).First(&category).Error
	if err == gorm.ErrRecordNotFound {
		return nil, resource_catalog.ErrNotFound
	}
	return &category, err
}

func (m *CategoryModelGorm) Update(ctx context.Context, data *resource_catalog.Category) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *CategoryModelGorm) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Delete(&resource_catalog.Category{}, id).Error
}

func (m *CategoryModelGorm) List(ctx context.Context, page, pageSize int) ([]*resource_catalog.Category, int64, error) {
	var categories []*resource_catalog.Category
	var total int64

	offset := (page - 1) * pageSize

	// 查询总数
	err := m.db.WithContext(ctx).Model(&resource_catalog.Category{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询列表
	err = m.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Order("sort ASC, id DESC").
		Find(&categories).Error

	return categories, total, err
}

func (m *CategoryModelGorm) FindByParentId(ctx context.Context, parentId int64) ([]*resource_catalog.Category, error) {
	var categories []*resource_catalog.Category
	err := m.db.WithContext(ctx).
		Where("parent_id = ?", parentId).
		Order("sort ASC").
		Find(&categories).Error
	return categories, err
}
