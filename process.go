package main

import (
	"goinject/w32"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type WindowsProcess struct {
	ProcessID       int
	ParentProcessID int
	Exe             string
}

func newWindowsProcess(e *windows.ProcessEntry32) WindowsProcess {
	// Find when the string ends for decoding
	end := 0
	for {
		if e.ExeFile[end] == 0 {
			break
		}
		end++
	}

	return WindowsProcess{
		ProcessID:       int(e.ProcessID),
		ParentProcessID: int(e.ParentProcessID),
		Exe:             syscall.UTF16ToString(e.ExeFile[:end]),
	}
}

func getProcesses() ([]WindowsProcess, error) {
	handle, err := windows.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(handle)

	var entry windows.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	// get the first process
	err = windows.Process32First(handle, &entry)
	if err != nil {
		return nil, err
	}

	results := make([]WindowsProcess, 0, 50)
	for {
		results = append(results, newWindowsProcess(&entry))

		err = windows.Process32Next(handle, &entry)
		if err != nil {
			if err == syscall.ERROR_NO_MORE_FILES {
				return results, nil
			}
			return nil, err
		}
	}
}

func getProcess(processes []WindowsProcess, name string) *WindowsProcess {
	for _, p := range processes {
		if strings.ToLower(p.Exe) == strings.ToLower(name) {
			return &p
		}
	}
	return nil
}

func getProcessId(processName string) int {
	processId := 0
	handle := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPPROCESS, 0)
	if handle != w32.INVALID_HANDLE {
		var processEntry w32.PROCESSENTRY32
		processEntry.Size = uint32(unsafe.Sizeof(processEntry))
		if w32.Process32First(handle, &processEntry) {
			for {
				if decodeUtf16ToString(processEntry.ExeFile) == processName {
					processId = int(processEntry.ProcessID)
					break
				}
				if !w32.Process32Next(handle, &processEntry) {
					break
				}
			}
		}
	}
	w32.CloseHandle(handle)
	return processId
}

func openProcess(processId int) w32.HANDLE {
	handle, _ := w32.OpenProcess(w32.PROCESS_ALL_ACCESS, false, uint32(processId))
	return handle
}

// returns False if 64-bit process is running on 64-bit OS
func getProcessArch() bool {
	handle := windows.CurrentProcess()
	var isWow64 bool
	err := windows.IsWow64Process(handle, &isWow64)
	if err != nil {
		panic(err)
	}

	return isWow64
}
