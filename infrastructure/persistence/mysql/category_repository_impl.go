package mysql

import (
	"context"
	"database/sql"
	"time"

	"idrm/domain/resource_catalog/entity"
	"idrm/domain/resource_catalog/repository"
	"idrm/domain/resource_catalog/service"
)

// CategoryRepositoryImpl Category仓储MySQL实现
type CategoryRepositoryImpl struct {
	db *sql.DB
}

// NewCategoryRepository 创建Category仓储
func NewCategoryRepository(db *sql.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

// Create 创建类别
func (r *CategoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) error {
	query := `INSERT INTO categories (name, code, parent_id, level, sort, description, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := time.Now()
	result, err := r.db.ExecContext(ctx, query,
		category.Name,
		category.Code,
		category.ParentID,
		category.Level,
		category.Sort,
		category.Description,
		category.Status,
		now,
		now,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	category.ID = id
	category.CreatedAt = now
	category.UpdatedAt = now

	return nil
}

// Update 更新类别
func (r *CategoryRepositoryImpl) Update(ctx context.Context, category *entity.Category) error {
	query := `UPDATE categories SET name = ?, code = ?, parent_id = ?, level = ?, sort = ?, 
		description = ?, status = ?, updated_at = ? WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query,
		category.Name,
		category.Code,
		category.ParentID,
		category.Level,
		category.Sort,
		category.Description,
		category.Status,
		time.Now(),
		category.ID,
	)

	return err
}

// Delete 删除类别
func (r *CategoryRepositoryImpl) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM categories WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// FindByID 根据ID查询
func (r *CategoryRepositoryImpl) FindByID(ctx context.Context, id int64) (*entity.Category, error) {
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at
		FROM categories WHERE id = ?`

	category := &entity.Category{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Code,
		&category.ParentID,
		&category.Level,
		&category.Sort,
		&category.Description,
		&category.Status,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, service.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return category, nil
}

// FindByCode 根据编码查询
func (r *CategoryRepositoryImpl) FindByCode(ctx context.Context, code string) (*entity.Category, error) {
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at
		FROM categories WHERE code = ?`

	category := &entity.Category{}
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&category.ID,
		&category.Name,
		&category.Code,
		&category.ParentID,
		&category.Level,
		&category.Sort,
		&category.Description,
		&category.Status,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, service.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return category, nil
}

// FindByParentID 根据父ID查询子类别
func (r *CategoryRepositoryImpl) FindByParentID(ctx context.Context, parentID int64) ([]*entity.Category, error) {
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at
		FROM categories WHERE parent_id = ? ORDER BY sort ASC`

	rows, err := r.db.QueryContext(ctx, query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Code,
			&category.ParentID,
			&category.Level,
			&category.Sort,
			&category.Description,
			&category.Status,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// List 分页查询
func (r *CategoryRepositoryImpl) List(ctx context.Context, page, pageSize int) ([]*entity.Category, int64, error) {
	offset := (page - 1) * pageSize

	// 查询总数
	var total int64
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM categories`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// 查询数据
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at
		FROM categories ORDER BY sort ASC, id DESC LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Code,
			&category.ParentID,
			&category.Level,
			&category.Sort,
			&category.Description,
			&category.Status,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		categories = append(categories, category)
	}

	return categories, total, nil
}

// FindAll 查询所有激活的类别
func (r *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]*entity.Category, error) {
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at
		FROM categories WHERE status = 1 ORDER BY sort ASC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		category := &entity.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Code,
			&category.ParentID,
			&category.Level,
			&category.Sort,
			&category.Description,
			&category.Status,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
