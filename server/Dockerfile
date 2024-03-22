FROM postgres:15-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev go

# Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# Build
COPY . ./
RUN go build -o ./bin/app cmd/main.go


FROM postgres:15-alpine

# Install additional dependencies if needed
# RUN apk --no-cache add <additional_packages>

# Copy built binary from builder stage
COPY --from=builder /usr/local/src/bin/app /

# Copy configuration files
COPY configuration/configuration.yml /configuration/configuration.yml
COPY .env /

# Start the application
CMD ["/app"]
