// main.go

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
	envPath := os.Getenv("ENV_PATH")
	fmt.Println("ENV_PATH:", envPath)
	if envPath == "" {
		return fmt.Errorf("A variável de ambiente ENV_PATH não foi definida.")
	}

	// Exemplo de como você pode obter variáveis específicas do ambiente
	os.Setenv("APP_PATH", os.Getenv("APP_PATH"))
	os.Setenv("CRONTAB", os.Getenv("CRONTAB"))
	os.Setenv("GIT_TOKEN_USER", os.Getenv("GIT_TOKEN_USER"))
	os.Setenv("GIT_USER", os.Getenv("GIT_USER"))
	os.Setenv("BRANCH", os.Getenv("BRANCH"))

	return nil
}

func executeScript() error {
	scriptContent := `
	#!/bin/bash
	PWD="$ENV_PATH"
	echo PWD: $PWD
	ENV_FILE="$PWD/.env"
	echo ENV_FILE: $ENV_FILE
	# Remova as linhas que leem do arquivo .env, pois agora as variáveis estão diretamente no ambiente

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
