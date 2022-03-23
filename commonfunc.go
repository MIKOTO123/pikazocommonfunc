package pikazocommonfunc

import (
	"fmt"
	"runtime"
)

// 获取正在运行的函数名
func RunFuncName(skip int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(skip, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func PrintEndInFunc(skip int) {
	fmt.Println("====================================", RunFuncName(skip), "====================================")
}

func PrintStartInFunc(skip int) {
	fmt.Println("====================================", RunFuncName(skip), "====================================")
}
