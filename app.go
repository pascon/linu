package main

import (
	"github.com/progrium/darwinkit/cocoa"
	"github.com/progrium/darwinkit/objc"
)

var appCtx AppCtx

func applicationCallback(n objc.Object) {

	appCtx = AppCtx{
		obj:              cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength),
		refreshActionSig: make(chan bool),
		portActionSig:    make(chan string),
		baudRateSig:      make(chan string),
		mainMenu:         cocoa.NSMenu_New(),
	}
	appCtx.initialize()
	go applicationControlPanel(appCtx)
	appCtx.refreshActionSig <- true
}

func createCocoaApplication() cocoa.NSApplication {
	app := cocoa.NSApp_WithDidLaunch(applicationCallback)
	return app
}
