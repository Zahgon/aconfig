package aconfighcl

import (
	"io/fs"
)

// Decoder of HCL files for aconfig.
type Decoder struct {
	fsys fs.FS
}

// New HCL decoder for aconfig.
func New() *Decoder {
	_ = "STUB: not implemented"

	// Format of the decoder.
	return nil
}

func (d *Decoder) Format() string {
	_ = "STUB: not implemented"

	// DecodeFile implements aconfig.FileDecoder.
	return ""
}

func (d *Decoder) DecodeFile(filename string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeFile implements aconfig.FileDecoder.
func (d *Decoder) Init(fsys fs.FS) { _ = "STUB: not implemented"; return }
