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

var WhatsappToken = newWhatsappTokenTable("simko_data", "whatsapp_token", "")

type whatsappTokenTable struct {
	mysql.Table

	// Columns
	StatusAktif mysql.ColumnBool
	NoWhatsapp  mysql.ColumnString
	Token       mysql.ColumnString
	URL         mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type WhatsappTokenTable struct {
	whatsappTokenTable

	NEW whatsappTokenTable
}

// AS creates new WhatsappTokenTable with assigned alias
func (a WhatsappTokenTable) AS(alias string) *WhatsappTokenTable {
	return newWhatsappTokenTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WhatsappTokenTable with assigned schema name
func (a WhatsappTokenTable) FromSchema(schemaName string) *WhatsappTokenTable {
	return newWhatsappTokenTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WhatsappTokenTable with assigned table prefix
func (a WhatsappTokenTable) WithPrefix(prefix string) *WhatsappTokenTable {
	return newWhatsappTokenTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WhatsappTokenTable with assigned table suffix
func (a WhatsappTokenTable) WithSuffix(suffix string) *WhatsappTokenTable {
	return newWhatsappTokenTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWhatsappTokenTable(schemaName, tableName, alias string) *WhatsappTokenTable {
	return &WhatsappTokenTable{
		whatsappTokenTable: newWhatsappTokenTableImpl(schemaName, tableName, alias),
		NEW:                newWhatsappTokenTableImpl("", "new", ""),
	}
}

func newWhatsappTokenTableImpl(schemaName, tableName, alias string) whatsappTokenTable {
	var (
		StatusAktifColumn = mysql.BoolColumn("status_aktif")
		NoWhatsappColumn  = mysql.StringColumn("no_whatsapp")
		TokenColumn       = mysql.StringColumn("token")
		URLColumn         = mysql.StringColumn("url")
		allColumns        = mysql.ColumnList{StatusAktifColumn, NoWhatsappColumn, TokenColumn, URLColumn}
		mutableColumns    = mysql.ColumnList{NoWhatsappColumn, TokenColumn, URLColumn}
	)

	return whatsappTokenTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StatusAktif: StatusAktifColumn,
		NoWhatsapp:  NoWhatsappColumn,
		Token:       TokenColumn,
		URL:         URLColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
