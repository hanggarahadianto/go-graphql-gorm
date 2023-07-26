package models

import "time"

type Car struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Size      string    `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}