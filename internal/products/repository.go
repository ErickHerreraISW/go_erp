package products

import "gorm.io/gorm"

type Repository interface {
	Create(p *Product) error
	FindAll() ([]Product, error)
	FindByID(id uint) (*Product, error)
	Update(p *Product) error
	Delete(id uint) error
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db: db} }

func (r *repo) Create(p *Product) error { return r.db.Create(p).Error }
func (r *repo) FindAll() ([]Product, error) {
	var list []Product
	return list, r.db.Find(&list).Error
}
func (r *repo) FindByID(id uint) (*Product, error) {
	var p Product
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *repo) Update(p *Product) error { return r.db.Save(p).Error }
func (r *repo) Delete(id uint) error    { return r.db.Delete(&Product{}, id).Error }
