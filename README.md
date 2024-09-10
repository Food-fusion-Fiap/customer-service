# Food Fusion - Tech Challenge - Customer Microservice

Microsserviço responsável pela parte de cliente (criar cliente e pesquisar cliente).

# Como rodar localmente
- `docker compose up` para rodar a base (postgres e pgadmin) (arquivo docker-compose.yml)
- em src/infra/db/gorm/gorm.go, commente a string de produção e descomente a string local
- rode a aplicação pela IDE ou pelo comando `go run ./main.go`

## Testes de estresse

Para realizar testes de estresse, certifique-se de ter o K6 instalado na sua máquina e execute o seguinte comando:

```bash
k6 run --duration 1m tests/stress.js
```

## Documentação da API
Ao importar a documentação presente em `docs/tech-challenge.json` de cada repositório, no Postman, terão valores de exemplos editáveis.


SonarCloud: https://sonarcloud.io/summary/new_code?id=Food-fusion-Fiap_customer-service
![image](https://github.com/user-attachments/assets/ab8acb89-bbbc-48be-b3cd-2c1eb74f8527)
