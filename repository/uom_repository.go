package repository

import (
	"database/sql"
	"go-api-enigma/model"
	"go-api-enigma/model/dto"
	"math"
)

type UomRepository interface {
	Save(uom model.Uom) error
	FindById(id string) (model.Uom, error)
	FindAll() ([]model.Uom, error)
	Update(uom model.Uom) error
	DeleteById(id string) error
	Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error)
}

type uomRepository struct {
	db *sql.DB
}

// Paging implements UomRepository.
func (u *uomRepository) Paging(payload dto.PageRequest) ([]model.Uom, dto.Paging, error) {
	if payload.Page <= 0 {
		payload.Page = 1
	}
	query := "SELECT id, type FROM m_uom LIMIT $2 OFFSET $1"
	rows, err := u.db.Query(query, (payload.Page-1)*payload.Size, payload.Size)
	if err != nil {
		return nil, dto.Paging{}, err
	}
	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Type)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		uoms = append(uoms, uom)
	}

	var count int
	row := u.db.QueryRow("SELECT COUNT(id) FROM m_uom")
	if err := row.Scan(&count); err != nil {
		return nil, dto.Paging{}, err
	}
	paging := dto.Paging{
		Page:       payload.Page,
		Size:       payload.Size,
		TotalRows:  count,
		TotalPages: int(math.Ceil(float64(count) / float64(payload.Size))),
	}
	return uoms, paging, nil
}

// DeleteById implements UomRepository.
func (u *uomRepository) DeleteById(id string) error {
	_, err := u.db.Exec("DELETE FROM m_uom WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements UomRepository.
func (u *uomRepository) FindAll() ([]model.Uom, error) {
	rows, err := u.db.Query("SELECT * FROM m_uom")
	if err != nil {
		return nil, err
	}
	var uoms []model.Uom
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Type)
		if err != nil {
			return nil, err
		}
		uoms = append(uoms, uom)
	}
	return uoms, nil
}

// FindById implements UomRepository.
func (u *uomRepository) FindById(id string) (model.Uom, error) {
	row := u.db.QueryRow("SELECT id, type FROM m_uom WHERE id = $1", id)
	var uom model.Uom
	err := row.Scan(&uom.Id, &uom.Type)
	if err != nil {
		return model.Uom{}, err
	}
	return uom, nil
}

// Save implements UomRepository.
func (u *uomRepository) Save(uom model.Uom) error {
	_, err := u.db.Exec("INSERT INTO m_uom VALUES($1, $2)", uom.Id, uom.Type)
	if err != nil {
		return err
	}
	return nil
}

// Update implements UomRepository.
func (u *uomRepository) Update(uom model.Uom) error {
	_, err := u.db.Exec("UPDATE m_uom SET type = $2 WHERE id=$1", uom.Id, uom.Type)
	if err != nil {
		return err
	}
	return nil
}

func NewUomRepository(db *sql.DB) UomRepository {
	return &uomRepository{db: db}
}
