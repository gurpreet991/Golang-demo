# golang-test

## Description

This project is a concurrent request processing handler.

## Project Structure

--> Main Logic(main.go): Handles the execution flow and request management.
--> Configuration(config.go): Stores the constants and settings.
--> Utilities(utils.go): Contains helper functions and types.


## Features

1. Processes multiple requests concurrently using Go routines.
2. Utilizes a priority queue to manage request order.
3. Manages concurrent operations with channels and WaitGroup.
4. Ensures safe concurrent access with sync.Mutex.
5. Configurable through environment variables.


## Setup

Clone the repository

git clone "REPOSITORY URL"


## Command to run project

go run main.go



## Test Case 1
m = 10
n = 600
k = 6000
u = 6
Request Per User = 4000


## Test Case 2
m = 30
n = 1800
k = 18000
u = 20
Request Per User = 1200

