package entity

type UserOption func(*User)

func NewUserEntity(opts ...UserOption) UserEntity {
	user := &User{}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func WithUserID(userID string) UserOption {
	return func(u *User) {
		u.UserID = userID
	}
}

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.Username = username
	}
}

func WithPassword(password string) UserOption {
	return func(u *User) {
		u.Password = password
	}
}

func WithRole(role uint8) UserOption {
	return func(u *User) {
		u.Role = role
	}
}
