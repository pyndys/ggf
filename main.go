package main

import (
	"fmt"
)

func main() {
	osName := OS()
	logo := getLogo(osName)

	reset := "\033[0m"

	fmt.Printf("%s %s %s@%s%s\n", logo.Colors[0], logo.Lines[0], user(), host(), reset)
	fmt.Printf("%s %s %sos%s       %s\n", logo.Colors[0], logo.Lines[1], logo.Colors[1], reset, osName)
	fmt.Printf("%s %s %skernel%s   %s\n", logo.Colors[0], logo.Lines[2], logo.Colors[1], reset, kernel())
	fmt.Printf("%s %s %sshell%s    %s\n", logo.Colors[0], logo.Lines[3], logo.Colors[1], reset, shell())
	fmt.Printf("%s %s %sterm%s     %s\n", logo.Colors[0], logo.Lines[4], logo.Colors[1], reset, term())
	fmt.Printf("%s %s %swm%s       %s\n", logo.Colors[0], logo.Lines[5], logo.Colors[1], reset, wm())
	fmt.Printf("%s %s %scpu%s      %s\n", logo.Colors[0], logo.Lines[6], logo.Colors[1], reset, cpu())
	fmt.Printf("%s %s %smemory%s   %s\n", logo.Colors[0], logo.Lines[7], logo.Colors[1], reset, memory())
	fmt.Printf("%s %s %suptime%s   %s\n", logo.Colors[0], logo.Lines[8], logo.Colors[1], reset, uptime())
}
