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

var Autoreply = newAutoreplyTable("simko_gateway", "autoreply", "")

type autoreplyTable struct {
	mysql.Table

	// Columns
	ID       mysql.ColumnInteger
	Keyword  mysql.ColumnString
	Response mysql.ColumnString
	Media    mysql.ColumnString
	Nomor    mysql.ColumnString
	MakeBy   mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AutoreplyTable struct {
	autoreplyTable

	NEW autoreplyTable
}

// AS creates new AutoreplyTable with assigned alias
func (a AutoreplyTable) AS(alias string) *AutoreplyTable {
	return newAutoreplyTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AutoreplyTable with assigned schema name
func (a AutoreplyTable) FromSchema(schemaName string) *AutoreplyTable {
	return newAutoreplyTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AutoreplyTable with assigned table prefix
func (a AutoreplyTable) WithPrefix(prefix string) *AutoreplyTable {
	return newAutoreplyTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AutoreplyTable with assigned table suffix
func (a AutoreplyTable) WithSuffix(suffix string) *AutoreplyTable {
	return newAutoreplyTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAutoreplyTable(schemaName, tableName, alias string) *AutoreplyTable {
	return &AutoreplyTable{
		autoreplyTable: newAutoreplyTableImpl(schemaName, tableName, alias),
		NEW:            newAutoreplyTableImpl("", "new", ""),
	}
}

func newAutoreplyTableImpl(schemaName, tableName, alias string) autoreplyTable {
	var (
		IDColumn       = mysql.IntegerColumn("id")
		KeywordColumn  = mysql.StringColumn("keyword")
		ResponseColumn = mysql.StringColumn("response")
		MediaColumn    = mysql.StringColumn("media")
		NomorColumn    = mysql.StringColumn("nomor")
		MakeByColumn   = mysql.StringColumn("make_by")
		allColumns     = mysql.ColumnList{IDColumn, KeywordColumn, ResponseColumn, MediaColumn, NomorColumn, MakeByColumn}
		mutableColumns = mysql.ColumnList{KeywordColumn, ResponseColumn, MediaColumn, NomorColumn, MakeByColumn}
	)

	return autoreplyTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Keyword:  KeywordColumn,
		Response: ResponseColumn,
		Media:    MediaColumn,
		Nomor:    NomorColumn,
		MakeBy:   MakeByColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
