package sqlx

import (
	"context"
	"fmt"

	"idrm/model/resource_catalog"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ resource_catalog.CategoryModel = (*CategoryModelSqlx)(nil)

type CategoryModelSqlx struct {
	conn  sqlx.SqlConn
	table string
}

func NewCategoryModel(conn sqlx.SqlConn) resource_catalog.CategoryModel {
	return &CategoryModelSqlx{
		conn:  conn,
		table: "`categories`",
	}
}

func (m *CategoryModelSqlx) Insert(ctx context.Context, data *resource_catalog.Category) error {
	query := fmt.Sprintf("insert into %s (`name`, `code`, `parent_id`, `level`, `sort`, `description`, `status`) values (?, ?, ?, ?, ?, ?, ?)", m.table)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.Code, data.ParentId, data.Level, data.Sort, data.Description, data.Status)
	if err != nil {
		return err
	}

	id, err := ret.LastInsertId()
	if err != nil {
		return err
	}

	data.Id = id
	return nil
}

func (m *CategoryModelSqlx) FindOne(ctx context.Context, id int64) (*resource_catalog.Category, error) {
	query := fmt.Sprintf("select * from %s where `id` = ? limit 1", m.table)
	var resp resource_catalog.Category
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, resource_catalog.ErrNotFound
	default:
		return nil, err
	}
}

func (m *CategoryModelSqlx) FindByCode(ctx context.Context, code string) (*resource_catalog.Category, error) {
	query := fmt.Sprintf("select * from %s where `code` = ? limit 1", m.table)
	var resp resource_catalog.Category
	err := m.conn.QueryRowCtx(ctx, &resp, query, code)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, resource_catalog.ErrNotFound
	default:
		return nil, err
	}
}

func (m *CategoryModelSqlx) Update(ctx context.Context, data *resource_catalog.Category) error {
	query := fmt.Sprintf("update %s set `name` = ?, `code` = ?, `parent_id` = ?, `level` = ?, `sort` = ?, `description` = ?, `status` = ? where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.Code, data.ParentId, data.Level, data.Sort, data.Description, data.Status, data.Id)
	return err
}

func (m *CategoryModelSqlx) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *CategoryModelSqlx) List(ctx context.Context, page, pageSize int) ([]*resource_catalog.Category, int64, error) {
	// 查询总数
	var total int64
	countQuery := fmt.Sprintf("select count(*) from %s", m.table)
	err := m.conn.QueryRowCtx(ctx, &total, countQuery)
	if err != nil {
		return nil, 0, err
	}

	// 查询列表
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select * from %s order by `sort` asc, `id` desc limit ? offset ?", m.table)
	var list []*resource_catalog.Category
	err = m.conn.QueryRowsCtx(ctx, &list, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (m *CategoryModelSqlx) FindByParentId(ctx context.Context, parentId int64) ([]*resource_catalog.Category, error) {
	query := fmt.Sprintf("select * from %s where `parent_id` = ? order by `sort` asc", m.table)
	var list []*resource_catalog.Category
	err := m.conn.QueryRowsCtx(ctx, &list, query, parentId)
	return list, err
}
