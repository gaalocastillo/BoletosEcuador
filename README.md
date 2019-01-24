# BoletosEcuador Distributed Systems
Project of the Second Evaluation of the Distributed Systems course at ESPOL, 2018-2s.

## Architecture  
![alt text](https://i.imgur.com/AWYaoxN.png "Diseno de Arquitectura")

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

### Installing Redis (Optional - Session Management if running locally)

Execute the following script on this link on your machine: https://gist.github.com/od3n/7914793   
This scripts installs REDIS 5.0 on your machine

Change the `redis.conf` file inside etc/redis/redis.conf directory, so that
redis port is able to receive external traffic (0.0.0.0)

Run the following command
`sudo service redis-server restart`

Your Redis is ready to go!

### Installing Dependencies (gin-gonic and GORM)
Run the following command in the terminal to install gin-gonic:

`go get -u github.com/gin-gonic/gin`

Run the following command in the terminal to install GORM:

`go get -u github.com/jinzhu/gorm`

Run the following command in the terminal to install other dependencies:

`go get -u "github.com/gin-contrib/sessions"`
`go get -u "github.com/gin-contrib/sessions/redis"`
`go get -u "github.com/jinzhu/gorm/dialects/postgres"`


### Running the app
Run in the project root directory:

`go build main.go`

`go run *.go`

Access to localhost:3001

## Performance Tests

![alt text](https://i.imgur.com/EfCgHx5.png "Performance Tests")
