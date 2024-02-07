#!/bin/bash

# Carrega as variáveis do arquivo .env
ENV_FILE=".env"
if [ -f "$ENV_FILE" ]; then
  while IFS= read -r line; do
    export "$line"
  done < "$ENV_FILE"
else
  echo "Arquivo .env não encontrado."
  exit 1
fi

# Verifica se o diretório do aplicativo existe
if [ ! -d "$APP_PATH" ]; then
  echo "O diretório do aplicativo não existe: $APP_PATH"
  exit 1
fi

# Navega até o diretório do aplicativo
cd "$APP_PATH" || exit

# Atualiza o repositório usando o Git
git config --global user.name "$GIT_USER"
git config --global user.email "$GIT_USER@example.com"
git config --global credential.helper "store --file=$HOME/.git-credentials"
echo "https://$GIT_USER:$GIT_TOKEN_USER@github.com" > "$HOME/.git-credentials"

git pull origin $BRANCH

# Certifica-se de que o script tem permissão de execução
# chmod +x script.sh
