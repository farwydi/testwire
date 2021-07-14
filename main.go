package main

import "testwire/dom"

func main() {

}

func newApp(rep dom.Rep) app {
	return app{rep: rep}
}

type app struct {
	rep dom.Rep
}
