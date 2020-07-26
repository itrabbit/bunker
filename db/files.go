package db

import (
	"github.com/itrabbit/bunker/models"
)

type filesMapper struct {
}

func (m *filesMapper) Delete(file *models.File) error {
	if file == nil {
		return nil
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	return db.Delete(file).Error
}

func (m *filesMapper) Exist(file *models.File) (bool, error) {
	if file == nil {
		return false, nil
	}
	if db == nil {
		return false, ErrInvalidDbConnection
	}
	count := 0
	if err := db.Model(&models.File{}).Where(
		"id != ? AND hash = ? AND mime_type = ? AND name_space_id = ?",
		file.ID,
		file.Hash,
		file.MimeType,
		file.NameSpaceID,
	).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *filesMapper) Save(file *models.File) error {
	if file == nil {
		return ErrInvalidData
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	if err := validate.Struct(file); err != nil {
		return err
	}
	if exist, err := m.Exist(file); err != nil {
		return err
	} else if exist {
		return ErrAlreadyExist
	}
	if db.NewRecord(file) {
		return db.Save(file).Error
	}
	return db.Create(file).Error
}

var FilesMapper = new(filesMapper)
