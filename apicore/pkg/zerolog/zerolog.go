package zerolog

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// panic (zerolog.PanicLevel, 5)
// fatal (zerolog.FatalLevel, 4)
// error (zerolog.ErrorLevel, 3)
// warn (zerolog.WarnLevel, 2)
// info (zerolog.InfoLevel, 1)
// debug (zerolog.DebugLevel, 0)
// trace (zerolog.TraceLevel, -1)

var (
	LOG_HOST    = os.Getenv("LOG_HOST")
	LOG_REGION  = os.Getenv("LOG_REGION")
	LOG_VERSION = os.Getenv("LOG_VERSION")
	LOG_LOGGER  = os.Getenv("LOG_LOGGER")
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
}
