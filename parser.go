package aconfig

import (
	"encoding"
	"flag"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

type structParser struct {
	cfg       Config
	fields    map[string]any
	flagSet   *flag.FlagSet
	envNames  map[string]struct{}
	flagNames map[string]struct{}
}

func newStructParser(cfg Config) *structParser { _ = "STUB: not implemented"; return nil }

type parsedField struct {
	name         string
	namefull     string
	value        any
	defaultValue any
	parent       *parsedField
	childs       map[string]any
	tags         map[string]string
	hasChilds    bool
	isRequired   bool
}

func (pf *parsedField) String() string { _ = "STUB: not implemented"; return "" }

func (sp *structParser) newParseField(parent *parsedField, field reflect.StructField) (*parsedField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: must be typed?

// TODO: must be typed

func (sp *structParser) parseStruct(x any) error { _ = "STUB: not implemented"; return nil }

// fmt.Printf("fields: %+v\n", fields)

func (sp *structParser) parseStructHelper(parent *parsedField, structValue reflect.Value, res map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// do not set defaultValue for struct or pointer type without a default value
// if fieldType.Kind() == reflect.Struct ||
// 	(fieldType.Kind() == reflect.Pointer && defaultTagValue == "") {
// 	pfield.defaultValue = nil
// }

// to have 'value' of type field

// if !sp.cfg.SkipDefaults {
// 	pv := fieldValue.Addr().Interface()
// 	if v, ok := pv.(encoding.TextUnmarshaler); ok {
// 		value = defaultTagValue
// 		err := v.UnmarshalText([]byte(fmt.Sprint(value)))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	pfield.value =
// 	res[pfield.name] = pfield
// 	continue
// }

// case reflect.Array:
// TODO: same as slice + check len?

// TODO: just assign?

// fmt.Printf("field: %+v got: %+v\n\n", pfield.name, values)

// byte-slice case

// TODO: if value is struct - parse
// value = parseSlice(fieldValue, map[string]any{})

// if !sp.cfg.SkipDefaults {
// 	pfield.value = value
// }

// if isPrimitive(field.Type.Elem()) {

// fmt.Printf("parts: %+v\n", parts)

// TODO: convert entry[1] to a primitive?

// } else {
// 	pfield.hasChilds = true
// }

// TODO: do not set pointer

// skip

// TODO: when WeaklyTypedInput will be false use decodePrimitive(...)

// we should not overwrite struct because there are childs

// fmt.Printf("def: %v %T '%+v'\n", fieldType.String(), value, value)

var fieldType = reflect.TypeOf(&parsedField{})

var hook = mapstructure.DecodeHookFuncType(func(from, to reflect.Type, data any) (any, error) {
	if from != fieldType {
		// fmt.Printf("hook: got %T (%+v) when %s\n", i, i, to.String())
		return data, nil
	}
	field := data.(*parsedField)

	ifaceTo := reflect.New(to).Interface()
	if unmarshaller, ok := ifaceTo.(encoding.TextUnmarshaler); ok {
		// TODO: only string can be here?
		b := []byte(field.value.(string))
		err := unmarshaller.UnmarshalText(b)
		return unmarshaller, err
	}
	// fmt.Printf("hook: when %s do '%+v' // %+v\n\n", to.String(), field.value, field)
	return field.value, nil
})

func (sp *structParser) apply(x any) error { _ = "STUB: not implemented"; return nil }

// TODO: temp fix?

func (sp *structParser) applyLevel(tag string, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *structParser) applyLevelHelper2(fields map[string]any, tag string, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *structParser) applyLevelHelper(fields map[string]any, tag string, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// fmt.Printf("got type %T (%v)\n", v, v)

// TODO: can be only for leaf nodes?

// fmt.Printf("got val %T (%v)\n", val, val)

// fmt.Printf("got map: %+v %T\n", vval, vval)

// no struct in childs - simple apply, mapstructure will take care

// TODO: reencode values?

func (sp *structParser) applyFlat(tag string, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func (sp *structParser) applyFlatHelper(fields map[string]any, tag string, values map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

func isPrimitive(v reflect.Type) bool { _ = "STUB: not implemented"; return false }
