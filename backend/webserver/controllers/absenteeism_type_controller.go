

package controllers

import (
	"net/http"
	"training/backend/services/entity"
	"training/backend/services/usecase/absenteeism_type"
	"training/package/log"
	"training/package/models"
	"training/package/trim"
	"training/package/util"
	"training/package/wrappers"

	"github.com/labstack/echo/v4"
)

func ListAbsenteeismType(c echo.Context) error {
	service := absenteeism_type.NewService()
	absenteeismTypes, err := service.ListAbsenteeismType()

	if util.IsError(err) {
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if absenteeismTypes == nil {
		return wrappers.MessageResponse(c, http.StatusOK, "no absenteeism types  data found")
	}

	absenteeismTypeResponse := make([]*models.AbsenteeismType, 0)
	for _, absenteeismType := range absenteeismTypes {
		absenteeismTypeResponse = append(absenteeismTypeResponse, &models.AbsenteeismType{
			ID:          absenteeismType.ID,
			Name:        absenteeismType.Name,
			Identifier:  absenteeismType.Identifier,
			Description: absenteeismType.Description,
			CreatedAt:   absenteeismType.CreatedAt,
			CreatedBy:   absenteeismType.CreatedBy,
			UpdatedAt:   absenteeismType.UpdatedAt,
			UpdatedBy:   absenteeismType.UpdatedBy,
		})
	}

	return wrappers.Response(c, http.StatusOK, absenteeismTypeResponse)
}

func GetAbsenteeismType(c echo.Context) error {

	modelID := &models.ID{}
	if err := c.Bind(&modelID); util.IsError(err) {
		log.Errorf("error binding absenteeism type id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(modelID); util.IsError(err) {
		log.Errorf("error validating absenteeism type id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "error validating absenteeism type id")
	}
	service := absenteeism_type.NewService()
	absenteeismType, err := service.GetAbsenteeismType(modelID.ID)

	if err != nil {
		log.Errorf("error getting absenteeism type: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	absenteeismTypeResponse := models.AbsenteeismType{
		ID:          absenteeismType.ID,
		Name:        absenteeismType.Name,
		Identifier:  absenteeismType.Identifier,
		Description: absenteeismType.Description,
		CreatedAt:   absenteeismType.CreatedAt,
		CreatedBy:   absenteeismType.CreatedBy,
		UpdatedAt:   absenteeismType.UpdatedAt,
		UpdatedBy:   absenteeismType.UpdatedBy,
	}
	return wrappers.Response(c, http.StatusOK, absenteeismTypeResponse)
}

func CreateAbsenteeismType(c echo.Context) error {
	absenteeismTypeModel := &models.AbsenteeismType{}
	if err := c.Bind(absenteeismTypeModel); util.IsError(err) {
		log.Errorf("error binding absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(absenteeismTypeModel); util.IsError(err) {
		log.Errorf("error validating absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "error validating absenteeism type")
	}

	service := absenteeism_type.NewService()

	name := trim.FormatText(absenteeismTypeModel.Name)

	_, err := service.CreateAbsenteeismType(0, name, absenteeismTypeModel.Description, absenteeismTypeModel.Identifier, absenteeismTypeModel.CreatedBy)

	if util.IsError(err) {
		if err.Error() == "absenteeism type is not found" {
			return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type is not found")
		}
		if err.Error() == "absenteeism type already exist" {
			return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type already exists")
		}
		log.Errorf("error creating new absenteeism type: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type created successfully")
}

func UpdateAbsenteeismType(c echo.Context) error {
	absenteeismTypeModel := models.AbsenteeismType{}
	if err := c.Bind(&absenteeismTypeModel); util.IsError(err) {
		log.Errorf("error binding absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(absenteeismTypeModel); util.IsError(err) {
		log.Errorf("error validating absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "error validating absenteeism type")
	}

	name := trim.FormatText(absenteeismTypeModel.Name)

	service := absenteeism_type.NewService()
	data := &entity.AbsenteeismType{
		ID:          absenteeismTypeModel.ID,
		Name:        name,
		Description: absenteeismTypeModel.Description,
		UpdatedBy:   absenteeismTypeModel.UpdatedBy,
	}
	_, err := service.UpdateAbsenteeismType(data)
	if util.IsError(err) {
		if err.Error() == "absenteeism type is not found" {
			return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type is not found")
		}
		if err.Error() == "absenteeism type already exist" {
			return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type already exists")
		}
		log.Errorf("error creating new absenteeism type: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type updated successfully")
}

func SoftDeleteAbsenteeismType(c echo.Context) error {

	absenteeismTypeDeletedBy := &models.DeletedBy{}

	if err := c.Bind(&absenteeismTypeDeletedBy); util.IsError(err) {
		log.Errorf("error binding absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if err := c.Validate(absenteeismTypeDeletedBy); util.IsError(err) {
		log.Errorf("error validating absenteeism type fields: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "error validating absenteeism type")
	}

	service := absenteeism_type.NewService()
	err := service.SoftDeleteAbsenteeismType(absenteeismTypeDeletedBy.ID, absenteeismTypeDeletedBy.DeletedBy)
	if util.IsError(err) {
		log.Errorf("error deleting absenteeism type: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type deleted successfully")
}

func DeleteAbsenteeismType(c echo.Context) error {

	modelID := &models.ID{}
	if err := c.Bind(&modelID); util.IsError(err) {
		log.Errorf("error binding absenteeism type id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if err := c.Validate(modelID); util.IsError(err) {
		log.Errorf("error validating absenteeism type id: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "error validating absenteeism type id")
	}

	service := absenteeism_type.NewService()
	err := service.HardDeleteAbsenteeismType(modelID.ID)
	if util.IsError(err) {
		log.Errorf("error deleting absenteeism type: %v", err)
		return wrappers.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	return wrappers.MessageResponse(c, http.StatusOK, "absenteeism type deleted successfully")
}
