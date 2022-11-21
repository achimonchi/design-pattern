package builder

type Customer struct {
	Name    string
	Email   string
	Address string
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (c *Customer) SetName(name string) *Customer {
	c.Name = name
	return c
}
