

package entity

import (
	"errors"
	"training/package/log"
	"time"
)

type AbsenteeismType struct {
	ID          int32
	SISID       int32
	Name        string
	Description string
	Identifier  string
	CreatedBy   int32
	CreatedAt   time.Time
	UpdatedBy   int32
	UpdatedAt   time.Time
	DeletedBy   int32
	DeletedAt   time.Time
}

func NewAbsenteeismType(sisID int32, name, description, identifier string, createdBy int32) (*AbsenteeismType, error) {
	absenteeismType := &AbsenteeismType{
		SISID:       sisID,
		Name:        name,
		Description: description,
		Identifier:  identifier,
		CreatedBy:   createdBy,
	}
	err := absenteeismType.ValidateNewAbsenteeismType()
	if err != nil {
		log.Errorf("error validating new AbsenteeismType entity %v", err)
		return &AbsenteeismType{}, err
	}

	return absenteeismType, err

}

func (r *AbsenteeismType) ValidateNewAbsenteeismType() error {
	if r.Name == "" {
		return errors.New("error validating AbsenteeismType entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating AbsenteeismType entity, created_by field required")
	}
	return nil
}

func (r *AbsenteeismType) ValidateUpdateAbsenteeismType() error {
	if r.ID <= 0 {
		return errors.New("error validating AbsenteeismType entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating AbsenteeismType entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating AbsenteeismType entity, updated_by field required")
	}
	return nil
}
