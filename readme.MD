# Backend Test
## Data Structures And Algorithms

> access each sub folder in `algs` directory and run main function to test code.
> replace code in main to test with your testcases.

1. [Gray Code](/algs/gray-code/main.go)

2. [Sum of Distances in Tree](/algs/max-length-of-repeated-subarray/main.go)

3. [Maximum Length of Repeated Subarray](/algs/sum-of-distance-in-tree/main.go)

## Hackathon
to run this project, please follow this steps:
1. install docker and docker-compose
2. run `docker-compose --env-file ./.env up` in root directory

note: please check `.env` file to change port and database config to your own

## Test Application
to run test, please follow this steps:
1. Access `http://127.0.0.1:{port}/swagger/index.html` to see documentation and test api (see swagger UI tool for more information)
2. For endpoints that require authentication, please add `Authorization:Bearer {token}` in header (using swagger UI tool or manually)
