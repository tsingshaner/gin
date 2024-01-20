package domain

func (*Domain) EqualPassword(password string, hashedPassword string) bool {
	return password == hashedPassword
}
