package models

import "github.com/itrabbit/bunker/pb"

type NameSpace struct {
	ID            uint64       `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	ApplicationID *uint64      `gorm:"index;null" json:"-"`
	Application   *Application `gorm:"foreignkey:ApplicationID" json:"application,omitempty"`
	Alias         string       `gorm:"unque_index;not null" json:"alias" validate:"required"`

	// Parts ['u1', 'welcome', '333']
	Parts StringList `gorm:"type:varchar(512);not null" json:"-"`

	// Need get access key
	IsPrivate bool `json:"isPrivate,omitempty"`

	// Timestamps
	CreatedAt Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *Time `gorm:"null" json:"updatedAt,omitempty"`
	DeletedAt *Time `gorm:"null" json:"deletedAt,omitempty"`
}

func (NameSpace) TableName() string {
	return "namespaces"
}

type NameSpaces []NameSpace

func (n *NameSpace) BeforeCreate() (err error) {
	now := Now()
	if n.CreatedAt.IsZero() {
		n.CreatedAt = now
	}
	n.UpdatedAt = &now
	return
}

func (n *NameSpace) BeforeUpdate() (err error) {
	n.UpdatedAt = NowPtr()
	return
}

func (n NameSpace) PBExport() (res *pb.NameSpace) {
	res = &pb.NameSpace{
		Id:        n.ID,
		Parts:     []string(n.Parts),
		IsPrivate: n.IsPrivate,
		CreatedAt: n.CreatedAt.Unix(),
	}
	if n.UpdatedAt != nil {
		res.UpdatedAt = n.UpdatedAt.Unix()
	}
	if n.DeletedAt != nil {
		res.DeletedAt = n.DeletedAt.Unix()
	}
	return
}

func (ns NameSpaces) PBExport() (res *pb.NameSpaces) {
	res = &pb.NameSpaces{
		Items: make([]*pb.NameSpace, len(ns), len(ns)),
	}
	for i, n := range ns {
		res.Items[i] = n.PBExport()
	}
	return
}
