package value

import "fmt"

type Address struct {
	Street     string
	City       string
	Prefecture string
	PostalCode string
}

func NewAddress(street string, city string, prefecture string, postalCode string) (*Address, error) {
	if street == "" || city == "" || prefecture == "" || postalCode == "" {
		return nil, fmt.Errorf("all address fields are required")
	}
	return &Address{
		Street:     street,
		City:       city,
		Prefecture: prefecture,
		PostalCode: postalCode,
	}, nil
}

func (a *Address) String() string {
	return fmt.Sprintf("%s, %s, %s %s", a.Street, a.City, a.Prefecture, a.PostalCode)
}

func (a *Address) Equals(other *Address) bool {
	return a.Street == other.Street &&
		a.City == other.City &&
		a.Prefecture == other.Prefecture &&
		a.PostalCode == other.PostalCode
}
