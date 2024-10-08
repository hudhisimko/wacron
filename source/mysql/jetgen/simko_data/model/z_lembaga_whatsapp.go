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

type ZLembagaWhatsapp struct {
	LembagaID     string
	NamaLembaga   string
	KodeCab       string
	Provider      string
	Data          string
	Header        string
	SleepAfter    int32
	SleepDuration int32
	SendDelay     int32
	LastSent      time.Time
	Idle          *int32
	LastSleep     time.Time
	Ready         *bool
	Sleeping      *bool
	RunDur        *int32
	DateExpired   *time.Time
}
