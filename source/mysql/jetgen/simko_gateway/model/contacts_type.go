//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type ContactsType string

const (
	ContactsType_Personal ContactsType = "Personal"
	ContactsType_Group    ContactsType = "Group"
	ContactsType_         ContactsType = ""
)

func (e *ContactsType) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "Personal":
		*e = ContactsType_Personal
	case "Group":
		*e = ContactsType_Group
	case "":
		*e = ContactsType_
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for ContactsType enum")
	}

	return nil
}

func (e ContactsType) String() string {
	return string(e)
}
