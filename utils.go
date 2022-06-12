package main

import (
	"fmt"
	"goinject/w32"
	"log"
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

func decodeUtf16ToString(encoded [260]uint16) string {
	end := 0
	for {
		if encoded[end] == 0 {
			break
		}
		end++
	}
	return windows.UTF16ToString(encoded[:end])
}

func injectDll(dllPath string, processId int) {
	handleProcess := openProcess(processId)
	dwSize := len(dllPath)

	loc, err := w32.VirtualAllocEx(handleProcess, 0, dwSize, w32.MEM_RESERVE|w32.MEM_COMMIT, w32.PAGE_EXECUTE_READWRITE)
	if err == nil {

		err := w32.WriteProcessMemory(handleProcess, uintptr(loc), []byte(dllPath), uint(dwSize))

		test, _ := w32.ReadProcessMemory(handleProcess, uintptr(loc), uint(dwSize))
		if string(test) != dllPath {
			fmt.Println(string(test))
			fmt.Println("DLL Path not found in target process memory")
		}

		moduleKernel, _ := syscall.LoadLibrary("kernel32.dll")
		lpLoadLibrary, err := syscall.GetProcAddress(moduleKernel, "LoadLibraryA")

		if err != nil {
			log.Panic(err)
		}

		handleThread, _, _ := w32.CreateRemoteThread(handleProcess, nil, 0, uintptr(lpLoadLibrary), loc, 0)
		time.Sleep(5 * time.Second)

		w32.CloseHandle(handleThread)
	}
}
