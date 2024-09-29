{
  description = "Go API Module";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
  };

  outputs =
    { self, nixpkgs }:
    let
      pkgs = nixpkgs.legacyPackages."x86_64-linux";
    in
    {
      # importing package example
      # packages."x86_64-linux".default = 
      #   pkgs.callPackage (import ./default.nix) {};

      devShells."x86_64-linux".default = pkgs.mkShell {

        packages = with pkgs; [
          gnumake
          go
          golangci-lint
          protobuf
          protoc-gen-go
          protoc-gen-go-grpc
        ];

      };
    };
}
