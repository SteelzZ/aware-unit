package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go func() {
		defer func() {
			if err := panicRecover(); err != nil {
				log.Panic("Nexus err: ", err)
			}
		}()

		// My awesome app
		for {
			fmt.Print(".")
			time.Sleep(time.Second)
		}
	}()

	signals := map[os.Signal]func(){
		syscall.SIGTERM: shutdown,
		syscall.SIGINT:  shutdown,
	}

	listenForSignals(signals)
}

func listenForSignals(sigmap map[os.Signal]func()) {
	sigchan := make(chan os.Signal, 1)

	for k := range sigmap {
		signal.Notify(sigchan, k)
	}

	for true {
		sig := <-sigchan
		handler, ok := sigmap[sig]
		if ok {
			handler()
		}
	}
}
func shutdown() {
	log.Print("Gracefully going down")
	os.Exit(0)
}

func panicRecover() error {
	if r := recover(); r != nil {
		var ok bool
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("pkg: %v", r)
		}
		return err
	}

	return nil
}
