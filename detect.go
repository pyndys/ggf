package main

import (
	"os"
	"path"
	"strings"
)

func user() string {
	return os.Getenv("USER")
}

func term() string {
	return os.Getenv("TERM")
}

func wm() string {
	if os.Getenv("XDG_SESSION_DESKTOP") == "" {
		return "unknown" + " (" + os.Getenv("XDG_SESSION_TYPE") + ")"
	} else {
		return os.Getenv("XDG_SESSION_DESKTOP") + " (" + os.Getenv("XDG_SESSION_TYPE") + ")"
	}
}

func host() string {
	hostname, _ := os.Hostname()
	return hostname
}

func kernel() string {
	kernelRaw, _ := os.ReadFile("/proc/sys/kernel/osrelease")
	kernelName := strings.TrimSpace(string(kernelRaw))
	return kernelName
}

func OS() string {
	osData, _ := os.ReadFile("/etc/os-release")
	osName := "unknown"
	for _, line := range strings.Split(string(osData), "\n") {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			osName = strings.Trim(line[12:], "\"")
			break
		}
	}
	return osName
}

func shell() string {
	shellPath := os.Getenv("SHELL")
	shellName := "unknown"
	if shellPath != "" {
		shellName = path.Base(shellPath)
	}
	return shellName
}
