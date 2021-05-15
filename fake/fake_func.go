package fake

import (
	"sync"
)

var FuncFake map[string]Func

var mul sync.Mutex

type Func struct {
	TagName string
	Call    interface{}
}

func RegisterFake(info Func) {
	if FuncFake == nil {
		FuncFake = make(map[string]Func)
	}
	mul.Lock()
	FuncFake[info.TagName] = info
	mul.Unlock()
}
func RegisterFakes(infos []Func) {
	if FuncFake == nil {
		FuncFake = make(map[string]Func)
	}
	mul.Lock()
	for _, info := range infos {
		FuncFake[info.TagName] = info
	}
	mul.Unlock()
}

func GetFake(functionName string) *Func {
	info, ok := FuncFake[functionName]
	if !ok {
		return nil
	}

	return &info
}
func RemoveFake(funcName string) {
	_, ok := FuncFake[funcName]
	if !ok {
		return
	}
	mul.Lock()
	delete(FuncFake, funcName)
	mul.Unlock()
}
