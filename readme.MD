# Backend Test
## Data Structures And Algorithms

> access each sub folder in `algs` directory and run main function to test code.
> replace code in main to test with your testcases.

1. [Gray Code](/algs/gray-code/main.go)

2. [Sum of Distances in Tree](/algs/max-length-of-repeated-subarray/main.go)

3. [Maximum Length of Repeated Subarray](/algs/sum-of-distance-in-tree/main.go)

## Hackathon
### prepare
- make sure that your machine have been installed `docker & docker compose`
- check if exist file .env in project root folder, if not, create one with below content
```
SERVER_PORT = 80
DB_PORT = 5432
```
- `SERVER_PORT` is port that web application will run on your localhost
- `DB_PORT` is port that database will run on your local env, this port is important for running unittest at repository layer
- please replace ports and make sure those ports are available on your machine, no application running on those ports.
### run application
- please run `docker-compose up` with `env file that stored in project root folder`, please stand at project root folder to run this command:
```
docker-compose --env-file ./.env up 
```
- after each time you run, to make sure that no issue will appear in the next time, please run `docker-compose --env-file ./.env down` to clean all related containers.

### test application
- access `http://127.0.0.1:{PORT}/swagger/index.htl` to access swagger documentation.
- all endpoints are already to test with swagger. see detail at `https://swagger.io/docs/`
### run unittest
- because of repository layer will test on real database, so please make sure that database container already started
- you can run unittest while `docker-compose up` are running, or run ` docker-compose --env-file ./.env up postgres` to start database container only.
- to run unittest, please stand at project root folder and run `go test ./...`
- notes: unittest is only implemented on layers: controller, service and repository 
