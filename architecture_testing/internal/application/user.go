package application

// Example comment.
type User struct {
	uuid    string
	name    string
	company string
}

func (u User) UUID() string {
	return u.uuid
}

func (u User) Name() string {
	return u.name
}

func (u User) IsEmpty() bool {
	return u.uuid == ""
}

func NewUser(UUID string, name string, company string) User {
	u := User{
		uuid:    UUID,
		name:    name,
		company: company,
	}
	return u
}
