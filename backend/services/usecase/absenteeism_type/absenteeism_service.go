
package absenteeism_type

import (
	"errors"
	"training/backend/services/entity"
	"training/backend/services/error_message"
	"training/backend/services/repository"
	"training/package/log"
	"strings"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewAbsenteeismType()
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateAbsenteeismType(sisID int32, name string, description, identifier string, createdBy int32) (int32, error) {
	var absenteeismTypeID int32

	exists, id, deletedAt, err := s.repo.CheckIfExist(identifier)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		return absenteeismTypeID, err
	}

	absenteeismType, err := entity.NewAbsenteeismType(sisID, strings.ToUpper(name), description, identifier, createdBy)
	if err != nil {
		return absenteeismTypeID, err
	}

	if !exists {
		absenteeismTypeID, err := s.repo.Create(absenteeismType)
		if err != nil {
			return absenteeismTypeID, err
		}
		return absenteeismTypeID, nil
	} else {
		if !deletedAt.IsZero() {
			absenteeismType.ID = id
			_, err := s.ActivateAbsenteeismType(absenteeismType)
			if err != nil {
				return absenteeismTypeID, err
			}
			return absenteeismTypeID, nil
		} else {
			err = errors.New("absenteeism type already exist")
			return absenteeismTypeID, err
		}
	}
}

func (s *Service) CheckAbsenteeismType(id int32) (bool, error) {
	exist, err := s.repo.Check(id)
	if err != nil {
		return exist, err
	}
	return exist, err
}

func (s *Service) ListAbsenteeismType() ([]*entity.AbsenteeismType, error) {
	absenteeismType, err := s.repo.List()

	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return absenteeismType, nil
		}
		return absenteeismType, err
	}
	return absenteeismType, err
}

func (s *Service) GetAbsenteeismType(id int32) (*entity.AbsenteeismType, error) {
	absenteeismType, err := s.repo.Get(id)

	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return absenteeismType, nil
		}
		return absenteeismType, err
	}
	return absenteeismType, err
}

func (s *Service) ActivateAbsenteeismType(e *entity.AbsenteeismType) (int32, error) {

	_, err := s.repo.Activate(e)
	if err != nil {
		return e.ID, err
	}
	return e.ID, err
}

func (s *Service) UpdateAbsenteeismType(e *entity.AbsenteeismType) (int32, error) {
	err := e.ValidateUpdateAbsenteeismType()
	if err != nil {
		log.Error(err)
		return e.ID, err
	}
	exists, id, deletedAt, err := s.repo.CheckIfExist(e.Name)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		return e.ID, err
	}

	if exists && id == e.ID && deletedAt.IsZero() {
		_, err = s.repo.Update(e)
		if err != nil {
			return e.ID, err
		}
	} else if exists && id != e.ID && deletedAt.IsZero() {
		err = errors.New("absenteeism type already exist")
		return e.ID, err
	} else if exists && id != e.ID && !deletedAt.IsZero() {
		err = s.HardDeleteAbsenteeismType(id)
		if err != nil {
			return e.ID, err
		}
		_, err = s.repo.Update(e)
		if err != nil {
			return e.ID, err
		}
	} else {
		_, err = s.repo.Update(e)
		if err != nil {
			return e.ID, err
		}

	}
	return e.ID, err
}

func (s *Service) SoftDeleteAbsenteeismType(id, deletedBy int32) error {
	_, err := s.GetAbsenteeismType(id)
	if err != nil {
		return err
	}
	err = s.repo.SoftDelete(id, deletedBy)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) HardDeleteAbsenteeismType(id int32) error {
	_, err := s.GetAbsenteeismType(id)
	if err != nil {
		return err
	}
	err = s.repo.HardDelete(id)
	if err != nil {
		return err
	}
	return err
}
