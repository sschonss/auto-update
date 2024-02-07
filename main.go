package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Configuração do log para gravar em um arquivo
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "app: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Carrega variáveis de um arquivo .env
	err = godotenv.Load()
	if err != nil {
		logger.Println("Erro ao carregar o arquivo .env:", err)
		return
	}

	// Obtém variáveis do ambiente
	appPath := os.Getenv("APP_PATH")
	crontab := os.Getenv("CRONTAB")
	gitUser := os.Getenv("GIT_USER")
	gitToken := os.Getenv("GIT_TOKEN")

	// Configuração do cron
	c := cron.New()
	c.AddFunc(crontab, func() {
		// Executa a lógica da aplicação aqui
		logger.Printf("Executando a aplicação no caminho %s\n", appPath)

		// Muda para o diretório especificado por appPath
		err := os.Chdir(appPath)
		if err != nil {
			logger.Printf("Erro ao mudar para o diretório %s: %s\n", appPath, err)
			return
		}

		cmd := exec.Command("git", "pull")
		cmd.Env = append(os.Environ(), fmt.Sprintf("GIT_USER=%s", gitUser), fmt.Sprintf("GIT_TOKEN=%s", gitToken))
		cmd.Env = append(cmd.Env, "PATH="+os.Getenv("PATH"))

		output, err := cmd.CombinedOutput()
		if err != nil {
			logger.Printf("Erro ao executar o comando git: %s\n", err)
			return
		}

		logger.Printf("Saída do git pull: %s\n", output)
	})

	// Inicia o cron
	c.Start()

	// Mantém a aplicação em execução
	select {}
}
