# Use a imagem base do Go
FROM golang:1.22-alpine

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos de dependências do Go
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copie o restante dos arquivos do projeto
COPY . .

# Exponha a porta que o aplicativo usará
EXPOSE 8080

# Defina o comando de inicialização do contêiner
CMD ["go", "run", "cmd/educ-off-api/main.go"]
