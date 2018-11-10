package main

import (
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"add": Add,
}

func Add(l *lua.LState) int {
	x := l.ToInt(1)
	y := l.ToInt(2)

	l.Push(lua.LNumber(x + y))
	return 1
}
func main() {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("mathmodule", Loader)
	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}
}
