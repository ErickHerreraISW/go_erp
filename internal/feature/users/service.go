package users

import (
	"errors"
	"time"

	"github.com/ErickHerreraISW/go_erp/internal/feature/erpinstance"
	"github.com/ErickHerreraISW/go_erp/internal/pkg/hash"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Service interface {
	Create(dto CreateUserDTO) (*User, error)
	List() ([]User, error)
	Get(id uint) (*User, error)
	Update(id uint, dto UpdateUserDTO) (*User, error)
	Delete(id uint) error
	Login(dto LoginDTO, token *jwtauth.JWTAuth) (string, error)
}

type svc struct {
	db      *gorm.DB
	repo    Repository
	erpRepo erpinstance.Repository
}

func NewService(r Repository, er erpinstance.Repository) Service { return &svc{repo: r, erpRepo: er} }

func (s *svc) Create(dto CreateUserDTO) (*User, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	ue, _ := s.repo.FindByEmail(dto.Email)

	if ue != nil {
		return nil, errors.New("user with this email already exists")
	}

	h, err := hash.HashPassword(dto.Password)

	if err != nil {
		return nil, err
	}
	u := &User{Name: dto.Name, Email: dto.Email, Password: h, Role: coalesce(dto.Role, "user")}
	return u, s.repo.Create(u)
}

func (s *svc) List() ([]User, error)      { return s.repo.FindAll() }
func (s *svc) Get(id uint) (*User, error) { return s.repo.FindByID(id) }

func (s *svc) Update(id uint, dto UpdateUserDTO) (*User, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if dto.Name != "" {
		u.Name = dto.Name
	}
	if dto.Role != "" {
		u.Role = dto.Role
	}
	return u, s.repo.Update(u)
}

func (s *svc) Delete(id uint) error { return s.repo.Delete(id) }

func (s *svc) Login(dto LoginDTO, token *jwtauth.JWTAuth) (string, error) {
	if err := dto.Validate(); err != nil {
		return "", err
	}
	u, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !hash.CheckPassword(u.Password, dto.Password) {
		return "", errors.New("invalid credentials")
	}
	_, tok, err := token.Encode(map[string]any{"sub": u.ID, "role": u.Role, "exp": time.Now().Add(24 * time.Hour).Unix()})
	if err != nil {
		log.Error().Err(err).Msg("jwt encode failed")
		return "", errors.New("could not create token")
	}
	return tok, nil
}

func coalesce[T comparable](v T, def T) T {
	var zero T
	if v == zero {
		return def
	}
	return v
}
