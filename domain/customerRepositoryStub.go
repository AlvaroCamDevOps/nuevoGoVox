package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Alvaro", "Limah", "123", "01-01-2023", "1"},
		{"1002", "Alva", "Limah", "123", "02-01-2023", "1"},
	}
	return CustomerRepositoryStub{customers}
}
