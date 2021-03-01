package cache

import (
	"database/sql"
)

import (
	"github.com/duolacloud/seata-golang/pkg/client/at/sql/schema"
)

type ITableMetaCache interface {
	GetTableMeta(tx *sql.Tx, tableName, resourceId string) (schema.TableMeta, error)

	Refresh(tx *sql.Tx, resourceId string)
}

var tableMetaCache ITableMetaCache

func SetTableMetaCache(tbMetaCache ITableMetaCache) {
	tableMetaCache = tbMetaCache
}

func GetTableMetaCache() ITableMetaCache {
	return tableMetaCache
}
