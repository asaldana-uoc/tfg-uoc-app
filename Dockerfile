FROM golang:1.16 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /app/tfg cmd/web/*.go

#FROM gcr.io/distroless/base-debian10
FROM centos
ARG GIT_COMMIT=unspecified
LABEL git_commit=$GIT_COMMIT

COPY --from=build-env /app /app

CMD ["/app/tfg"]