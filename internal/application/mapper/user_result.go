package mapper

import (
	"github.com/artsadert/TaskFlow/internal/application/common"
	"github.com/artsadert/TaskFlow/internal/domain/entities"
)

func NewUserResultFromEntity(entity *entities.User) *common.UserResult {
	return &common.UserResult{Id: entity.Id,
		Name:      entity.Name,
		Email:     entity.Email,
		Create_at: entity.Create_at,
		Update_at: entity.Update_at,
	}
}

func NewUsersResultFromEntities(user_entities []*entities.User) []*common.UserResult {
	res := []*common.UserResult{}

	for _, entity := range user_entities {
		res = append(res, NewUserResultFromEntity(entity))
	}

	return res
}
