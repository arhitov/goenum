package goenum

import (
	"database/sql/driver"
)

type StringMeta[VMap StringValueMetaMap[VMeta], VMeta any] string

type StringValueMetaMap[VMeta any] interface {
	ValueMapper
	MetaMap() map[string]VMeta
}

func (e StringMeta[VMap, VMeta]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e StringMeta[VMap, VMeta]) Key() string {
	return string(e)
}

func (e StringMeta[VMap, VMeta]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e StringMeta[VMap, VMeta]) String() string {
	return e.Key()
}

func (e StringMeta[VMap, VMeta]) Title() string {
	switch title := any(e.Meta()).(type) {
	case interface{ Title() string }:
		return title.Title()
	default:
		return e.String()
	}
}

func (e StringMeta[VMap, VMeta]) Meta() *VMeta {
	l := newMap[VMap]().MetaMap()
	if meta, ok := l[string(e)]; ok {
		return &meta
	}
	return nil
}

func (e *StringMeta[VMap, VMeta]) Scan(value any) error {
	t, err := scanValueMapper[StringMeta[VMap, VMeta]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e StringMeta[VMap, VMeta]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}
