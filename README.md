# github-stats

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![CircleCI](https://circleci.com/gh/worldofprasanna/github-stats.svg?style=svg)](https://circleci.com/gh/worldofprasanna/github-stats)
[![Maintainability](https://api.codeclimate.com/v1/badges/ca9aa9f54f9df2ac62b8/maintainability)](https://codeclimate.com/github/worldofprasanna/github-stats/maintainability)

> Extract info about commits from github

## Table of Contents

- [github-stats](#github-stats)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Tech specifics](#tech-specifics)
  - [Assumptions](#assumptions)
  - [Maintainers](#maintainers)
  - [Contributing](#contributing)
  - [License](#license)

## Installation 

There are 3 ways to install this command line utility,

1. Install from source code. This needs golang to be installed (> 1.11) and GOPATH to be set with properly with GOPATH added to the PATH.
```
/bin/install
```

2. If you prefer Docker,
```
# Build the docker container,
docker build -t github-stats .

# To test the functionality,
a. Fetch the active day of week along with average commit
docker run github-stats activeDay --weeks=20 kubernetes/kubernetes

b. List the average commit for week
docker run github-stats listAverageCommits kubernetes/kubernetes
```

3. Github releases contains binaries targeting difference operating systems. This is generated using [goreleaser](https://github.com/goreleaser/goreleaser)

[Download Binary here](https://github.com/worldofprasanna/github-stats/releases/tag/v1.0.0)


## Usage

```
# To know about the command, use --help option. After installation is successful,
./bin/run --help
./bin/run activeDay --weeks=20 kubernetes/kubernetes
./bin/run listAverageCommits --sort=desc kubernetes/kubernetes

```
```
# To run the unit test
./bin/test

# To run the linter
./bin/lint

# To get the binary
./bin/build

```

## Tech specifics

- This uses CircleCI to run unit tests, linters, build and publish the package. You can see the pipelines [here](https://circleci.com/workflow-run/3e861475-a6e1-46de-a664-5395783c92c9)
- Code climate integration is done to see if there are any code smells
- This repo depends on Go modules and hence it needs golang > 1.11

## Assumptions

- Number of weeks would be less than or equal to 52
- Listing the average commits per day can be a separate sub command
- When listing the average commits per day, it will always fetch for last 52 weeks

## Maintainers

[@worldofprasanna](https://github.com/worldofprasanna)

## Contributing

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © 2019 Prasanna
