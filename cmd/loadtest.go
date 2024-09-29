package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thiagohmm/DesafioLoadBalancer/loadtest"
)

var (
	url         string
	requests    int
	concurrency int
	token       string
)

var rootCmd = &cobra.Command{Use: "app"}

// loadtestCmd representa o comando para executar o teste de carga
var loadtestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Executa um teste de carga em uma URL",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			log.Fatal("Por favor, forneça uma URL válida.")
		}

		tester := loadtest.NewTester(url, requests, concurrency, token)
		tester.Run()
	},
}

func init() {
	loadtestCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	loadtestCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Número total de requests")
	loadtestCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Número de chamadas simultâneas")
	loadtestCmd.Flags().StringVarP(&token, "token", "t", "", "Token de autenticação")

	rootCmd.AddCommand(loadtestCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
