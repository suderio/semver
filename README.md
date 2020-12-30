# semver

Command-line tool for semantic versioning.
## Reference

https://semver.org/

Used regex:

```
^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$
```

## Directory layout

```bash
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── app
│   ├── app.go
│   └── app_test.go
├── config
│   ├── config.go
│   └── config_test.go
├── dev-compose.yml
├── go.mod
├── go.sum
├── main.go
├── parser
│   ├── parser.go
│   └── parser_test.go

```
