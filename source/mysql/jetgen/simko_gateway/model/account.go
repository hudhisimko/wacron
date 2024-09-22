//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Account struct {
	ID           int32 `sql:"primary_key"`
	Username     string
	Password     string
	APIKey       string
	Level        AccountLevel // 1 = ADMIN 2 = USER
	Chunk        int32
	Photo        string
	Whatsapp     string
	Aktif        AccountAktif
	Token        string
	Port         string
	Logo         string
	URLDownload  string
	DateDownload *time.Time
	ScanQr       int32
}
