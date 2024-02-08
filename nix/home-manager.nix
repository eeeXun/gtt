{ config, lib, pkgs, ... }:
with lib;
let
  cfg = config.programs.gtt;
  yamlFormat = pkgs.formats.yaml { };
in
{
  options.programs.gtt = {
    enable = mkEnableOption "gtt";

    settings = mkOption {
      type = yamlFormat.type;
      default = { };
      example = literalExpression ''
        {
          api_key = {
            deepl = "ctiegsrn-mcgcmdunw:l984-cnulfmuz";
          };
        }
      '';

      description = ''
        Configuration writte to
        {file}`$XDG_CONFIG_HOME/.config/gtt/server.yaml`.
      '';
    };
  };

  config = mkIf cfg.enable
    {
      home.packages = with pkgs; [ gtt ];

      xdg.configFile."gtt/server.yaml" = mkIf (cfg.settings != { }) {
        source = yamlFormat.generate "server.yaml" cfg.settings;
      };
    };
}
