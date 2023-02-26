# Plugin

```bash
Manage plugins

Usage:
  gitscan plugin [command]

Aliases:
  plugin, p

Available Commands:
  info        Show information about the specified plugin
  install     Install a plugin
  list        List installed plugin
  run         Run a plugin on the fly
  uninstall   Uninstall a plugin
  update      Update an existing plugin

Flags:
  -h, --help   help for plugin

Global Flags:
      --cache-dir string          cache directory (default "/Users/teppei/Library/Caches/gitscan")
  -c, --config string             config path (default "gitscan.yaml")
  -d, --debug                     debug mode
      --generate-default-config   write the default config to gitscan-default.yaml
      --insecure                  allow insecure server connections when using TLS
  -q, --quiet                     suppress progress bar and log output
      --timeout duration          timeout (default 5m0s)
  -v, --version                   show version

Use "gitscan plugin [command] --help" for more information about a command.
```