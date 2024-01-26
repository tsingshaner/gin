package entity

func (u *User) GetUserID() string {
	return u.UserID
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetRole() uint8 {
	return u.Role
}
