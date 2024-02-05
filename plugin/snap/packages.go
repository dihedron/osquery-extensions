package snap

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/dihedron/osquery-extensions/plugin"
	"github.com/osquery/osquery-go/plugin/table"
)

var Packages = &plugin.Table{
	Name: "snap_packages",
	Columns: func() []table.ColumnDefinition {
		return []table.ColumnDefinition{
			table.TextColumn("name"),
			table.TextColumn("version"),
			table.TextColumn("revision"),
			table.TextColumn("tracking"),
			table.TextColumn("publisher"),
			table.TextColumn("notes"),
		}
	},
	Data: listPackages,
}

// listPackages will be called whenever the table is queried; it should return
// a full table scan.
func listPackages(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	cmd := exec.Command("/usr/bin/snap", "list", "--all")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s:\n%s\n", err, stderr.String())
	}

	result := []map[string]string{}

	scanner := bufio.NewScanner(bytes.NewReader(stdout.Bytes()))
	for scanner.Scan() {
		tokens := []string{}
		for _, token := range strings.Split(scanner.Text(), " ") {
			if len(strings.TrimSpace(token)) > 0 {
				tokens = append(tokens, strings.TrimSpace(token))
			}
		}
		fmt.Printf("line: %v\n", strings.Join(tokens, ","))
		result = append(result, map[string]string{
			"name":      tokens[0],
			"version":   tokens[1],
			"revision":  tokens[2],
			"tracking":  tokens[3],
			"publisher": tokens[4],
			"notes":     tokens[5],
		})
	}

	return result, nil
}
