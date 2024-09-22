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

var WhatsappGetToken = newWhatsappGetTokenTable("simko_data", "whatsapp_get_token", "")

type whatsappGetTokenTable struct {
	mysql.Table

	// Columns
	Memberid   mysql.ColumnString
	Nama       mysql.ColumnString
	NoWhatsapp mysql.ColumnString
	Token      mysql.ColumnString
	URL        mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type WhatsappGetTokenTable struct {
	whatsappGetTokenTable

	NEW whatsappGetTokenTable
}

// AS creates new WhatsappGetTokenTable with assigned alias
func (a WhatsappGetTokenTable) AS(alias string) *WhatsappGetTokenTable {
	return newWhatsappGetTokenTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WhatsappGetTokenTable with assigned schema name
func (a WhatsappGetTokenTable) FromSchema(schemaName string) *WhatsappGetTokenTable {
	return newWhatsappGetTokenTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WhatsappGetTokenTable with assigned table prefix
func (a WhatsappGetTokenTable) WithPrefix(prefix string) *WhatsappGetTokenTable {
	return newWhatsappGetTokenTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WhatsappGetTokenTable with assigned table suffix
func (a WhatsappGetTokenTable) WithSuffix(suffix string) *WhatsappGetTokenTable {
	return newWhatsappGetTokenTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWhatsappGetTokenTable(schemaName, tableName, alias string) *WhatsappGetTokenTable {
	return &WhatsappGetTokenTable{
		whatsappGetTokenTable: newWhatsappGetTokenTableImpl(schemaName, tableName, alias),
		NEW:                   newWhatsappGetTokenTableImpl("", "new", ""),
	}
}

func newWhatsappGetTokenTableImpl(schemaName, tableName, alias string) whatsappGetTokenTable {
	var (
		MemberidColumn   = mysql.StringColumn("memberid")
		NamaColumn       = mysql.StringColumn("nama")
		NoWhatsappColumn = mysql.StringColumn("no_whatsapp")
		TokenColumn      = mysql.StringColumn("token")
		URLColumn        = mysql.StringColumn("url")
		allColumns       = mysql.ColumnList{MemberidColumn, NamaColumn, NoWhatsappColumn, TokenColumn, URLColumn}
		mutableColumns   = mysql.ColumnList{MemberidColumn, NamaColumn, NoWhatsappColumn, TokenColumn, URLColumn}
	)

	return whatsappGetTokenTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Memberid:   MemberidColumn,
		Nama:       NamaColumn,
		NoWhatsapp: NoWhatsappColumn,
		Token:      TokenColumn,
		URL:        URLColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
