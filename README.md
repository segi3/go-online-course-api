### notes

init go hot air reload

```bash
# install
go install github.com/cosmtrek/air@latest

# init config
air init
```

change directory of `main.go` inside air conf file `.air.toml`

```
cmd = "go build -o ./tmp/main.exe ./cmd/api/"
```

install golang-migrate
```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

google wire-gen
```bash
# example
wire gen internal/oauth/injector/wire.go
```