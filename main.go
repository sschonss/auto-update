package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/cron/v3"
)

func main() {
	// Carrega as variáveis do arquivo .env
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Inicializa o cron
	c := cron.New()

	// Adiciona a tarefa cron
	_, err = c.AddFunc(os.Getenv("CRONTAB"), func() {
		fmt.Println("Executando tarefa cron...")
		err := executeScript()
		if err != nil {
			log.Println("Erro ao executar o script:", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// Inicia o cron
	c.Start()

	// Aguarda indefinidamente
	select {}
}

func loadEnv() error {
	envFile := ".env"
	file, err := os.Open(envFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func executeScript() error {
	cmd := exec.Command("/bin/bash", "-c", "./update.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
