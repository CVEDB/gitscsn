# CircleCI

```
$ cat .circleci/config.yml
jobs:
  build:
    docker:
      - image: docker:stable-git
    steps:
      - checkout
      - setup_remote_docker
      - run:
          name: Build image
          command: docker build -t gitscan-ci-test:${CIRCLE_SHA1} .
      - run:
          name: Install gitscan
          command: |
            apk add --update-cache --upgrade curl
            curl -sfL https://raw.githubusercontent.com/aquasecurity/gitscan/main/contrib/install.sh | sh -s -- -b /usr/local/bin
      - run:
          name: Scan the local image with gitscan
          command: gitscan image --exit-code 0 --no-progress gitscan-ci-test:${CIRCLE_SHA1}
workflows:
  version: 2
  release:
    jobs:
      - build
```

[Example][example]
[Repository][repository]

[example]: https://circleci.com/gh/aquasecurity/gitscan-ci-test
[repository]: https://github.com/aquasecurity/gitscan-ci-test
