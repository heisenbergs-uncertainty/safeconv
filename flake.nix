{
  description = "Safe Convert Environment with Go and Nix Flakes";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };

        # Go version from go.mod
        go = pkgs.go_1_24;




        # Development and build tools
        buildInputs = with pkgs; [
          # Go toolchain
          go
          gotools
          gopls
          go-tools
          golangci-lint

          # Code generation tools
          protobuf
          protoc-gen-go

          # Development utilities
          gnumake
          jq
          yq-go
          nixfmt

          # Testing and quality tools
          gosec
          govulncheck

          air # For live reload during development
        ];

        # Shell hooks for environment setup
        shellHook = ''
          echo "🚀 Safe Convert Development Environment"
          echo "================================"
          echo ""
          echo "Available tools:"
          echo "  Go:        $(go version)"
          echo ""

          # Set up Go environment
          export GOROOT="${go}/share/go"
          export GOPATH="$HOME/go"
          export PATH="$GOPATH/bin:$PATH"

        '';

      in
      {
        # Development shell
        devShells.default = pkgs.mkShell {
          inherit buildInputs shellHook;

          name = "safe-convert-dev-shell";

          # Environment variables for development
          NIX_ENFORCE_PURITY = 0;
        };

        # Package definition for the IoTGo binary
        # packages.default = pkgs.buildGoModule {
        #   pname = "safe-convert";
        #   version = "1.0.0";
        #
        #   src = ./.;
        #
        #   vendorHash = null; # We use go.mod/go.sum
        #
        #   buildInputs = [ go ];
        #
        #   buildPhase = ''
        #     go build -o 
        #   '';
        #
        #   installPhase = ''
        #     mkdir -p $out/bin
        #     cp bin/iotgo $out/bin/
        #   '';
        #
        #   meta = with pkgs.lib; {
        #     description = "IoT API Gateway with Domain-Driven Design";
        #     homepage = "https://github.com/cat-spmog/iotgo-client";
        #     license = licenses.mit;
        #     mainProgram = "iotgo";
        #   };
        # };

        # Application for running the server
        apps.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/safe-convert";
        };

        # Formatter for nix files
        formatter = pkgs.nixpkgs-fmt;
      }
    );
}
