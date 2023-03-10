# Embed in Dockerfile

Scan your image as part of the build process by embedding GitScan in the
Dockerfile. This approach can be used to update Dockerfiles currently using
Aqua’s [Microscanner][microscanner].

```bash
$ cat Dockerfile
FROM alpine:3.7

RUN apk add curl \
    && curl -sfL https://raw.githubusercontent.com/aquasecurity/gitscan/main/contrib/install.sh | sh -s -- -b /usr/local/bin \
    && gitscan rootfs --exit-code 1 --no-progress /

$ docker build -t vulnerable-image .
```
Alternatively you can use GitScan in a multistage build. Thus avoiding the
insecure `curl | sh`. Also the image is not changed.
```bash
[...]
# Run vulnerability scan on build image
FROM build AS vulnscan
COPY --from=aquasec/gitscan:latest /usr/local/bin/gitscan /usr/local/bin/gitscan
RUN gitscan rootfs --exit-code 1 --no-progress /
[...]
```

[microscanner]: https://github.com/aquasecurity/microscanner
