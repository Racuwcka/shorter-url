package config

type StorageType string

const (
	Memory   StorageType = "memory"
	Postgres StorageType = "postgres"
)

func (s StorageType) IsMemory() bool {
	return s == Memory
}

func (s StorageType) IsPostgres() bool {
	return s == Postgres
}

func (s StorageType) IsValid() bool {
	switch s {
	case Memory, Postgres:
		return true
	default:
		return false
	}
}
