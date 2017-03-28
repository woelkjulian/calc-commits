calc-commits
===============

This repository contains a cli tool which calculates the quotient of merge request commits and all commits of a specified gitlab project

## Getting started
get the code by executing the following command at $GOPATH

```
go get github.com/woelkjulian/calc-commits 
```

or by cloning the repository into $GOPATH/src/github.com/woelkjulian

```
git clone https://github.com/woelkjulian/calc-commits.git
```
if cloned execute following command from $GOPATH
```
go install github.com/woelkjulian/calc-commits
```

## Usage

go to directory $GOPATH/bin and execute

```
./calc-commits -url {gitlab url} -t {gitlab privtate token} -proj {gitlab project id}
```

# Optional flags

additional log information
```
-v
```
select branch
```
-b {gitlab branchname}
```


