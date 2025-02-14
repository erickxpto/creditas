## Descrição

Creditas API é um serviço de simulação de empréstimos que permite calcular o valor total, parcelas mensais e juros totais com base em diferentes tipos de taxas de juros (fixa e variável) e moedas. O serviço também envia um e-mail com os resultados da simulação ao final.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para desenvolver o serviço.
- **Docker**: Utilizado para containerizar a aplicação e facilitar o desenvolvimento e a implantação.
- **Swagger**: Utilizado para documentar a API e fornecer uma interface interativa para testar os endpoints.
- **SMTP**: Utilizado para enviar e-mails com os resultados da simulação. [Opcional]

## Padrões de Design Utilizados

### Strategy Pattern

O padrão Strategy é utilizado para definir diferentes estratégias de cálculo de taxa de juros. Isso permite que o código seja facilmente extensível para adicionar novos tipos de taxas de juros no futuro.

**Motivo da Utilização**: O padrão Strategy permite que o comportamento de cálculo de taxa de juros seja encapsulado em diferentes classes, tornando o código mais modular e fácil de manter.

### Factory Pattern

O padrão Factory é utilizado para criar instâncias das estratégias de taxa de juros com base no tipo de taxa de juros fornecido na requisição.

**Motivo da Utilização**: O padrão Factory facilita a criação de objetos complexos e permite que o código de criação de objetos seja centralizado em um único lugar, melhorando a legibilidade e a manutenibilidade do código.

## Estrutura do Projeto

```
docker-compose.yml
Dockerfile
docs/
go.mod
go.sum
main.go
pkg/
    entities/
    factory/
    handlers/
    services/
    tests/
    utils/
```

- **docs/**: Contém a documentação Swagger da API.
- **pkg/entities/**: Contém as definições das entidades `SimulationRequest` e `SimulationResponse`.
- **pkg/factory/**: Contém a implementação da fábrica de estratégias de taxa de juros.
- **pkg/handlers/**: Contém o handler para a rota de simulação.
- **pkg/services/**: Contém as implementações das estratégias de taxa de juros, serviço de simulação, conversor de moedas e serviço de e-mail.
- **pkg/tests/**: Contém os testes unitários e de integração.
- **pkg/utils/**: Contém funções utilitárias para cálculos.

## Como Executar

### Pré-requisitos

- Docker
- Docker Compose

### Passos

1. Clone o repositório:

```sh
git clone https://github.com/erickgledson/creditas.git
cd creditas
```

2. [Opcional] Configure as variáveis de ambiente no arquivo docker-compose.yml:

```yml
environment:
  - PORT=8080
  - SMTP_HOST=smtp.example.com
  - SMTP_PORT=587
  - SMTP_USERNAME=your_username
  - SMTP_PASSWORD=your_password
```

3. Construa e inicie os containers:

```sh
docker-compose up --build
```

4. Acesse a documentação Swagger em [http://localhost:8080/swagger/](http://localhost:8080/swagger/).

## Testes

Para executar os testes unitários e de integração, utilize o comando:

```sh
go test ./...
```

## Bônus

1. Foi implementando um service para envio de e-mails, bastando adicionar a configuração necessária para funcionar (dados de SMTP). O código está abstraído então a simulação funcionará mesmo que não seja adicionado os dados para envio de e-mail.

2. Há abstração para diferentes taxas de juros. A idéia foi adicionar o padrão strategy para diferentes taxas. Porém, todo o cálculo está fixo numa variável por não ter requisitos explícitos para a implementação.

3. O projeto pode rodar com docker/docker-compose.

4. Foi adicionado suporte para conversão de diversas moedas para BRL. Foi criado um service, com dados fixos. O ideal aqui, seria ter uma API responsável por obter esses dados e fazer a devida conversão para diferentes moedas (ex.: BRL -> EUR ou EUR -> BRL).


## Outras Informações

Não foi adicionado questões como cache, mensageria ou libs externas (a não ser o swagger).
Tentei focar ao máximo no que era pedido, separando responsabilidades no código e deixando o mais simples possível para que pudesse ser extensível. 

## Requisitando a API

Requisitando com 1 simulação
```sh
curl -X POST http://localhost:8080/simulate \
-H "Content-Type: application/json" \
-d '[
  {
    "amount": 1000,
    "birthday": "2000-01-01",
    "payment_term": 12,
    "email": "test1@example.com",
    "interest_rate_type": "default",
    "currency": "USD"
  }
]'
```

Requisitando com 5 simulações
```sh
curl -X POST http://localhost:8080/simulate \
-H "Content-Type: application/json" \
-d '[
  {
    "amount": 1000,
    "birthday": "2000-01-01",
    "payment_term": 12,
    "email": "test1@example.com",
    "interest_rate_type": "default",
    "currency": "USD"
  },
  {
    "amount": 2000,
    "birthday": "1990-06-15",
    "payment_term": 24,
    "email": "test2@example.com",
    "interest_rate_type": "variable",
    "currency": "EUR"
  },
  {
    "amount": 1500,
    "birthday": "1980-12-31",
    "payment_term": 36,
    "email": "test3@example.com",
    "interest_rate_type": "default",
    "currency": "BRL"
  },
  {
    "amount": 2500,
    "birthday": "1970-05-20",
    "payment_term": 48,
    "email": "test4@example.com",
    "interest_rate_type": "variable",
    "currency": "USD"
  },
  {
    "amount": 3000,
    "birthday": "1960-11-10",
    "payment_term": 60,
    "email": "test5@example.com",
    "interest_rate_type": "default",
    "currency": "EUR"
  }
]'
```

Requisitando com 10 mil simulações

Entre na pasta `support` e execute o comando
```sh
curl -X POST http://localhost:8080/simulate \
-H "Content-Type: application/json" \
-d @payload.json
```
