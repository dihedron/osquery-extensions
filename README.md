# osquery-extensions

This package uses `github.com/osquery/osquery-go` bindings to create a set of virtual tables for OSQuery.

## How to install

In order to install the extension, run the following commands:

```bash
$> make && sudo make install
```

`make` will compile the extension for the `linux/amd64` target; if your architecture is different, you can run `make <os>/<arch>` for one of the supported operating system and architecture combinations supported by the Go compiler (see `go tool dist list` for the list of supported combinations).

`sudo make install` will install the extension under `/usr/lib/osquery/extensions` and set it permissions as expected by the OSQuery plugin system (see [the relevant documentation](https://osquery.readthedocs.io/en/stable/deployment/extensions/#auto-loading-extensions)); moreover it will register the extension under `/etc/osquery/extensions.load` so it can be autoloaded by `osqueryd`.

In order to **uninstall** the extension, simply run 

```bash
$> sudo make uninstall
```

It will undo the install process by removing the binary from the extensions directory and unregistering it from the auto-load configuration file.

## How to add more tables

Each table is defined by the following `struct`:

```golang
type Table struct {
	Name    string
	Columns func() []table.ColumnDefinition
	Data    table.GenerateFunc
}
```

where the `Columns` function returns the list of columns in the table, and `Data` returns the list of records (each as a `map[string]string`).

In order to register the table, add it it the following variable in `main.go`:

```golang
var tables = []*plugin.Table{
	snap.Packages,
	// add your tables here...
}
```


See `plugin/snap/Packages` for an example.