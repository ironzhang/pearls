package config_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/ironzhang/pearls/config"
)

type Service struct {
	Name string
	Host string
}

type DB struct {
	Hostname string
	Username string
	Password string
}

type Sub struct {
	I8  int8
	I16 int16
	U8  uint8
	U16 uint16
}

type Misc struct {
	I32 int32
	I64 int64
	U32 uint32
	U64 uint64
	F32 float32
	F64 float64
	Sub Sub
}

type Config struct {
	Environment string
	Service     Service
	DB          DB
	Misc        Misc
}

var example = Config{
	Environment: "test",
	Service: Service{
		Name: "config",
		Host: "127.0.0.1:2000",
	},
	DB: DB{
		Hostname: "localhost:3306",
		Username: "root",
		Password: "123456",
	},
	Misc: Misc{
		I32: 32,
		I64: 64,
		U32: 32,
		U64: 64,
		F32: 3.14,
		F64: 3.1415926,
		Sub: Sub{
			I8:  8,
			I16: 16,
			U8:  8,
			U16: 16,
		},
	},
}

func TestJSONConfig(t *testing.T) {
	filename := "test.json"
	got, want := Config{}, example
	if err := config.JSON.WriteToFile(filename, want); err != nil {
		t.Fatalf("write to file: %v", err)
	}
	defer os.Remove(filename)
	if err := config.JSON.LoadFromFile(filename, &got); err != nil {
		t.Fatalf("load from file: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%+v != %+v", got, want)
	}
	//fmt.Printf("%+v\n", got)
}

func TestTOMLConfig(t *testing.T) {
	filename := "test.cfg"
	got, want := Config{}, example
	if err := config.TOML.WriteToFile(filename, want); err != nil {
		t.Fatalf("write to file: %v", err)
	}
	defer os.Remove(filename)
	if err := config.TOML.LoadFromFile(filename, &got); err != nil {
		t.Fatalf("load from file: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%+v != %+v", got, want)
	}
	//fmt.Printf("%+v\n", got)
}

func ExampleLoadFromFile() {
	type Config struct {
		A int
		B string
		C float64
	}

	if err := config.WriteToFile("example.json", &Config{A: 1, B: "hello", C: 3.14}); err != nil {
		fmt.Printf("write to file: %v", err)
		return
	}
	defer os.Remove("example.json")

	var cfg Config
	if err := config.LoadFromFile("example.json", &cfg); err != nil {
		fmt.Printf("load from file: %v", err)
		return
	}
	fmt.Println(cfg)

	// output:
	// {1 hello 3.14}
}
