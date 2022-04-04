package env

import (
	"github.com/ennetech/go-common/logz"
	"os"
)

func Get(params ...string) string {
	val := os.Getenv(params[0])
	if val == "" && len(params) == 1 {
		logz.Debug("ENV", "Environment variable "+params[0]+" is missing")
		os.Exit(1)
	} else if val == "" && len(params) == 2 {
		logz.Debug("ENV", "Environment variable "+params[0]+" is missing, using default: "+params[1])
		return params[1]
	} else {
		logz.Debug("ENV", "Environment variable "+params[0]+" is present: "+val)
		return val
	}
	return ""
}
