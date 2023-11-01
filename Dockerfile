FROM golang:1.21.3-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o /build/project-api .

FROM gcr.io/distroless/static-debian11

COPY --from=builder /build/project-api /project-api

ENTRYPOINT [ "/project-api" ]
