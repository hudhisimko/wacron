//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type DaftarProduk struct {
	KdProduk   string `sql:"primary_key"`
	NamaProduk string
	Harga      float64
	Kuota      int32
	MasaAktif  int32
	Jenis      string
}
