package goenum

import (
	"database/sql/driver"
)

type StringNamed[VMap StringValueNamedMap] string

type StringValueNamedMap interface {
	ValueMapper
	NameMap() map[string]string
}

func (e StringNamed[VMap]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e StringNamed[VMap]) Key() string {
	return string(e)
}

func (e StringNamed[VMap]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e StringNamed[VMap]) String() string {
	return e.Key()
}

func (e StringNamed[VMap]) Name() string {
	l := newMap[VMap]().NameMap()
	if name, ok := l[string(e)]; ok {
		return name
	}
	return e.Key()
}

func (e *StringNamed[VMap]) Scan(value any) error {
	t, err := scanValueMapper[StringNamed[VMap]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e StringNamed[VMap]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}
