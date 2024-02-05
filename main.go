package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/dihedron/osquery-extensions/plugin"
	"github.com/dihedron/osquery-extensions/plugin/snap"
	"github.com/osquery/osquery-go"
)

var tables = []*plugin.Table{
	snap.Packages,
	// add your tables here...
}

func init() {
	options := &slog.HandlerOptions{
		Level:     slog.LevelWarn,
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stderr, options)
	slog.SetDefault(slog.New(handler))
}

func main() {
	socket := flag.String("socket", "", "Path to OSquery socket file")
	timeout := flag.String("timeout", "", "OSquery timeout (for exponential backoff)")
	interval := flag.String("interval", "", "OSquery timeout (for exponential backoff)")
	verbose := flag.Bool("verbose", false, "OSquery extensions log verbosity")
	flag.Parse()

	slog.Debug("application starting", "socket", socket, "timeout", timeout, "interval", interval, "verbose", verbose)

	// start the extensions manager
	server, err := osquery.NewExtensionManagerServer("extensions", *socket)
	if err != nil {
		slog.Error("error creating extension", "error", err)
		os.Exit(1)
	}

	slog.Debug("extension manager ready")

	// register the tables
	for _, table := range tables {
		slog.Debug("registering table...", "name", table.Name)
		table.Register(server)
	}

	// run the server
	if err := server.Run(); err != nil {
		slog.Error("error running the extension manager", "error", err)
	}
}
