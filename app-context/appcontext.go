package appcontext

import (
	"github.com/gambitier/gorest/database"
	"github.com/gambitier/gorest/services"
)

type AppContext struct {
	Services     *services.Service
	DatabaseRepo *database.DatabaseRepository
}

func InitAppContext() (*AppContext, error) {
	db, err := database.NewDatabaseRepository()
	if err != nil {
		return nil, err
	}

	service, err := services.NewService(db)
	if err != nil {
		return nil, err
	}

	return &AppContext{
		Services:     service,
		DatabaseRepo: db,
	}, nil
}
