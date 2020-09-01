package queries

type GetTakerHello struct {
	name string
}

func NewGetTakerHello(name string) GetTakerHello {
	return GetTakerHello{
		name: name,
	}
}

func (c GetTakerHello) Name() string {
	return c.name
}

func (GetTakerHello) QueryName() string {
	return "GetTakerHello"
}
