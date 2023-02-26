# CI/CD Integrations

## GitHub Actions
[GitHub Actions](https://github.com/features/actions) is GitHub's native CI/CD and job orchestration service.

### gitscan-action (Official)

GitHub Action for integrating GitScan into your GitHub pipeline

👉 Get it at: <https://github.com/aquasecurity/gitscan-action>

### gitscan-action (Community)

GitHub Action to scan vulnerability using GitScan. If vulnerabilities are found by GitScan, it creates a GitHub Issue.

👉 Get it at: <https://github.com/marketplace/actions/gitscan-action>

### gitscan-github-issues (Community)

In this action, GitScan scans the dependency files such as package-lock.json and go.sum in your repository, then create GitHub issues according to the result.

👉 Get it at: <https://github.com/marketplace/actions/gitscan-github-issues>

## Azure DevOps (Official)
[Azure Devops](https://azure.microsoft.com/en-us/products/devops/#overview) is Microsoft Azure cloud native CI/CD service.

GitScan has a "Azure Devops Pipelines Task" for GitScan, that lets you easily introduce security scanning into your workflow, with an integrated Azure Devops UI.

👉 Get it at: <https://github.com/aquasecurity/gitscan-azure-pipelines-task>

## Semaphore (Community)
[Semaphore](https://semaphoreci.com/) is a CI/CD service.

You can use GitScan in Semaphore for scanning code, containers, infrastructure, and Kubernetes in Semaphore workflow.

👉 Get it at: <https://semaphoreci.com/blog/continuous-container-vulnerability-testing-with-gitscan>

## CircleCI (Community)
[CircleCI](https://circleci.com/) is a CI/CD service.

You can use the GitScan Orb for Circle CI to introduce security scanning into your workflow.

👉 Get it at: <https://circleci.com/developer/orbs/orb/fifteen5/gitscan-orb>
Source: <https://github.com/15five/gitscan-orb>

## Woodpecker CI (Community)

Example GitScan step in pipeline

```yml
pipeline:
  securitycheck:
    image: aquasec/gitscan:latest
    commands:
      # use any gitscan command, if exit code is 0 woodpecker marks it as passed, else it assumes it failed
      - gitscan fs --exit-code 1 --skip-dirs web/ --skip-dirs docs/ --severity MEDIUM,HIGH,CRITICAL .
```

Woodpecker does use GitScan itself so you can [see it in use there](https://github.com/woodpecker-ci/woodpecker/pull/1163).

## Concourse CI (Community)
[Concourse CI](https://concourse-ci.org/) is a CI/CD service.

You can use GitScan Resource in Concourse for scanning containers and introducing security scanning into your workflow.
It has capabilities to fail the pipeline, create issues, alert communication channels (using respective resources) based on GitScan scan output.

👉 Get it at: <https://github.com/Comcast/gitscan-resource/>
