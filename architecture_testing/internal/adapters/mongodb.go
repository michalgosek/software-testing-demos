package adapters

import (
	"context"

	"architecture_testing/internal/application"
)

type MongoDB struct {
	lookup map[string]UserDocument
}

func (m *MongoDB) QueryUserWithUUID(ctx context.Context, UUID string) (application.User, error) {
	doc, ok := m.lookup[UUID]
	if !ok {
		return application.User{}, nil
	}
	appUser := doc.ParseToApplicationUser()
	return appUser, nil
}

type UserDocument struct {
	UUID    string `bson:"uuid"`
	Name    string `bson:"name"`
	Company string `bson:"company"`
}

func (u UserDocument) ParseToApplicationUser() application.User {
	ap := application.NewUser(u.UUID, u.Name, u.Company)
	return ap
}

type Interfacer interface {
	Akki() string
}
