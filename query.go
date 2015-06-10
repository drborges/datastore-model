package db

import (
	"appengine/datastore"
)

func From(e Entity) *Query {
	metadata := &Metadata{}
	MetadataExtractorChain{KindExtractor{metadata}}.ExtractFrom(e)
	return &Query{datastore.NewQuery(metadata.Kind)}
}

type Query struct {
	*datastore.Query
}

func (this *Query) Filter(filter string, value interface{}) *Query {
	this.Query = this.Query.Filter(filter, value)
	return this
}
