package application

type UserRepository interface {
	QueryUserWithUUID(UUID string) (User, error)
}

type Service struct {
	repository UserRepository
}

func (s *Service) GetUserWithUUID(UUID string) (User, error) {
	user, err := s.repository.QueryUserWithUUID(UUID)
	if err != nil {
		return User{}, nil
	}
	return user, nil
}
