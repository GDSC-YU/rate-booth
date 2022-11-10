// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Booth struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	LogoUrl      string    `json:"logo_url"`
	Description  string    `json:"description"`
	TwitterUrl   string    `json:"twitter_url"`
	InstagramUrl string    `json:"instagram_url"`
}

type Rating struct {
	ID        uuid.UUID `json:"id"`
	Price     int32     `json:"price"`
	Quality   int32     `json:"quality"`
	CreatedAt time.Time `json:"created_at"`
	BoothID   uuid.UUID `json:"booth_id"`
}
