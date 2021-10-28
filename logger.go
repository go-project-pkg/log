package log

import (
	"go.uber.org/zap"
)

// Logger ...
type Logger struct {
	zapLogger *zap.Logger
}

// Create a new logger instance.
func newLogger(l *zap.Logger) *Logger {
	return &Logger{
		zapLogger: l,
	}
}

// WithName specify logger name
func (l *Logger) WithName(name string) *Logger {
	logger := l.zapLogger.Named(name)
	return newLogger(logger)
}

// WithFields ...
func (l *Logger) WithFields(fields ...Field) *Logger {
	logger := l.zapLogger.With(fields...)
	return newLogger(logger)
}

// Debug ...
func (l *Logger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Debugf ...
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Debugf(format, v...)
}

// Debugw ...
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info ...
func (l *Logger) Info(msg string, fields ...Field) {
	l.zapLogger.Info(msg, fields...)
}

// Infof ...
func (l *Logger) Infof(format string, v ...interface{}) {
	l.zapLogger.Sugar().Infof(format, v...)
}

// Infow ...
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn ...
func (l *Logger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Warnf ...
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Warnf(format, v...)
}

// Warnw ...
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error ...
func (l *Logger) Error(msg string, fields ...Field) {
	l.zapLogger.Error(msg, fields...)
}

// Errorf ...
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Errorf(format, v...)
}

// Errorw ...
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// DPanic ...
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.zapLogger.DPanic(msg, fields...)
}

// DPanicf ...
func (l *Logger) DPanicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().DPanicf(format, v...)
}

// DPanicw ...
func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().DPanicw(msg, keysAndValues...)
}

// Panic ...
func (l *Logger) Panic(msg string, fields ...Field) {
	l.zapLogger.Panic(msg, fields...)
}

// Panicf ...
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Panicf(format, v...)
}

// Panicw ...
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal ...
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Fatalf ...
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.zapLogger.Sugar().Fatalf(format, v...)
}

// Fatalw ...
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

// Sync memory data to log files.
func (l *Logger) Sync() {
	_ = l.zapLogger.Sync()
}
