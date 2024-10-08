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

var MyanggotaTrans = newMyanggotaTransTable("simko_data", "myanggota_trans", "")

type myanggotaTransTable struct {
	mysql.Table

	// Columns
	KodeRek    mysql.ColumnString
	NoRekening mysql.ColumnString
	Modul      mysql.ColumnString
	KdLembaga  mysql.ColumnString
	NoHp       mysql.ColumnString
	StatusData mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MyanggotaTransTable struct {
	myanggotaTransTable

	NEW myanggotaTransTable
}

// AS creates new MyanggotaTransTable with assigned alias
func (a MyanggotaTransTable) AS(alias string) *MyanggotaTransTable {
	return newMyanggotaTransTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MyanggotaTransTable with assigned schema name
func (a MyanggotaTransTable) FromSchema(schemaName string) *MyanggotaTransTable {
	return newMyanggotaTransTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MyanggotaTransTable with assigned table prefix
func (a MyanggotaTransTable) WithPrefix(prefix string) *MyanggotaTransTable {
	return newMyanggotaTransTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MyanggotaTransTable with assigned table suffix
func (a MyanggotaTransTable) WithSuffix(suffix string) *MyanggotaTransTable {
	return newMyanggotaTransTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMyanggotaTransTable(schemaName, tableName, alias string) *MyanggotaTransTable {
	return &MyanggotaTransTable{
		myanggotaTransTable: newMyanggotaTransTableImpl(schemaName, tableName, alias),
		NEW:                 newMyanggotaTransTableImpl("", "new", ""),
	}
}

func newMyanggotaTransTableImpl(schemaName, tableName, alias string) myanggotaTransTable {
	var (
		KodeRekColumn    = mysql.StringColumn("kode_rek")
		NoRekeningColumn = mysql.StringColumn("no_rekening")
		ModulColumn      = mysql.StringColumn("modul")
		KdLembagaColumn  = mysql.StringColumn("kd_lembaga")
		NoHpColumn       = mysql.StringColumn("no_hp")
		StatusDataColumn = mysql.BoolColumn("status_data")
		allColumns       = mysql.ColumnList{KodeRekColumn, NoRekeningColumn, ModulColumn, KdLembagaColumn, NoHpColumn, StatusDataColumn}
		mutableColumns   = mysql.ColumnList{NoRekeningColumn, ModulColumn, KdLembagaColumn, NoHpColumn, StatusDataColumn}
	)

	return myanggotaTransTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		KodeRek:    KodeRekColumn,
		NoRekening: NoRekeningColumn,
		Modul:      ModulColumn,
		KdLembaga:  KdLembagaColumn,
		NoHp:       NoHpColumn,
		StatusData: StatusDataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
