# wake

## Install

Download the version for your operating system on the [Releases page here on GitHub](https://github.com/pheebcodes/wake/releases) or [install Go](https://golang.org/) and run the below command to install `wake` using `go get`.

```bash
$ go get github.com/pheebcodes/wake/cmd/wake
```

## Configure

Write any number of configuration files to `~/.wake/[name of computer].toml`.

```toml
ip_address = "192.168.1.255" # IP address you want to broadcast to. Required!
mac_address = "01:23:34:56:78:90" # MAC address of the computer you want to wake. Required!
port = 2 # The port you want to send the packet on. Optional, defaults to 9.
```

## Run

```bash
$ wake [name of computer you want to wake]
```
