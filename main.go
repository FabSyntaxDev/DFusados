package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// 2. Busca as variáveis do ambiente
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// 3. Monta a string de conexão
	// Usamos 'localhost' porque o Go está no Windows e o Banco no Docker
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", user, pass, port, name)

	// 4. Estabelece a conexão com o banco
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer conn.Close(ctx) // Garante que a conexão feche ao terminar o programa

	fmt.Println("Conexão estabelecida com sucesso!")

	// 5. Dados para inserir (Simulando uma entrada de usuário)
	novoNome := "FabSyntaxDev"
	novoEmail := "contato@fabsyntax.dev"

	// 6. Executa o comando INSERT
	// O $1 e $2 são substituídos pelos valores de novoNome e novoEmail com segurança
	sql := "INSERT INTO usuarios (nome, email) VALUES ($1, $2)"

	_, err = conn.Exec(ctx, sql, novoNome, novoEmail)
	if err != nil {
		log.Fatalf("Erro ao inserir dados: %v", err)
	}

	fmt.Printf("Sucesso! Usuário %s inserido.\n", novoNome)
}
