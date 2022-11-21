package builder

import (
	"fmt"
	"testing"
)

func TestCustomerBuilder(t *testing.T) {
	customer := NewCustomerBuilder().
		SetAddress("Jakarta").
		SetName("customer A").
		Build()

	fmt.Println(customer.Name)
}
