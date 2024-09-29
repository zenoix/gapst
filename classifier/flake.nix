{
  description = "Python gRPC Environment";

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

        packages = [
          pkgs.gnumake
          pkgs.gnused
          (pkgs.python3.withPackages (
            python-pkgs: with python-pkgs; [
              grpcio
              grpcio-tools
              numpy
              pandas
              pandas-stubs
              scikit-learn
            ]
          ))
        ];

      };
    };
}
