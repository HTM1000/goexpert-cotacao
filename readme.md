# GoExpert - Cotação de Moeda

Este projeto consiste em um sistema simples para obter e salvar cotações do dólar (USD) em relação ao real (BRL). O sistema é composto por um servidor em Go que consome a API externa de cotações e armazena os dados em um banco SQLite, e por um cliente que consome a API local para exibir e salvar a cotação em um arquivo `.txt`.

## Estrutura do Projeto

### Arquivos principais:

- **server.go**: Implementa o servidor HTTP que:
  - Consome a cotação do dólar de uma API externa (`https://economia.awesomeapi.com.br/json/last/USD-BRL`).
  - Salva a cotação obtida no banco de dados SQLite.
  - Expõe um endpoint `/cotacao` para fornecer a cotação mais recente.

- **client.go**: Implementa o cliente HTTP que:
  - Consome o endpoint `/cotacao` do servidor.
  - Salva a cotação retornada em um arquivo de texto chamado `cotacao.txt`.

- **database.go**: Contém o script para criação da tabela `cotacoes` no banco de dados SQLite.

- **go.mod** e **go.sum**: Gerenciam as dependências do projeto, incluindo o driver `github.com/mattn/go-sqlite3` para integração com o SQLite.

## Funcionalidades

### Servidor

1. Obtém a cotação do dólar através de uma API externa.
2. Armazena a cotação em um banco de dados SQLite.
3. Disponibiliza a cotação mais recente por meio do endpoint `/cotacao`.

### Cliente

1. Consome o endpoint `/cotacao` do servidor.
2. Salva a cotação obtida em um arquivo `cotacao.txt` no formato:  

## Dependências

- Go (versão 1.23.4 ou superior)
- Driver SQLite para Go: `github.com/mattn/go-sqlite3`

## Como Executar

### Passo 1: Clonar o Repositório
```plaintext
bash
git clone https://github.com/HTM1000/goexpert-cotacao.git
cd goexpert-cotacao
```

### Passo 2: Instalar Dependências

Certifique-se de ter o Go instalado e execute:
`go mod tidy`

### Passo 3: Executar o Servidor

Inicie o servidor:
`go run server.go`
O servidor estará disponível em `http://localhost:8080`.

### Passo 4: Executar o Cliente

Com o servidor rodando, execute o cliente para consumir a cotação:
`go run client.go`
O cliente salvará a cotação em um arquivo `cotacao.txt`.

## Estrutura do Banco de Dados

A tabela `cotacoes` é criada automaticamente no banco SQLite. Estrutura da tabela:
```plaintext
CREATE TABLE IF NOT EXISTS cotacoes (
  id INTEGER PRIMARY KEY,
  valor TEXT
);
```

## Observações

 - Timeouts: Foram configurados timeouts para as operações HTTP e de banco de dados, garantindo resiliência em situações de lentidão.
 - Banco de Dados: O arquivo do banco SQLite é chamado cotacoes.db e é criado na raiz do projeto.