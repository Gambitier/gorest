package database

type DatabaseRepository struct {
	UserRepo *UserRepoHandler
}

func NewDatabaseRepository() (*DatabaseRepository, error) {
	return &DatabaseRepository{
		UserRepo: NewUserRepoHandler(),
	}, nil
}
