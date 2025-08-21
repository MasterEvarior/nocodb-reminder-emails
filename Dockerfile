FROM golang:1.25-alpine AS build-stage

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ndbre

FROM gcr.io/distroless/base-debian12 AS release-stage
WORKDIR /
COPY --from=build-stage /app/ndbre /app/ndbre
USER nonroot:nonroot

CMD ["/app/ndbre"]