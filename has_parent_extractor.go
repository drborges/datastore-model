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

func (this HasParentExtractor) Extract(e Entity, f reflect.StructField, v reflect.Value) error {
	metadata := f.Tag.Get("db")
	values := strings.Split(metadata, ",")
	for _, value := range values {
		if strings.TrimSpace(value) == "has_parent" {
			this.Metadata.Parent = e.Parent()
			this.Metadata.HasParent = true
		}
	}

	if this.Metadata.HasParent && e.Parent() == nil {
		return ErrMissingParentKey
	}

	return nil
}
