package env

import "os"

//GetEnv : get env value
func GetEnv(key string) string {
	return os.Getenv(key)
}
