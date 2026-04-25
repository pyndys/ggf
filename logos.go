package main

import "strings"

type Logo struct {
	Lines []string
}

func getLogo(osName string) Logo {
	name := strings.ToLower(osName)
	var l []string

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
		}
	}

	return Logo{Lines: l}
}
