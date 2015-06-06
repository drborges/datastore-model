package db

import (
	"reflect"
	"errors"
)

var (
	ErrMissingStringId = errors.New(`Entity is missing StringId. String field tagged with db:"id" cannot be empty.`)
	ErrMissingIntId    = errors.New(`Entity is missing IntId. Integer field tagged with db:"id" cannot be zero.`)
)

// ExtractEntityKeyIDs extracts the StringID and IntID
// datastore key components from struct tags
//
// e.g.:
//
// The following struct declares an id tag on a field
// of type string, thus its StringID.
//
// type Person struct {
//   db.Model    `db:"People"`
//   Name string `db:"id"`
// }
//
// The following struct declares an id tag on a field
// of type int, thus its IntID.
//
// type BankAccount struct {
//   db.Model   `db:"Accounts"`
//   Number int `db:"id"`
// }
//
// If multiple id tags are used on a struct fields
// only the first tag from top to bottom is considered
func ExtractEntityKeyIDs(e entity) (string, int64, error) {
	elem := reflect.TypeOf(e).Elem()
	elemValue := reflect.ValueOf(e).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		tag := field.Tag.Get("db")
		value := elemValue.Field(i)
		if tag == "id" {
			switch field.Type.Kind() {
			case reflect.String:
				v := value.String()
				if v == "" {
					return "", 0, ErrMissingStringId
				}
				return value.String(), 0, nil
			case reflect.Int,
				reflect.Int8,
				reflect.Int16,
				reflect.Int32,
				reflect.Int64:
				v := value.Int()
				if v == 0 {
					return "", 0, ErrMissingIntId
				}
				return "", v, nil
			}
		}
	}

	// Default key values for auto generated keys
	return "", 0, nil
}

// ExtractEntityKind extracts entity kind from struct tag
// applied to db.Model field
//
// e.g.:
//
// type Person struct {
//   db.Model `db:"People"`
//   Name     string
// }
//
func ExtractEntityKind(e entity) string {
	elem := reflect.TypeOf(e).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.Type.Name() == reflect.TypeOf(Model{}).Name() {
			if kind := field.Tag.Get("db"); kind != "" {
				return kind
			}
		}
	}

	return elem.Name()
}
