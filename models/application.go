package models

import "github.com/itrabbit/bunker/pb"

type Application struct {
	ID   uint64 `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Name string `gorm:"not null" json:"name"`
	Key  string `gorm:"unique;not null" json:"-"`

	// Timestamps
	CreatedAt Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *Time `gorm:"null" json:"updatedAt,omitempty"`
	DeletedAt *Time `gorm:"null" json:"deletedAt,omitempty"`
}

func (Application) TableName() string {
	return "applications"
}

type Applications []Application

func (a *Application) BeforeCreate() (err error) {
	now := Now()
	if a.CreatedAt.IsZero() {
		a.CreatedAt = now
	}
	a.UpdatedAt = &now
	return
}

func (a *Application) BeforeUpdate() (err error) {
	a.UpdatedAt = NowPtr()
	return
}

func (a Application) PBExport() (res *pb.Application) {
	res = &pb.Application{
		Id:        a.ID,
		Key:       a.Key,
		Name:      a.Name,
		CreatedAt: a.CreatedAt.Unix(),
	}
	if a.UpdatedAt != nil {
		res.UpdatedAt = a.UpdatedAt.Unix()
	}
	if a.DeletedAt != nil {
		res.DeletedAt = a.DeletedAt.Unix()
	}
	return
}

func (as Applications) PBExport() (res *pb.Applications) {
	res = &pb.Applications{
		Items: make([]*pb.Application, len(as), len(as)),
	}
	for i, a := range as {
		res.Items[i] = a.PBExport()
	}
	return
}
