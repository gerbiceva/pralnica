FROM docker.io/library/golang:1.23 AS build

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

COPY --from=build /backend /backend
# COPY --from=build /rso-stats/defaults.env /defaults.env

EXPOSE 1234

USER nonroot:nonroot

ENTRYPOINT [ "/backend" ]