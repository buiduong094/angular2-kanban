# Setting up
```
$ cd $GOPATH
$ mkdir -p src/github.com/octoberstorm && cd $_
$ git clone https://github.com/yourname/angular2-kanban.git
$ go get -u github.com/tools/godep
```

# Add dependencies
```
$ go get some-golang-dep
$ godep save (or godep save ./...)
```

# Build & Run Server
```
$ go build
$ ./angular2-kanban
```

