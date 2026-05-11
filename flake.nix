{
  description = "Great Go Fetch (ggf) - a fast system fetch tool written in Go";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = {
    self,
    nixpkgs,
  }: let
    systems = ["x86_64-linux" "aarch64-linux"];
    forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
  in {
    packages = forAllSystems (pkgs: {
      default = pkgs.buildGoModule {
        pname = "ggf";
        version = "0.2.0";
        src = ./.;
        vendorHash = null;

        ldflags = [
          "-s"
          "-w"
        ];
      };
    });

    ## For nix run
    apps = forAllSystems (pkgs: {
      default = {
        type = "app";
        program = "${self.packages.${pkgs.stdenv.hostPlatform.system}.default}/bin/ggf";
      };
    });
  };
}
