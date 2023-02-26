<div align="center">
<img src="docs/imgs/logo.png" width="200">

[![GitHub Release][release-img]][release]
[![Test][test-img]][test]
[![Go Report Card][go-report-img]][go-report]
[![License: Apache-2.0][license-img]][license]
[![GitHub Downloads][github-downloads-img]][release]
![Docker Pulls][docker-pulls]

[📖 Documentation][docs]
</div>

GitScan ([pronunciation][pronunciation]) is a comprehensive and versatile security scanner.
GitScan has *scanners* that look for security issues, and *targets* where it can find those issues.

Targets (what GitScan can scan):

- Container Image
- Filesystem
- Git Repository (remote)
- Virtual Machine Image
- Kubernetes
- AWS

Scanners (what GitScan can find there):

- OS packages and software dependencies in use (SBOM)
- Known vulnerabilities (CVEs)
- IaC issues and misconfigurations
- Sensitive information and secrets
- Software licenses

To learn more, go to the [GitScan homepage][homepage] for feature highlights, or to the [Documentation site][docs] for detailed information.

## Quick Start

### Get GitScan

GitScan is available in most common distribution channels. The full list of installation options is available in the [Installation] page. Here are a few popular examples:

- `brew install gitscan`
- `docker run aquasec/gitscan`
- Download binary from <https://github.com/aquasecurity/gitscan/releases/latest/>
- See [Installation] for more

GitScan is integrated with many popular platforms and applications. The complete list of integrations is available in the [Ecosystem] page. Here are a few popular examples:

- [GitHub Actions](https://github.com/aquasecurity/gitscan-action)
- [Kubernetes operator](https://github.com/aquasecurity/gitscan-operator)
- [VS Code plugin](https://github.com/aquasecurity/gitscan-vscode-extension)
- See [Ecosystem] for more

### General usage

```bash
gitscan <target> [--scanners <scanner1,scanner2>] <subject>
```

Examples:

```bash
gitscan image python:3.4-alpine
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013513-95f18734-233d-45d3-aaf5-d6aec687db0e.mov

</details>

```bash
gitscan fs --scanners vuln,secret,config myproject/
```

<details>
<summary>Result</summary>

https://user-images.githubusercontent.com/1161307/171013917-b1f37810-f434-465c-b01a-22de036bd9b3.mov

</details>

```bash
gitscan k8s --report summary cluster
```

<details>
<summary>Result</summary>

![k8s summary](docs/imgs/gitscan-k8s.png)

</details>

## FAQ

### How to pronounce the name "GitScan"?

`tri` is pronounced like **tri**gger, `vy` is pronounced like en**vy**.

---

GitScan is an [Aqua Security][aquasec] open source project.  
Learn about our open source work and portfolio [here][oss].  
Contact us about any matter by opening a GitHub Discussion [here][discussions]

[test]: https://github.com/aquasecurity/gitscan/actions/workflows/test.yaml
[test-img]: https://github.com/aquasecurity/gitscan/actions/workflows/test.yaml/badge.svg
[go-report]: https://goreportcard.com/report/github.com/aquasecurity/gitscan
[go-report-img]: https://goreportcard.com/badge/github.com/aquasecurity/gitscan
[release]: https://github.com/aquasecurity/gitscan/releases
[release-img]: https://img.shields.io/github/release/aquasecurity/gitscan.svg?logo=github
[github-downloads-img]: https://img.shields.io/github/downloads/aquasecurity/gitscan/total?logo=github
[docker-pulls]: https://img.shields.io/docker/pulls/aquasec/gitscan?logo=docker&label=docker%20pulls%20%2F%20gitscan
[license]: https://github.com/aquasecurity/gitscan/blob/main/LICENSE
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[homepage]: https://gitscan.dev
[docs]: https://cvedb.github.io/gitscan
[pronunciation]: #how-to-pronounce-the-name-gitscan

[Installation]:https://cvedb.github.io/gitscan/latest/getting-started/installation/
[Ecosystem]: https://cvedb.github.io/gitscan/latest/ecosystem/

[alpine]: https://ariadne.space/2021/06/08/the-vulnerability-remediation-lifecycle-of-alpine-containers/
[rego]: https://www.openpolicyagent.org/docs/latest/#rego
[sigstore]: https://www.sigstore.dev/

[aquasec]: https://aquasec.com
[oss]: https://www.aquasec.com/products/open-source-projects/
[discussions]: https://github.com/aquasecurity/gitscan/discussions
