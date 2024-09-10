# Fast Food FIAP - Tech Challenge - Customer Microservice

Microsserviço responsável pela parte de cliente (criar cliente e pesquisar cliente).

## Instruções para Rodar a Aplicação Local

1. **Configuração do Ambiente:**

      Certifique-se de ter o Docker e o Kubernetes configurados corretamente na sua máquina.
      
2. **Build da Imagem Docker:**
      
      Execute o seguinte comando para construir a imagem Docker da aplicação:
      
      ```bash
      docker build . -t customer-service -f Dockerfile
      ```

3. **Deploy com Kubernetes:**

      Aplique os recursos do Kubernetes utilizando o seguinte comando:

      ```bash
      kubectl apply -f infra/
      ```
      
4. **Verificação do Status dos Pods:**

      Aguarde até que todos os pods estejam rodando com o comando:

      ```bash
      kubectl get pods --watch
      ```

5. **Pronto.**

## Remoção de Recursos

Para apagar todos os recursos criados, utilize o comando:

```bash
kubectl delete -f infra/
```

## Testes de estresse

Para realizar testes de estresse, certifique-se de ter o K6 instalado na sua máquina e execute o seguinte comando:

```bash
k6 run --duration 1m tests/stress.js
```

## Documentação da API
Ao importar a documentação presente em `docs/tech-challenge.json` de cada repositório, no Postman, terão valores de exemplos editáveis.


SonarCloud: https://sonarcloud.io/summary/new_code?id=Food-fusion-Fiap_customer-service
![image](https://github.com/user-attachments/assets/ab8acb89-bbbc-48be-b3cd-2c1eb74f8527)
