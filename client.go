package main

import (
  "net"
  "log"
  "os"
  "io"
)

func push(serverAddr string, localFile string, remoteFile string) {
  localFp, err := os.OpenFile(localFile, os.O_RDONLY, 0)
  if err != nil {
    log.Print("Error: ", err)
    return
  }
  defer localFp.Close()
  conn, err := net.Dial("tcp", serverAddr)
  if err != nil {
    log.Print("Error: ", err)
    return
  }
  defer conn.Close()
  SendPacket(conn, []byte(">" + remoteFile))
  data, err := ReadPacket(conn)
  if err != nil {
    log.Print("Push Error: ", err)
  } else if(!(len(data) > 0 && data[0] == '.')) {
    log.Print("Push Error: ", string(data[1:]))
  } else {
		io.Copy(conn, localFp)
		log.Print("Transfer Ended...")
	}
}

func pull(serverAddr string, remoteFile string, localFile string) {
  localFp, err := os.OpenFile(localFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0)
  if err != nil {
    log.Print("Error: ", err)
    return
  }
  defer localFp.Close()

  conn, err := net.Dial("tcp", serverAddr)
  if err != nil {
    log.Print("Error: ", err)
    return
  }
  defer conn.Close()

  SendPacket(conn, []byte("<" + remoteFile))
  data, err := ReadPacket(conn)
  if err != nil {
    log.Print("Pull Error: ", err)
  } else if(!(len(data) > 0 && data[0] == '.')) {
    log.Print("Pull Error: ", string(data[1:]))
  } else {
		io.Copy(localFp, conn)
		log.Print("Transfer Ended...")
  }
}
