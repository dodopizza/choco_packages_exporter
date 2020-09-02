# Continuous Integration and Release Process

This project has adopted [GitHub Flow](https://guides.github.com/introduction/flow/index.html) for development lifecycle.

Application versioning: https://semver.org/

Also Continuous Integration (CI) and some routine actions are achieved using [GitHub Actions](https://github.com/features/actions).

## Workflows

There are several workflows to react on different GitHub events:

- [Continuous Integration and Pull Requests](./on-push-and-pull-request.yml)
  - _Purpose_: Run applicatios unit tests to ensure that changes doesn't broke anything.
  - _Run conditions_: Runs on every `push` and `pull request` event to any branch except `master`.

- [Build draft-release](./on-push-to-master.yml)
  - _Purpose_: Run unit tests and build application on `master` branch and create draft for the release.
  - _Run conditions_: Runs on every `push` event to `master` branch.

## How to publish new release

1. On every `push` event to `master` branch there is created draft for the future release (automated with `Build draft-release` workflow);
2. Check the `Build draft-release` workflow file and change `appVersion` env for major or minor releases if it is necessary;
3. You have to check release notes in the release draft. It is good practices to describe all changes in the release and add links to the issues for each change.
4. Publish the release
