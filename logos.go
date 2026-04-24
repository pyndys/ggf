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
			" /  __ \\\\ ",
			"|  /    | ",
			"|  \\\\___- ",
			"-_        ",
			" --_      ",
			"          ",
		}
	default:
		l = []string{"", "", "", "", "", "", ""}
	}

	return Logo{Lines: l}
}
