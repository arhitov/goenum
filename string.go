package goenum

import (
	"database/sql/driver"
)

type String[VMap StringValueMap] string

type StringValueMap interface {
	ValueMapper
}

// Valid проверяет валидность статуса
func (e String[VMap]) Valid() bool {
	_, ok := e.ValueMap()[string(e)]
	return ok
}

func (e String[VMap]) Key() string {
	return string(e)
}

func (e String[VMap]) Value() (driver.Value, error) {
	return enumSqlValue(&e)
}

func (e String[VMap]) String() string {
	return e.Key()
}

func (e *String[VMap]) Scan(value any) error {
	t, err := scanValueMapper[String[VMap]](value, e.ValueMap())
	tPtr := &t
	*e = *tPtr
	return err
}

func (e String[VMap]) ValueMap() map[string]any {
	return newMap[VMap]().ValueMap()
}
