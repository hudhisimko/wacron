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

type Lembaga struct {
	Memberid            string `sql:"primary_key"`
	Jenis               *string
	Nama                *string
	Alamat              *string
	Kota                *string
	Email               *string
	TglMasuk            *time.Time
	Manager             *string
	Telp                *string
	Db                  *string
	Web                 *string
	URLLembaga          string
	ClientService       string
	AuthKey             string
	ContentType         string
	Aktif               *int32
	Whatsapp            int32
	ModeWa              string
	IjinkanSmsTag       bool
	IjinkanHapusMundur  bool
	UpdateProgramKhusus bool
	BmtcheckingJepara   bool
	ServerWhatsapp      string
	ServerOtp           string
	LisensiSistem       LembagaLisensiSistem
	TglExpSistem        time.Time
	TglExpWa            time.Time
	TglExpSms           time.Time
	KuotaAwalWa         int32
	KuotaTransWa        int32
	KuotaAkhirWa        int32
	KuotaAwalSms        int32
	KuotaTransSms       int32
	KuotaAkhirSms       int32
	PaketSistem         string
	PaketWa             string
	PaketSms            string
	IjinkanTransMundur  string
	SmsStatus           int8
	SmsServer           string
	ApikeyWadoh         string
	ApikeyCab           string
}
