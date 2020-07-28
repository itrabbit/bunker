package db

import (
	"github.com/itrabbit/bunker/models"
)

type nameSpaceMapper struct {
}

func (m *nameSpaceMapper) FindOne(alias string) (*models.NameSpace, error) {
	if db == nil {
		return nil, ErrInvalidDbConnection
	}
	var ns models.NameSpace
	if err := db.
		Preload("Application").
		Where("alias = ?", alias).
		First(&ns).
		Error; err != nil {
		return nil, err
	}
	if ns.ID < 1 {
		return nil, ErrNotFound
	}
	return &ns, nil
}

func (m *nameSpaceMapper) Delete(ns *models.NameSpace) error {
	if ns == nil {
		return nil
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	return db.Delete(ns).Error
}

func (m *nameSpaceMapper) Exist(ns *models.NameSpace) (bool, error) {
	if ns == nil {
		return false, nil
	}
	if db == nil {
		return false, ErrInvalidDbConnection
	}
	count := 0
	if err := db.Model(&models.NameSpace{}).Where(
		"id != ? AND alias = ?",
		ns.ID,
		ns.Alias,
	).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *nameSpaceMapper) Save(ns *models.NameSpace) error {
	if ns == nil {
		return ErrInvalidData
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	if err := validate.Struct(ns); err != nil {
		return err
	}
	if exist, err := m.Exist(ns); err != nil {
		return err
	} else if exist {
		return ErrAlreadyExist
	}
	if !db.NewRecord(ns) {
		return db.Save(ns).Error
	}
	return db.Create(ns).Error
}

var NameSpaceMapper = new(nameSpaceMapper)
