package main

import (
	"fmt"
)

func main() {
	// Colors
	blue := "\033[34m"
	purple := "\033[35m"
	reset := "\033[0m"

	osName := OS()
	logo := getLogo(osName)

	// Print ggf
	fmt.Printf("%s %s %s%s%s@%s%s\n", blue, logo.Lines[0], reset, purple, user(), host(), reset)
	fmt.Printf("%s %s %sos       %s\n", blue, logo.Lines[1], reset, osName)
	fmt.Printf("%s %s %skernel   %s\n", blue, logo.Lines[2], reset, kernel())
	fmt.Printf("%s %s %sshell    %s\n", blue, logo.Lines[3], reset, shell())
	fmt.Printf("%s %s %suptime   %s\n", blue, logo.Lines[4], reset, uptime())
	fmt.Printf("%s %s %sterm     %s\n", blue, logo.Lines[5], reset, term())
	fmt.Printf("%s %s %swm       %s\n", blue, logo.Lines[6], reset, wm())
	fmt.Printf("%s %s %smemory   %s\n", blue, logo.Lines[7], reset, memory())
}
