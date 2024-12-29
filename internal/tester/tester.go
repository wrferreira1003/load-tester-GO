package tester

import (
	"net/http"
	"sync"
	"time"

	"github.com/wrferreira1003/load-tester-Go/internal/config"
)

type Tester struct {
	config *config.Config
	client *http.Client // Cliente HTTP com timeout configurado
}

// NewTester cria uma nova instância do Tester.
func NewTester(config *config.Config) *Tester {
	client := &http.Client{
		Timeout: 10 * time.Second, // Timeout para cada requisição
	}
	return &Tester{config: config, client: client}
}

// RunLoadTest executa o teste de carga.
func (t *Tester) RunLoadTest() error {
	var wg sync.WaitGroup
	results := make(chan int, t.config.RequestNumber) // Canal para armazenar resultados
	startTime := time.Now()

	// Função para enviar uma requisição
	sendRequest := func() {
		defer wg.Done()
		resp, err := t.client.Get(t.config.URL)
		if err != nil {
			results <- -1 // -1 para indicar falha na requisição
			return
		}
		results <- resp.StatusCode
		resp.Body.Close()
	}

	// Gerar todas as requisições
	for i := 0; i < t.config.RequestNumber; i++ {
		wg.Add(1)
		go sendRequest()
	}

	// Aguardar que todas as requisições sejam concluídas
	go func() {
		wg.Wait()
		close(results) // Fechar o canal quando todas as goroutines terminarem
	}()

	duration := time.Since(startTime)
	return GenerateReport(results, duration, t.config.RequestNumber)
}
