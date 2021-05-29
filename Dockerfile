FROM golang:1.16 as build-env
ARG GIT_COMMIT=buit

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -ldflags "-X github.com/asaldana-uoc/tfg-uoc-app/internal/handlers/handlers.gitCommit=$GIT_COMMIT" \
    -o /app/tfg cmd/web/*.go

FROM gcr.io/distroless/base
COPY --from=build-env /app /app

CMD ["/app/tfg"]