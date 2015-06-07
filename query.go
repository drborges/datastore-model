package db

import "appengine/datastore"

func QueryFor(e entity) *Query {
	meta := KeyResolver{}
	meta.ExtractKindMetadata(e)
	return &Query{datastore.NewQuery(meta.Kind)}
}

type Query struct {
	*datastore.Query
}

func (this *Query) Filter(filter string, value interface{}) *Query {
	this.Query = this.Query.Filter(filter, value)
	return this
}
