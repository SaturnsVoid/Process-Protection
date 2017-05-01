// Process Protection project main.go
package main

import (
	"fmt"
	"syscall"
)

var (
	ntdll                   = syscall.MustLoadDLL("ntdll.dll")
	NtSetInformationProcess = ntdll.MustFindProc("NtSetInformationProcess")
)

func SetInformationProcess(hProcess uintptr, processInformationClass int, processInformation int, processInformationLength int) {
	_, _, _ = NtSetInformationProcess.Call(hProcess, uintptr(processInformationClass), uintptr(processInformation), uintptr(processInformationLength))
}

func main() {
	fmt.Println("This program will protect its-self.")
	me, _ := syscall.GetCurrentProcess()
	SetInformationProcess(uintptr(me), 29, 1, 4)
	fmt.Println("Try killing me with a unelevated tool.")
	for {
	}
}
