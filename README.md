# BoletosEcuador Distributed Systems
Project of the Second Evaluation of the Distributed Systems course at ESPOL, 2018-2s.



## Group Members:
María Belén Guaranda   
Galo Castillo   
Leonardo Kuffó

## How to run the application
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

### Installing gin-gonic and GORM
Run the following command in the terminal to install gin-gonic:

`go get -u github.com/gin-gonic/gin`

Run the following command in the terminal to install GORM:

`go get -u github.com/jinzhu/gorm`

### Running the app
Run in the project root directory:

`go build main.go`

`go run *.go`




