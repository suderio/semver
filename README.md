# semver

Command-line tool for semantic versioning.

This is just a toy project to start learning go in a more idiomatic way.

## Usage
```bash
semver 1.2.3
1.2.3

semver up 1.2.3
1.2.4

semver set 1.2.3 --minor 3
1.3.0

semver 1.2.3 --verbose
Major: 1 Minor: 2 Patch: 3
```

## References

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
├── dev-compose.yml
├── go.mod
├── go.sum
├── main.go
├── parser
│   ├── parser.go
│   └── parser_test.go

```
