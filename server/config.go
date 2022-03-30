package server

import (
	"log"

	"github.com/BurntSushi/toml"
)

type GrpcServerConfig struct {
	Addr    string `toml:"addr"`
	UseTLS  bool   `toml:"use_tls"`
	CrtFile string `toml:"crt_file"`
	KeyFile string `toml:"key_file"`
}

type HttpServerConfig struct {
	Addr string `toml:"addr"`
}

type TomlConfig struct {
	GRPCServer GrpcServerConfig `toml:"grpc_server"`
	HTTPServer HttpServerConfig `toml:"http_server"`
}

var (
	config TomlConfig
)

func LoadConfig() {
	var lConfig TomlConfig
	if _, err := toml.DecodeFile("server/config.toml", &lConfig); err != nil {
		log.Fatalln("Error decoding config file", err)
	}

	config = lConfig
}
