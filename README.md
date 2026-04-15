# MCOS CLI (`ha`)

中文文档: [`README.zh-CN.md`](./README.zh-CN.md)

Command line interface to the **Supervisor** HTTP API for **MCOS** stacks (P0: GitHub org **`muthur-command`**, module **`github.com/muthur-command/cli`**).

## Usage

- `ha help`
- `ha <subcommand> <action> [<options>]`

Example:

- `ha core info --raw-json`

### Global flags (summary)

```text
      --api-token string   Supervisor API token
      --config string      Optional config file (default is $HOME/.muthurcommand.yaml)
      --endpoint string    Endpoint for Supervisor (default is 'supervisor')
  -h, --help               help for ha
      --log-level string   Log level (defaults to Warn)
      --no-progress        Disable the progress spinner
      --raw-json           Output raw JSON from the API
```

Environment variables use the `SUPERVISOR_` prefix (e.g. `SUPERVISOR_LOG_LEVEL`).

### Subcommands

Run `ha help` for the authoritative list generated from this binary. Typical groups include **apps**, **supervisor**, **core**, **dns**, **audio**, **os**, **backups**, etc.

## Installation

On MCOS / Supervisor-managed hosts, the **`ha`** binary is usually shipped in the **CLI plugin** container or equivalent terminal access path. Release assets on this repo use names **`ha_amd64`** and **`ha_aarch64`** for compatibility with downstream image builds.

## Developing

### Remote Supervisor

For development against a live Supervisor, use a **`remote_api`**-style add-on or equivalent and set:

```shell
export SUPERVISOR_ENDPOINT=http://192.168.1.2
export SUPERVISOR_API_TOKEN=replace_this_with_token
go run main.go info
```

### Build

```bash
CGO_ENABLED=0 go build -ldflags="-s -w" -o "ha"
```

Cross-build matrix and release asset upload: [`.github/workflows/build.yml`](.github/workflows/build.yml) on branch **`mc`**.

### Contributing

1. Branch from **`mc`**.
2. Commit with clear messages.
3. `go test ./...`, `gofmt -s`.
4. Open a PR to **`muthur-command/cli`**.

## License

See **LICENSE** (Apache-2.0; retain upstream copyright / NOTICE where required).
