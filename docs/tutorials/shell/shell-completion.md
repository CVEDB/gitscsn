# Enable shell completion

Below is example steps to enable shell completion feature for `gitscan` cli:

### 1. Know your current shell

```bash
$ echo $SHELL
/bin/zsh # For this example it is zsh, but will be vary depend on your $SHELL, maybe /bin/bash or /bin/fish
```

### 2. Run `completion` command to get sub-commands

``` bash
$ gitscan completion zsh -h
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(gitscan completion zsh); compdef _gitscan gitscan

To load completions for every new session, execute once:

#### Linux:

	gitscan completion zsh > "${fpath[1]}/_gitscan"

#### macOS:

	gitscan completion zsh > $(brew --prefix)/share/zsh/site-functions/_gitscan

You will need to start a new shell for this setup to take effect.
```

### 3. Run the sub-commands following the instruction

```bash
echo "autoload -U compinit; compinit" >> ~/.zshrc
source <(gitscan completion zsh); compdef _gitscan gitscan
gitscan completion zsh > "${fpath[1]}/_gitscan"
```

### 4. Start a new shell and you can see the shell completion

```bash
$ gitscan [tab]
aws         -- scan aws account
completion  -- Generate the autocompletion script for the specified shell
config      -- Scan config files for misconfigurations
filesystem  -- Scan local filesystem
help        -- Help about any command
image       -- Scan a container image
kubernetes  -- scan kubernetes cluster
module      -- Manage modules
plugin      -- Manage plugins
repository  -- Scan a remote repository
rootfs      -- Scan rootfs
sbom        -- Scan SBOM for vulnerabilities
server      -- Server mode
version     -- Print the version
```