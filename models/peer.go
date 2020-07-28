package models

import (
	"github.com/itrabbit/bunker/pb"
	"github.com/itrabbit/rid"
)

type Peer struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Addresses StringList `gorm:"type:varchar(256);not null" json:"addresses" validate:"required"`
	Name      string     `gorm:"unque_index;not null" validate:"required"`

	// Timestamps
	CreatedAt Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *Time `gorm:"null" json:"updatedAt,omitempty"`
	DeletedAt *Time `gorm:"null" json:"deletedAt,omitempty"`
}

func (Peer) TableName() string {
	return "peers"
}

type Peers []Peer

func (p *Peer) BeforeCreate() (err error) {
	now := Now()
	if len(p.Name) < 1 {
		p.Name = rid.New().String()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	p.UpdatedAt = &now
	return
}

func (p *Peer) BeforeUpdate() (err error) {
	p.UpdatedAt = NowPtr()
	return
}

func (p Peer) PBExport() (res *pb.Peer) {
	res = &pb.Peer{
		Id:        p.ID,
		Name:      p.Name,
		Addresses: p.Addresses,
		CreatedAt: p.CreatedAt.Unix(),
	}
	if p.UpdatedAt != nil {
		res.UpdatedAt = p.UpdatedAt.Unix()
	}
	if p.DeletedAt != nil {
		res.DeletedAt = p.DeletedAt.Unix()
	}
	return
}

func (ps Peers) PBExport() (res *pb.Peers) {
	res = &pb.Peers{
		Items: make([]*pb.Peer, len(ps), len(ps)),
	}
	for i, p := range ps {
		res.Items[i] = p.PBExport()
	}
	return
}
