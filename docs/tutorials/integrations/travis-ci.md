# Travis CI

```
$ cat .travis.yml
services:
  - docker

env:
  global:
    - COMMIT=${TRAVIS_COMMIT::8}

before_install:
  - docker build -t gitscan-ci-test:${COMMIT} .
  - export VERSION=$(curl --silent "https://api.github.com/repos/aquasecurity/gitscan/releases/latest" | grep '"tag_name":' | sed -E 's/.*"v([^"]+)".*/\1/')
  - wget https://github.com/aquasecurity/gitscan/releases/download/v${VERSION}/gitscan_${VERSION}_Linux-64bit.tar.gz
  - tar zxvf gitscan_${VERSION}_Linux-64bit.tar.gz
script:
  - ./gitscan image --exit-code 0 --severity HIGH --no-progress gitscan-ci-test:${COMMIT}
  - ./gitscan image --exit-code 1 --severity CRITICAL --no-progress gitscan-ci-test:${COMMIT}
cache:
  directories:
    - $HOME/.cache/gitscan
```

[Example][example]
[Repository][repository]

[example]: https://travis-ci.org/aquasecurity/gitscan-ci-test
[repository]: https://github.com/aquasecurity/gitscan-ci-test
