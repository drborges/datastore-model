package db

import "appengine/datastore"

func QueryFor(e entity) *Query {
	kind, _ := ExtractEntityKindMetadata(e)
	return &Query{datastore.NewQuery(kind)}
}

type Query struct {
	*datastore.Query
}

func (this *Query) Filter(filter string, value interface{}) *Query {
	this.Query = this.Query.Filter(filter, value)
	return this
}
