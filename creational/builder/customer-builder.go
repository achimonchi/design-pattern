package builder

type CustomerBuilder struct {
	Name    string
	Email   string
	Address string
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{}
}

func (c *CustomerBuilder) SetName(name string) *CustomerBuilder {
	c.Name = name
	return c
}
func (c *CustomerBuilder) SetEmail(email string) *CustomerBuilder {
	c.Email = email
	return c
}
func (c *CustomerBuilder) SetAddress(address string) *CustomerBuilder {
	c.Address = address
	return c
}

func (c *CustomerBuilder) Build() *Customer {
	return &Customer{
		Name:    c.Name,
		Email:   c.Email,
		Address: c.Address,
	}
}
