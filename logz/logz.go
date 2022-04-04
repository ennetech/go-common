package logz

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

var DebugOn = false

// fmt.Sprintf(Blue+"[%s] %s"+Reset+"\n", tag, content)

func line(color string, tag string, content string) {
	fmt.Println(color + "[" + tag + "] " + content + reset)
}

func Debug(tag string, content string) {
	if DebugOn {
		line(purple, tag, content)
	}
}

func DebugHex(tag string, content []byte) {
	if DebugOn {
		line(purple, tag, hex.EncodeToString(content))
	}
}

func DebugHexString(tag string, content []byte) {
	if DebugOn {
		Debug("STR: "+tag, strings.Trim(string(content), " \n\r\t"))
		DebugHex("HEX: "+tag, content)
	}
}

func Info(tag string, content string) {
	line(blue, tag, content)
}

func Warn(tag string, content string) {
	line(yellow, tag, content)
}

func Error(tag string, content string) {
	line(red, tag, content)
}

func Start(tag string, content string) {
	line(green, tag, content)
}

func End(tag string, content string) {
	line(gray, tag, content)
}

func Die(err error, what string, ok string) {
	if err != nil {
		Error("FATAL - "+what, err.Error())
		os.Exit(1)
	} else {
		Debug(what, ok)
	}
}
