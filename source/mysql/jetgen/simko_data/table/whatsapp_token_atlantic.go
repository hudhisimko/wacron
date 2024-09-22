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

var WhatsappTokenAtlantic = newWhatsappTokenAtlanticTable("simko_data", "whatsapp_token_atlantic", "")

type whatsappTokenAtlanticTable struct {
	mysql.Table

	// Columns
	NoWhatsapp  mysql.ColumnString
	Token       mysql.ColumnString
	URL         mysql.ColumnString
	StatusToken mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type WhatsappTokenAtlanticTable struct {
	whatsappTokenAtlanticTable

	NEW whatsappTokenAtlanticTable
}

// AS creates new WhatsappTokenAtlanticTable with assigned alias
func (a WhatsappTokenAtlanticTable) AS(alias string) *WhatsappTokenAtlanticTable {
	return newWhatsappTokenAtlanticTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WhatsappTokenAtlanticTable with assigned schema name
func (a WhatsappTokenAtlanticTable) FromSchema(schemaName string) *WhatsappTokenAtlanticTable {
	return newWhatsappTokenAtlanticTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WhatsappTokenAtlanticTable with assigned table prefix
func (a WhatsappTokenAtlanticTable) WithPrefix(prefix string) *WhatsappTokenAtlanticTable {
	return newWhatsappTokenAtlanticTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WhatsappTokenAtlanticTable with assigned table suffix
func (a WhatsappTokenAtlanticTable) WithSuffix(suffix string) *WhatsappTokenAtlanticTable {
	return newWhatsappTokenAtlanticTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWhatsappTokenAtlanticTable(schemaName, tableName, alias string) *WhatsappTokenAtlanticTable {
	return &WhatsappTokenAtlanticTable{
		whatsappTokenAtlanticTable: newWhatsappTokenAtlanticTableImpl(schemaName, tableName, alias),
		NEW:                        newWhatsappTokenAtlanticTableImpl("", "new", ""),
	}
}

func newWhatsappTokenAtlanticTableImpl(schemaName, tableName, alias string) whatsappTokenAtlanticTable {
	var (
		NoWhatsappColumn  = mysql.StringColumn("no_whatsapp")
		TokenColumn       = mysql.StringColumn("token")
		URLColumn         = mysql.StringColumn("url")
		StatusTokenColumn = mysql.BoolColumn("status_token")
		allColumns        = mysql.ColumnList{NoWhatsappColumn, TokenColumn, URLColumn, StatusTokenColumn}
		mutableColumns    = mysql.ColumnList{NoWhatsappColumn, TokenColumn, URLColumn, StatusTokenColumn}
	)

	return whatsappTokenAtlanticTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		NoWhatsapp:  NoWhatsappColumn,
		Token:       TokenColumn,
		URL:         URLColumn,
		StatusToken: StatusTokenColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
