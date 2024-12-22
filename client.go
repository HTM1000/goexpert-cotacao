package main

import (
	"context"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"encoding/json"
)

func getCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Resposta do servidor:", string(body))

	// Extrair o valor da cotação do JSON
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	cotacao, ok := result["cotacao"].(string)
	if !ok {
		return "", fmt.Errorf("cotação não encontrada")
	}

	return cotacao, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotacao, err := getCotacao(ctx)
	if err != nil {
		fmt.Println("Erro ao obter cotação:", err)
		return
	}

	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %s", cotacao)), 0644)
	if err != nil {
		fmt.Println("Erro ao salvar cotação:", err)
		return
	}

	fmt.Println("Cotação salva com sucesso!")
}
