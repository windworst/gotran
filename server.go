package main

import (
  "net"
  "log"
  "os"
  "io"
)

func handleConnInternal(conn net.Conn) {
  defer conn.Close()
  for {
    data, err := ReadPacket(conn)
    if err != nil || len(data) <= 0 {
      return
    }
    switch data[0] {
      case '<': {
        path := string(data[1:])
        fp, err := os.OpenFile(path, os.O_RDONLY, 0)
        if err != nil {
          log.Print("[READ] '" + path + "' Failed: ", err)
          SendPacket(conn, []byte("!" + err.Error()))
          return
        }
        log.Print("[READ] '" + path + "'")
        defer fp.Close()
        SendPacket(conn, []byte("."))
        io.Copy(conn, fp)
        return
      }
      case '>': {
        path := string(data[1:])
        fp, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0)
        if err != nil {
          log.Print("[SAVE] '" + path + "' Failed: ", err)
          SendPacket(conn, []byte("!" + err.Error()))
          return
        }
        log.Print("[SAVE] '" + path + "'")
        defer fp.Close()
        SendPacket(conn, []byte("."))
        io.Copy(fp, conn)
        return
      }
      default: {
        SendPacket(conn, []byte("!" + "Invalid command"))
      }
    }
  }
}

func handleConnection(conn net.Conn) {
  log.Print("Client In: ", conn.RemoteAddr())
  handleConnInternal(conn)
  log.Print("Client Out: ", conn.RemoteAddr())
}

func server(bindAddr string) {
  ln, err := net.Listen("tcp", bindAddr)
  if err != nil {
    log.Print("Listen Error: ", err)
    return
  }
  log.Print("Listening: ", ln.Addr())
  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Print("Accept Error: ", err)
      continue
    }
    go handleConnection(conn)
  }
}
