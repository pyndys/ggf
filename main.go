package main

import (
	"flag"
	"fmt"
)

func main() {
	osName := OS()
	resetColor := "\033[0m"

	var logoOverride string
	flag.StringVar(&logoOverride, "logo", "", "override distro logo")
	flag.StringVar(&logoOverride, "l", "", "override distro logo (shorthand)")

	var noLogo bool
	flag.BoolVar(&noLogo, "no-logo", false, "print without logo")
	flag.BoolVar(&noLogo, "nl", false, "print without logo (shorthand)")

	var noColor bool
	flag.BoolVar(&noColor, "no-color", false, "print without colors")
	flag.BoolVar(&noColor, "nc", false, "print without colors (shorthand)")

	logo := getLogo(osName)

	flag.Parse()
	if len(logoOverride) > 0 {
		logo = getLogo(logoOverride)
	}
	if noLogo == true {
		for i := range logo.Lines {
			logo.Lines[i] = ""
		}
	}
	if noColor == true {
		for i := range logo.Colors {
			logo.Colors[i] = resetColor
		}
	}

	fmt.Printf("%s %s %s@%s%s\n", logo.Colors[0], logo.Lines[0], user(), host(), resetColor)
	fmt.Printf("%s %s %sos%s       %s\n", logo.Colors[0], logo.Lines[1], logo.Colors[1], resetColor, osName)
	fmt.Printf("%s %s %skernel%s   %s\n", logo.Colors[0], logo.Lines[2], logo.Colors[1], resetColor, kernel())
	fmt.Printf("%s %s %sshell%s    %s\n", logo.Colors[0], logo.Lines[3], logo.Colors[1], resetColor, shell())
	fmt.Printf("%s %s %sterm%s     %s\n", logo.Colors[0], logo.Lines[4], logo.Colors[1], resetColor, term())
	fmt.Printf("%s %s %swm%s       %s\n", logo.Colors[0], logo.Lines[5], logo.Colors[1], resetColor, wm())
	fmt.Printf("%s %s %scpu%s      %s\n", logo.Colors[0], logo.Lines[6], logo.Colors[1], resetColor, cpu())
	fmt.Printf("%s %s %smemory%s   %s\n", logo.Colors[0], logo.Lines[7], logo.Colors[1], resetColor, memory())
	fmt.Printf("%s %s %suptime%s   %s\n", logo.Colors[0], logo.Lines[8], logo.Colors[1], resetColor, uptime())
}
