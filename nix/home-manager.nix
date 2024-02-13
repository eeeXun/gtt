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
            deepl.value = "this-is-an-api-key";
            chatgpt.file = "$\{builtins.readFile /run/agenix/chatgpt}";
          };
        }
      '';

      description = ''
        Configuration written to
        {file}`$XDG_CONFIG_HOME/.config/gtt/server.yaml`.
      '';
    };

    theme = mkOption {
      type = yamlFormat.type;
      default = { };
      example = literalExpression ''
        {
          tokyonight = {
            bg = "0x1a1b26";
            fg = "0xc0caf5";
            gray = "0x414868";
            red = "0xf7768e";
            green = "0x9ece6a";
            yellow = "0xe0af68";
            blue = "0x7aa2f7";
            purple = "0xbb9af7";
            cyan = "0x7dcfff";
            orange = "0xff9e64";         
          };
        }
      '';

      description = ''
        Theme written to
        {file}`$XDG_CONFIG_HOME/.config/gtt/theme.yaml`.
      '';
    };


    keymap = mkOption {
      type = yamlFormat.type;
      default = { };
      example = literalExpression ''
        {
          exit = "C-c";
          translate = "C-j";
        }
      '';
    };
  };

  config = mkIf cfg.enable
    {
      home.packages = with pkgs; [ gtt ];

      xdg.configFile = {
        "gtt/server.yaml" = mkIf (cfg.settings != { }) {
          source = yamlFormat.generate "server.yaml" cfg.settings;
        };

        "gtt/theme.yaml" = mkIf (cfg.theme != { }) {
          source = yamlFormat.generate "theme.yaml" cfg.theme;
        };

        "gtt/keymap.yaml" = mkIf (cfg.keymap != { }) {
          source = yamlFormat.generate "keymap.yaml" cfg.theme;
        };
      };
    };
}
