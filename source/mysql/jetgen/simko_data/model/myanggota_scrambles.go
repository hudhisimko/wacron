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

type MyanggotaScrambles struct {
	Phone     string
	Scramble  []byte `sql:"primary_key"`
	CreatedAt time.Time
	Age       *int32
}
