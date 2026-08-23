{ lib, buildGoModule }:
buildGoModule {
  pname = "a";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-m8o6hGP1nB6fyCf9sLty2lbTN8AoszRZrs2M+4ak9xE=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    description = "";
    homepage = "https://github.com/theobori/a";
    license = lib.licenses.mit;
    mainProgram = "a";
  };
}
