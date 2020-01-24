# GoTran
Go语言编写的文件传输Demo

### 运行 ###
```
$ go build
$ ./gotran
        Simple File Transfer in Golang
Usage:
  gotran server  <ListenPort>                                              # As Server
  gotran push    <ServerAddr> <LocalFilePath>    <RemoteFilePath>          # Upload File
  gotran pull    <ServerAddr> <RemotePFilePath>  <LocalFilePath>           # Down File
$
```

### 服务器 ###
```
$ gotran server 8888
2020/01/24 21:31:23 Listening: [::]:8888
```

### 客户端 ###
```
# 此处 123.456.78.90 指代运行服务器的IP地址

# 上传
$ gotran push 123.456.78.90:8888 local_file.zip remote_file.zip
Push: local_file.zip -> remote_file.zip
2020/01/24 21:46:27 Transfer Ended...

# 下载
gotran.exe pull 123.456.78.90:8888 remote_file.zip local_file.zip
Pull: remote_file.zip -> local_file.zip
2020/01/24 21:46:27 Transfer Ended...
```
