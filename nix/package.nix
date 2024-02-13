{ lib
, buildGoModule
, pkg-config
, alsa-lib
, wl-clipboard
, xclip
}:

buildGoModule {
  pname = "gtt";
  version = "8";

  src = ../.;

  vendorHash = "sha256-5Uwi1apowHoUtvkSgmUV9WbfpVQFTqJ9GA2sRnC5nFw=";

  nativeBuildInputs = [
    pkg-config
  ];

  buildInputs = [
    alsa-lib
    xclip
    wl-clipboard
  ];

  ldflags = [ "-s" "-w" ];

  meta = with lib; {
    description = "Google Translate TUI (Originally). Now support Apertium, Argos, Bing, ChatGPT, DeepL, Google, Reverso";
    homepage = "https://github.com/eeeXun/gtt";
    license = licenses.mit;
    maintainers = with maintainers; [ linuxissuper ];
    mainProgram = "gtt";
  };
}
