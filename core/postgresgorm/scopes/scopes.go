package scopes

import (
	"context"
	"fmt"
	typeMapper "github.com/duongbui2002/core-package/core/reflection/typemapper"
	"github.com/duongbui2002/core-package/core/utils"
	"gorm.io/gorm"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func FilterAllItemsWithSoftDeleted(db *gorm.DB) *gorm.DB {
	// https://gorm.io/docs/delete.html#Find-soft-deleted-records
	return db.Unscoped()
}
func SoftDeleted(db *gorm.DB) *gorm.DB {
	return db.Unscoped().Where("deleted_at IS NOT NULL")
}
func FilterByID(id uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func FilterPaginate[TDataModel any](
	ctx context.Context,
	listQuery *utils.ListQuery,
) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var totalRows int64

		dataModel := typeMapper.GenericInstanceByT[TDataModel]()
		// https://gorm.io/docs/advanced_query.html
		db.WithContext(ctx).Model(dataModel).Count(&totalRows)

		// generate where query
		query := db.WithContext(ctx).
			Model(dataModel).
			Offset(listQuery.GetOffset()).
			Limit(listQuery.GetLimit()).
			Order(listQuery.GetOrderBy())

		if listQuery.Filters != nil {
			for _, filter := range listQuery.Filters {
				column := filter.Field
				action := filter.Comparison
				value := filter.Value

				switch action {
				case "equals":
					whereQuery := fmt.Sprintf("%s = ?", column)
					query = query.Where(whereQuery, value)
				case "contains":
					whereQuery := fmt.Sprintf("%s LIKE ?", column)
					query = query.Where(whereQuery, "%"+value+"%")
				case "in":
					whereQuery := fmt.Sprintf("%s IN (?)", column)
					queryArray := strings.Split(value, ",")
					query = query.Where(whereQuery, queryArray)
				}
			}
		}

		return query
	}
}
