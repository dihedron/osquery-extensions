package plugin

import (
	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

type Table struct {
	Name    string
	Columns func() []table.ColumnDefinition
	Data    table.GenerateFunc
}

func (t *Table) Register(server *osquery.ExtensionManagerServer) {
	server.RegisterPlugin(table.NewPlugin(t.Name, t.Columns(), t.Data))
}
