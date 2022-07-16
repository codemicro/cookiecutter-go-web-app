package config

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func InitLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

var HTTP = struct {
	Host   string
	Port   int
	Secure bool
}{
	Host:   asString(withDefault(fetchFromFile("http.host"), "0.0.0.0")),
	Port:   asInt(withDefault(fetchFromFile("http.port"), 8080)),
	Secure: asBool(required(fetchFromFile("http.secure"))),
}
