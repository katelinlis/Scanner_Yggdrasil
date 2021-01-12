package model

import (
	"github.com/jinzhu/gorm"
)

type Peer struct {
	gorm.Model
	Coords string
	Addr   string `gorm:"unique;not null"`
}

type PeerLinks struct {
	gorm.Model
	NodeIDSecond  uint
	NodeIDPrimary uint
}
