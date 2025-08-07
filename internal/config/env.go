package config

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)

func (e Env) IsLocal() bool {
	return e == EnvLocal
}

func (e Env) IsDev() bool {
	return e == EnvDev
}

func (e Env) IsProd() bool {
	return e == EnvProd
}

func (e Env) IsValid() bool {
	switch e {
	case EnvLocal, EnvDev, EnvProd:
		return true
	default:
		return false
	}
}
