//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/mysql"

var AccountLevel = &struct {
	AccountLevel1 mysql.StringExpression
	AccountLevel2 mysql.StringExpression
}{
	AccountLevel1: mysql.NewEnumValue("1"),
	AccountLevel2: mysql.NewEnumValue("2"),
}
