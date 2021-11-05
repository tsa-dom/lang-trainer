FROM node:16-alpine AS client

WORKDIR /usr/app/

COPY ./client .

ARG REACT_APP_API_URI
ENV REACT_APP_API_URI=${REACT_APP_API_URI}

RUN npm ci --production && \
    npm run build

FROM golang:alpine AS server

WORKDIR /usr/app/

COPY . .

RUN go get ./... && \
    go build -o server cmd/lang-trainer/main.go

FROM alpine

WORKDIR /go/bin/

COPY --from=client /usr/app/build /go/bin/build
COPY --from=server /usr/app/server /go/bin/server
COPY schema.sql /go/bin/schema.sql

EXPOSE 8080

CMD ["./server"]