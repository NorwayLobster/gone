package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func zerolog_demo() {
	// no_log_level()
	// test_demo1()
	// local_logger()
	// sublogger()
	// colorful_log()
	// colorful_log_1()
	// simplify()
	// set_caller()
	hook_demo()
}

func debug_demo() {
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Debug().Msg("This message appears only when log level set to debug")
	log.Info().Msg("This message appears when log level set to debug or info")

	if e := log.Debug(); e.Enabled() {
		e.Str("foo", "bar").Msg("some debug message")
	}
}

func test_demo1() {
	log.Print("hello world")
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()

	log.Info().
		Dict("dict", zerolog.Dict().
			Str("bar", "baz").
			Int("n", 1),
		).Msg("hello world")
}

func no_log_level() {
	log.Log().
		Str("foo", "bar").
		Msg("")
}

func local_logger() {
	logger := zerolog.New(os.Stderr)
	logger.Info().
		Str("bar", "baz").
		Str("foo", "bar").
		Msg("hello world")
}

func sublogger() {
	logger := zerolog.New(os.Stderr)
	sublogger := logger.With().
		Str("foo", "bar").
		Logger()
	sublogger.Info().Msg("hello world")
}

//note: ConsoleWriter is for dev not for prod
func colorful_log() {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	logger.Info().
		Str("bar", "baz").
		Str("foo", "bar").
		Msg("hello world")
}

func colorful_log_1() {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	logger := log.Output(output).With().Timestamp().Logger()
	logger.Info().Str("foo", "bar").Msg("hello world")
}

func simplify() {
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "m"

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("hello world")
}

func set_caller() {
	logger := zerolog.New(os.Stderr).With().Caller().Logger()
	logger.Info().Msg("hello world")
}

type AddFieldHook struct {
}

func (AddFieldHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.DebugLevel {
		e.Str("name", "dj")
	}
}

func hook_demo() {
	hooked := log.Hook(AddFieldHook{})
	hooked.Debug().Msg("")
}
