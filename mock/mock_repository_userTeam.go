// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/repository/userTeam.go

// Package mock is a generated GoMock package.
package mock

import (
	entity "CurlARC/internal/domain/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserTeamRepository is a mock of UserTeamRepository interface.
type MockUserTeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTeamRepositoryMockRecorder
}

// MockUserTeamRepositoryMockRecorder is the mock recorder for MockUserTeamRepository.
type MockUserTeamRepositoryMockRecorder struct {
	mock *MockUserTeamRepository
}

// NewMockUserTeamRepository creates a new mock instance.
func NewMockUserTeamRepository(ctrl *gomock.Controller) *MockUserTeamRepository {
	mock := &MockUserTeamRepository{ctrl: ctrl}
	mock.recorder = &MockUserTeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTeamRepository) EXPECT() *MockUserTeamRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserTeamRepository) Delete(userId, teamId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, teamId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserTeamRepositoryMockRecorder) Delete(userId, teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserTeamRepository)(nil).Delete), userId, teamId)
}

// FindInvitedTeamsByUserId mocks base method.
func (m *MockUserTeamRepository) FindInvitedTeamsByUserId(userId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInvitedTeamsByUserId", userId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInvitedTeamsByUserId indicates an expected call of FindInvitedTeamsByUserId.
func (mr *MockUserTeamRepositoryMockRecorder) FindInvitedTeamsByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInvitedTeamsByUserId", reflect.TypeOf((*MockUserTeamRepository)(nil).FindInvitedTeamsByUserId), userId)
}

// FindMembersByTeamId mocks base method.
func (m *MockUserTeamRepository) FindMembersByTeamId(teamId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMembersByTeamId", teamId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMembersByTeamId indicates an expected call of FindMembersByTeamId.
func (mr *MockUserTeamRepositoryMockRecorder) FindMembersByTeamId(teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMembersByTeamId", reflect.TypeOf((*MockUserTeamRepository)(nil).FindMembersByTeamId), teamId)
}

// FindTeamsByUserId mocks base method.
func (m *MockUserTeamRepository) FindTeamsByUserId(userId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTeamsByUserId", userId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindTeamsByUserId indicates an expected call of FindTeamsByUserId.
func (mr *MockUserTeamRepositoryMockRecorder) FindTeamsByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTeamsByUserId", reflect.TypeOf((*MockUserTeamRepository)(nil).FindTeamsByUserId), userId)
}

// FindUsersByTeamId mocks base method.
func (m *MockUserTeamRepository) FindUsersByTeamId(teamId string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUsersByTeamId", teamId)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUsersByTeamId indicates an expected call of FindUsersByTeamId.
func (mr *MockUserTeamRepositoryMockRecorder) FindUsersByTeamId(teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUsersByTeamId", reflect.TypeOf((*MockUserTeamRepository)(nil).FindUsersByTeamId), teamId)
}

// IsMember mocks base method.
func (m *MockUserTeamRepository) IsMember(userId, teamId string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMember", userId, teamId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMember indicates an expected call of IsMember.
func (mr *MockUserTeamRepositoryMockRecorder) IsMember(userId, teamId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMember", reflect.TypeOf((*MockUserTeamRepository)(nil).IsMember), userId, teamId)
}

// Save mocks base method.
func (m *MockUserTeamRepository) Save(userTeam *entity.UserTeam) (*entity.UserTeam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", userTeam)
	ret0, _ := ret[0].(*entity.UserTeam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockUserTeamRepositoryMockRecorder) Save(userTeam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserTeamRepository)(nil).Save), userTeam)
}

// UpdateState mocks base method.
func (m *MockUserTeamRepository) UpdateState(userTeam *entity.UserTeam) (*entity.UserTeam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateState", userTeam)
	ret0, _ := ret[0].(*entity.UserTeam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateState indicates an expected call of UpdateState.
func (mr *MockUserTeamRepositoryMockRecorder) UpdateState(userTeam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateState", reflect.TypeOf((*MockUserTeamRepository)(nil).UpdateState), userTeam)
}