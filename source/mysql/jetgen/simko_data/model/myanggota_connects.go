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

type MyanggotaConnects struct {
	ConnectID          []byte `sql:"primary_key"`
	Phone              string
	Challenge          *[]byte
	ChallengeSuccessAt *time.Time
	UpdatedAt          time.Time
	CreatedAt          time.Time
}
