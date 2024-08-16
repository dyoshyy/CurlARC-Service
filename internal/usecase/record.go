package usecase

import (
	"CurlARC/internal/domain/model"
	"CurlARC/internal/domain/repository"
	"errors"
	"time"

	"gorm.io/datatypes"
)

type RecordUsecase interface {
	CreateRecord(userId, teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error)
	AppendEndData(recordId, userId string, endsData datatypes.JSON) (*model.Record, error)
	GetRecordsByTeamId(teamId string) (*[]model.Record, error)
	UpdateRecord(recordId, userId string, updates model.RecordUpdate) (*model.Record, error)
	DeleteRecord(id string) error

	SetVisibility(recordId, userId string, isPublic bool) (*model.Record, error)
}

type recordUsecase struct {
	recordRepo   repository.RecordRepository
	userTeamRepo repository.UserTeamRepository
	teamRepo     repository.TeamRepository
}

func NewRecordUsecase(recordRepo repository.RecordRepository, userTeamRepo repository.UserTeamRepository, teamRepo repository.TeamRepository) RecordUsecase {
	return &recordUsecase{recordRepo: recordRepo, userTeamRepo: userTeamRepo, teamRepo: teamRepo}
}

func (u *recordUsecase) CreateRecord(userId, teamId, enemyTeamName, place string, result model.Result, date time.Time) (*model.Record, error) {

	// check if the user is a member of the team
	if _, err := u.userTeamRepo.IsMember(userId, teamId); err != nil {
		return nil, err
	}

	// check if the team exists
	if _, err := u.teamRepo.FindById(teamId); err != nil {
		return nil, err
	}

	return u.recordRepo.Create(teamId, enemyTeamName, place, result, date)
}

func (u *recordUsecase) AppendEndData(recordId, userId string, endsData datatypes.JSON) (*model.Record, error) {
	// Get the record by ID
	record, err := u.recordRepo.FindById(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("appender is not a member of the team")
	}

	// Initialize the existing endsData
	var existingEndsData datatypes.JSON
	if record.EndsData != nil {
		existingEndsData = record.EndsData
	}

	// Merge or append the new data
	// Assuming endsData is a JSON array; otherwise, adjust the merging logic
	updatedEndsData := append(existingEndsData, endsData...)

	// Prepare the update struct
	updateFields := model.RecordUpdate{
		EndsData: &updatedEndsData,
	}

	// Update the record with the new endsData
	updatedRecord, err := u.recordRepo.Update(recordId, updateFields)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) GetRecordsByTeamId(teamId string) (*[]model.Record, error) {
	return u.recordRepo.FindByTeamId(teamId)
}

func (u *recordUsecase) UpdateRecord(recordId, userId string, updates model.RecordUpdate) (*model.Record, error) {
	// Get the record by ID
	record, err := u.recordRepo.FindById(recordId)
	if err != nil {
		return nil, err
	}

	// Check if the user is a member of the team
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("updater is not a member of the team")
	}

	// Update the record with only the fields provided in the updates
	updatedRecord, err := u.recordRepo.Update(recordId, updates)
	if err != nil {
		return nil, err
	}

	return updatedRecord, nil
}

func (u *recordUsecase) DeleteRecord(id string) error {
	return u.recordRepo.Delete(id)
}

func (u *recordUsecase) SetVisibility(recordId, userId string, isPublic bool) (*model.Record, error) {

	// check if the record exists
	record, err := u.recordRepo.FindById(recordId)
	if err != nil {
		return nil, err
	}

	// check if the user is the member of the record
	isMember, err := u.userTeamRepo.IsMember(userId, record.TeamId)
	if err != nil {
		return nil, err
	}
	if !isMember {
		return nil, errors.New("inviter is not a member of the team")
	}

	// Prepare the update struct
	updateFields := model.RecordUpdate{
		IsPublic: &isPublic,
	}

	return u.recordRepo.Update(recordId, updateFields)
}
