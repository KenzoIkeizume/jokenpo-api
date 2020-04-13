package model

import "time"

type User struct {
	ID        uint       `gorm:"PRIMARY_KEY" json:"id"`
	Name      string     `gorm:"NOT NULL" json:"name"`
	Age       int        `gorm:"type:int; NOT NULL" json:"age"`
	CreatedAt time.Time  `gorm:"type:datetime; NOT NULL; ASSOCIATION_AUTOCREATE"`
	UpdatedAt time.Time  `gorm:"type:datetime; NOT NULL; ASSOCIATION_AUTOUPDATE"`
	DeletedAt *time.Time `gorm:"type:datetime;"`
}
