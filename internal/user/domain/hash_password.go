package domain

func (*Domain) HashPassword(password string) (string, error) {
	return password, nil
}
