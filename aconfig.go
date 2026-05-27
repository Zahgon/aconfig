package aconfig

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

// Loader of user configuration.
type Loader struct {
	config  Config
	dst     any
	parser  *structParser
	fields  []*fieldData
	fsys    fs.FS
	flagSet *flag.FlagSet
	errInit error
}

// Config to configure configuration loader.
type Config struct {
	// NewParser set to true enables a new and better struct parser.
	// Default is false because there might be bugs.
	// In the future new parser will be enabled by default.
	NewParser bool

	SkipDefaults bool // SkipDefaults set to true will not load config from 'default' tag.
	SkipFiles    bool // SkipFiles set to true will not load config from files.
	SkipEnv      bool // SkipEnv set to true will not load config from environment variables.
	SkipFlags    bool // SkipFlags set to true will not load config from flag parameters.

	EnvPrefix  string // EnvPrefix for environment variables.
	FlagPrefix string // FlagPrefix for flag parameters.

	// envDelimiter for environment variables. Is always "_" due to env-var format.
	// Also unexported cause there is no sense to change it.
	envDelimiter string

	FlagDelimiter string // FlagDelimiter for flag parameters. If not set - default is ".".

	// AllFieldsRequired set to true will fail config loading if one of the fields was not set.
	// File, environment, flag must provide a value for the field.
	// If default is set and this option is enabled (or required tag is set) there will be an error.
	AllFieldRequired bool

	// AllowDuplicates set to true will not fail on duplicated names on fields (env, flag, etc...)
	AllowDuplicates bool

	// AllowUnknownFields set to true will not fail on unknown fields in files.
	AllowUnknownFields bool

	// AllowUnknownEnvs set to true will not fail on unknown environment variables ().
	// When false error is returned only when EnvPrefix isn't empty.
	AllowUnknownEnvs bool

	// AllowUnknownFlags set to true will not fail on unknown flag parameters ().
	// When false error is returned only when FlagPrefix isn't empty.
	AllowUnknownFlags bool

	// DontGenerateTags disables tag generation for JSON, YAML, TOML file formats.
	DontGenerateTags bool

	// FailOnFileNotFound will stop Loader on a first not found file from Files field in this structure.
	FailOnFileNotFound bool

	// FileSystem from which files will be loaded. Default is nil (OS file system).
	FileSystem fs.FS

	// MergeFiles set to true will collect all the entries from all the given files.
	// Easy wat to cobine base.yaml with prod.yaml
	MergeFiles bool

	// FileFlag the name of the flag that defines the path to the configuration file passed through the CLI.
	// (To make it easier to transfer the config file via flags.)
	FileFlag string

	// Files from which config should be loaded.
	Files []string

	// Envs hold the environment variable from which envs will be parsed.
	// By default is nil and then os.Environ() will be used.
	Envs []string

	// Args hold the command-line arguments from which flags will be parsed.
	// By default is nil and then os.Args will be used.
	// Unless loader.Flags() will be explicitly parsed by the user.
	Args []string

	// FileDecoders to enable other than JSON file formats and prevent additional dependencies.
	// Add required submodules to the go.mod and register them in this field.
	// Example:
	//	FileDecoders: map[string]aconfig.FileDecoder{
	//		".yaml": aconfigyaml.New(),
	//		".toml": aconfigtoml.New(),
	//		".env": aconfigdotenv.New(),
	// 	}
	FileDecoders map[string]FileDecoder

	// SliceSeparator hold the separator for slice values. Default is ",".
	SliceSeparator string
}

// FileDecoder is used to read config from files. See aconfig submodules.
type FileDecoder interface {
	Format() string
	DecodeFile(filename string) (map[string]any, error)
	// Init(fsys fs.FS)
}

// Field of the user configuration structure.
// Done as an interface to export less things in lib.
type Field interface {
	// Name of the field.
	Name() string

	// Tag returns a given tag for a field.
	Tag(tag string) string

	// Parent of the current node.
	Parent() (Field, bool)
}

// LoaderFor creates a new Loader based on a given configuration structure.
// Supports only non-nil structures.
func LoaderFor(dst any, cfg Config) *Loader { _ = "STUB: not implemented"; return nil }

func (l *Loader) init() {
	l.config.envDelimiter = "_"

	if l.config.FlagDelimiter == "" {
		l.config.FlagDelimiter = "."
	}

	if l.config.EnvPrefix != "" {
		l.config.EnvPrefix += l.config.envDelimiter
	}
	if l.config.FlagPrefix != "" {
		l.config.FlagPrefix += l.config.FlagDelimiter
	}

	l.fsys = &fsOrOS{l.config.FileSystem}

	if _, ok := l.config.FileDecoders[".json"]; !ok {
		if l.config.FileDecoders == nil {
			l.config.FileDecoders = map[string]FileDecoder{}
		}
		l.config.FileDecoders[".json"] = &jsonDecoder{}
	}
	for _, dec := range l.config.FileDecoders {
		dec, ok := dec.(interface{ Init(fs.FS) })
		if !ok {
			continue
		}
		dec.Init(l.fsys)
	}

	if l.config.Envs == nil {
		l.config.Envs = os.Environ()
	}
	if l.config.Args == nil {
		l.config.Args = os.Args[1:]
	}

	if l.config.NewParser {
		l.parser = newStructParser(l.config)
		if err := l.parser.parseStruct(l.dst); err != nil {
			l.errInit = err
			return
		}
	} else {
		l.fields = l.getFields(l.dst)
	}

	l.flagSet = flag.NewFlagSet(l.config.FlagPrefix, flag.ContinueOnError)
	if !l.config.SkipFlags {
		names := make(map[string]bool, len(l.fields))
		if l.config.NewParser {
			l.flagSet = l.parser.flagSet
		} else {
			for _, field := range l.fields {
				flagName := l.fullTag(l.config.FlagPrefix, field, "flag")
				if flagName == "" {
					continue
				}
				if names[flagName] && !l.config.AllowDuplicates {
					l.errInit = fmt.Errorf("duplicate flag %q", flagName)
					return
				}
				names[flagName] = true
				l.flagSet.String(flagName, field.Tag("default"), field.Tag("usage"))
			}
		}
	}

	if l.config.FileFlag != "" {
		// TODO: should be prefixed ?
		l.flagSet.String(l.config.FileFlag, "", "config file param")
	}

	if l.config.SliceSeparator == "" {
		l.config.SliceSeparator = ","
	}
}

// Flags returngs flag.FlagSet to create your own flags.
// FlagSet name is Config.FlagPrefix and error handling is set to ContinueOnError.
func (l *Loader) Flags() *flag.FlagSet {
	_ = "STUB: not implemented"

	// WalkFields iterates over configuration fields.
	// Easy way to create documentation or user-friendly help.
	return nil
}

func (l *Loader) WalkFields(fn func(f Field) bool) { _ = "STUB: not implemented"; return }

// Load configuration into a given param.
func (l *Loader) Load() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadConfig() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) parseFlags() error {
	_ = "STUB: not implemented"
	// TODO: too simple?
	return nil
}

func (l *Loader) loadSources() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) checkRequired() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadDefaults() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadFiles() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadFile(file string) error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadFileFlag() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) loadEnvironment() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) postEnvCheck(values map[string]any, dupls map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (l *Loader) loadFlags() error { _ = "STUB: not implemented"; return nil }

func (l *Loader) postFlagCheck(values map[string]any, dupls map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(cristaloleg): revisit.
func (l *Loader) setField(field *fieldData, name string, values map[string]any, dupls map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}
