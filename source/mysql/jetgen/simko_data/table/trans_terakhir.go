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

var TransTerakhir = newTransTerakhirTable("simko_data", "trans_terakhir", "")

type transTerakhirTable struct {
	mysql.Table

	// Columns
	TglTrans mysql.ColumnDate
	Modul    mysql.ColumnString
	NoUrut   mysql.ColumnInteger
	KodeCab  mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type TransTerakhirTable struct {
	transTerakhirTable

	NEW transTerakhirTable
}

// AS creates new TransTerakhirTable with assigned alias
func (a TransTerakhirTable) AS(alias string) *TransTerakhirTable {
	return newTransTerakhirTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransTerakhirTable with assigned schema name
func (a TransTerakhirTable) FromSchema(schemaName string) *TransTerakhirTable {
	return newTransTerakhirTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransTerakhirTable with assigned table prefix
func (a TransTerakhirTable) WithPrefix(prefix string) *TransTerakhirTable {
	return newTransTerakhirTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransTerakhirTable with assigned table suffix
func (a TransTerakhirTable) WithSuffix(suffix string) *TransTerakhirTable {
	return newTransTerakhirTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransTerakhirTable(schemaName, tableName, alias string) *TransTerakhirTable {
	return &TransTerakhirTable{
		transTerakhirTable: newTransTerakhirTableImpl(schemaName, tableName, alias),
		NEW:                newTransTerakhirTableImpl("", "new", ""),
	}
}

func newTransTerakhirTableImpl(schemaName, tableName, alias string) transTerakhirTable {
	var (
		TglTransColumn = mysql.DateColumn("tgl_trans")
		ModulColumn    = mysql.StringColumn("modul")
		NoUrutColumn   = mysql.IntegerColumn("no_urut")
		KodeCabColumn  = mysql.StringColumn("kode_cab")
		allColumns     = mysql.ColumnList{TglTransColumn, ModulColumn, NoUrutColumn, KodeCabColumn}
		mutableColumns = mysql.ColumnList{TglTransColumn, ModulColumn, NoUrutColumn, KodeCabColumn}
	)

	return transTerakhirTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		TglTrans: TglTransColumn,
		Modul:    ModulColumn,
		NoUrut:   NoUrutColumn,
		KodeCab:  KodeCabColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
