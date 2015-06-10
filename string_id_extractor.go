package db

import (
"reflect"
)

type StringIdExtractor struct {
	Metadata *Metadata
}

func (this StringIdExtractor) Accept(f reflect.StructField) bool {
	return f.Tag.Get("db") != "" && f.Type.Kind() == reflect.String
}

func (this StringIdExtractor) Extract(e Entity, f reflect.StructField, v reflect.Value) error {
	value := v.String()
	if value == "" {
		return ErrMissingStringId
	}
	this.Metadata.StringID = value
	return nil
}
