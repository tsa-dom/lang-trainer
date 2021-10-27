FROM node:16-alpine AS frontend

WORKDIR /usr/app/

COPY ./frontend .

RUN npm ci --production && \
    npm run build

FROM golang:alpine AS backend

WORKDIR /usr/app/

COPY ./backend .

RUN go get ./... && \
    go build -o server

FROM alpine

WORKDIR /go/bin/

COPY --from=frontend /usr/app/build /go/bin/build
COPY --from=backend /usr/app/server /go/bin/server

EXPOSE 8080

CMD ["./server"]