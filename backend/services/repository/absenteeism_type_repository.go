

package repository

import (
	"context"
	"fmt"
	"os"
	"training/backend/services/database"
	"training/backend/services/entity"
	"training/package/log"
	"training/package/util"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AbsenteeismTypeConn struct {
	conn *pgxpool.Pool
}

func NewAbsenteeismType() *AbsenteeismTypeConn {
	conn, err := database.Connect()
	if util.IsError(err) {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return &AbsenteeismTypeConn{
		conn: conn,
	}
}

func getAbsenteeismTypeQuery() string {
	return "SELECT id, name,description,identifier, created_by, created_at, updated_by, updated_at, deleted_by, deleted_at FROM absenteeism_type"
}

func (con *AbsenteeismTypeConn) Create(e *entity.AbsenteeismType) (int32, error) {
	var absenteeismTypeID int32
	query := "INSERT INTO absenteeism_type (sis_id, name, identifier, description, created_by, created_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	err := con.conn.QueryRow(context.Background(), query, e.SISID, e.Name, e.Identifier, e.Description, e.CreatedBy, time.Now()).Scan(&absenteeismTypeID)
	if util.IsError(err) {
		log.Errorf("error creating absenteeism type: %v", err)
	}
	return absenteeismTypeID, err
}

func (con *AbsenteeismTypeConn) CheckIfExist(identifier string) (bool, int32, time.Time, error) {
	var exists bool
	var id pgtype.Int4
	var deletedAt pgtype.Timestamp

	query := "SELECT EXISTS(SELECT 1 FROM absenteeism_type WHERE identifier =$1) AS absenteeism_type_exist,id,deleted_at FROM absenteeism_type WHERE  identifier =$1"

	err := con.conn.QueryRow(context.Background(), query, identifier).Scan(&exists,&id, &deletedAt)
	if err != nil {
		log.Errorf("error checking assessment type existence %v", err)
	}

	return exists, id.Int, deletedAt.Time, err
}

func (con *AbsenteeismTypeConn) Check(id int32) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM absenteeism_type WHERE id  = $1)"
	err := con.conn.QueryRow(context.Background(), query, id).Scan(&exists)
	if util.IsError(err) {
		log.Errorf("error checking absenteeism type by id: %v", err)
	}
	return exists, err
}

func (con *AbsenteeismTypeConn) List() ([]*entity.AbsenteeismType, error) {
	var id int32
	var name, description, identifier pgtype.GenericText
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	var absenteeismTypes []*entity.AbsenteeismType
	query := getAbsenteeismTypeQuery() + ` WHERE deleted_at IS NULL`

	rows, err := con.conn.Query(context.Background(), query)
	if util.IsError(err) {
		log.Errorf("error querying absenteeism type %v", err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &description, &identifier, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt); util.IsError(err) {
			log.Errorf("error scanning absenteeism type %v", err)
			return nil, err
		}
		absenteeismType := &entity.AbsenteeismType{
			ID:          id,
			Name:        name.String,
			Description: description.String,
			Identifier:  identifier.String,
			CreatedBy:   createdBy.Int,
			CreatedAt:   createdAt.Time,
			UpdatedBy:   updatedBy.Int,
			UpdatedAt:   updatedAt.Time,
			DeletedBy:   deletedBy.Int,
			DeletedAt:   deletedAt.Time,
		}
		absenteeismTypes = append(absenteeismTypes, absenteeismType)
	}
	return absenteeismTypes, err
}

func (con *AbsenteeismTypeConn) Get(id int32) (*entity.AbsenteeismType, error) {
	var name, description, identifier pgtype.Text
	var createdBy, updatedBy, deletedBy pgtype.Int4
	var createdAt, updatedAt, deletedAt pgtype.Timestamp
	query := getAbsenteeismTypeQuery() + ` WHERE deleted_at IS NULL AND id = $1`
	err := con.conn.QueryRow(context.Background(), query, id).Scan(&id, &name, &description, &identifier, &createdBy, &createdAt, &updatedBy, &updatedAt, &deletedBy, &deletedAt)

	if util.IsError(err) {
		log.Errorf("error getting absenteeism type %v", err)
		return nil, err
	}
	absenteeismType := &entity.AbsenteeismType{
		ID:          id,
		Name:        name.String,
		Description: description.String,
		Identifier:  identifier.String,
		CreatedBy:   createdBy.Int,
		CreatedAt:   createdAt.Time,
		UpdatedBy:   updatedBy.Int,
		UpdatedAt:   updatedAt.Time,
		DeletedBy:   deletedBy.Int,
		DeletedAt:   deletedAt.Time,
	}

	return absenteeismType, err

}


func (con *AbsenteeismTypeConn) Activate(e *entity.AbsenteeismType) (int32, error) {
	query := `UPDATE absenteeism_type SET 
							  name = $1, 
							  description = $2, 
							  created_by = $3, 
							  created_at = $4, 
							  updated_by = $5, 
							  updated_at = $6,   
							  deleted_by = $7, 
							  deleted_at = $8 
							  WHERE id = $9`
	_, err := con.conn.Exec(context.Background(), query, e.Name,e.Description, e.CreatedBy, time.Now(),nil,nil,nil,nil, e.ID)
	if util.IsError(err) {
		log.Errorf("error activating absenteeism type: %v", err)
	}
	return e.ID, err
}

func (con *AbsenteeismTypeConn) Update(e *entity.AbsenteeismType) (int32, error) {
	query := `UPDATE absenteeism_type SET 
									  name = $1, 
									  description = $2, 
									  updated_by = $3, 
									  updated_at = $4 
									  WHERE id = $5`
	_, err := con.conn.Exec(context.Background(), query, e.Name, e.Description, e.UpdatedBy, time.Now(), e.ID)
	if util.IsError(err) {
		log.Errorf("error updating absenteeism type by id: %v", err)
	}
	return e.ID, err
}

func (con *AbsenteeismTypeConn) SoftDelete(id, deletedBy int32) error {
	query := `UPDATE absenteeism_type SET 
									  deleted_by = $1, 
									  deleted_at = $2 
									  WHERE id = $3`
	_, err := con.conn.Exec(context.Background(), query, deletedBy, time.Now(), id)
	if util.IsError(err) {
		log.Errorf("error soft delete absenteeism type by id: %v", err)
	}
	return err
}

func (con *AbsenteeismTypeConn) HardDelete(id int32) error {
	query := "DELETE FROM absenteeism_type WHERE id = $1"
	_, err := con.conn.Exec(context.Background(), query, id)
	if util.IsError(err) {
		log.Errorf("error hard delete absenteeism type by id: %v", err)
	}
	return err
}

func (con *AbsenteeismTypeConn) GetLastIdentifier() (int32, error) {

	var maxSn pgtype.Int4
	query := `SELECT MAX(CAST(NULLIF(SPLIT_PART(identifier, '-', 2), '') AS INTEGER)) AS last_identifier FROM absenteeism_type`
	err := con.conn.QueryRow(context.Background(), query).Scan(&maxSn)

	if util.IsError(err) {
		log.Errorf("error getting last absenteeism type identifier number %v", err)
		return 0, err
	}

	return int32(maxSn.Int), nil
}
