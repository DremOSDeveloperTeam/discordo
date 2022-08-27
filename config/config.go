package config

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
)

var Cfg []byte

// Config initializes a new Lua state, loads a configuration file, and defines essential micellaneous fields.
type Config struct {
	// Path is the path of the configuration file. Its value is the path to the configuration directory until Load is called.
	Path string
	// Plugin is the Go plugin. Its value is nil until Load is called.
	Plugin *plugin.Plugin
}

func NewConfig(path string) *Config {
	return &Config{
		Path: path,
	}
}

func (cfg *Config) Load() error {
	// Create the configuration directory if it does not exist already, recursively.
	err := os.MkdirAll(cfg.Path, os.ModePerm)
	if err != nil {
		return err
	}

	cfg.Path = filepath.Join(cfg.Path, "config.go")
	_, err = os.Stat(cfg.Path)
	// Create a new configuration file if it does not exist already with the read-write flag.
	if os.IsNotExist(err) {
		f, err := os.Create(cfg.Path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(Cfg)
		if err != nil {
			return err
		}

		return f.Sync()
	}

	outputPath := filepath.Join(filepath.Dir(cfg.Path), "config.so")
	// Build the configuration plugin file using "go build" command.
	err = cfg.Build(outputPath)
	if err != nil {
		return err
	}

	// Open the configuration plugin file located in the configuration directory.
	p, err := plugin.Open(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Plugin = p
	return nil
}

func (cfg *Config) Build(outputPath string) error {
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputPath, cfg.Path)
	return cmd.Run()
}

func (cfg *Config) String(name string) string {
	sym, err := cfg.Plugin.Lookup(name)
	if err != nil {
		return ""
	}

	return *sym.(*string)
}

func (cfg *Config) Bool(name string) bool {
	sym, err := cfg.Plugin.Lookup(name)
	if err != nil {
		log.Fatal(err)
	}

	return *sym.(*bool)
}

func (cfg *Config) Int(name string) int {
	sym, err := cfg.Plugin.Lookup(name)
	if err != nil {
		return 0
	}

	return *sym.(*int)
}
