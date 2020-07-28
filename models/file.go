package models

type File struct {
	ID          uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Alias       string     `gorm:"unque_index;not null" json:"alias" validate:"required"`
	OwnerID     *uint64    `gorm:"index;null" json:"-"`
	NameSpaceID *uint64    `gorm:"index;null" json:"-"`
	NameSpace   *NameSpace `gorm:"foreignkey:NameSpaceID" json:"namespace"`
	PeerID      *uint64    `gorm:"index;null" json:"-"`
	Peer        *Peer      `gorm:"foreignkey:PeerID" json:"-"`

	// Info
	Hash         string `gorm:"not null" json:"hash" validate:"required"`
	Size         uint64 `gorm:"not null" json:"size" validate:"required"`
	MimeType     string `gorm:"index;not null" json:"mimeType" validate:"required"`
	OriginalName string `gorm:"null" json:"originalName,omitempty"`

	// Variants get content
	Variants   StringList `gorm:"type:varchar(1024);null" json:"variants,omitempty"`
	Processing bool       `json:"processing,omitempty"`

	// Metadata
	Metadata HashMap `gorm:"type:longtext;null" json:"metadata,omitempty"`

	// Timestamps
	CreatedAt Time  `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *Time `gorm:"null" json:"updatedAt,omitempty"`
	DeletedAt *Time `gorm:"null" json:"deletedAt,omitempty"`
}

type Files []File

func (File) TableName() string {
	return "files"
}

func (f *File) BeforeCreate() (err error) {
	now := Now()
	if f.CreatedAt.IsZero() {
		f.CreatedAt = now
	}
	f.UpdatedAt = &now
	return
}

func (f *File) BeforeUpdate() (err error) {
	f.UpdatedAt = NowPtr()
	return
}
