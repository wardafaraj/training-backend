
package absenteeism_type

import (
	"training/backend/services/entity"
	"time"
)

type Reader interface {
	Get(id int32) (*entity.AbsenteeismType, error)
	List() ([]*entity.AbsenteeismType, error)
	Check(id int32) (bool, error)
	CheckIfExist(name string) (bool,int32,time.Time, error)
	GetLastIdentifier() (int32, error)
}

type Writer interface {
	Create(e *entity.AbsenteeismType) (int32, error)
	Update(e *entity.AbsenteeismType) (int32, error)
	Activate(e *entity.AbsenteeismType) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateAbsenteeismType(sisID int32, name string, description, identifier string,  createdBy int32) (int32, error)
	CheckAbsenteeismType(id int32) (bool, error)
	ListAbsenteeismType() ([]*entity.AbsenteeismType, error)
	GetAbsenteeismType(id int32) (*entity.AbsenteeismType, error)
	UpdateAbsenteeismType(e *entity.AbsenteeismType) (int32, error)
	SoftDeleteAbsenteeismType(id, deletedBy int32) error
	HardDeleteAbsenteeismType(id int32) error
}
