# Bamboo

[![Build Status](https://travis-ci.com/dapperlabs/bamboo-node.svg?token=MYJ5scBoBxhZRGvDecen&branch=master)](https://travis-ci.com/dapperlabs/bamboo-node)

Bamboo is a highly-performant blockchain designed to power the next generation of decentralized applications.

## Getting started

* Read through the [project setup](/docs/setup.md) instructions to install required tools
* Read the documentation pertaining to [your stream](#work-streams)
* Familiarize yourself with the [workflow](#workflow) below
* Browse the rest of this README to get up to speed on concepts like testing, code style, and common code patterns
* Contact your stream owner to receive your first task!

## Documentation

You can find a high-level overview of the Bamboo architecture on the [documentation website](https://bamboo-docs.herokuapp.com/). Application-level documentation lives [within the packages of this repository](#code-documentation).

## Installation

### Setting up your environment

#### Install Go
- Download and install [Go 1.12](https://golang.org/doc/install)
- Create your workspace `$GOPATH` directory and update your bash_profile to contain the following:

```bash
export `$GOPATH=$HOME/path-to-your-go-workspace/`
```

It's also a good idea to update your `$PATH` to use third party GO binaries: 

```bash
export PATH="$PATH:$GOPATH/bin"
```

- Test that Go was installed correctly: https://golang.org/doc/install#testing
- Clone this repository to `$GOPATH/src/github.com/dapperlabs/bamboo-node/`

_Note: since we are using go modules and we prepend every `go` command with `GO111MODULE=on`, you can also clone this repo anywhere you want._

#### Install Docker
- Download and install [Docker CE](https://docs.docker.com/install/)
- Test Docker by running the integration tests for this repository:
```bash
./test.sh
```

The first run will take a while because some base layers will be downloaded and built for the first time. See our [testing instructions](#testing) for more details.

### Building binaries

This project includes several binaries defined in the `/cmd` directory:

```
$ GO111MODULE=on go build -o donotcommit ./cmd/execute/
$ GO111MODULE=on go build -o donotcommit ./cmd/security/
$ GO111MODULE=on go build -o donotcommit ./cmd/testhelpers/
```

TODO: move to Makefile

### Generating code

#### Dependency injection using Wire

Install wire:

```bash
GO111MODULE=on go get -u github.com/google/wire/cmd/wire
```

```
$ GO111MODULE=on wire ./internal/execute/
$ GO111MODULE=on wire ./internal/security/
$ GO111MODULE=on wire ./internal/access/
```
TODO: move to Makefile

#### Generate gRPC stubs from protobuf files

1. Install prototool https://github.com/uber/prototool#installation  
2. `go get -u github.com/golang/protobuf/protoc-gen-go`

```
$ prototool generate proto/
```
TODO: move to Makefile

## Testing

Run:

```bash
./test.sh
```

If iterating just on failed test, then we can do so without rebuilding the system:

```bash
docker-compose up --build --no-deps test
```

Cleanup:

```bash
docker-compose down
```

TODO: move to Makefile (remove also shell script)


## Contributing

This guide provides a comprehensive overview of our development processes, guidelines, and steps required to contribute to the project.

### Work streams

Bamboo development is divided across several streams of work with the goal of separating concerns and facilitating rapid development. 

Each stream is owned by a Bamboo core team member who oversees and directs all development within that stream. As a contributor, you will communicate primarily with your stream owner.

Stream owners will assign tasks to contributors and ensure that all TODOs are tracked.

| Stream         | Owner(s)                    | Home directory  |
| -------------- | --------------------------- | --------- |
| Collection  | [Peter Siemens](https://github.com/psiemens]) | [/internal/protocol/collect](/internal/protocol/collect) |
| Consensus | [Alexander Hentschel](https://github.com/AlexHentschel) | [/internal/protocol/consensus](/internal/protocol/consensus) |
| Execution      | [Bastian Müller](https://github.com/turbolent) | [/internal/protocol/execute](/internal/protocol/execute) |
| Verifcation | [Moar Zamski](https://github.com/pazams) | [/internal/protocol/verify](/internal/protocol/verify) |
| Sealing | [Moar Zamski](https://github.com/pazams) | [/internal/protocol/seal](/internal/protocol/seal) |
| Networking | [Yahya Hassanzadeh](https://github.com/yhassanzadeh)     | [/pkg/network](/pkg/network) |
| Cryptography | [Tarak Ben Youssef](https://github.com/tarakby)     | [/pkg/crypto](/pkg/crypto) |
| Emulator | [Brian Ho](https://github.com/mrbrianhobo), [Peter Siemens](https://github.com/psiemens)     | [/internal/emulator](/internal/emulator) |
| Client Library | [Brian Ho](https://github.com/mrbrianhobo), [Peter Siemens](https://github.com/psiemens)     | [/client](/client), [/internal/cli](/internal/cli), [/cmd/bamboo](/cmd/bamboo) |
| Observation | [Peter Siemens](https://github.com/psiemens)     | [/internal/protocol/observe](/internal/protocol/observe) |
| Ops & Performance | [Timofey Smirnov](https://github.com/tsmirnov) | |
| Language & Runtime | [Bastian Müller](https://github.com/turbolent) | [/language](/language) |

### Workflow

### Issues

Development tasks are assigned using GitHub issues. Each issue will contain a breakdown of the required task and any necessary background information, as well as an esitmate of the required work. You are expected to track the progress of issues assigned to you and provide updates if needed, in the form of issue comments.

If you need to create a new issue, please use the provided issue templates to ensure that all necessary information is included.

#### Branches

Work for a specific task should be completed in a separate branch corresponding to the issue for that task.

When creating a new branch, use the following convention: `<your-name>/<issue-number>-<issue-description>`

For example, `peter/125-update-transaction` is the name of a branch Peter is working on, and corresponds to issue 125 regarding transaction updates.

##### Feature Branches

When working on a larger feature, feel free to create a feature branch with the following format: `feature/<feature-name>`.

#### Pull Requests

You should open a pull request when you have completed work for a task and would like to receive a review from teammates and stream owners. Please use the provided pull request template when opening a PR.

##### Reviews

You should request a review from any relevant team members who are also working within your stream. The stream owner will automatically be requested for review.

A PR can be merged once all CI checks pass and it is approved by at least two people, including the stream owner.

If you are reviewing another team member's PR, please keep feedback constructive and friendly.

##### Work-In-Progress PRs

You can open a WIP pull request to track ongoing work for a task.

#### Testing

Each PR that you open should include necessary tests to ensure the correctness and stability of your code. The specific testing requirements for each task will be defined in the issue itself.

### Code standards

The Bamboo project has a high standard for code quality and expects all submitted PRs to meet the guidelines outlined in our [code style guide](code-style.md).

TODO: add style guide

### Code documentation

The application-level documentation for Bamboo lives inside each of the sub-packages of this repository.

#### Documentation instructions for stream owners

Stream owners are responsible for ensuring that all code owned by their stream is well-documented. Documentation for a stream should accomplish the following:

1. Provide an overview of all stream functions
2. Outline the different packages used by the stream
3. Highlight dependencies on other streams

Each stream should contain a README in its home directory. This page, which acts as a jumping-off point for new contributors, should list each function of the stream along with a short description and links to relevant packages.

Here's an example: [internal/protocol/collect/README.md](internal/protocol/collect/README.md)

#### Stream package documentation 

All packages owned by a stream should be documented using `godoc`.

Here's an example: [internal/protocol/collect/clusters](internal/protocol/collect/clusters)

##### Auto-generated READMEs

A `README.md` can be generated from the `godoc` output by updating [godoc.sh](/godoc.sh) with the path of your package. The above example was generated by this line:

```bash
godoc2md github.com/dapperlabs/bamboo-node/internal/protocol/collect/clusters > internal/protocol/collect/clusters/README.md
```

Once your package is added to that file, running `go generate` in the root of this repo will generate a new `README.md`.

#### Documentation instructions for contributors

TODO: describe documentation standards for all code
