{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs = inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; }
      {
        systems = [
          "x86_64-linux"
          "aarch64-linux"
        ];

        perSystem = { self', pkgs, ... }: {
          apps.default = {
            type = "app";
            program = self'.packages.default;
          };

          packages.default = pkgs.callPackage (import ./nix/package.nix) { };
        };

        flake = {
          homeManagerModules.default = import ./nix/home-manager.nix;
        };
      };
}
