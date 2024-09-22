//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var PengajuanAktif = newPengajuanAktifTable("simko_data", "pengajuan_aktif", "")

type pengajuanAktifTable struct {
	mysql.Table

	// Columns
	KodePengajuan    mysql.ColumnString
	TglPengajuan     mysql.ColumnDate
	NoHp             mysql.ColumnString
	Keterangan       mysql.ColumnString
	Nominal          mysql.ColumnFloat
	NominalDisetujui mysql.ColumnFloat
	KdLembaga        mysql.ColumnString
	KdCabang         mysql.ColumnString
	NoAnggota        mysql.ColumnString
	NoRekening       mysql.ColumnString
	StatusPengajuan  mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type PengajuanAktifTable struct {
	pengajuanAktifTable

	NEW pengajuanAktifTable
}

// AS creates new PengajuanAktifTable with assigned alias
func (a PengajuanAktifTable) AS(alias string) *PengajuanAktifTable {
	return newPengajuanAktifTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PengajuanAktifTable with assigned schema name
func (a PengajuanAktifTable) FromSchema(schemaName string) *PengajuanAktifTable {
	return newPengajuanAktifTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PengajuanAktifTable with assigned table prefix
func (a PengajuanAktifTable) WithPrefix(prefix string) *PengajuanAktifTable {
	return newPengajuanAktifTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PengajuanAktifTable with assigned table suffix
func (a PengajuanAktifTable) WithSuffix(suffix string) *PengajuanAktifTable {
	return newPengajuanAktifTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPengajuanAktifTable(schemaName, tableName, alias string) *PengajuanAktifTable {
	return &PengajuanAktifTable{
		pengajuanAktifTable: newPengajuanAktifTableImpl(schemaName, tableName, alias),
		NEW:                 newPengajuanAktifTableImpl("", "new", ""),
	}
}

func newPengajuanAktifTableImpl(schemaName, tableName, alias string) pengajuanAktifTable {
	var (
		KodePengajuanColumn    = mysql.StringColumn("kode_pengajuan")
		TglPengajuanColumn     = mysql.DateColumn("tgl_pengajuan")
		NoHpColumn             = mysql.StringColumn("no_hp")
		KeteranganColumn       = mysql.StringColumn("keterangan")
		NominalColumn          = mysql.FloatColumn("nominal")
		NominalDisetujuiColumn = mysql.FloatColumn("nominal_disetujui")
		KdLembagaColumn        = mysql.StringColumn("kd_lembaga")
		KdCabangColumn         = mysql.StringColumn("kd_cabang")
		NoAnggotaColumn        = mysql.StringColumn("no_anggota")
		NoRekeningColumn       = mysql.StringColumn("no_rekening")
		StatusPengajuanColumn  = mysql.BoolColumn("status_pengajuan")
		allColumns             = mysql.ColumnList{KodePengajuanColumn, TglPengajuanColumn, NoHpColumn, KeteranganColumn, NominalColumn, NominalDisetujuiColumn, KdLembagaColumn, KdCabangColumn, NoAnggotaColumn, NoRekeningColumn, StatusPengajuanColumn}
		mutableColumns         = mysql.ColumnList{TglPengajuanColumn, NoHpColumn, KeteranganColumn, NominalColumn, NominalDisetujuiColumn, KdLembagaColumn, KdCabangColumn, NoAnggotaColumn, NoRekeningColumn, StatusPengajuanColumn}
	)

	return pengajuanAktifTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		KodePengajuan:    KodePengajuanColumn,
		TglPengajuan:     TglPengajuanColumn,
		NoHp:             NoHpColumn,
		Keterangan:       KeteranganColumn,
		Nominal:          NominalColumn,
		NominalDisetujui: NominalDisetujuiColumn,
		KdLembaga:        KdLembagaColumn,
		KdCabang:         KdCabangColumn,
		NoAnggota:        NoAnggotaColumn,
		NoRekening:       NoRekeningColumn,
		StatusPengajuan:  StatusPengajuanColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
