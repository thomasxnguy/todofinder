package app

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"errors"
	"fmt"
)

// Logger defines the logger interface that is exposed via ContextScope.
type Logger interface {
	// adds a field that should be added to every message being logged
	SetField(name, value string)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// logger wraps logrus.Logger so that it can log messages sharing a common set of fields.
type logger struct {
	fields logrus.Fields
}

// NewLogger creates a logger object with the specified logrus.Logger and the fields that should be added to every message.
func NewLogger(fields logrus.Fields) Logger {
	return &logger{
		fields: fields,
	}
}

func (l *logger) SetField(name, value string) {
	l.fields[name] = value
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.tagged().Debugf(format, args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.tagged().Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.tagged().Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.tagged().Errorf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	l.tagged().Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.tagged().Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.tagged().Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.tagged().Error(args...)
}

func (l *logger) tagged() *logrus.Entry {
	return appLogger.WithFields(l.fields)
}

// Application logger that should be initialized by server.
var appLogger *logrus.Logger

// LoadAppLogger initialise the global application logger that all
// logger depends on to write messages.
func LoadAppLogger(c *Configuration) (*logrus.Logger, error) {
	if c == nil {
		return nil, fmt.Errorf("try to initialize AppLogger without configuration")
	}
	output := c.LogOutput
	level := c.LogLevel

	logrus.Debugf("Logger settings with output: %v, level: %v", output, level)
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	log := &logrus.Logger{
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logLevel,
	}

	switch output {
	case "stderr":
		log.Out = os.Stderr
	case "stdout":
		log.Out = os.Stdout
	case "file":
		logFile := c.LogFile
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err == nil {
			log.Out = file
		} else {
			return nil, err
		}
	case "discard":
		log.Out = ioutil.Discard
	default:
		return nil, errors.New("invalid value for log output")
	}

	appLogger = log
	return appLogger, nil
}
