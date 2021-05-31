# Utilitzem com a contenidor temporal la imatge oficial de Golang versió 1.15
FROM golang:1.16 as build-env
# Definit un argument d'entrada que es passarà per la linea de comandes al
# invocar la construcció de la imatge. Aquest valor, s'injectarà al codi Go
# per mostrar la versió del codi en format git commit hash
ARG GIT_COMMIT=buit

# Establim el directori de treball
WORKDIR /go/src/app
# Copiem tot el contingut del repositori en el directori destí /go/src/app
ADD . /go/src/app

# Descarreguem les dependències que necessita el codi
RUN go get -d -v ./...

# Compilem l'aplicació passant com a paràmetre l'argument GIT_COMMIT que s'injectarà en el codi.
# Ubiquem el binari generat en la ruta /app/tfg
RUN go build -ldflags "-X 'github.com/asaldana-uoc/tfg-uoc-app/internal/handlers.gitCommit=$GIT_COMMIT'" \
    -o /app/tfg cmd/web/*.go

# Creem un nou contenidor net fent servir el projecte d'imatges distroless
FROM gcr.io/distroless/base
# Copiem del contenidor previ el binari generat a /app/tfg
COPY --from=build-env /app /app

# És el procés que s'iniciarà per defecte quan s'aixequi un nou contenidor a partir de la imatge creada
CMD ["/app/tfg"]