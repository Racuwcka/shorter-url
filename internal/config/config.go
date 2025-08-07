package config

import (
	"flag"
	"log"
	"time"

	"github.com/Racuwcka/shorter-url/pkg/config"
)

type Config struct {
	Env         `env-default:"local"`
	StorageType `env-default:"memory"`
	HTTPServer
	Cache
}

type HTTPServer struct {
	Host string
	Port string
}

type Cache struct {
	Capacity        int           `env-default:"1000"`
	ShutdownTimeout time.Duration `env-default:"5"`
}

const fileEnv = ".env"

func MustLoadConfig() *Config {
	envConfig, err := config.LoadEnvFile(fileEnv)
	if err != nil {
		log.Fatal(err)
	}

	defaultEnv := envConfig.GetEnvString("env", "local")
	defaultStorageType := envConfig.GetEnvString("storage_type", "memory")

	defaultHost := envConfig.GetEnvString("host", "localhost")
	defaultPort := envConfig.GetEnvString("port", "8080")
	defaultShutdownTimeout := envConfig.GetEnvInt("shutdown_timeout", 5)
	defaultCapacityCache := envConfig.GetEnvInt("capacity_cache", 1000)

	envFlag := flag.String("env", defaultEnv, "the app env")
	StorageTypeFlag := flag.String("storage_type", defaultStorageType, "the storage type")

	host := flag.String("host", defaultHost, "the host address to connect to")
	port := flag.String("port", defaultPort, "the port")
	shutdownTimeout := flag.Int("shutdown_timeout", defaultShutdownTimeout, "shutdownTimeout time")
	capacityCache := flag.Int("capacity_cache", defaultCapacityCache, "capacity of cache")

	flag.Parse()

	env := Env(*envFlag)
	if !env.IsValid() {
		log.Fatalf("invalid env value: %s", env)
	}

	storageType := StorageType(*StorageTypeFlag)
	if !storageType.IsValid() {
		log.Fatalf("invalid storage type value: %s", storageType)
	}

	return &Config{
		Env:         env,
		StorageType: storageType,
		HTTPServer: HTTPServer{
			Host: *host,
			Port: *port,
		},
		Cache: Cache{
			ShutdownTimeout: time.Duration(*shutdownTimeout) * time.Second,
			Capacity:        *capacityCache,
		},
	}
}
