# Auto Update

Auto Update is a simple script to update your git repository with the latest changes from the remote repository. It is useful when you have a repository that you want to keep up to date with the latest changes from the remote repository.

## Requirements

To run Auto Update, you need to have the following software installed on your machine:

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)


## Installation

To install Auto Update, simply clone the repository and configure the script to run on a schedule.

The following steps will guide you through the installation process:

1. Clone the repository to your local machine:

```bash
git clone <repository-url>
```

2. Change into the directory of the cloned repository:

```bash
cd <repository-name>
```

3. Make the script executable:

```bash
go build -o update
```

4. Configure the script to run on a schedule:

```bash
sudo nano /etc/environment
```
In the file, set the following environment variables:

```bash
APP_PATH=/path/to/your/app
ENV_PATH=/path/to/your/env
CRONTAB=* * * * *
GIT_TOKEN_USER=your_token
GIT_USER=your_user
BRANCH=main
```

You can set the CRONTAB variable to any valid cron expression. For example, to run the script every day at 3:00 AM, you can set the CRONTAB variable to:

```bash
CRONTAB=0 3 * * *
```

Another way to save the environment variables is used zshrc or bashrc.

```bash
nano ~/.zshrc
```

```bash
export APP_PATH=/path/to/your/app
export ENV_PATH=/path/to/your/env
export CRONTAB=* * * * *
export GIT_TOKEN_USER=your_token
export GIT_USER=your_user
export BRANCH=main
```

```bash
source ~/.zshrc
```

5. Create a new systemd service:

```bash
sudo nano /etc/systemd/system/auto-update.service
```

In the auto-update.service file, add the following configuration:

```bash
# /etc/systemd/system/auto-update.service

[Unit]
Description=Auto Update Service

[Service]
ExecStart=/caminho/para/seu/executavel
WorkingDirectory=/caminho/do/seu/executavel

[Install]
WantedBy=default.target

```

6. Reload the systemd daemon:

```bash
sudo systemctl daemon-reload
```

7. Enable the service:

```bash
sudo systemctl enable auto-update.service
```

8. Start the service:

```bash
sudo systemctl start auto-update.service
```

9. Check the status of the service:

```bash
sudo systemctl status auto-update.service
```

If everything is set up correctly, you should see the status of the service as active.

## Recommended Usage

Auto Update is best used in a repository that you want to keep up to date with the latest changes from the remote repository. It is especially use for production environments where you want to ensure that the latest changes are always deployed to the server.

## Why Auto Update?

It alreadys exists some tools to do this, but I wanted to create a simple and easy to use tool to update my repositories. I also wanted to learn more about Go and how to create a simple script to automate tasks.

Sometimes, the hardwares are not so powerful to run a lot of services, so I wanted to create a simple tool to do this.

## Contributing

If you have any ideas on how to improve Auto Update, feel free to open an issue or submit a pull request. I am always open to new ideas and improvements.

## License

Auto Update is licensed under the [MIT License](LICENSE).

