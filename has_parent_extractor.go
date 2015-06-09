package db

import (
	"reflect"
	"strings"
)

type HasParentExtractor struct {
	Metadata *Metadata
}

func (this HasParentExtractor) Accept(f reflect.StructField) bool {
	return f.Type.Name() == reflect.TypeOf(Model{}).Name()
}

func (this HasParentExtractor) Extract(f reflect.StructField) error {
	metadata := f.Tag.Get("db")
	values := strings.Split(metadata, ",")
	for _, value := range values {
		if strings.TrimSpace(value) == "has_parent" {
			this.Metadata.HasParent = true
		}
	}

	return nil
}