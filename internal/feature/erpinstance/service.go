package erpinstance

import (
	"context"

	"github.com/ErickHerreraISW/go_erp/internal/utils"
)

type Service interface {
	Create() (*ErpInstance, error)
	FindByID(id uint) (*ErpInstance, error)
	FindByKey(key string) (*ErpInstance, error)
}

type svc struct{ repo Repository }

func NewService(r Repository) Service { return &svc{repo: r} }

func (s *svc) Create() (*ErpInstance, error) {

	erpKey, _ := utils.RandomKey(50)

	e := &ErpInstance{ErpKey: erpKey}
	return e, s.repo.Create(e)
}

func (s *svc) FindByID(id uint) (*ErpInstance, error) { return s.repo.FindByID(id) }

func (s *svc) FindByKey(key string) (*ErpInstance, error) {
	ctx := context.Background()
	return s.repo.FindByKey(ctx, key)
}
