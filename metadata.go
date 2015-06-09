package db

import (
	"appengine/datastore"
	"reflect"
)

type Metadata struct {
	Kind      string
	StringID string
	IntID    int64
	HasParent bool
	Parent   *datastore.Key
}

type MetadataExtractor interface {
	Accept(reflect.StructField) bool
	Extract(entity, reflect.StructField) error
}