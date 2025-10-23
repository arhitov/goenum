package goenum

import (
	"database/sql/driver"
)

type StringMetaSorted[VMap StringValueMetaSortedMap[VMeta], VMeta any] string

type StringValueMetaSortedMap[VMeta any] interface {
	ValueMapper
	MetaMap() map[string]VMeta
	Sorted() []string
}

func (e StringMetaSorted[VMap, VMeta]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e StringMetaSorted[VMap, VMeta]) Key() string {
	return string(e)
}

func (e StringMetaSorted[VMap, VMeta]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e StringMetaSorted[VMap, VMeta]) String() string {
	return e.Key()
}

func (e StringMetaSorted[VMap, VMeta]) Title() string {
	switch title := any(e.Meta()).(type) {
	case interface{ Title() string }:
		return title.Title()
	default:
		return e.String()
	}
}

func (e StringMetaSorted[VMap, VMeta]) Meta() *VMeta {
	l := newMap[VMap]().MetaMap()
	if meta, ok := l[string(e)]; ok {
		return &meta
	}
	return nil
}

func (e *StringMetaSorted[VMap, VMeta]) Scan(value any) error {
	t, err := scanValueMapper[StringMetaSorted[VMap, VMeta]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e StringMetaSorted[VMap, VMeta]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}

func (e StringMetaSorted[VMap, VMeta]) Sorted() []string {
	return newMap[VMap]().Sorted()
}
