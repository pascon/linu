package main

import (
	"github.com/progrium/darwinkit/cocoa"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	cocoa.TerminateAfterWindowsClose = false
	app := createCocoaApplication()
	app.Run()
}
