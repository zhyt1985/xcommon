package verify

import (
	"sync"
)

var FuncVerify map[string]Func

var mul sync.Mutex

type Func struct {
	Name        string
	Description string
	CallParam   func([]string, interface{}) bool // 带参数的调用
	Call        func(interface{}) bool           // 没有参数调用
}

func RegisterVerify(info Func) {
	if FuncVerify == nil {
		FuncVerify = make(map[string]Func)
	}
	mul.Lock()
	FuncVerify[info.Name] = info
	mul.Unlock()
}
func RegisterVerifies(infos []Func) {
	if FuncVerify == nil {
		FuncVerify = make(map[string]Func)
	}
	mul.Lock()
	for _, info := range infos {
		FuncVerify[info.Name] = info
	}
	mul.Unlock()
}

func GetVerify(functionName string) *Func {
	info, ok := FuncVerify[functionName]
	if !ok {
		return nil
	}

	return &info
}
func RemoveVerify(funcName string) {
	_, ok := FuncVerify[funcName]
	if !ok {
		return
	}
	mul.Lock()
	delete(FuncVerify, funcName)
	mul.Unlock()
}
