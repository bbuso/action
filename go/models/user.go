package models

type User struct {
	id       int
	address  Address
	phone    string
	email    string
	district District
}

func (u *User) Register() {
	address := Address{}
	err := u.validateEmail("brbuso@gmail.com")
	if err != nil {
		return
	}
	err = u.getDistrict(address)
	if err != nil {
		return
	}
	err = u.setAddress(address)
	if err != nil {
		return
	}
}

func (u User) Post() {
	msg := Message{}
	u.district.Post(u, msg)

}

func (u *User) validateEmail(email string) error {
	u.email = email
	return nil
}

func (u *User) getDistrict(a Address) error {
	u.district = District{}
	return nil
}

func (u *User) setAddress(a Address) error {
	if err := a.ValidateAddress(); err != nil {
		u.address = a
		return nil
	} else {
		return err
	}
}
