package products

type Service interface {
	Create(dto CreateProductDTO) (*Product, error)
	List() ([]Product, error)
	Get(id uint) (*Product, error)
	Update(id uint, dto UpdateProductDTO) (*Product, error)
	Delete(id uint) error
}

type svc struct{ repo Repository }

func NewService(r Repository) Service { return &svc{repo: r} }

func (s *svc) Create(dto CreateProductDTO) (*Product, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	p := &Product{Name: dto.Name, SKU: dto.SKU, Price: dto.Price, Stock: dto.Stock, Description: dto.Description}
	return p, s.repo.Create(p)
}

func (s *svc) List() ([]Product, error)      { return s.repo.FindAll() }
func (s *svc) Get(id uint) (*Product, error) { return s.repo.FindByID(id) }

func (s *svc) Update(id uint, dto UpdateProductDTO) (*Product, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if dto.Name != "" {
		p.Name = dto.Name
	}
	if dto.SKU != "" {
		p.SKU = dto.SKU
	}
	if dto.Price != 0 {
		p.Price = dto.Price
	}
	if dto.Stock != 0 {
		p.Stock = dto.Stock
	}
	if dto.Description != "" {
		p.Description = dto.Description
	}
	return p, s.repo.Update(p)
}

func (s *svc) Delete(id uint) error { return s.repo.Delete(id) }
