package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// Field ...
	Field = zapcore.Field
	// Level ...
	Level = zapcore.Level
)

// Alias for zap level.
var (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	PanicLevel  = zapcore.PanicLevel
	DPanicLevel = zapcore.DPanicLevel
	FatalLevel  = zapcore.FatalLevel
)

// Alias for zap type functions.
var (
	Any         = zap.Any
	Array       = zap.Array
	Object      = zap.Object
	Binary      = zap.Binary
	Bool        = zap.Bool
	Bools       = zap.Bools
	ByteString  = zap.ByteString
	ByteStrings = zap.ByteStrings
	Complex64   = zap.Complex64
	Complex64s  = zap.Complex64s
	Complex128  = zap.Complex128
	Complex128s = zap.Complex128s
	Duration    = zap.Duration
	Durations   = zap.Durations
	Err         = zap.Error
	Errors      = zap.Errors
	Float32     = zap.Float32
	Float32s    = zap.Float32s
	Float64     = zap.Float64
	Float64s    = zap.Float64s
	Int         = zap.Int
	Ints        = zap.Ints
	Int8        = zap.Int8
	Int8s       = zap.Int8s
	Int16       = zap.Int16
	Int16s      = zap.Int16s
	Int32       = zap.Int32
	Int32s      = zap.Int32s
	Int64       = zap.Int64
	Int64s      = zap.Int64s
	Namespace   = zap.Namespace
	Reflect     = zap.Reflect
	Stack       = zap.Stack
	String      = zap.String
	Stringer    = zap.Stringer
	Strings     = zap.Strings
	Time        = zap.Time
	Times       = zap.Times
	Uint        = zap.Uint
	Uints       = zap.Uints
	Uint8       = zap.Uint8
	Uint8s      = zap.Uint8s
	Uint16      = zap.Uint16
	Uint16s     = zap.Uint16s
	Uint32      = zap.Uint32
	Uint32s     = zap.Uint32s
	Uint64      = zap.Uint64
	Uint64s     = zap.Uint64s
	Uintptr     = zap.Uintptr
	Uintptrs    = zap.Uintptrs
)

// User can directly use package level functions
var (
	Debug  = std.Debug
	Debugf = std.Debugf
	Debugw = std.Debugw

	Info  = std.Debug
	Infof = std.Debugf
	Infow = std.Debugw

	Warn  = std.Debug
	Warnf = std.Debugf
	Warnw = std.Debugw

	Error  = std.Debug
	Errorf = std.Debugf
	Errorw = std.Debugw

	DPanic  = std.Debug
	DPanicf = std.Debugf
	DPanicw = std.Debugw

	Panic  = std.Debug
	Panicf = std.Debugf
	Panicw = std.Debugw

	Fatal  = std.Debug
	Fatalf = std.Debugf
	Fatalw = std.Debugw

	Sync = std.Sync

	WithName   = std.WithName
	WithFields = std.WithFields
)
