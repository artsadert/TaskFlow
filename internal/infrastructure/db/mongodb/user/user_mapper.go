package user

import "github.com/artsadert/TaskFlow/internal/domain/entities"

func toDBUser(user *entities.User) *DBUser {
	return &DBUser{
		Uuid:      user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Create_at: user.Create_at,
		Update_at: user.Update_at,
	}
}

func fromDBUser(user *DBUser) *entities.User {
	return &entities.User{
		Id:        user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		Create_at: user.Create_at,
		Update_at: user.Update_at,
	}
}
