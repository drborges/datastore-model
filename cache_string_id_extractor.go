package db

import (
	"reflect"
)

type CacheStringIdExtractor struct {
	Metadata *Metadata
}

func (this CacheStringIdExtractor) Accept(f reflect.StructField) bool {
	return f.Tag.Get("cache") == "id" && f.Type.Kind() == reflect.String
}

func (this CacheStringIdExtractor) Extract(e Entity, f reflect.StructField, v reflect.Value) error {
	value := v.String()
	if value == "" {
		return ErrMissingCacheStringId
	}
	this.Metadata.CacheStringID = value
	return nil
}
