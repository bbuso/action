package models

type Cong struct{}

func (c Cong) Post(u User, msg Message) error {
	return nil
}
