# Installing Go on Mac

With [brew](https://brew.sh/) installed, open your terminal and run the following commands

* `brew update`
* `brew install golang`

# Setup the workspace

Go has a different approach of managing code, you'll need to create a [single Workspace for all your Go projects](https://golang.org/doc/code.html#Workspaces).

First, you'll need to tell Go the location of your workspace by adding some environmental variables to your shell config (`bash_profile`, `bashrc`, `zshrc`, etc).

* `vi .bashrc`

And add the following to your config

```
export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

# Run Project

* `go run main.go`
