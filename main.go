// main.go

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
	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	c := cron.New()

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

	c.Start()

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

	scriptContent := `
	ENV_FILE=".env"
	if [ -f "$ENV_FILE" ]; then
  		while IFS= read -r line; do
    	export "$line"
  	done < "$ENV_FILE"
	else
  		echo "Arquivo .env não encontrado."
  		exit 1
	fi

	if [ ! -d "$APP_PATH" ]; then
  		echo "O diretório do aplicativo não existe: $APP_PATH"
  		exit 1
	fi

	cd "$APP_PATH" || exit

	git config --global user.name "$GIT_USER"
	git config --global user.email "$GIT_USER@example.com"
	git config --global credential.helper "store --file=$HOME/.git-credentials"
	echo "https://$GIT_USER:$GIT_TOKEN_USER@github.com" > "$HOME/.git-credentials"

	git pull origin $BRANCH

	# Verifica se há mudanças no repositório antes de executar npm run build
	if git diff --quiet; then
  		echo "Não há mudanças no repositório. npm run build não será executado."
	else
  		./vendor/bin/sail npm run build
	fi
`

	cmd := exec.Command("/bin/bash", "-c", scriptContent)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
