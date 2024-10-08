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

var MyanggotaWhatsappIn = newMyanggotaWhatsappInTable("simko_data", "myanggota_whatsapp_in", "")

type myanggotaWhatsappInTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnInteger
	Message   mysql.ColumnString
	From      mysql.ColumnString
	Image     mysql.ColumnString
	CreatedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type MyanggotaWhatsappInTable struct {
	myanggotaWhatsappInTable

	NEW myanggotaWhatsappInTable
}

// AS creates new MyanggotaWhatsappInTable with assigned alias
func (a MyanggotaWhatsappInTable) AS(alias string) *MyanggotaWhatsappInTable {
	return newMyanggotaWhatsappInTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new MyanggotaWhatsappInTable with assigned schema name
func (a MyanggotaWhatsappInTable) FromSchema(schemaName string) *MyanggotaWhatsappInTable {
	return newMyanggotaWhatsappInTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new MyanggotaWhatsappInTable with assigned table prefix
func (a MyanggotaWhatsappInTable) WithPrefix(prefix string) *MyanggotaWhatsappInTable {
	return newMyanggotaWhatsappInTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new MyanggotaWhatsappInTable with assigned table suffix
func (a MyanggotaWhatsappInTable) WithSuffix(suffix string) *MyanggotaWhatsappInTable {
	return newMyanggotaWhatsappInTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newMyanggotaWhatsappInTable(schemaName, tableName, alias string) *MyanggotaWhatsappInTable {
	return &MyanggotaWhatsappInTable{
		myanggotaWhatsappInTable: newMyanggotaWhatsappInTableImpl(schemaName, tableName, alias),
		NEW:                      newMyanggotaWhatsappInTableImpl("", "new", ""),
	}
}

func newMyanggotaWhatsappInTableImpl(schemaName, tableName, alias string) myanggotaWhatsappInTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		MessageColumn   = mysql.StringColumn("message")
		FromColumn      = mysql.StringColumn("from")
		ImageColumn     = mysql.StringColumn("image")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		allColumns      = mysql.ColumnList{IDColumn, MessageColumn, FromColumn, ImageColumn, CreatedAtColumn}
		mutableColumns  = mysql.ColumnList{MessageColumn, FromColumn, ImageColumn, CreatedAtColumn}
	)

	return myanggotaWhatsappInTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Message:   MessageColumn,
		From:      FromColumn,
		Image:     ImageColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
