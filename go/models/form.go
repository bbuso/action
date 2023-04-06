package models

type Form struct {
	Prefix    []string
	FirstName string
	LastName  string
	Address   string
	City      string
	State     []string
	Zip       []string
	Topic     []string
	Response  bool
}
