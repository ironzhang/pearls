package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type jsoncfg struct{}

func (jsoncfg) LoadFromFile(filename string, cfg interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return err
	}
	return nil
}

func (jsoncfg) WriteToFile(filename string, cfg interface{}) error {
	data, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}

type tomlcfg struct{}

func (tomlcfg) LoadFromFile(filename string, cfg interface{}) error {
	_, err := toml.DecodeFile(filename, cfg)
	return err
}

func (tomlcfg) WriteToFile(filename string, cfg interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	enc.Indent = "\t"
	return enc.Encode(cfg)
}

var JSON jsoncfg
var TOML tomlcfg

func LoadFromFile(filename string, cfg interface{}) error {
	return JSON.LoadFromFile(filename, cfg)
}

func WriteToFile(filename string, cfg interface{}) error {
	return JSON.WriteToFile(filename, cfg)
}
