package db

import (
	"appengine"
	"reflect"
)

type Querier struct {
	c appengine.Context
	q *Query
}

func (this Querier) All(slice interface{}) error {
	keys, err := this.q.GetAll(this.c, slice)

	if err != nil {
		return err
	}

	for i, key := range keys {
		EntityAt(slice, i).SetKey(key)
	}

	return nil
}

func (this Querier) First(entity entity) error {
	i := this.q.Run(this.c)
	key, err := i.Next(entity)
	if err != nil {
		return err
	}

	entity.SetKey(key)
	return nil
}

func EntityAt(slice interface{}, i int) entity {
	s := reflect.ValueOf(slice)

	if s.Kind() == reflect.Slice {
		return s.Index(i).Interface().(entity)
	}

	if s.Type().Elem().Kind() == reflect.Slice {
		return s.Elem().Index(i).Interface().(entity)
	}

	panic("Querier.toEntitySlice given a non-slice type")
}
