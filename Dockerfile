#imagem do golang
FROM golang:1.24 as build
WORKDIR /app
# copia todo o projeto para dentro da imagem
COPY . .
# roda o comando go build para compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun3

# As proximas linhas, copia em uma imagem vazia o build rodado a cima
# A imagem scratch é a menor imagem possível para rodar uma aplicação
# Nas seguintes linhas, estamos copiando a imagem da aplicação do GO que chamamos de build para a imagem scratch, isso é feito para reduzir o tamanho da imagem
FROM scratch
# Define o diretório de trabalho da imagem
WORKDIR /app
#copio a pasta app/cloudrun da imagem build para a pasta app da imagem scratch
COPY --from=build /app/cloudrun3 .
# Define o comando de entrada da imagem
ENTRYPOINT ["./cloudrun3"]