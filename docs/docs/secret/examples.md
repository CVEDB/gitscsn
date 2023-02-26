# Examples
Also see [quick start][quick-start].

## Skip Directories
GitScan traversals directories and scans all files except those matching the built-in allow rules by default.
If your have a lot of files in your container image or project, the scanning takes time.
To make it faster, you can skip traversal in the specific directory.
Also, it would be helpful if your project contains secrets and certificates for testing.

``` shell
$ gitscan image --skip-dirs /var/lib --skip-dirs /var/log YOUR_IMAGE
```

``` shell
$ gitscan fs --skip-dirs ./my-test-dir --skip-dirs ./my-testing-cert/ /path/to/your_project
```

`--skip-files` also works similarly.

## Filter by severity

Use `--severity` option.

``` shell
$ gitscan fs --severity CRITICAL ./

app/secret.sh (secrets)
=======================
Total: 1 (CRITICAL: 1)

+----------+-------------------+----------+---------+--------------------------------+
| CATEGORY |    DESCRIPTION    | SEVERITY | LINE NO |             MATCH              |
+----------+-------------------+----------+---------+--------------------------------+
|   AWS    | AWS Access Key ID | CRITICAL |   10    | export AWS_ACCESS_KEY_ID=***** |
+----------+-------------------+----------+---------+--------------------------------+
```

## Filter by RuleID

Use `.gitscanignore`.

```bash
$ cat .gitscanignore

# Ignore these rules
generic-unwanted-rule
aws-account-id
```

## Disable secret scanning
If you need vulnerability scanning only, you can disable secret scanning via the `--scanners` flag.

``` shell
$ gitscan image --scanners vuln alpine:3.15
```

## With configuration file
`gitscan-secret.yaml` in the working directory is loaded by default.

``` yaml
$ cat gitscan-secret.yaml
rules:
  - id: rule1
    category: general
    title: Generic Rule
    severity: HIGH
    regex: (?i)(?P<key>(secret))(=|:).{0,5}['"](?P<secret>[0-9a-zA-Z\-_=]{8,64})['"]
allow-rules:
  - id: social-security-number
    description: skip social security number
    regex: 219-09-9999
  - id: log-dir
    description: skip log directory
    path: ^\/var\/log\/
disable-rules:
  - slack-access-token
  - slack-web-hook
disable-allow-rules:
  - markdown

# The following command automatically loads the above configuration.
$ gitscan image YOUR_IMAGE
```

Also, you can customize the config file path via `--secret-config`.

``` yaml
$ cat ./secret-config/gitscan.yaml
rules:
  - id: rule1
    category: general
    title: Generic Rule
    severity: HIGH
    regex: (?i)(?P<key>(secret))(=|:).{0,5}['"](?P<secret>[0-9a-zA-Z\-_=]{8,64})['"]
    allow-rules:
      - id: skip-text
        description: skip text files
        path: .*\.txt
enable-builtin-rules:
  - aws-access-key-id
  - aws-account-id
  - aws-secret-access-key
disable-allow-rules:
  - usr-dirs

# Pass the above config with `--secret-config`.
$ gitscan fs --secret-config ./secret-config/gitscan.yaml /path/to/your_project
```

[quick-start]: ./scanning.md#quick-start
