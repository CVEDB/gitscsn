# GitScan Scanner

GitScan vulnerability scanner standalone installation.

## TL;DR;

```
$ helm install gitscan . --namespace gitscan --create-namespace
```

## Introduction

This chart bootstraps a GitScan deployment on a [Kubernetes](http://kubernetes.io) cluster using the
[Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.12+
- Helm 3+

## Installing from the Aqua Chart Repository

```
helm repo add aquasecurity https://cvedb.github.io/helm-charts/
helm repo update
helm search repo gitscan
helm install my-gitscan aquasecurity/gitscan
```

## Installing the Chart

To install the chart with the release name `my-release`:

```
$ helm install my-release .
```

The command deploys GitScan on the Kubernetes cluster in the default configuration. The [Parameters](#parameters)
section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`.

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Parameters

The following table lists the configurable parameters of the GitScan chart and their default values.

|                 Parameter             |                                Description                              |    Default     |
|---------------------------------------|-------------------------------------------------------------------------|----------------|
| `image.registry`                      | Image registry                                                          | `docker.io`    |
| `image.repository`                    | Image name                                                              | `aquasec/gitscan` |
| `image.tag`                           | Image tag                                                               | `{TAG_NAME}`   |
| `image.pullPolicy`                    | Image pull policy                                                       | `IfNotPresent` |
| `image.pullSecret`                    | The name of an imagePullSecret used to pull gitscan image from e.g. Docker Hub or a private registry  | |
| `replicaCount`                        | Number of GitScan Pods to run                                   | `1`            |
| `gitscan.debugMode`                     | The flag to enable or disable GitScan debug mode                          | `false` |
| `gitscan.gitHubToken`                   | The GitHub access token to download GitScan DB. More info: https://github.com/aquasecurity/gitscan#github-rate-limiting                          |      |
| `gitscan.registryUsername`              | The username used to log in at dockerhub. More info: https://cvedb.github.io/gitscan/dev/advanced/private-registries/docker-hub/ |      |
| `gitscan.registryPassword`              | The password used to log in at dockerhub. More info: https://cvedb.github.io/gitscan/dev/advanced/private-registries/docker-hub/ |      |
| `gitscan.registryCredentialsExistingSecret` | Name of Secret containing dockerhub credentials. Alternative to the 2 parameters above, has precedence if set.                    |      |
| `gitscan.serviceAccount.annotations`        | Additional annotations to add to the Kubernetes service account resource |     |
| `gitscan.skipUpdate`                    | The flag to enable or disable GitScan DB downloads from GitHub            | `false`        |
| `gitscan.dbRepository`                  | OCI repository to retrieve the gitscan vulnerability database from        | `ghcr.io/aquasecurity/trivy-db`        |
| `gitscan.cache.redis.enabled`           | Enable Redis as caching backend                                         | `false` |
| `gitscan.cache.redis.url`               | Specify redis connection url, e.g. redis://redis.redis.svc:6379         | `` |
| `gitscan.cache.redis.ttl`               | Specify redis TTL, e.g. 3600s or 24h                                    | `` |
| `gitscan.serverToken`                   | The token to authenticate GitScan client with GitScan server                | `` |
| `gitscan.existingSecret`                | existingSecret if an existing secret has been created outside the chart. Overrides gitHubToken, registryUsername, registryPassword, serverToken | `` |
| `gitscan.podAnnotations`                | Annotations for pods created by statefulset                             | `{}` |
| `gitscan.extraEnvVars`                  | extraEnvVars to be set on the container                                 | `{}` |
| `service.name`                        | If specified, the name used for the GitScan service                       |     |
| `service.type`                        | Kubernetes service type                                                 | `ClusterIP` |
| `service.port`                        | Kubernetes service port                                                 | `4954`      |
| `httpProxy`                           | The URL of the HTTP proxy server                                        |     |
| `httpsProxy`                          | The URL of the HTTPS proxy server                                       |     |
| `noProxy`                             | The URLs that the proxy settings do not apply to                        |     |
| `nodeSelector`                        | Node labels for pod assignment                                              |     |
| `affinity`                            | Affinity settings for pod assignment                                              |     |
| `tolerations`                         | Tolerations for pod assignment                                              |     |
| `podAnnotations`                      | Annotations for pods created by statefulset                             | `{}` |

The above parameters map to the env variables defined in [gitscan](https://github.com/aquasecurity/gitscan#configuration).

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`.

```
$ helm install my-release . \
       --namespace my-namespace \
       --set "service.port=9090" \
       --set "gitscan.vulnType=os\,library"
```

## Storage

This chart uses a PersistentVolumeClaim to reduce the number of database downloads between POD restarts or updates. The storageclass should have the reclaim policy  `Retain`.

## Caching

You can specify a Redis server as cache backend. This Redis server has to be already present. You can use the [bitnami chart](https://bitnami.com/stack/redis/helm).
More Information about the caching backends can be found [here](https://github.com/aquasecurity/gitscan#specify-cache-backend).
