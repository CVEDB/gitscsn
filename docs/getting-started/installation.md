# Installing GitScan

In this section you will find an aggregation of the different ways to install GitScan. installations are listed as either "official" or "community". Official integrations are developed by the core GitScan team and supported by it. Community integrations are integrations developed by the community, and collected here for your convenience. For support or questions about community integrations, please contact the original developers.

## Install using Package Manager

### RHEL/CentOS (Official)

=== "Repository"
    Add repository setting to `/etc/yum.repos.d`.

    ``` bash
    RELEASE_VERSION=$(grep -Po '(?<=VERSION_ID=")[0-9]' /etc/os-release)
    cat << EOF | sudo tee -a /etc/yum.repos.d/gitscan.repo
    [gitscan]
    name=GitScan repository
    baseurl=https://cvedb.github.io/gitscan-repo/rpm/releases/$RELEASE_VERSION/\$basearch/
    gpgcheck=0
    enabled=1
    EOF
    sudo yum -y update
    sudo yum -y install gitscan
    ```

=== "RPM"

    ``` bash
    rpm -ivh https://github.com/aquasecurity/gitscan/releases/download/{{ git.tag }}/gitscan_{{ git.tag[1:] }}_Linux-64bit.rpm
    ```

### Debian/Ubuntu (Official)

=== "Repository"
    Add repository setting to `/etc/apt/sources.list.d`.

    ``` bash
    sudo apt-get install wget apt-transport-https gnupg lsb-release
    wget -qO - https://cvedb.github.io/gitscan-repo/deb/public.key | gpg --dearmor | sudo tee /usr/share/keyrings/gitscan.gpg > /dev/null
    echo "deb [signed-by=/usr/share/keyrings/gitscan.gpg] https://cvedb.github.io/gitscan-repo/deb $(lsb_release -sc) main" | sudo tee -a /etc/apt/sources.list.d/gitscan.list
    sudo apt-get update
    sudo apt-get install gitscan
    ```

=== "DEB"

    ``` bash
    wget https://github.com/aquasecurity/gitscan/releases/download/{{ git.tag }}/gitscan_{{ git.tag[1:] }}_Linux-64bit.deb
    sudo dpkg -i gitscan_{{ git.tag[1:] }}_Linux-64bit.deb
    ```

### Homebrew (Official)

Homebrew for MacOS and Linux.

```bash
brew install gitscan
```

### Arch Linux (Community)

Arch Community Package Manager.

```bash
pacman -S gitscan
```

References: 
- <https://archlinux.org/packages/community/x86_64/gitscan/>
- <https://github.com/archlinux/svntogit-community/blob/packages/gitscan/trunk/PKGBUILD>


### MacPorts (Community)

[MacPorts](https://www.macports.org) for MacOS.

```bash
sudo port install gitscan
```

References:
- <https://ports.macports.org/port/gitscan/details/>

### Nix/NixOS (Community)

Nix package manager for Linux and MacOS.

=== "Command line"

`nix-env --install -A nixpkgs.gitscan`

=== "Configuration"

```nix
  # your other config ...
  environment.systemPackages = with pkgs; [
    # your other packages ...
    gitscan
  ];
```

=== "Home Manager"

```nix
  # your other config ...
  home.packages = with pkgs; [
    # your other packages ...
    gitscan
  ];
```

References: 
-  <https://github.com/NixOS/nixpkgs/blob/master/pkgs/tools/admin/gitscan/default.nix>

## Install from GitHub Release (Official)

### Download Binary

1. Download the file for your operating system/architecture from [GitHub Release assets](https://github.com/aquasecurity/gitscan/releases/tag/{{ git.tag }}) (`curl -LO https://url.to/gitscan.tar.gz`).  
2. Unpack the downloaded archive (`tar -xzf ./gitscan.tar.gz`).
3. Put the binary somewhere in your `$PATH` (e.g `mv ./gitscan /usr/local/bin/`).
4. Make sure the binary has execution bit turned on (`chmod +x ./gitscan`).

### Install Script

The process above can be automated by the following script:

```bash
curl -sfL https://raw.githubusercontent.com/aquasecurity/gitscan/main/contrib/install.sh | sh -s -- -b /usr/local/bin {{ git.tag }}
```

### Install from source

```bash
git clone --depth 1 --branch {{ git.tag }} https://github.com/aquasecurity/gitscan
cd gitscan
go install
```

## Use container image

1. Pull GitScan image (`docker pull aquasec/gitscan:{{ git.tag[1:] }}`)
2. It is advisable to mount a consistent [cache dir](https://cvedb.github.io/gitscan/{{ git.tag }}/docs/vulnerability/examples/cache/) on the host into the GitScan container.
3. For scanning container images with GitScan, mount `docker.sock` from the host into the GitScan container.

Example:

``` bash
docker run -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/Library/Caches:/root/.cache/ aquasec/gitscan:{{ git.tag[1:] }} image python:3.4-alpine
```

Registry | Repository | Link | Supportability
Docker Hub | `docker.io/aquasec/gitscan` | https://hub.docker.com/r/aquasec/gitscan | Official
GitHub Container Registry (GHCR) | `ghcr.io/aquasecurity/gitscan` | https://github.com/orgs/aquasecurity/packages/container/package/gitscan | Official
AWS Elastic Container Registry (ECR) | `public.ecr.aws/aquasecurity/gitscan` | https://gallery.ecr.aws/aquasecurity/gitscan | Official

## Other Tools to use and deploy GitScan

For additional tools and ways to install and use GitScan in different environments such as in IDE, Kubernetes or CI/CD, see [Ecosystem section](../ecosystem/index.md).
