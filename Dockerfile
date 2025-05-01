FROM golang:1.24 AS baseimage

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o urge

CMD ["/urge"]

FROM scratch
COPY --from=baseimage /app/urge /urge
ENTRYPOINT ["/urge"]

