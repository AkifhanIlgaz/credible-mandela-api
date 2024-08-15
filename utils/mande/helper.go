package mande

const (
	MinCredToRegister float64 = 10.0
)

func IsEnoughCredToRegister(cred float64) bool {
	return cred >= MinCredToRegister
}
