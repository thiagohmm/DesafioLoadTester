package loadtest

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/thiagohmm/DesafioLoadBalancer/models"
)

type Tester struct {
	url           string
	totalRequests int
	concurrency   int
	token         string
}

func NewTester(url string, totalRequests int, concurrency int, token string) *Tester {
	return &Tester{
		url:           url,
		totalRequests: totalRequests,
		concurrency:   concurrency,
	}
}

func (t *Tester) Run() {
	results := make(chan models.Result, t.totalRequests)
	var wg sync.WaitGroup
	sem := make(chan struct{}, t.concurrency)

	startTime := time.Now()

	for i := 0; i < t.totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			t.makeRequest(results)
			<-sem
		}()
	}

	wg.Wait()
	close(results)

	duration := time.Since(startTime)
	generateReport(results, t.totalRequests, duration)
}

func (t *Tester) makeRequest(results chan<- models.Result) {
	req, err := http.NewRequest("GET", t.url, nil)
	if err != nil {
		log.Println("Erro ao criar request:", err)
		results <- models.Result{StatusCode: 0}
		return
	}
	if t.token != "" {
		req.Header.Add("Authorization", t.token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao realizar request:", err)
		results <- models.Result{StatusCode: 0}
		return
	}
	defer resp.Body.Close()

	results <- models.Result{StatusCode: resp.StatusCode}
}
