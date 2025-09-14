package erpinstanceuser

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	Create(p *ErpInstanceUser) error
	FindByID(id uint) (*ErpInstanceUser, error)
	FindByErpInstanceId(ctx context.Context, erpInstanceId uint) (*[]ErpInstanceUser, error)
	FindByUserId(ctx context.Context, userId uint) (*ErpInstanceUser, error)
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db: db} }

func (r *repo) WithTx(tx *gorm.DB) Repository { return &repo{db: tx} }

func (r *repo) Create(e *ErpInstanceUser) error { return r.db.Create(e).Error }

func (r *repo) FindByID(id uint) (*ErpInstanceUser, error) {
	var e ErpInstanceUser
	if err := r.db.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *repo) FindByErpInstanceId(ctx context.Context, erpInstanceId uint) (*[]ErpInstanceUser, error) {
	var e []ErpInstanceUser
	if err := r.db.WithContext(ctx).
		Where("ErpInstanceID = ?", erpInstanceId).
		Find(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // o tu error de negocio NotFound
		}
		return nil, err
	}
	return &e, nil
}

func (r *repo) FindByUserId(ctx context.Context, userId uint) (*ErpInstanceUser, error) {
	var e ErpInstanceUser
	if err := r.db.WithContext(ctx).
		Where("userId = ?", userId).
		Take(&e).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // o tu error de negocio NotFound
		}
		return nil, err
	}
	return &e, nil
}
