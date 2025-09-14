package users

import "gorm.io/gorm"

type Repository interface {
	WithTx(tx *gorm.DB) Repository
	Create(u *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(u *User) error
	Delete(id uint) error
}

type repo struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repo{db: db} }

func (r *repo) WithTx(tx *gorm.DB) Repository { return &repo{db: tx} }

func (r *repo) Create(u *User) error { return r.db.Create(u).Error }

func (r *repo) FindAll() ([]User, error) {
	var list []User
	return list, r.db.Find(&list).Error
}

func (r *repo) FindByID(id uint) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) FindByEmail(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) Update(u *User) error { return r.db.Save(u).Error }

func (r *repo) Delete(id uint) error { return r.db.Delete(&User{}, id).Error }
