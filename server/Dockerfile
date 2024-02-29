FROM postgres:15-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext go


#dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download


#build
COPY . ./
RUN go build -o ./bin/app cmd/main.go

#run
FROM alpine AS runner
COPY --from=builder /usr/local/src/bin/app /
COPY configuration/configuration.yml /configuration/configuration.yml
COPY .env /
CMD ["/app"]
