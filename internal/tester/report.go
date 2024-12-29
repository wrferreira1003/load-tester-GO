package tester

import (
	"fmt"
	"time"
)

// GenerateReport gera e imprime o relatório com base nos resultados.
// GenerateReport gera e imprime o relatório com base nos resultados.
func GenerateReport(results chan int, duration time.Duration, totalRequests int) error {
	statusCount := make(map[int]int)

	// Contar os status HTTP recebidos
	for status := range results {
		statusCount[status]++
	}

	// Relatório detalhado
	fmt.Println("\n==== Relatório de Teste de Carga ====")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Total de requests enviados: %d\n", totalRequests)
	fmt.Printf("Requests bem-sucedidos (HTTP 200): %d (%.2f%%)\n", statusCount[200], (float64(statusCount[200])/float64(totalRequests))*100)
	fmt.Printf("Requests falhos: %d (%.2f%%)\n", statusCount[-1], (float64(statusCount[-1])/float64(totalRequests))*100)

	fmt.Println("\nDistribuição de status HTTP:")
	for status, count := range statusCount {
		if status != -1 { // Excluir erros já relatados acima
			fmt.Printf("- HTTP %d: %d (%.2f%%)\n", status, count, (float64(count)/float64(totalRequests))*100)
		}
	}

	fmt.Println("=====================================")

	return nil
}
