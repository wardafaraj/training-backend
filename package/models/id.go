package models

type ID struct {
	ID int32 `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
}

type DeletedBy struct {
	ID        int32 `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	DeletedBy int32 `json:"deleted_by"`
}
