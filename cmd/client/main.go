package main

import (
	"challeng-client-server-api/internal/dependencies"
	"challeng-client-server-api/internal/handler"
	quotationintegration "challeng-client-server-api/internal/integration/quotation"
	quotationsqliterepository "challeng-client-server-api/internal/repository/file"
	"challeng-client-server-api/internal/usecase"
	"errors"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

const serverAddr = "127.0.0.1:8080"

func isServerRunning() bool {
	conn, err := net.DialTimeout("tcp", serverAddr, 100*time.Millisecond)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

func startServer() (*exec.Cmd, error) {
	if isServerRunning() {
		return nil, errors.New("server is already running on " + serverAddr)
	}

	cmd := exec.Command("go", "run", "cmd/server/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	for i := 0; i < 6; i++ {
		if isServerRunning() {
			return cmd, nil
		}
		time.Sleep(3000 * time.Millisecond)
	}

	if cmd.Process != nil {
		cmd.Process.Kill()
	}
	return nil, errors.New("server failed to start in time")
}

func main() {
	serverCmd, err := startServer()
	if err != nil {
		log.Printf("Note: %v", err)
	} else {
		defer func() {
			if serverCmd != nil && serverCmd.Process != nil {
				serverCmd.Process.Kill()
			}
		}()
	}

	dep := dependencies.InitClient()

	quotationIntegration := quotationintegration.NewQuotationIntegration(dep.HttpClient)
	quotationRepository := quotationsqliterepository.NewQuotationRepository()

	fetchQuotationUseCase := usecase.NewFetchQuotationUseCase(quotationRepository, quotationIntegration)

	fetchQuotationHandler := handler.NewFetchQuotationHandler(fetchQuotationUseCase)

	err = fetchQuotationHandler.Handle()

	if err != nil {
		log.Fatalf("Error handling fetch quotation: %v", err)
	}

	log.Println("Data fetched and saved successfully.")
}
