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

var MyanggotaScrambles = newMyanggotaScramblesTable("simko_data", "myanggota_scrambles", "")

type myanggotaScramblesTable struct {
	mysql.Table

	// Columns
	Phone     mysql.ColumnString
	Scramble  mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp
	Age       mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MyanggotaScramblesTable struct {
	myanggotaScramblesTable

	NEW myanggotaScramblesTable
}

// AS creates new MyanggotaScramblesTable with assigned alias
func (a MyanggotaScramblesTable) AS(alias string) *MyanggotaScramblesTable {
	return newMyanggotaScramblesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MyanggotaScramblesTable with assigned schema name
func (a MyanggotaScramblesTable) FromSchema(schemaName string) *MyanggotaScramblesTable {
	return newMyanggotaScramblesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MyanggotaScramblesTable with assigned table prefix
func (a MyanggotaScramblesTable) WithPrefix(prefix string) *MyanggotaScramblesTable {
	return newMyanggotaScramblesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MyanggotaScramblesTable with assigned table suffix
func (a MyanggotaScramblesTable) WithSuffix(suffix string) *MyanggotaScramblesTable {
	return newMyanggotaScramblesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMyanggotaScramblesTable(schemaName, tableName, alias string) *MyanggotaScramblesTable {
	return &MyanggotaScramblesTable{
		myanggotaScramblesTable: newMyanggotaScramblesTableImpl(schemaName, tableName, alias),
		NEW:                     newMyanggotaScramblesTableImpl("", "new", ""),
	}
}

func newMyanggotaScramblesTableImpl(schemaName, tableName, alias string) myanggotaScramblesTable {
	var (
		PhoneColumn     = mysql.StringColumn("phone")
		ScrambleColumn  = mysql.StringColumn("scramble")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		AgeColumn       = mysql.IntegerColumn("age")
		allColumns      = mysql.ColumnList{PhoneColumn, ScrambleColumn, CreatedAtColumn, AgeColumn}
		mutableColumns  = mysql.ColumnList{PhoneColumn, CreatedAtColumn, AgeColumn}
	)

	return myanggotaScramblesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Phone:     PhoneColumn,
		Scramble:  ScrambleColumn,
		CreatedAt: CreatedAtColumn,
		Age:       AgeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
