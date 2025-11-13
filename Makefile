build:
	go build -o bin/main ./cmd/

link:
	ln -sf $(PWD)/bin/main ~/.local/bin/sqlchelper

