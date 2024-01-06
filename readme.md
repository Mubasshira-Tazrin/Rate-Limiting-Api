## API Rate Limiter with Token Bucket Algorithm and Redis Integration

## Project-Structure

```

├── .github                    
│   └── workflows              
│       └── ci.yaml            <-- Github Actions CI configuration file(for build, lint, test).
├── cmd                        
│   └── app                    
│       ├── app.go             <-- Server setup, init slogger, handlers.
│       └── handler.go         <-- Net/http handler methods.
│       └── handler_test.go    <-- Handler helpers test, as in, TestWriteResponse.
│       └── helpers.go         <-- Slogger Initialization.
├── internal                   
│   ├── constants                 
│   │   ├── constants.go       <-- Constants, as in, server address, port, redis command.
│   ├── redis                 
│   │   ├── redis.go           <-- Redigo library implementation.
│   ├── slogconf                 
│   │   ├── slogconf.go        <-- Structured logging configuration.
│   ├── utils                 
│   │   ├── utils.go           <-- Essential app utils, as in construct auth token etc.
├── .golangci.yml              <-- Configuration for golangci-lint tool.
├── go.mod                     <-- Go module file, tracking dependencies.
├── main.go                    <-- Main application entry point.
├── Makefile                   <-- Defines set of tasks to be executed.
└── readme.md                  <-- Project README file.

```

<!-- GETTING STARTED -->

### Getting Started

###### Clone using ssh protocol `git clone git@github.com:Mubasshira-Tazrin/Rate-Limiting-Api.git`


#### Redis Setup and Config

* (Good For Now) Local redis setup and config.
* (Recommended )Use docker compose.


#### Run-the-application

* Run the application with `make run` command from project root.


