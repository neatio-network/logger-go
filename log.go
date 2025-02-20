package logger

import (
	"os"

	. "github.com/neatio-net/common-go"
	"github.com/neatio-net/log15-go"
)

var mainHandler log15.Handler
var bypassHandler log15.Handler

func init() {
	resetWithLogLevel("debug")
}

func SetLogLevel(logLevel string) {
	resetWithLogLevel(logLevel)
}

func resetWithLogLevel(logLevel string) {	
	mainHandler = log15.LvlFilterHandler(
		getLevel(logLevel),
		log15.StreamHandler(os.Stdout, log15.TerminalFormat()),
	)
	bypassHandler = log15.StreamHandler(os.Stdout, log15.TerminalFormat())
	log15.Root().SetHandler(mainHandler)
}

func MainHandler() log15.Handler {
	return mainHandler
}

func BypassHandler() log15.Handler {
	return bypassHandler
}

func New(ctx ...interface{}) log15.Logger {
	return NewMain(ctx...)
}

func NewMain(ctx ...interface{}) log15.Logger {
	return log15.Root().New(ctx...)
}

func NewBypass(ctx ...interface{}) log15.Logger {
	bypass := log15.New(ctx...)
	bypass.SetHandler(bypassHandler)
	return bypass
}

func getLevel(lvlString string) log15.Lvl {
	lvl, err := log15.LvlFromString(lvlString)
	if err != nil {
		Exit(Fmt("Invalid log level %v: %v", lvlString, err))
	}
	return lvl
}

var LvlFilterHandler = log15.LvlFilterHandler
var LvlDebug = log15.LvlDebug
var LvlInfo = log15.LvlInfo
var LvlNotice = log15.LvlNotice
var LvlWarn = log15.LvlWarn
var LvlError = log15.LvlError
