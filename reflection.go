package aconfig

import (
	"reflect"
)

type fieldData struct {
	name       string
	parent     *fieldData
	field      reflect.StructField
	value      reflect.Value
	isSet      bool
	isRequired bool
	tags       map[string]string
}

func (f *fieldData) Name() string { _ = "STUB: not implemented"; return "" }

func (f *fieldData) Tag(tag string) string { _ = "STUB: not implemented"; return "" }

func (f *fieldData) Parent() (Field, bool) { _ = "STUB: not implemented"; return *new(Field), false }

func (l *Loader) newSimpleFieldData(value reflect.Value) *fieldData {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) newFieldData(field reflect.StructField, value reflect.Value, parent *fieldData) *fieldData {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) tagsForField(field reflect.StructField) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) fullTag(prefix string, f *fieldData, tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func (l *Loader) getFields(x interface{}) []*fieldData { _ = "STUB: not implemented"; return nil }

func (l *Loader) getFieldsHelper(valueObject reflect.Value, parent *fieldData) []*fieldData {
	_ = "STUB: not implemented"
	return nil
}

// if it's a struct - expand and process it's fields

func (l *Loader) setFieldData(field *fieldData, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// unwrap pointers

func (*Loader) setBool(field *fieldData, value string) error { _ = "STUB: not implemented"; return nil }

func (*Loader) setInt(field *fieldData, value string) error { _ = "STUB: not implemented"; return nil }

func (l *Loader) setInt64(field *fieldData, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Loader) setUint(field *fieldData, value string) error { _ = "STUB: not implemented"; return nil }

func (*Loader) setFloat(field *fieldData, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Loader) setString(field *fieldData, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Loader) setInterface(field *fieldData, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) setSlice(field *fieldData, value string) error {
	_ = "STUB: not implemented"
	// Special case for []byte
	return nil
}

func (l *Loader) setMap(field *fieldData, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) m2s(m map[string]interface{}, structValue reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// first try to find field by name

// if tag is set - use it

// if tag is not set - try to find field by name

func mii(m interface{}) map[string]interface{} { _ = "STUB: not implemented"; return nil }
