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

type PortalpulsaCallback struct {
	ID          int64 `sql:"primary_key"`
	Trxid       *string
	Code        *string
	Phone       *string
	Idcust      *string
	Sequence    *string
	Status      *int32
	Sn          *string
	Note        *string
	Price       *string
	TrxidAPI    *string
	DateInsert  *time.Time
	DateUpdate  *time.Time
	LastBalance *string
}
