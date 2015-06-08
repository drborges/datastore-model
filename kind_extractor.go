package db

import (
	"reflect"
	"strings"
)

type KindExtractor struct {
	Entity entity
	Metadata *Metadata
}

func (this KindExtractor) Accept(f reflect.StructField) bool {
	return f.Type.Name() == reflect.TypeOf(Model{}).Name()
}

func (this KindExtractor) Extract(f reflect.StructField) error {
	elem := reflect.TypeOf(this.Entity).Elem()
	this.Metadata.Kind = elem.Name()

	kindMetadata := f.Tag.Get("db")
	values := strings.Split(kindMetadata, ",")
	if strings.TrimSpace(values[0]) != "" {
		this.Metadata.Kind = values[0]
	}

	return nil
}