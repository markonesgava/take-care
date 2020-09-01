package domain

type Repositorier interface {
	GetOne(id string) (*Entitier, error)
	GetAll() ([]*Entitier, error)
	Add(entity *Entitier) error
	Update(id string, entity *Entitier) error
}
