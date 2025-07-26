package goenum

import (
	"database/sql/driver"
)

type StringNamedMeta[VMap StringValueNamedMetaMap[VMeta], VMeta any] string

type StringValueNamedMetaMap[VMeta any] interface {
	ValueMapper
	NameMap() map[string]string
	MetaMap() map[string]VMeta
}

func (e StringNamedMeta[VMap, VMeta]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e StringNamedMeta[VMap, VMeta]) Key() string {
	return string(e)
}

func (e StringNamedMeta[VMap, VMeta]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e StringNamedMeta[VMap, VMeta]) String() string {
	return e.Key()
}

func (e StringNamedMeta[VMap, VMeta]) Name() string {
	l := newMap[VMap]().NameMap()
	if name, ok := l[string(e)]; ok {
		return name
	}
	return e.Key()
}

func (e StringNamedMeta[VMap, VMeta]) Meta() *VMeta {
	l := newMap[VMap]().MetaMap()
	if meta, ok := l[string(e)]; ok {
		return &meta
	}
	return nil
}

func (e *StringNamedMeta[VMap, VMeta]) Scan(value any) error {
	t, err := scanValueMapper[StringNamedMeta[VMap, VMeta]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e StringNamedMeta[VMap, VMeta]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}
