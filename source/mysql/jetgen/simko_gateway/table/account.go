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

var Account = newAccountTable("simko_gateway", "account", "")

type accountTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	Username     mysql.ColumnString
	Password     mysql.ColumnString
	APIKey       mysql.ColumnString
	Level        mysql.ColumnString // 1 = ADMIN 2 = USER
	Chunk        mysql.ColumnInteger
	Photo        mysql.ColumnString
	Whatsapp     mysql.ColumnString
	Aktif        mysql.ColumnString
	Token        mysql.ColumnString
	Port         mysql.ColumnString
	Logo         mysql.ColumnString
	URLDownload  mysql.ColumnString
	DateDownload mysql.ColumnTimestamp
	ScanQr       mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type AccountTable struct {
	accountTable

	NEW accountTable
}

// AS creates new AccountTable with assigned alias
func (a AccountTable) AS(alias string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountTable with assigned schema name
func (a AccountTable) FromSchema(schemaName string) *AccountTable {
	return newAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountTable with assigned table prefix
func (a AccountTable) WithPrefix(prefix string) *AccountTable {
	return newAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountTable with assigned table suffix
func (a AccountTable) WithSuffix(suffix string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountTable(schemaName, tableName, alias string) *AccountTable {
	return &AccountTable{
		accountTable: newAccountTableImpl(schemaName, tableName, alias),
		NEW:          newAccountTableImpl("", "new", ""),
	}
}

func newAccountTableImpl(schemaName, tableName, alias string) accountTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		UsernameColumn     = mysql.StringColumn("username")
		PasswordColumn     = mysql.StringColumn("password")
		APIKeyColumn       = mysql.StringColumn("api_key")
		LevelColumn        = mysql.StringColumn("level")
		ChunkColumn        = mysql.IntegerColumn("chunk")
		PhotoColumn        = mysql.StringColumn("photo")
		WhatsappColumn     = mysql.StringColumn("whatsapp")
		AktifColumn        = mysql.StringColumn("aktif")
		TokenColumn        = mysql.StringColumn("token")
		PortColumn         = mysql.StringColumn("port")
		LogoColumn         = mysql.StringColumn("logo")
		URLDownloadColumn  = mysql.StringColumn("url_download")
		DateDownloadColumn = mysql.TimestampColumn("date_download")
		ScanQrColumn       = mysql.IntegerColumn("scan_qr")
		allColumns         = mysql.ColumnList{IDColumn, UsernameColumn, PasswordColumn, APIKeyColumn, LevelColumn, ChunkColumn, PhotoColumn, WhatsappColumn, AktifColumn, TokenColumn, PortColumn, LogoColumn, URLDownloadColumn, DateDownloadColumn, ScanQrColumn}
		mutableColumns     = mysql.ColumnList{UsernameColumn, PasswordColumn, APIKeyColumn, LevelColumn, ChunkColumn, PhotoColumn, WhatsappColumn, AktifColumn, TokenColumn, PortColumn, LogoColumn, URLDownloadColumn, DateDownloadColumn, ScanQrColumn}
	)

	return accountTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		Username:     UsernameColumn,
		Password:     PasswordColumn,
		APIKey:       APIKeyColumn,
		Level:        LevelColumn,
		Chunk:        ChunkColumn,
		Photo:        PhotoColumn,
		Whatsapp:     WhatsappColumn,
		Aktif:        AktifColumn,
		Token:        TokenColumn,
		Port:         PortColumn,
		Logo:         LogoColumn,
		URLDownload:  URLDownloadColumn,
		DateDownload: DateDownloadColumn,
		ScanQr:       ScanQrColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
