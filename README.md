# BoletosEcuador Distributed Systems
Project of the Second Evaluation of the Distributed Systems course at ESPOL, 2018-2s.

## Description
***Coming Soon***

## Group Members:
María Belén Guaranda   
Galo Castillo   
Leonardo Kuffó   

## Installing th
### Instalación de Golang
Follow the next steps for configuring the database
1. `CREATE USER cucaracha WITH PASSWORD 'cucarachaAdmin';`
2. `CREATE DATABASE boletos_ecuador_db;`
3. `GRANT ALL ON DATABASE boletos_ecuador_db TO cucaracha;;`

# How to run the application
### Installing Golang on Ubuntu
Run the following commands:
`sudo apt-get update;`
`sudo apt-get -y upgrade;`
`wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz`
`sudo tar -xvf go1.11.4.linux-amd64.tar.gz
sudo mv go /usr/local`

Setup Go Environment:
`export GOROOT=/usr/local/go`
`export GOPATH=$HOME/Projects/Proj1`
`export PATH=$GOPATH/bin:$GOROOT/bin:$PATH`


go build main.go
go run *.go

