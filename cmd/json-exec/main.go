package main

import (
	"fmt"
	"os"

	"go.sophtrust.dev/json-exec/internal/cli"
	"go.sophtrust.dev/pkg/zerolog/v2"
	"go.sophtrust.dev/pkg/zerolog/v2/log"
)

func main() {
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05.000Z07:00"
	stdoutLevelWriter := zerolog.NewFilteredLevelWriter([]zerolog.Level{
		zerolog.DebugLevel, zerolog.InfoLevel, zerolog.WarnLevel,
	}, os.Stdout)
	stderrLevelWriter := zerolog.NewFilteredLevelWriter([]zerolog.Level{
		zerolog.ErrorLevel, zerolog.FatalLevel, zerolog.PanicLevel,
	}, os.Stderr)
	writer := zerolog.MultiLevelWriter(stdoutLevelWriter, stderrLevelWriter)
	l := zerolog.New(writer).With().Timestamp().Logger()
	l.SetLevel(zerolog.InfoLevel)
	log.ReplaceGlobal(l)

	cmd := cli.NewRootCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	os.Exit(cmd.GetExitCode())
}
