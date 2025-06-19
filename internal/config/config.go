package config

import (
	"flag"
	"time"

	"github.com/Racuwcka/shorter-url/pkg/config"
)

type Config struct {
	Env string `env-default:"local"`
	HTTPServer
	Cache
}

type HTTPServer struct {
	Addr string `env-default:"localhost:8080"`
}

type Cache struct {
	Capacity        int           `env-default:"1000"`
	ShutdownTimeout time.Duration `env-default:"5"`
}

func LoadConfig() *Config {
	envConfig := config.LoadEnvFile(".env")

	defaultEnv := envConfig.GetEnvString("env", "local")
	defaultAddr := envConfig.GetEnvString("addr", "localhost:8080")
	defaultShutdownTimeout := envConfig.GetEnvInt("shutdownTimeout", 5)
	defaultCapacityCache := envConfig.GetEnvInt("capacityCache", 1000)

	env := flag.String("env", defaultEnv, "the app env")
	addr := flag.String("addr", defaultAddr, "the address to connect to")
	shutdownTimeout := flag.Int("shutdownTimeout", defaultShutdownTimeout, "shutdownTimeout time")
	capacityCache := flag.Int("capacityCache", defaultCapacityCache, "capacity of cache")

	flag.Parse()

	return &Config{
		Env:        *env,
		HTTPServer: HTTPServer{Addr: *addr},
		Cache: Cache{
			ShutdownTimeout: time.Duration(*shutdownTimeout) * time.Second,
			Capacity:        *capacityCache,
		},
	}
}
