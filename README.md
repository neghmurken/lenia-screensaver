# xscreensavers #

Lenia, a continuous cellular automata, implemented as a screensaver for X with Go and OpenGL

## Requirements

 - inotify-tools (optional)

## How to dev

Build the binary with

```shell
$ make build
```

This will produce a `lenia` binary in the `/bin` folder

Use the project watcher to automatically rebuild the binary if a go file is modified

```shell
$ make watch
```

## How to run

```shell
$ make run

... or

$ /bin/lenia
```
