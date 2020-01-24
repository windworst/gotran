package main

import (
  "fmt"
  "os"
)

func help() {
  fmt.Printf(
`        Simple File Transfer in Golang
Usage:
  %[1]s server  <ListenPort>                                              # As Server
  %[1]s push    <ServerAddr> <LocalFilePath>    <RemoteFilePath>          # Upload File
  %[1]s pull    <ServerAddr> <RemotePFilePath>  <LocalFilePath>           # Down File
`, os.Args[0])
}

func main() {
  args := os.Args[1:]
  if len(args) < 1 {
    help()
    return
	}
  action := args[0]
  switch action {
  case "server": {
    if len(args) <= 1 {
      fmt.Println("Action '" + action + "' missing args: need specify <port>")
      return
    }
    server("0.0.0.0:" + args[1])
  }
  case "push": {
    if len(args) <= 3 {
      fmt.Println("Action '" + action + "' missing args: need <ServerAddr> <LocalFilePath> <RemoteFilePath>")
      return
		}
		fmt.Printf("Push: %s -> %s\n", args[2], args[3])
    push(args[1], args[2], args[3])
  }
  case "pull": {
    if len(args) <= 3 {
      fmt.Println("Action '" + action + "' missing args: need <ServerAddr> <RemoteFilePath>  <LocalFilePath>")
      return
		}
		fmt.Printf("Pull: %s -> %s\n", args[2], args[3])
    pull(args[1], args[2], args[3])
  }
	default:
    fmt.Println("Action '" + action + "' not defined")
  }
}
