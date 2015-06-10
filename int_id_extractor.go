package db

import (
"reflect"
)

type IntIdExtractor struct {
	Metadata *Metadata
}

func (this IntIdExtractor) Accept(f reflect.StructField) bool {
	if f.Tag.Get("db") == "" {
		return false
	}

	switch f.Type.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64: return true
	default: return false
	}
}

func (this IntIdExtractor) Extract(e Entity, f reflect.StructField, v reflect.Value) error {
	value := v.Int()
	if value == 0 {
		return ErrMissingIntId
	}
	this.Metadata.IntID = value
	return nil
}
