package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os/exec"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

func main() {
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		//log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
		SnapPackagesGenerate(context.Background(), table.QueryContext{})
	}

	server, err := osquery.NewExtensionManagerServer("snaps", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	// Create and register a new table plugin with the server.
	// table.NewPlugin requires the table plugin name,
	// a slice of Columns and a Generate function.
	server.RegisterPlugin(table.NewPlugin("snap_packages", SnapPackagesColumns(), SnapPackagesGenerate))
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

// FoobarColumns returns the columns that our table will return.
func SnapPackagesColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("name"),
		table.TextColumn("version"),
		table.TextColumn("revision"),
		table.TextColumn("tracking"),
		table.TextColumn("publisher"),
		table.TextColumn("notes"),
	}
}

// SnapPackagesGenerate will be called whenever the table is queried. It should return
// a full table scan.
func SnapPackagesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {

	cmd := exec.Command("/usr/bin/snap", "list", "--all")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s:\n%s\n", err, stderr.String())
	}

	scanner := bufio.NewScanner(bytes.NewReader(stdout.Bytes()))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//strings.Split(scanner.Text())
	}

	return []map[string]string{
		{
			"foo": "bar",
			"baz": "baz",
		},
		{
			"foo": "bar",
			"baz": "baz",
		},
	}, nil
}
