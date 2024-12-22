package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type CotacaoAPI struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func fetchCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var cotacao CotacaoAPI
	err = json.NewDecoder(resp.Body).Decode(&cotacao)
	if err != nil {
		return "", err
	}

	return cotacao.USDBRL.Bid, nil
}

func saveCotacao(ctx context.Context, db *sql.DB, valor string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO cotacoes (valor) VALUES (?)", valor)
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := fetchCotacao(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao obter cotação: %v", err), http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = saveCotacao(ctx, db, cotacao)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao salvar cotação: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"cotacao": "%s"}`, cotacao)))
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, valor TEXT)")
	if err != nil {
		fmt.Println("Erro ao criar tabela:", err)
		return
	}

	http.HandleFunc("/cotacao", cotacaoHandler)
	fmt.Println("Servidor rodando na porta 8080...")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	} else {
		fmt.Println("Servidor iniciado com sucesso!")
	}
}
