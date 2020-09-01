package commands

type CreateTaker struct {
	name string
}

func NewCreateTaker(name string) CreateTaker {
	return CreateTaker{
		name: name,
	}
}

func (c CreateTaker) Name() string {
	return c.name
}

func (CreateTaker) CommandName() string {
	return "CreateTaker"
}
