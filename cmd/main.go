package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/wrferreira1003/load-tester-Go/internal/config"
	"github.com/wrferreira1003/load-tester-Go/internal/tester"
)

func main() {
	// Definir as flags CLI
	url := flag.String("url", "", "URL do serviço a ser testado (obrigatório)")
	requests := flag.Int("requests", 0, "Número total de requests (obrigatório)")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas (obrigatório)")

	flag.Parse()

	// Validar os parâmetros obrigatórios
	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Erro: Todos os parâmetros obrigatórios devem ser fornecidos.")
		fmt.Println("Uso: --url=<URL> --requests=<NÚMERO> --concurrency=<NÚMERO>")
		os.Exit(1)
	}

	// Exibir as configurações carregadas
	fmt.Println("==== Configurações Carregadas ====")
	fmt.Printf("URL: %s\n", *url)
	fmt.Printf("Número de Requests: %d\n", *requests)
	fmt.Printf("Concorrência: %d\n", *concurrency)
	fmt.Println("=================================")

	// Criar a configuração manualmente
	config := &config.Config{
		URL:           *url,
		RequestNumber: *requests,
		Concurrency:   *concurrency,
	}

	// Executar o teste de carga
	fmt.Println("Iniciando teste de carga...")
	test := tester.NewTester(config)
	err := test.RunLoadTest()
	if err != nil {
		log.Fatalf("Erro durante o teste de carga: %v", err)
	}

	fmt.Println("Teste de carga concluído com sucesso!")
}
