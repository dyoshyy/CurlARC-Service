package repository

import (
	"CurlARC/internal/domain/entity"
	"CurlARC/internal/handler/response"
)

type RecordRepository interface {
	Save(entity.Record) (*entity.Record, error)
	FindByRecordId(recordId string) (*entity.Record, error)
	FindIndicesByTeamId(teamId string) (*[]response.RecordIndex, error)
	FindByTeamId(teamId string) (*[]entity.Record, error)
	Update(record entity.Record) (*entity.Record, error)
	Delete(recordId string) error
}
