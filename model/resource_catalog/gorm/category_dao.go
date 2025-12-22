package gorm

import (
	"context"
	"errors"

	"idrm/model/resource_catalog"

	"gorm.io/gorm"
)

var _ resource_catalog.CategoryModel = (*CategoryDao)(nil)

type CategoryDao struct {
	db *gorm.DB
}

// NewCategoryDao 创建CategoryDao实例
func NewCategoryDao(db *gorm.DB) resource_catalog.CategoryModel {
	return &CategoryDao{db: db}
}

// Insert 插入类别
func (d *CategoryDao) Insert(ctx context.Context, data *resource_catalog.Category) (*resource_catalog.Category, error) {
	if err := d.db.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

// FindOne 根据ID查找类别
func (d *CategoryDao) FindOne(ctx context.Context, id int64) (*resource_catalog.Category, error) {
	var category resource_catalog.Category
	err := d.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}
	return &category, nil
}

// FindByCode 根据code查找类别
func (d *CategoryDao) FindByCode(ctx context.Context, code string) (*resource_catalog.Category, error) {
	var category resource_catalog.Category
	err := d.db.WithContext(ctx).Where("code = ?", code).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // code不存在返回nil
		}
		return nil, err
	}
	return &category, nil
}

// Update 更新类别
func (d *CategoryDao) Update(ctx context.Context, data *resource_catalog.Category) error {
	return d.db.WithContext(ctx).Updates(data).Error
}

// Delete 删除类别
func (d *CategoryDao) Delete(ctx context.Context, id int64) error {
	return d.db.WithContext(ctx).Delete(&resource_catalog.Category{}, id).Error
}

// FindAll 查找所有类别
func (d *CategoryDao) FindAll(ctx context.Context) ([]*resource_catalog.Category, error) {
	var categories []*resource_catalog.Category
	err := d.db.WithContext(ctx).Order("sort ASC, id ASC").Find(&categories).Error
	return categories, err
}

// FindByParentId 根据父ID查找子类别
func (d *CategoryDao) FindByParentId(ctx context.Context, parentId int64) ([]*resource_catalog.Category, error) {
	var categories []*resource_catalog.Category
	err := d.db.WithContext(ctx).
		Where("parent_id = ?", parentId).
		Order("sort ASC, id ASC").
		Find(&categories).Error
	return categories, err
}

// List 分页查询类别列表
func (d *CategoryDao) List(ctx context.Context, page, pageSize int) ([]*resource_catalog.Category, int64, error) {
	var categories []*resource_catalog.Category
	var total int64

	// 计算总数
	if err := d.db.WithContext(ctx).Model(&resource_catalog.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := d.db.WithContext(ctx).
		Offset(offset).
		Limit(pageSize).
		Order("sort ASC, id ASC").
		Find(&categories).Error

	return categories, total, err
}

// WithTx 返回带事务的DAO实例
func (d *CategoryDao) WithTx(tx interface{}) resource_catalog.CategoryModel {
	if gormTx, ok := tx.(*gorm.DB); ok {
		return &CategoryDao{db: gormTx}
	}
	// 如果不是gorm事务，返回自身
	return d
}

// Trans 执行事务
func (d *CategoryDao) Trans(ctx context.Context, fn func(ctx context.Context, model resource_catalog.CategoryModel) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txModel := &CategoryDao{db: tx}
		return fn(ctx, txModel)
	})
}
