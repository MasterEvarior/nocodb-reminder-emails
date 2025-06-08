{
  description = "Development flake for nocodb-reminder-emails";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-25.05";
  };

  outputs =
    { nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
    in
    {
      devShells."${x86}".default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          # Golang
          go
          golangci-lint

          # Formatters
          treefmt
          mdformat
          yamlfmt
          jsonfmt
          deadnix
          nixfmt-rfc-style
        ];

        shellHook = ''
          git config --local core.hooksPath .githooks/
        '';

        # Environment Variables
        NDBRE_BASE_URL = "https://nocodb.mydomain.com";
        NDBRE_API_TOKEN = "xyz";
        NDBRE_EMAIL_FROM = "reminders@mydomain.com";
        NDBRE_SMTP_SERVER = "localhost:2525";
        NDBRE_EMAIL_TO = "reminders@mydomain.com";
      };
    };
}
