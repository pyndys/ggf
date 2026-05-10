# ggf (Great Go Fetch)

**English** | [Русский](README_RU.md)

A fast system info fetch utility for Linux-based systems, written in Go. Displays: os, kernel, shell, uptime, term, wm, memory.

## Features

- **High Performance:** Direct data retrieval from `/proc`, `/etc/os-release`, and environment variables.
- **Lightweight:** Optimized binary with zero external dependencies.
- **Distribution Support:** Adaptive ASCII logos for popular Linux distros.

## Installation and Usage

### On NixOS (Flakes)

You can run the utility instantly without installation:
```sh
nix run github:pyndys/ggf
```

For permanent installation, add this to your `flake.nix` inputs:
```nix
ggf.url = "github:pyndys/ggf";
```

Then add the package to your `systemPackages`:
```nix
inputs.ggf.packages.${system}.default
```

### Build from Source

Requires **Go >= 1.18** and **git**.

```sh
git clone https://github.com/pyndys/ggf.git
cd ggf
go build -ldflags="-s -w" -o ggf .
```

*The `-s -w` flags are used to strip the symbol table and debug information, significantly reducing the binary size.*

## License
This project is licensed under the [MIT License](LICENSE).

## Credits
* ASCII Logos: Adapted from the [pfetch](https://github.com/dylanaraps/pfetch) project.
