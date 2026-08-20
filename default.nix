{ lib, buildGoModule }:
buildGoModule {
  pname = "a";
  version = "0.0.1";

  src = ./.;

  vendorHash = "sha256-THOTzWban6ZdGYV+qEH8AAyddvtHhPOekCamehXxHLY=";

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
