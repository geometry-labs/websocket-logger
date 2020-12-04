FROM golang:1.15-buster AS builder

# GO ENV VARS
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# GO BUILD PREP
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# DO GO BUILD
RUN go build -o main .
WORKDIR /dist
RUN cp /build/main .

# MAKE SMOL FINAL IMAGE

FROM scratch
COPY --from=builder /build/main /
# EXPOSE 1234
CMD ["/main"]