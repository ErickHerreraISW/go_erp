package erpinstance

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	Create(p *ErpInstance) error
	FindByID(id uint) (*ErpInstance, error)
	FindByKey(ctx context.Context, key string) (*ErpInstance, error)
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db: db} }

func (r *repo) WithTx(tx *gorm.DB) Repository { return &repo{db: tx} }

func (r *repo) Create(e *ErpInstance) error { return r.db.Create(e).Error }

func (r *repo) FindByID(id uint) (*ErpInstance, error) {
	var e ErpInstance
	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *repo) FindByKey(ctx context.Context, key string) (*ErpInstance, error) {
	var e ErpInstance
	if err := r.db.WithContext(ctx).
		Where("erpkey = ?", key).
		Take(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // o tu error de negocio NotFound
		}
		return nil, err
	}
	return &e, nil
}
