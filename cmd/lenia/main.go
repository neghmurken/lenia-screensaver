package main

import (
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)

    <-sigc

    os.Exit(0)
}
