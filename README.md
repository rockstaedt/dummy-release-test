# dummy-release-test

[![build](https://github.com/rockstaedt/dummy-release-test/actions/workflows/CI.yml/badge.svg)](https://github.com/rockstaedt/dummy-release-test/actions/workflows/CI.yml)
[![codecov](https://codecov.io/gh/rockstaedt/dummy-release-test/branch/main/graph/badge.svg?token=VW245SMVP5)](https://codecov.io/gh/rockstaedt/dummy-release-test)
[![Latest tag](https://img.shields.io/github/v/tag/rockstaedt/dummy-release-test)](https://github.com/rockstaedt/dummy-release-test/releases)

This is just a repo to play with release workflows.

## Outcome

This repo has the following workflows:
1. If a push is made, the CI workflow is triggered to test and build the project.
2. If a pull request is opened, all labels from the linked issue are copied.
3. If a pull request is merged, the release type is set depending on the labels:
   - bug/patch label initiates a patch release
   - feature/minor label initiates a minor release
4. If a release was created, the binary workflow is triggered and attaches all binaries to the release.

## Learnings

- A workflow cannot be triggered by another workflow. This is on purpose to prevent infinite workflow loops.
  - Can be avoided by using a personal access token.
  - Always make sure that no loop occurs.

