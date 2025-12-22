package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	// ORM选择：sqlx 或 gorm
	ORM string `json:",default=sqlx"`

	// 多数据库配置
	DataSources struct {
		ResourceCatalog struct {
			Source string
		}
		DataView struct {
			Source string
		}
		DataUnderstanding struct {
			Source string
		}
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
