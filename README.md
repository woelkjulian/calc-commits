calc-commits
===============

This repository contains a go cli tool which calculates the quotient of merge request commits and all commits of a specified gitlab project

## Install go
If not already done, install go to get started

Follow the instructions at https://golang.org/doc/install to install and setup go

## Getting started

get the code by executing the following command at $GOPATH/src

```
go get github.com/woelkjulian/calc-commits
```

or by cloning the repository into $GOPATH/src/github.com/woelkjulian

```
git clone https://github.com/woelkjulian/calc-commits.git
```
if cloned execute following command from $GOPATH/src
```
go install github.com/woelkjulian/calc-commits
```

## Usage

go to directory $GOPATH/bin

Printout help/info text by executing:

```
./calc-commits -usage

```

To start the tool you need:
- gitlab url
- gitlab private token
- project namespace/name or project id

execute:

```
// with project namespace/name
./calc-commits -url {gitlab url} -t {gitlab private token} -projname {project namespace/name}

// with project id
./calc-commits -url {gitlab url} -t {gitlab private token} -projid {project id}

// for example
./calc-commits -url https://gitlab.example.com -t abcdef123456 -projname tools/calc-commits

```

## Optional flags

```
// additional log information
-v

// select branch (default master)
-b {gitlab branchname}

// set gitlab api version (default 3)
-vapi {gitlab api version number}
```
