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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", user, pass, port, name)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer conn.Close(ctx)

	fmt.Println("Conexão estabelecida com sucesso!")

	novoNome := "test"
	novoEmail := "contato@fabsyntax.dev"

	sql := "INSERT INTO usuarios (nome, email) VALUES ($1, $2)"

	_, err = conn.Exec(ctx, sql, novoNome, novoEmail)
	if err != nil {
		log.Fatalf("Erro ao inserir dados: %v", err)
	}

	fmt.Printf("Sucesso! Usuário %s inserido.\n", novoNome)
}
