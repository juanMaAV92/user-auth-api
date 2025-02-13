package database

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("recor not found")
)

type Database struct {
	DB *gorm.DB
}

func (r *Database) Create(ctx context.Context, dst interface{}) error {
	tx := r.DB.Create(dst)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Database) Update(ctx context.Context, dst1 interface{}, dst2 map[string]interface{}) error {
	tx := r.DB.Model(dst1).Updates(dst2)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Database) GetByCondition(ctx context.Context, dst interface{}, field, value string) error {
	results := r.DB.Find(dst, field+"= ?", value)

	if results.Error != nil {
		if errors.Is(results.Error, gorm.ErrRecordNotFound) {
			return ErrRecordNotFound
		}
		return results.Error
	}

	return nil
}

func (r *Database) FindByConditions(ctx context.Context, dst interface{}, conditions map[string]interface{}) error {

	query := r.DB.Model(dst)
	for field, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	result := query.Find(dst)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ErrRecordNotFound
		}
		return result.Error
	}

	return nil
}
