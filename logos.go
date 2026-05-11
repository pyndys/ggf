package main

import "strings"

type Logo struct {
	Lines, Colors []string
}

func getLogo(osName string) Logo {
	name := strings.ToLower(osName)
	var l, c []string

	switch {
	case strings.Contains(name, "nixos"):
		l = []string{
			"  \\\\  \\\\ //   ",
			" ==\\\\__\\\\/ // ",
			"   //   \\\\//  ",
			"==//     //== ",
			" //\\\\___//    ",
			"// /\\\\  \\\\==  ",
			"  // \\\\  \\\\   ",
			"              ",
		}
		c = []string{
			"\033[34m",
			"\033[35m",
		}
	case strings.Contains(name, "arch"):
		l = []string{
			"      /\\      ",
			"     /  \\     ",
			"    /\\  \\\\    ",
			"   /      \\   ",
			"  /   ,,   \\  ",
			" /   |  |  -\\ ",
			"/_-''    ''-_\\",
			"              ",
		}
		c = []string{
			"\033[34m",
			"\033[36m",
		}
	case strings.Contains(name, "debian"):
		l = []string{
			"  _____   ",
			" /  __ \\  ",
			"|  /    | ",
			"|  \\___-  ",
			"-_        ",
			" --_      ",
			"          ",
			"          ",
		}
		c = []string{
			"\033[31m",
			"\033[33m",
		}
	case strings.Contains(name, "ubuntu"):
		l = []string{
			"        _   ",
			"    ---(_)  ",
			" _/  ---  \\ ",
			"(_) |   |   ",
			"  \\  --- _/ ",
			"    ---(_)  ",
			"            ",
			"            ",
		}
		c = []string{
			"\033[33m",
			"\033[31m",
		}
	case strings.Contains(name, "suse"):
		l = []string{
			"  _______  ",
			"__|   __ \\ ",
			"     / .\\ \\",
			"     \\__/ |",
			"   _______|",
			"   \\_______",
			"__________/",
			"           ",
		}
		c = []string{
			"\033[32m",
			"\033[36m",
		}
	case strings.Contains(name, "gentoo"):
		l = []string{
			" _-----_   ",
			"(       \\  ",
			"\\    0   \\ ",
			" \\        )",
			" /      _/ ",
			"(     _-   ",
			"\\____-     ",
			"           ",
		}
		c = []string{
			"\033[35m",
			"\033[90m",
		}
	case strings.Contains(name, "alpine"):
		l = []string{
			"    /\\ /\\    ",
			"   // \\  \\   ",
			"  //   \\  \\  ",
			" ///    \\  \\ ",
			" //      \\  \\",
			"          \\  ",
			"             ",
			"             ",
		}
		c = []string{
			"\033[34m",
			"\033[0m",
		}
	default:
		l = []string{
			"    ___   ",
			"   (.· |  ",
			"   (<> |  ",
			"  / __  \\ ",
			" ( /  \\ /|",
			"_/\\ __)/_)",
			"\\/-____\\/ ",
			"          ",
		}
		c = []string{
			"\033[90m",
			"\033[33m",
		}
	}

	return Logo{Lines: l, Colors: c}
}
