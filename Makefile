BIN = $(GOPATH)/bin
BASE = $(GOPATH)/src/$(PACKAGE)
PKGS = go list ./... | grep -v "^vendor/"

default: build

#Tools
DEP = $(BIN)/dep
$(BIN)/dep:
	go get -u github.com/golang/dep/cmd/dep

GINKGO = $(BIN)/ginkgo
$(BIN)/ginkgo:
	go get -u github.com/onsi/ginkgo/ginkgo

MIGRATE = $(BIN)/migrate
$(BIN)/migrate:
	go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
	go build -tags 'postgres' -o $(BIN)/migrate github.com/mattes/migrate/cli

LINT = $(BIN)/golint
$(BIN)/golint:
	go get github.com/golang/lint/golint

METALINT = $(BIN)/gometalinter.v2
$(BIN)/gometalinter.v2:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	$(METALINT) --install

.PHONY: installtools
installtools: | $(LINT) $(MIGRATE) $(GINKGO) $(DEP)
	echo "Installing tools"

.PHONY: metalint
metalint: | $(METALINT)
	$(METALINT) ./... --vendor \
	--fast \
	--exclude="exported (function)|(var)|(method)|(type).*should have comment or be unexported" \
	--format="{{.Path.Abs}}:{{.Line}}:{{if .Col}}{{.Col}}{{end}}:{{.Severity}}: {{.Message}} ({{.Linter}})"

.PHONY: lint
lint:
	$(LINT) $$($(PKGS)) | grep -v -E "exported (function)|(var)|(method)|(type).*should have comment or be unexported"

.PHONY: test
test: | $(GINKGO) $(LINT)
	go vet ./...
	go fmt ./...
	$(GINKGO) -r

.PHONY: dep
dep: | $(DEP)
	$(DEP) ensure

build: dep fixlibcrypto
	go fmt ./...
	go build

.PHONY: setup
setup: checkdbvars
	curl https://raw.githubusercontent.com/vulcanize/vulcanizedb/master/db/schema.sql > vulcanize_schema.sql
	createdb $(NAME)
	psql $(NAME) < vulcanize_schema.sql

# have to copy over crypto related headers https://github.com/ethereum/go-ethereum/issues/2738
.PHONY: fixlibcrypto
fixlibcrypto:
	mkdir tmp-go-ethereum
	git clone https://github.com/ethereum/go-ethereum.git tmp-go-ethereum
	cp -r "tmp-go-ethereum/crypto/secp256k1/libsecp256k1" "vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/"
	rm -rf "tmp-go-ethereum"

#Database
HOST_NAME = localhost
PORT = 5432
NAME =
CONNECT_STRING=postgresql://$(HOST_NAME):$(PORT)/$(NAME)?sslmode=disable

.PHONY: checkdbvars
checkdbvars:
	test -n "$(HOST_NAME)" # $$HOST_NAME 
	test -n "$(PORT)" # $$PORT
	test -n "$(NAME)" # $$NAME
	@echo $(CONNECT_STRING)

.PHONY: rollback
rollback: $(MIGRATE) checkdbvars
	$(MIGRATE) -database $(CONNECT_STRING) -path ./db/migrations down 1
	pg_dump -O -s $(CONNECT_STRING) > db/schema.sql

.PHONY: migrate
migrate: $(MIGRATE) checkdbvars
	$(MIGRATE) -database $(CONNECT_STRING) -path ./db/migrations up
	pg_dump -O -s $(CONNECT_STRING) > db/schema.sql

.PHONY: import
import:
	test -n "$(NAME)" # $$NAME
	psql $(NAME) < db/schema.sql
