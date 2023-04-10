package models

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

// Insert 插入数据
// @param m 插入的表数据
func Insert[T any](ctx context.Context, db *gorm.DB, m *T) (int64, error) {
	cRe := db.WithContext(ctx).Omit(clause.Associations).Create(&m)
	if cRe.Error != nil {
		return 0, cRe.Error
	}
	return cRe.RowsAffected, nil
}

// Update 根据条件更新数据
// @param whereMap 查询条件
// @param updateMap 更新数据
func Update[T any](ctx context.Context, db *gorm.DB, whereMap, updateMap map[string]interface{}) (int64, error) {
	if len(whereMap) == 0 {
		return 0, errors.New("更新必须带条件！")
	}
	m := new(T)
	query := db.WithContext(ctx).Model(m).Debug()
	if whereMap != nil {
		for k, v := range whereMap {
			if strings.Count(k, "?") > 1 {
				nv := v.([]interface{})
				query = query.Where(k, nv...)
			} else {
				query = query.Where(k, v)
			}
		}
	}
	re := query.Updates(&updateMap)
	if re.Error != nil {
		return 0, re.Error
	}
	return re.RowsAffected, nil
}

// GetDetail 根据条件获取单个详情
// @param whereMap 查询条件
// @param field 要查询字段
func GetDetail[T any](ctx context.Context, db *gorm.DB, whereMap map[string]interface{}, field []string) (*T, error) {
	var re *T
	query := db.WithContext(ctx).Select(field).Debug()
	if whereMap != nil {
		for k, v := range whereMap {
			if strings.Count(k, "?") > 1 {
				nv := v.([]interface{})
				query = query.Where(k, nv...)
			} else {
				query = query.Where(k, v)
			}
		}
	}
	err := query.Take(&re)
	if err.Error != nil {
		return nil, err.Error
	}
	return re, nil
}

// GetList 根据条件获取多个信息
// @param whereMap 查询条件
// @param field 要查询的字段
// @param offset 偏移位置,-1：不使用该条件
// @param limit  获取多少条
// @param order 排序
func GetList[T any](ctx context.Context, db *gorm.DB, whereMap map[string]interface{}, field []string, offset, limit int, order string) ([]*T, error) {
	var reSlice []*T
	query := db.WithContext(ctx).Select(field).Debug()
	for k, v := range whereMap {
		//xx like ? or xx like ?
		if strings.Count(k, "?") > 1 {
			nv := v.([]interface{})
			query = query.Where(k, nv...)
		} else {
			query = query.Where(k, v)
		}
	}
	if offset != -1 {
		query = query.Offset(offset)
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	if order != "" {
		query = query.Order(order)
	}
	err := query.Debug().Find(&reSlice).Error
	if err != nil {
		return nil, err
	}
	return reSlice, nil
}

// Count 根据条件统计数量
// @param whereMap 查询条件
func Count[T any](ctx context.Context, db *gorm.DB, whereMap map[string]interface{}) (int64, error) {
	m := new(T)
	query := db.WithContext(ctx).Model(m).Debug()
	if whereMap != nil {
		for k, v := range whereMap {
			if strings.Count(k, "?") > 1 {
				nv := v.([]interface{})
				query = query.Where(k, nv...)
			} else {
				query = query.Where(k, v)
			}
		}
	}
	var num int64
	err := query.Count(&num).Error
	if err != nil {
		return 0, err
	}
	return num, nil
}

// Delete 根据条件删除数据
// @param whereMap 查询条件
// @param limit 要删除的数量：0 不限制
func Delete[T any](ctx context.Context, db *gorm.DB, whereMap map[string]interface{}, limit int) (int64, error) {
	if len(whereMap) == 0 {
		return 0, errors.New("删除必须带条件！")
	}
	m := new(T)
	query := db.WithContext(ctx).Model(m).Debug()
	if whereMap != nil {
		for k, v := range whereMap {
			if strings.Count(k, "?") > 1 {
				nv := v.([]interface{})
				query = query.Where(k, nv...)
			} else {
				query = query.Where(k, v)
			}
		}
	}
	if limit != 0 {
		query = query.Limit(limit)
	}
	var num int64
	re := query.Delete(&num)
	if re.Error != nil {
		return 0, re.Error
	}
	return re.RowsAffected, nil
}
