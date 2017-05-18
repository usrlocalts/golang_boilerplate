.PHONY: all
all: build fmt vet lint test

APP=golang_boilerplate
DB_USER=boilerplate
GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
UNIT_TEST_PACKAGES=$(shell glide novendor | grep -v "featuretests")
DB_NAME=$(APP)_dev
TEST_DB_NAME="$(APP)_test"
TEST_DB_URL_VAR="DB_URL= dbname=$(TEST_DB_NAME) user=$(DB_USER) password='' host=localhost sslmode=disable"

APP_EXECUTABLE="./out/$(APP)"

setup:
	go get -u github.com/golang/lint/golint
	go get github.com/DATA-DOG/godog/cmd/godog
	glide install
	createuser -s $(DB_USER) || true
	cp application.yml.sample application.yml
	@echo "Created application.yml. Please modify values as needed"
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME) || true
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)
	ENVIRONMENT=development $(APP_EXECUTABLE) migrate
	@echo "Golang Boilerplate is setup!! Run make test to run tests"

build-deps:
	glide install

update-deps:
	glide update

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

build: build-deps compile fmt vet lint

install:
	go install ./...

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

test: compile testdb.drop testdb.create testdb.migrate
	ENVIRONMENT=test go test $(UNIT_TEST_PACKAGES) -p=1

featuretest: testdb.drop testdb.create testdb.migrate
	cd featuretests && ENVIRONMENT=test godog

db.setup: db.create db.migrate

db.create:
	createdb -O$(DB_USER) -Eutf8 $(DB_NAME)

db.migrate:
	ENVIRONMENT=development $(APP_EXECUTABLE) migrate

db.rollback:
	$(APP_EXECUTABLE) rollback

db.drop:
	dropdb --if-exists -U$(DB_USER) $(DB_NAME)

db.reset: db.drop db.create db.migrate

db.create-migration:
	migrate -url "postgres://$(DB_USER)@localhost:5432/$(DB_NAME)?sslmode=disable" -path ./migrations create $(MIGRATION_NAME)

testdb.create: testdb.drop
	createdb -O$(DB_USER) -Eutf8 $(TEST_DB_NAME)

testdb.migrate:
	DB_URL="postgres://$(DB_USER)@localhost:5432/$(TEST_DB_NAME)?sslmode=disable" "./out/golang_boilerplate" migrate $(APP_EXECUTABLE) migrate

testdb.drop:
	dropdb --if-exists -U$(DB_USER) $(TEST_DB_NAME)

test-coverage: compile testdb.drop testdb.create testdb.migrate
	@echo "mode: count" > out/coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=out/coverage.out -covermode=count $(pkg);\
	tail -n +2 out/coverage.out >> out/coverage-all.out;)
	go tool cover -html=out/coverage-all.out -o out/coverage.html

copy-config:
	cp application.yml.sample application.yml
