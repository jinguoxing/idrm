package category

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ Model = (*CategoryModel)(nil)

type CategoryModel struct {
	conn sqlx.SqlConn
}

// NewModel 创建Model实例
func NewCategoryModel(conn *sql.DB) Model {
	return &CategoryModel{
		conn: sqlx.NewSqlConnFromDB(conn),
	}
}

// Insert 插入类别
func (m *CategoryModel) Insert(ctx context.Context, data *Category) (*Category, error) {
	query := `INSERT INTO category (name, code, parent_id, level, sort, description, status) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := m.conn.ExecCtx(ctx, query,
		data.Name, data.Code, data.ParentId, data.Level, data.Sort, data.Description, data.Status)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	data.Id = id
	return data, nil
}

// FindOne 根据ID查找类别
func (m *CategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	var category Category
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at 
              FROM category WHERE id = ? LIMIT 1`

	err := m.conn.QueryRowCtx(ctx, &category, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, nil
}

// FindByCode 根据code查找类别
func (m *CategoryModel) FindByCode(ctx context.Context, code string) (*Category, error) {
	var category Category
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at 
              FROM category WHERE code = ? LIMIT 1`

	err := m.conn.QueryRowCtx(ctx, &category, query, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // code不存在返回nil
		}
		return nil, err
	}

	return &category, nil
}

// Update 更新类别
func (m *CategoryModel) Update(ctx context.Context, data *Category) error {
	query := `UPDATE category SET name = ?, code = ?, parent_id = ?, level = ?, sort = ?, 
              description = ?, status = ? WHERE id = ?`

	_, err := m.conn.ExecCtx(ctx, query,
		data.Name, data.Code, data.ParentId, data.Level, data.Sort, data.Description, data.Status, data.Id)
	return err
}

// Delete 删除类别
func (m *CategoryModel) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM category WHERE id = ?`
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

// FindAll 查找所有类别
func (m *CategoryModel) FindAll(ctx context.Context) ([]*Category, error) {
	var categories []*Category
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at 
              FROM category ORDER BY sort ASC, id ASC`

	err := m.conn.QueryRowsCtx(ctx, &categories, query)
	return categories, err
}

// FindByParentId 根据父ID查找子类别
func (m *CategoryModel) FindByParentId(ctx context.Context, parentId int64) ([]*Category, error) {
	var categories []*Category
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at 
              FROM category WHERE parent_id = ? ORDER BY sort ASC, id ASC`

	err := m.conn.QueryRowsCtx(ctx, &categories, query, parentId)
	return categories, err
}

// List 分页查询类别列表
func (m *CategoryModel) List(ctx context.Context, page, pageSize int) ([]*Category, int64, error) {
	var categories []*Category
	var total int64

	// 计算总数
	countQuery := `SELECT COUNT(*) FROM category`
	err := m.conn.QueryRowCtx(ctx, &total, countQuery)
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	query := `SELECT id, name, code, parent_id, level, sort, description, status, created_at, updated_at 
              FROM category ORDER BY sort ASC, id ASC LIMIT ? OFFSET ?`

	err = m.conn.QueryRowsCtx(ctx, &categories, query, pageSize, offset)
	return categories, total, err
}

// WithTx 返回带事务的Model实例
func (m *CategoryModel) WithTx(tx interface{}) Model {
	if sqlxConn, ok := tx.(sqlx.SqlConn); ok {
		return &CategoryModel{conn: sqlxConn}
	}
	// 如果不是sqlx连接，返回自身
	return m
}

// Trans 执行事务
func (m *CategoryModel) Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		txConn := sqlx.NewSqlConnFromSession(session)
		txModel := &CategoryModel{conn: txConn}
		return fn(ctx, txModel)
	})
}

// init 注册sqlx工厂
func init() {
	RegisterSqlxFactory(func(db interface{}) Model {
		if sqlDB, ok := db.(*sql.DB); ok {
			return NewCategoryModel(sqlDB)
		}
		panic("invalid database type for sqlx factory")
	})
}
