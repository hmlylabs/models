package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" sql:"index"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New()
	if base.ID == uuid.Nil {
		return errors.New("failed to save data invalid uuid")
	}
	return nil
}
