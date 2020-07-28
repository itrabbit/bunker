package db

import (
	"github.com/itrabbit/bunker/models"
)

type peersMapper struct {
}

func (m *peersMapper) All() (peers models.Peers, err error) {
	if db == nil {
		err = ErrInvalidDbConnection
		return
	}
	err = db.Find(&peers).Error
	return
}

func (m *peersMapper) Exist(peer *models.Peer) (bool, error) {
	if peer == nil {
		return false, nil
	}
	if db == nil {
		return false, ErrInvalidDbConnection
	}
	count := 0
	if err := db.Model(&models.Peer{}).Where(
		"id != ? AND name = ?",
		peer.ID,
		peer.Name,
	).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *peersMapper) Delete(peer *models.Peer) error {
	if peer == nil {
		return nil
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	return db.Delete(peer).Error
}

func (m *peersMapper) Save(peer *models.Peer) error {
	if peer == nil {
		return ErrInvalidData
	}
	if db == nil {
		return ErrInvalidDbConnection
	}
	if exist, err := m.Exist(peer); err != nil {
		return err
	} else if exist {
		return ErrAlreadyExist
	}
	if !db.NewRecord(peer) {
		return db.Save(peer).Error
	}
	return db.Create(peer).Error
}

var PeersMapper = new(peersMapper)
