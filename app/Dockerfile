FROM golang AS builder
RUN go get github.com/lib/pq
COPY app.go /src/app.go
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-extldflags "-static"' -o app app.go

FROM scratch
COPY --from=builder /src/app /
WORKDIR /
ENV PSQL_HOST PSQL_PORT PSQL_USER PSQL_PASSWORD PSQL_DB
EXPOSE 8080
ENTRYPOINT [ "/app" ]
CMD []
