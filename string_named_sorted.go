package goenum

import (
	"database/sql/driver"
)

type StringNamedSorted[VMap StringValueNamedSortedMap] string

type StringValueNamedSortedMap interface {
	ValueMapper
	NameMap() map[string]string
	Sorted() []string
}

func (e StringNamedSorted[VMap]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e StringNamedSorted[VMap]) Key() string {
	return string(e)
}

func (e StringNamedSorted[VMap]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e StringNamedSorted[VMap]) String() string {
	return e.Key()
}

func (e StringNamedSorted[VMap]) Name() string {
	l := newMap[VMap]().NameMap()
	if name, ok := l[string(e)]; ok {
		return name
	}
	return e.Key()
}

func (e *StringNamedSorted[VMap]) Scan(value any) error {
	t, err := scanValueMapper[StringNamedSorted[VMap]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e StringNamedSorted[VMap]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}

func (e StringNamedSorted[VMap]) Sorted() []string {
	return newMap[VMap]().Sorted()
}
