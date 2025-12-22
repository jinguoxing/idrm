package resource_catalog

import "time"

// Category 类别实体（sqlx和gorm共用同一个结构）
type Category struct {
	Id          int64     `db:"id" gorm:"column:id;primaryKey"`
	Name        string    `db:"name" gorm:"column:name;type:varchar(100);not null"`
	Code        string    `db:"code" gorm:"column:code;type:varchar(50);uniqueIndex;not null"`
	ParentId    int64     `db:"parent_id" gorm:"column:parent_id;index;default:0"`
	Level       int       `db:"level" gorm:"column:level;default:1"`
	Sort        int       `db:"sort" gorm:"column:sort;default:0"`
	Description string    `db:"description" gorm:"column:description;type:text"`
	Status      int       `db:"status" gorm:"column:status;index;default:1"`
	CreatedAt   time.Time `db:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `db:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

// TableName gorm表名
func (Category) TableName() string {
	return "categories"
}
