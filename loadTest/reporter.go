package loadtest

import (
	"fmt"
	"time"

	"github.com/thiagohmm/DesafioLoadBalancer/models"
)

func generateReport(results <-chan models.Result, totalRequests int, duration time.Duration) {
	var total200, total404, total500, others int

	for result := range results {
		switch result.StatusCode {
		case 200:
			total200++
		case 404:
			total404++
		case 500:
			total500++
		default:
			others++
		}
	}

	fmt.Println("\n=== Relatório de Teste de Carga ===")
	fmt.Printf("Tempo total: %v\n", duration)
	fmt.Printf("Requests realizados: %d\n", totalRequests)
	fmt.Printf("Status 200: %d\n", total200)
	fmt.Printf("Status 404: %d\n", total404)
	fmt.Printf("Status 500: %d\n", total500)
	if others > 0 {
		fmt.Printf("Outros status: %d\n", others)
	}
}
