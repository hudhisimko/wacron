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

var AProsesServer = newAProsesServerTable("simko_data", "a_proses_server", "")

type aProsesServerTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	User         mysql.ColumnString
	Host         mysql.ColumnString
	Db           mysql.ColumnString
	Command      mysql.ColumnString
	Time         mysql.ColumnInteger
	State        mysql.ColumnString
	Info         mysql.ColumnString
	TimeMs       mysql.ColumnFloat
	Stage        mysql.ColumnInteger
	MaxStage     mysql.ColumnInteger
	Progress     mysql.ColumnFloat
	MemoryUsed   mysql.ColumnInteger
	ExaminedRows mysql.ColumnInteger
	QueryID      mysql.ColumnInteger
	InfoBinary   mysql.ColumnString
	Tid          mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AProsesServerTable struct {
	aProsesServerTable

	NEW aProsesServerTable
}

// AS creates new AProsesServerTable with assigned alias
func (a AProsesServerTable) AS(alias string) *AProsesServerTable {
	return newAProsesServerTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AProsesServerTable with assigned schema name
func (a AProsesServerTable) FromSchema(schemaName string) *AProsesServerTable {
	return newAProsesServerTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AProsesServerTable with assigned table prefix
func (a AProsesServerTable) WithPrefix(prefix string) *AProsesServerTable {
	return newAProsesServerTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AProsesServerTable with assigned table suffix
func (a AProsesServerTable) WithSuffix(suffix string) *AProsesServerTable {
	return newAProsesServerTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAProsesServerTable(schemaName, tableName, alias string) *AProsesServerTable {
	return &AProsesServerTable{
		aProsesServerTable: newAProsesServerTableImpl(schemaName, tableName, alias),
		NEW:                newAProsesServerTableImpl("", "new", ""),
	}
}

func newAProsesServerTableImpl(schemaName, tableName, alias string) aProsesServerTable {
	var (
		IDColumn           = mysql.IntegerColumn("ID")
		UserColumn         = mysql.StringColumn("USER")
		HostColumn         = mysql.StringColumn("HOST")
		DbColumn           = mysql.StringColumn("DB")
		CommandColumn      = mysql.StringColumn("COMMAND")
		TimeColumn         = mysql.IntegerColumn("TIME")
		StateColumn        = mysql.StringColumn("STATE")
		InfoColumn         = mysql.StringColumn("INFO")
		TimeMsColumn       = mysql.FloatColumn("TIME_MS")
		StageColumn        = mysql.IntegerColumn("STAGE")
		MaxStageColumn     = mysql.IntegerColumn("MAX_STAGE")
		ProgressColumn     = mysql.FloatColumn("PROGRESS")
		MemoryUsedColumn   = mysql.IntegerColumn("MEMORY_USED")
		ExaminedRowsColumn = mysql.IntegerColumn("EXAMINED_ROWS")
		QueryIDColumn      = mysql.IntegerColumn("QUERY_ID")
		InfoBinaryColumn   = mysql.StringColumn("INFO_BINARY")
		TidColumn          = mysql.IntegerColumn("TID")
		allColumns         = mysql.ColumnList{IDColumn, UserColumn, HostColumn, DbColumn, CommandColumn, TimeColumn, StateColumn, InfoColumn, TimeMsColumn, StageColumn, MaxStageColumn, ProgressColumn, MemoryUsedColumn, ExaminedRowsColumn, QueryIDColumn, InfoBinaryColumn, TidColumn}
		mutableColumns     = mysql.ColumnList{IDColumn, UserColumn, HostColumn, DbColumn, CommandColumn, TimeColumn, StateColumn, InfoColumn, TimeMsColumn, StageColumn, MaxStageColumn, ProgressColumn, MemoryUsedColumn, ExaminedRowsColumn, QueryIDColumn, InfoBinaryColumn, TidColumn}
	)

	return aProsesServerTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		User:         UserColumn,
		Host:         HostColumn,
		Db:           DbColumn,
		Command:      CommandColumn,
		Time:         TimeColumn,
		State:        StateColumn,
		Info:         InfoColumn,
		TimeMs:       TimeMsColumn,
		Stage:        StageColumn,
		MaxStage:     MaxStageColumn,
		Progress:     ProgressColumn,
		MemoryUsed:   MemoryUsedColumn,
		ExaminedRows: ExaminedRowsColumn,
		QueryID:      QueryIDColumn,
		InfoBinary:   InfoBinaryColumn,
		Tid:          TidColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
