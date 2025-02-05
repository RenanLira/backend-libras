# Usa a imagem oficial do Golang como base
FROM golang:1.23.2-alpine

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos do projeto para o container
COPY . .

# Baixa dependências e compila o binário
RUN go mod tidy && go build -C cmd/api -o ../../build/main

# Define a porta padrão como variável de ambiente
ENV PORT=8080

# Expõe a porta
EXPOSE 8080

# Executa o binário
CMD ["./build/main"]


