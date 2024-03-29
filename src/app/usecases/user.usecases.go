package usecases

import (
	"fmt"

	"github.com/go-api/src/app/dto"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

type UserUsecases struct {
	UserRepository repository.UserRepository
}

// CREATE USER
func (uu *UserUsecases) CreateUser(rawUser *entities.UserEntity) (interface{}, error) {
	// BUILD DE USER ID
	newUser := entities.NewUser(rawUser)

	// CREATE THE USER
	response, err := uu.UserRepository.Create(newUser)

	// RETURN ERROR IF STATUS IS FALSE
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// MAP ENTITY TO DTO
	if userEntity, ok := response.(*entities.UserEntity); ok {
		return dto.NewUser(userEntity), nil
	}

	return response, nil

}

// GET ALL USERS
func (uu *UserUsecases) FindAllUsers() ([]dto.UsersDTO, error) {
	// GET USERS FROM DB
	usersEntities, err := uu.UserRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	// MAP ENTITY TO DTO
	return dto.NewUsersDtos(usersEntities), nil
}
