package main

import (
	"flag"
	"log"
	"net"

	"github.com/sitano/throttle"

	"github.com/rjeczalik/netxtest"
)

func limit(l net.Listener, limitGlobal, limitPerConn int) net.Listener {
	limited := throttle.WrapListener(l)
	limited.SetCapacity(uint64(limitGlobal))
	limited.SetConnCapacity(uint64(limitPerConn))
	return limited
}

func main() {
	var test netxtest.LimitListenerTest

	test.RegisterFlags(flag.CommandLine)
	flag.Parse()

	if err := test.Run(limit); err != nil {
		log.Fatal(err)
	}
}
