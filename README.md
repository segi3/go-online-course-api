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
```bash
# install
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# migrate
migrate -database "mysql://root@tcp(localhost:3306)/go_online_course" -path database/migrations/ up
```

google wire-gen
```bash
# install
go get github.com/google/wire/cmd/wire@latest

# example
wire gen internal/oauth/injector/wire.go
```