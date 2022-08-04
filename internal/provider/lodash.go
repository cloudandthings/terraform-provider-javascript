package provider

import (
	"github.com/dop251/goja"
)

func installLodash(js *goja.Runtime) {
	lodash, err := Asset("js/lodash-4.17.21.min.js")
	if err != nil {
		panic(err)
	}
	var lodashProgram = goja.MustCompile("lodash.min.js", string(lodash), false)
	js.RunProgram(lodashProgram)
}
