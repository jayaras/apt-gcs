build:
	env GOOS=linux GOARCH=amd64 go build -o gs github.com/jayaras/apt-gcs/cmd
