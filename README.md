# github-stats

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> Extract info about commits from github

## Table of Contents

- [github-stats](#github-stats)
  - [Table of Contents](#table-of-contents)
  - [Usage](#usage)
  - [Install the binary in GOPATH/bin](#install-the-binary-in-gopathbin)
  - [Assumptions](#assumptions)
  - [Maintainers](#maintainers)
  - [Contributing](#contributing)
  - [License](#license)

## Usage

```
# To run the unit test
./bin/test

# To run the linter
./bin/lint

```
```
# To know about the command, use --help option
./bin/run --help
./bin/run activeDay --weeks=20 kubernetes/kubernetes
./bin/run listAverageCommits --sort=desc kubernetes/kubernetes

```
```
# Docker Setup
docker build -t github-stats .

# To test the functionality

1. Fetch the active day of week along with average commit
docker run github-stats activeDay --weeks=20 kubernetes/kubernetes

2. List the average commit for week
docker run github-stats listAverageCommits kubernetes/kubernetes

```
## Install the binary in GOPATH/bin

```
/bin/install
```

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
