package worker

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

type SchedulerConfig struct {
	Addr string `toml:"addr"`
}

type TomlConfig struct {
	GRPCServer GrpcServerConfig `toml:"grpc_server"`
	Scheduler  SchedulerConfig  `toml:"scheduler"`
}

var (
	config TomlConfig
)

func LoadConfig() {
	var localConfig TomlConfig
	if _, err := toml.DecodeFile("worker/config.toml", &localConfig); err != nil {
		log.Fatalln("Error decoding config file", err)
	}

	config = localConfig
}
