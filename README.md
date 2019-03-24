Tarpit command application that may slow down malicious attempts to scan a system.

Supported Protocols:

- SSH

# Building and Running

```
$ make linux-amd64
GOOS=linux GOARCH=amd64 go build -o "build/tarpit-linux-amd64"
```

This command will build the application in `build/tarpit-linux-amd64`. We can then run the tarpit on an unprivileged port (e.g. 2222).

`$ ./build/tarpit-linux-amd64 -p 2222`

Or you can run it on a privileged port with `sudo`.

`$ sudo ./build/tarpit-linux-amd64 -p 22`

## Options

```
Usage of ./tarpit:
  -b, --bind-address string   address to bind the socket to
  -d, --delay string          delay between the tarpit keep-alive data packets (default "10s")
  -g, --gid uint16            setgid, after creating a listening socket
  -p, --port uint16           TCP port (default 22)
  -P, --proto string          protocol to tarpit (default "ssh")
  -u, --uid uint16            setuid, after creating a listening socket
  -v, --version               show current version
```

## Using privileged ports

The ports `< 1024` require superuser privileges. The command allows to drop superuser privileges (using setuid/setgid), right after it acquires a listening socket. Thus, allowing to bind to a privileged port and start serving as a regular unprivileged user. This is done by running the command as a superuser (e.g. with `sudo`) and setting the `-u/--uid <uid>` and `-g/--gid <gid>` command line flags.

`$ sudo ./build/tarpit-linux-amd64 -p 22 -u "$(id -u)" -g "$(id -g)"`

# Acknowledgements

Thanks to nullprogram.com / @skeeto for the [article about tarpits](https://nullprogram.com/blog/2019/03/22/) that served as an inspiration for this application.
