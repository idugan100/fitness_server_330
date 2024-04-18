#! /bin/zsh
# make sure to update db path before compiling (/home/ubuntu/fitness_server_330/database)
GOOS=linux GOARCH=amd64 go build -o prod_server
