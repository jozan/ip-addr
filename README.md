# ip-addr
A super simple tool to get local or public IP address.

#### Usage
```
$ ip -h
Usage of ip:
  -l, --local=false: Print local IP address
  -p, --public=false: Print public IP address
```

Copy public IP into clipboard (on Mac):
```
$ ip -p | pbcopy
```

#### How to install

1. Download and install [Go](https://golang.org/)
2. Clone this repo
3. Run `go build ip.go` in the cloned directory
4. Move binary to somewhere f.e. 
