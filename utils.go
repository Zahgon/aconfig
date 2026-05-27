package aconfig

import (
	"flag"
	"io/fs"
	"reflect"
)

func assertStruct(x interface{}) { _ = "STUB: not implemented"; return }

func getEnv(env []string) map[string]interface{} { _ = "STUB: not implemented"; return nil }

func getFlags(flagSet *flag.FlagSet) map[string]interface{} { _ = "STUB: not implemented"; return nil }

func getActualFlag(name string, flagSet *flag.FlagSet) *flag.Flag {
	_ = "STUB: not implemented"
	return nil
}

func makeName(name string, parent *fieldData) string { _ = "STUB: not implemented"; return "" }

func (l *Loader) makeTagValue(field reflect.StructField, tag string, words []string) string {
	_ = "STUB: not implemented"
	return ""
}

// based on https://github.com/fatih/camelcase
func splitNameByWords(src string) []string { _ = "STUB: not implemented"; return nil }

// split into fields based on class of unicode character

// handle upper case -> lower case sequences, e.g.
// "PDFL", "oader" -> "PDF", "Loader"

// copy-paste until https://github.com/golang/go/issues/46336 is fixed
// returns: before, after, isFound
func cut(s, sep string) (_, _ string, _ bool) { _ = "STUB: not implemented"; return "", "", false }

var _ fs.FS = &fsOrOS{}

type fsOrOS struct{ fs.FS }

func (f *fsOrOS) Open(name string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

type jsonDecoder struct {
	fsys fs.FS
}

func (d *jsonDecoder) Init(fsys fs.FS) {
	_ = "STUB: not implemented"

	// Format of the decoder.
	return
}

func (d *jsonDecoder) Format() string {
	_ = "STUB: not implemented"

	// DecodeFile implements FileDecoder.
	return ""
}

func (d *jsonDecoder) DecodeFile(filename string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *Loader) sliceToString(curr interface{}) string { _ = "STUB: not implemented"; return "" }

func find(actualFields map[string]interface{}, name string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
