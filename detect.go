package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
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
		return fmt.Sprintf("unknown (%s)", os.Getenv("XDG_SESSION_TYPE"))
	} else {
		return fmt.Sprintf("%s (%s)", os.Getenv("XDG_SESSION_DESKTOP"), os.Getenv("XDG_SESSION_TYPE"))
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

func uptime() string {
	secondsRaw, _ := os.ReadFile("/proc/uptime")
	seconds := strings.Split(string(secondsRaw), " ")[0]

	secondsFloat, _ := strconv.ParseFloat(seconds, 64)
	secondsInt := int(secondsFloat)

	days := int(secondsInt / 86400)
	hours := int((secondsInt % 86400) / 3600)
	minutes := int((secondsInt % 3600) / 60)

	switch {
	case secondsInt < 60:
		return "less than minute"
	case 60 <= secondsInt && secondsInt < 3600:
		return fmt.Sprintf("%d minutes", minutes)
	case 3600 <= secondsInt && secondsInt < 86400:
		return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
	default:
		return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
	}
}

func memory() string {
	memData, _ := os.ReadFile("/proc/meminfo")
	var s string
	var memTotalKb, memAvailableKb float64

	for _, line := range strings.Split(string(memData), "\n") {
		if strings.HasPrefix(line, "MemTotal:") {
			s = strings.TrimPrefix(line, "MemTotal:")
			s = strings.TrimSuffix(s, "kB")
			s = strings.TrimSpace(s)

			memTotalKb, _ = strconv.ParseFloat(s, 64)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			s = strings.TrimPrefix(line, "MemAvailable:")
			s = strings.TrimSuffix(s, "kB")
			s = strings.TrimSpace(s)

			memAvailableKb, _ = strconv.ParseFloat(s, 64)
		}
	}

	memTotalGb := memTotalKb / (1024 * 1024)
	memUsingGb := (memTotalKb - memAvailableKb) / (1024 * 1024)

	return fmt.Sprintf("%.2f GB / %.2f GB", memUsingGb, memTotalGb)
}

func cpu() string {
	cpuData, _ := os.ReadFile("/proc/cpuinfo")
	cpuName := "unknown"
	for _, line := range strings.Split(string(cpuData), "\n") {
		if strings.HasPrefix(line, "model name") {
			cpuName = strings.Trim(line[13:], "\"")
			break
		}
	}
	return cpuName
}
