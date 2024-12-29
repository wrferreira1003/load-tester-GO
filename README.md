# Load Tester

## **Objetivo**
O **Load Tester** é uma aplicação CLI desenvolvida em Go para realizar testes de carga em um serviço web.  
A aplicação simula múltiplas requisições simultâneas e gera relatórios com informações detalhadas, como:
- Tempo total de execução.
- Total de requisições enviadas.
- Status HTTP retornados (200, 500, etc.).
- Requisições bem-sucedidas e falhas.

---

## **Entradas**
A aplicação aceita os seguintes parâmetros via linha de comando:

- `--url`: URL do serviço a ser testado (obrigatório).
- `--requests`: Número total de requisições a serem feitas (obrigatório).
- `--concurrency`: Número de requisições simultâneas (obrigatório).

---

## **Saída**
Ao final do teste, a aplicação gera um relatório no terminal com:

- **Tempo total de execução do teste.**
- **Total de requisições enviadas.**
- **Número de requisições bem-sucedidas (HTTP 200).**
- **Número de requisições falhas ou outros códigos HTTP.**
- **Distribuição percentual de status HTTP.**

Exemplo de saída:

==== Relatório de Teste de Carga ==== Tempo total: 1.234s Total de requests enviados: 100 Requests bem-sucedidos (HTTP 200): 98 (98.00%) Requests falhos: 2 (2.00%)

Distribuição de status HTTP:

HTTP 200: 98 (98.00%)
HTTP 500: 2 (2.00%) =====================================
yaml
Copiar código

---

## **Como Executar**

### **1. Localmente**
1. Certifique-se de ter o Go instalado.
2. Compile e execute o programa:
   ```bash
   go build -o load-tester ./cmd/main.go
   ./load-tester --url=https://example.com --requests=100 --concurrency=10

### **2. Docker**
1. Certifique-se de ter o Docker instalado.
2. Build da imagem:
   ```bash
   docker build -t load-tester .
   ```
3. Execute o comando:
   ```bash
   docker run --rm load-tester --url=https://example.com --requests=100 --concurrency=10
