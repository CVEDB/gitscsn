# Plugins
GitScan provides a plugin feature to allow others to extend the GitScan CLI without the need to change the GitScancode base.
This plugin system was inspired by the plugin system used in [kubectl][kubectl], [Helm][helm], and [Conftest][conftest].

## Overview
GitScan plugins are add-on tools that integrate seamlessly with GitScan.
They provide a way to extend the core feature set of GitScan, but without requiring every new feature to be written in Go and added to the core tool.

- They can be added and removed from a GitScan installation without impacting the core GitScan tool.
- They can be written in any programming language.
- They integrate with GitScan, and will show up in GitScan help and subcommands.

!!! warning
    GitScan plugins available in public are not audited for security.
    You should install and run third-party plugins at your own risk, since they are arbitrary programs running on your machine.


## Installing a Plugin
A plugin can be installed using the `gitscan plugin install` command.
This command takes a url and will download the plugin and install it in the plugin cache.

GitScan adheres to the XDG specification, so the location depends on whether XDG_DATA_HOME is set.
GitScan will now search XDG_DATA_HOME for the location of the GitScan plugins cache.
The preference order is as follows:

- XDG_DATA_HOME if set and .gitscan/plugins exists within the XDG_DATA_HOME dir
- ~/.gitscan/plugins

Under the hood GitScan leverages [go-getter][go-getter] to download plugins.
This means the following protocols are supported for downloading plugins:

- OCI Registries
- Local Files
- Git
- HTTP/HTTPS
- Mercurial
- Amazon S3
- Google Cloud Storage

For example, to download the Kubernetes GitScan plugin you can execute the following command:

```bash
$ gitscan plugin install github.com/aquasecurity/gitscan-plugin-kubectl
```
## Using Plugins
Once the plugin is installed, GitScan will load all available plugins in the cache on the start of the next GitScan execution.
A plugin will be made in the GitScan CLI based on the plugin name.
To display all plugins, you can list them by `gitscan --help`

```bash
$ gitscan --help
NAME:
   gitscan - A simple and comprehensive vulnerability scanner for containers

USAGE:
   gitscan [global options] command [command options] target

VERSION:
   dev

COMMANDS:
   image, i          scan an image
   filesystem, fs    scan local filesystem
   repository, repo  scan remote repository
   client, c         client mode
   server, s         server mode
   plugin, p         manage plugins
   kubectl           scan kubectl resources
   help, h           Shows a list of commands or help for one command
```

As shown above, `kubectl` subcommand exists in the `COMMANDS` section.
To call the kubectl plugin and scan existing Kubernetes deployments, you can execute the following command:

```
$ gitscan kubectl deployment <deployment-id> -- --ignore-unfixed --severity CRITICAL
```

Internally the kubectl plugin calls the kubectl binary to fetch information about that deployment and passes the using images to GitScan.
You can see the detail [here][gitscan-plugin-kubectl].

If you want to omit even the subcommand, you can use `TRIVY_RUN_AS_PLUGIN` environment variable.

```bash
$ TRIVY_RUN_AS_PLUGIN=kubectl gitscan job your-job -- --format json
```

## Installing and Running Plugins on the fly
`gitscan plugin run` installs a plugin and runs it on the fly.
If the plugin is already present in the cache, the installation is skipped.

```bash
gitscan plugin run github.com/aquasecurity/gitscan-plugin-kubectl pod your-pod -- --exit-code 1
```

## Uninstalling Plugins
Specify a plugin name with `gitscan plugin uninstall` command.

```bash
$ gitscan plugin uninstall kubectl
```

## Building Plugins
Each plugin has a top-level directory, and then a plugin.yaml file.

```bash
your-plugin/
  |
  |- plugin.yaml
  |- your-plugin.sh
```

In the example above, the plugin is contained inside of a directory named `your-plugin`.
It has two files: plugin.yaml (required) and an executable script, your-plugin.sh (optional).

The core of a plugin is a simple YAML file named plugin.yaml.
Here is an example YAML of gitscan-plugin-kubectl plugin that adds support for Kubernetes scanning.

```yaml
name: "kubectl"
repository: github.com/aquasecurity/gitscan-plugin-kubectl
version: "0.1.0"
usage: scan kubectl resources
description: |-
  A GitScan plugin that scans the images of a kubernetes resource.
  Usage: gitscan kubectl TYPE[.VERSION][.GROUP] NAME
platforms:
  - selector: # optional
      os: darwin
      arch: amd64
    uri: ./gitscan-kubectl # where the execution file is (local file, http, git, etc.)
    bin: ./gitscan-kubectl # path to the execution file
  - selector: # optional
      os: linux
      arch: amd64
    uri: https://github.com/aquasecurity/gitscan-plugin-kubectl/releases/download/v0.1.0/gitscan-kubectl.tar.gz
    bin: ./gitscan-kubectl
```

The `plugin.yaml` field should contain the following information:

- name: The name of the plugin. This also determines how the plugin will be made available in the GitScan CLI. For example, if the plugin is named kubectl, you can call the plugin with `gitscan kubectl`. (required)
- version: The version of the plugin. (required)
- usage: A short usage description. (required)
- description: A long description of the plugin. This is where you could provide a helpful documentation of your plugin. (required)
- platforms: (required)
  - selector: The OS/Architecture specific variations of a execution file. (optional)
    - os: OS information based on GOOS (linux, darwin, etc.) (optional)
    - arch: The architecture information based on GOARCH (amd64, arm64, etc.) (optional)
  - uri: Where the executable file is. Relative path from the root directory of the plugin or remote URL such as HTTP and S3. (required)
  - bin: Which file to call when the plugin is executed. Relative path from the root directory of the plugin. (required)

The following rules will apply in deciding which platform to select:

- If both `os` and `arch` under `selector` match the current platform, search will stop and the platform will be used.
- If `selector` is not present, the platform will be used.
- If `os` matches and there is no more specific `arch` match, the platform will be used.
- If no `platform` match is found, GitScan will exit with an error.

After determining platform, GitScan will download the execution file from `uri` and store it in the plugin cache.
When the plugin is called via GitScan CLI, `bin` command will be executed.

The plugin is responsible for handling flags and arguments. Any arguments are passed to the plugin from the `gitscan` command.

## Example
https://github.com/aquasecurity/gitscan-plugin-kubectl

[kubectl]: https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/
[helm]: https://helm.sh/docs/topics/plugins/
[conftest]: https://www.conftest.dev/plugins/
[go-getter]: https://github.com/hashicorp/go-getter
[gitscan-plugin-kubectl]: https://github.com/aquasecurity/gitscan-plugin-kubectl

