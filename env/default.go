package env

import (
	"os"
)

func SetDefaultEnv() {
	os.Setenv("ENV", "Default")

}
