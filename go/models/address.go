package models

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip5     string
}

func (a Address) ValidateAddress() error {
	return nil
}

func (a Address) GetCongs() {}
