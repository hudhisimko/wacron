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

var PortalpulsaCallback = newPortalpulsaCallbackTable("simko_data", "portalpulsa_callback", "")

type portalpulsaCallbackTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	Trxid       mysql.ColumnString
	Code        mysql.ColumnString
	Phone       mysql.ColumnString
	Idcust      mysql.ColumnString
	Sequence    mysql.ColumnString
	Status      mysql.ColumnInteger
	Sn          mysql.ColumnString
	Note        mysql.ColumnString
	Price       mysql.ColumnString
	TrxidAPI    mysql.ColumnString
	DateInsert  mysql.ColumnTimestamp
	DateUpdate  mysql.ColumnTimestamp
	LastBalance mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type PortalpulsaCallbackTable struct {
	portalpulsaCallbackTable

	NEW portalpulsaCallbackTable
}

// AS creates new PortalpulsaCallbackTable with assigned alias
func (a PortalpulsaCallbackTable) AS(alias string) *PortalpulsaCallbackTable {
	return newPortalpulsaCallbackTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PortalpulsaCallbackTable with assigned schema name
func (a PortalpulsaCallbackTable) FromSchema(schemaName string) *PortalpulsaCallbackTable {
	return newPortalpulsaCallbackTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PortalpulsaCallbackTable with assigned table prefix
func (a PortalpulsaCallbackTable) WithPrefix(prefix string) *PortalpulsaCallbackTable {
	return newPortalpulsaCallbackTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PortalpulsaCallbackTable with assigned table suffix
func (a PortalpulsaCallbackTable) WithSuffix(suffix string) *PortalpulsaCallbackTable {
	return newPortalpulsaCallbackTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPortalpulsaCallbackTable(schemaName, tableName, alias string) *PortalpulsaCallbackTable {
	return &PortalpulsaCallbackTable{
		portalpulsaCallbackTable: newPortalpulsaCallbackTableImpl(schemaName, tableName, alias),
		NEW:                      newPortalpulsaCallbackTableImpl("", "new", ""),
	}
}

func newPortalpulsaCallbackTableImpl(schemaName, tableName, alias string) portalpulsaCallbackTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		TrxidColumn       = mysql.StringColumn("trxid")
		CodeColumn        = mysql.StringColumn("code")
		PhoneColumn       = mysql.StringColumn("phone")
		IdcustColumn      = mysql.StringColumn("idcust")
		SequenceColumn    = mysql.StringColumn("sequence")
		StatusColumn      = mysql.IntegerColumn("status")
		SnColumn          = mysql.StringColumn("sn")
		NoteColumn        = mysql.StringColumn("note")
		PriceColumn       = mysql.StringColumn("price")
		TrxidAPIColumn    = mysql.StringColumn("trxid_api")
		DateInsertColumn  = mysql.TimestampColumn("date_insert")
		DateUpdateColumn  = mysql.TimestampColumn("date_update")
		LastBalanceColumn = mysql.StringColumn("last_balance")
		allColumns        = mysql.ColumnList{IDColumn, TrxidColumn, CodeColumn, PhoneColumn, IdcustColumn, SequenceColumn, StatusColumn, SnColumn, NoteColumn, PriceColumn, TrxidAPIColumn, DateInsertColumn, DateUpdateColumn, LastBalanceColumn}
		mutableColumns    = mysql.ColumnList{TrxidColumn, CodeColumn, PhoneColumn, IdcustColumn, SequenceColumn, StatusColumn, SnColumn, NoteColumn, PriceColumn, TrxidAPIColumn, DateInsertColumn, DateUpdateColumn, LastBalanceColumn}
	)

	return portalpulsaCallbackTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Trxid:       TrxidColumn,
		Code:        CodeColumn,
		Phone:       PhoneColumn,
		Idcust:      IdcustColumn,
		Sequence:    SequenceColumn,
		Status:      StatusColumn,
		Sn:          SnColumn,
		Note:        NoteColumn,
		Price:       PriceColumn,
		TrxidAPI:    TrxidAPIColumn,
		DateInsert:  DateInsertColumn,
		DateUpdate:  DateUpdateColumn,
		LastBalance: LastBalanceColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
