package logging

import (
	"github.com/natefinch/lumberjack"
	"io"
	"log/slog"
	"os"
)

// Config general configuration struct
type Config struct {
	Directory string `yaml:"directory"`
	Filename  string `yaml:"filename"`
}

// New constructor for configuration
func New(conf Config) *slog.Logger {

	lw := lumberjack.Logger{
		Filename: conf.Filename,
		MaxSize:  20,
		MaxAge:   365,
	}

	mw := io.MultiWriter(&lw, os.Stdout)

	l := slog.New(
		slog.NewTextHandler(mw, nil),
	)

	return l
}
