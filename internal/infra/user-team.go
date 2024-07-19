package infra

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
)

type UserTeamRepository struct {
	SqlHandler
}

func NewUserTeamRepository(sqlHandler SqlHandler) repository.UserTeamRepository {
	userTeamRepository := UserTeamRepository{SqlHandler: sqlHandler}
	return &userTeamRepository
}

func (userTeamRepo *UserTeamRepository) Save(userId, teamId string) error {
	userTeam := &model.UserTeam{UserId: userId, TeamId: teamId}
	result := userTeamRepo.SqlHandler.Conn.Create(userTeam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) Delete(userId, teamId string) error {
	userTeam := &model.UserTeam{UserId: userId, TeamId: teamId}
	result := userTeamRepo.SqlHandler.Conn.Delete(userTeam)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (userTeamRepo *UserTeamRepository) FindUsersByTeamId(teamId string) ([]string, error) {

	var userTeams []*model.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("team_id = ?", teamId).Find(&userTeams)
	if result.Error != nil {
		return nil, result.Error
	}

	var userIds []string
	for _, userTeam := range userTeams {
		userIds = append(userIds, userTeam.UserId)
	}

	return userIds, nil
}

func (userTeamRepo *UserTeamRepository) FindTeamsByUserId(userId string) ([]string, error) {

	var userTeams []*model.UserTeam
	result := userTeamRepo.SqlHandler.Conn.Where("user_id = ?", userId).Find(&userTeams)
	if result.Error != nil {
		return nil, result.Error
	}

	var teamIds []string
	for _, userTeam := range userTeams {
		teamIds = append(teamIds, userTeam.TeamId)
	}

	return teamIds, nil
}
