//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/mysql"
)

var ViewDataNasabah = newViewDataNasabahTable("simko_data", "view_data_nasabah", "")

type viewDataNasabahTable struct {
	mysql.Table

	// Columns
	NamaLembaga    mysql.ColumnString
	KodeNasabah    mysql.ColumnString
	KodeLembaga    mysql.ColumnString
	NasabahID      mysql.ColumnString
	NamaNasabah    mysql.ColumnString
	AlamatNasabah  mysql.ColumnString
	JenisIdentitas mysql.ColumnString
	NoIdentitas    mysql.ColumnString
	KodeCabang     mysql.ColumnString
	URLLembaga     mysql.ColumnString
	ClientService  mysql.ColumnString
	AuthKey        mysql.ColumnString
	ContentType    mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ViewDataNasabahTable struct {
	viewDataNasabahTable

	NEW viewDataNasabahTable
}

// AS creates new ViewDataNasabahTable with assigned alias
func (a ViewDataNasabahTable) AS(alias string) *ViewDataNasabahTable {
	return newViewDataNasabahTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ViewDataNasabahTable with assigned schema name
func (a ViewDataNasabahTable) FromSchema(schemaName string) *ViewDataNasabahTable {
	return newViewDataNasabahTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ViewDataNasabahTable with assigned table prefix
func (a ViewDataNasabahTable) WithPrefix(prefix string) *ViewDataNasabahTable {
	return newViewDataNasabahTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ViewDataNasabahTable with assigned table suffix
func (a ViewDataNasabahTable) WithSuffix(suffix string) *ViewDataNasabahTable {
	return newViewDataNasabahTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newViewDataNasabahTable(schemaName, tableName, alias string) *ViewDataNasabahTable {
	return &ViewDataNasabahTable{
		viewDataNasabahTable: newViewDataNasabahTableImpl(schemaName, tableName, alias),
		NEW:                  newViewDataNasabahTableImpl("", "new", ""),
	}
}

func newViewDataNasabahTableImpl(schemaName, tableName, alias string) viewDataNasabahTable {
	var (
		NamaLembagaColumn    = mysql.StringColumn("nama_lembaga")
		KodeNasabahColumn    = mysql.StringColumn("kode_nasabah")
		KodeLembagaColumn    = mysql.StringColumn("kode_lembaga")
		NasabahIDColumn      = mysql.StringColumn("nasabah_id")
		NamaNasabahColumn    = mysql.StringColumn("nama_nasabah")
		AlamatNasabahColumn  = mysql.StringColumn("alamat_nasabah")
		JenisIdentitasColumn = mysql.StringColumn("jenis_identitas")
		NoIdentitasColumn    = mysql.StringColumn("no_identitas")
		KodeCabangColumn     = mysql.StringColumn("kode_cabang")
		URLLembagaColumn     = mysql.StringColumn("url_lembaga")
		ClientServiceColumn  = mysql.StringColumn("CLIENT_SERVICE")
		AuthKeyColumn        = mysql.StringColumn("AUTH_KEY")
		ContentTypeColumn    = mysql.StringColumn("CONTENT_TYPE")
		allColumns           = mysql.ColumnList{NamaLembagaColumn, KodeNasabahColumn, KodeLembagaColumn, NasabahIDColumn, NamaNasabahColumn, AlamatNasabahColumn, JenisIdentitasColumn, NoIdentitasColumn, KodeCabangColumn, URLLembagaColumn, ClientServiceColumn, AuthKeyColumn, ContentTypeColumn}
		mutableColumns       = mysql.ColumnList{NamaLembagaColumn, KodeNasabahColumn, KodeLembagaColumn, NasabahIDColumn, NamaNasabahColumn, AlamatNasabahColumn, JenisIdentitasColumn, NoIdentitasColumn, KodeCabangColumn, URLLembagaColumn, ClientServiceColumn, AuthKeyColumn, ContentTypeColumn}
	)

	return viewDataNasabahTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		NamaLembaga:    NamaLembagaColumn,
		KodeNasabah:    KodeNasabahColumn,
		KodeLembaga:    KodeLembagaColumn,
		NasabahID:      NasabahIDColumn,
		NamaNasabah:    NamaNasabahColumn,
		AlamatNasabah:  AlamatNasabahColumn,
		JenisIdentitas: JenisIdentitasColumn,
		NoIdentitas:    NoIdentitasColumn,
		KodeCabang:     KodeCabangColumn,
		URLLembaga:     URLLembagaColumn,
		ClientService:  ClientServiceColumn,
		AuthKey:        AuthKeyColumn,
		ContentType:    ContentTypeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
