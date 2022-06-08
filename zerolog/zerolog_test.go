package zerolog_test

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestContextualLogger(t *testing.T) {
	log := zerolog.New(os.Stdout)
	log.Info().Str("content", "hello world").Int("count", 3).Msg("TestContextualLogger")

	// 添加上下wen
	log = log.With().Caller().Str("foo", "bar").Logger()
	log.Info().Msg("hello world")
}

func TestZerologLevel(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log := zerolog.New(os.Stdout)

	log.Trace().Msg("Trace")
	log.Debug().Msg("Debug")
	log.Info().Msg("Info")
	log.Warn().Msg("Warn")
	log.Error().Msg("Error")
	log.Log().Msg("Log")
}
