if [ $1 = "start" ]; then
    export MODE_ENV="development"
    eval go run cmd/lang-trainer/main.go
fi

if [ $1 = "test" ]; then
    export MODE_ENV="test"
    export DB_HOST=host.docker.internal
    export DB_PORT=5432
    export DB_USERNAME=gotest
    export DB_PASSWORD=tester
    export DB_DATABASE=testdb
    export ACK_GINKGO_RC=true
    eval test go -v ./app/...
fi

if [ $1 = "test-ci" ]; then
    export MODE_ENV="test"
    export DB_HOST=localhost
    export DB_PORT=5432
    export DB_USERNAME=postgres
    export DB_PASSWORD=postgres
    export DB_DATABASE=testdb
    eval ginkgo -v ./app/...
fi