if [ $1 = "start" ]; then
    eval go run cmd/lang-trainer/main.go
fi

if [ $1 = "dep" ]; then
    eval go get github.com/onsi/ginkgo/ginkgo
    eval go mod download
fi

if [[ $1 = "test" && $2 = "" ]]; then
    export DB_HOST=host.docker.internal
    export DB_PORT=5432
    export DB_USERNAME=gotest
    export DB_PASSWORD=tester
    export DB_DATABASE=testdb
    export DB_SSLMODE=disable
    export ACK_GINKGO_RC=true
    export SECRET=huippusalainen
    eval ginkgo -v ./app/...
fi

if [[ $1 = "test" && $2 != "" ]]; then
    export DB_HOST=host.docker.internal
    export DB_PORT=5432
    export DB_USERNAME=gotest
    export DB_PASSWORD=tester
    export DB_DATABASE=testdb
    export DB_SSLMODE=disable
    export ACK_GINKGO_RC=true
    export SECRET=huippusalainen
    eval ginkgo --focus $2 -v ./app/...
fi

if [[ $1 = "cover" ]]; then
    export DB_HOST=host.docker.internal
    export DB_PORT=5432
    export DB_USERNAME=gotest
    export DB_PASSWORD=tester
    export DB_DATABASE=testdb
    export DB_SSLMODE=disable
    export ACK_GINKGO_RC=true
    export SECRET=huippusalainen
    eval ginkgo -race $2 -v ./app/...
fi

if [ $1 = "test-ci" ]; then
    export DB_HOST=localhost
    export DB_PORT=5432
    export DB_USERNAME=postgres
    export DB_PASSWORD=postgres
    export DB_DATABASE=testdb
    export DB_SSLMODE=disable
    export SECRET=huippusalainen
    eval ginkgo -v ./app/...
fi