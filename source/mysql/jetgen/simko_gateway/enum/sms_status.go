//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/mysql"

var SmsStatus = &struct {
	MenungguJadwal mysql.StringExpression
	Gagal          mysql.StringExpression
	Terkirim       mysql.StringExpression
	Habis          mysql.StringExpression
}{
	MenungguJadwal: mysql.NewEnumValue("MENUNGGU JADWAL"),
	Gagal:          mysql.NewEnumValue("GAGAL"),
	Terkirim:       mysql.NewEnumValue("TERKIRIM"),
	Habis:          mysql.NewEnumValue("HABIS"),
}
