# diginex-micro-service-1

REST API for reversing a string

## How to use it:

* See usage with:

```
dignex-micro-service-1 help
```
```
Diginex Micro Service 1.

Usage:

  diginex-micro-service-1 [commands|flags]

The commands & flags are:

  help              prints help

  --port              port on which application will run (default: 8080)
  -- diginex_micro_sercvice_2_base_url		Base URL of diginex micro sercvice 2 for accessing the reverse string API

Examples:

  # prints help:
  diginex-micro-service-1 help

  # sample usage
  diginex-micro-service-1 --diginex_micro_sercvice_2_base_url=http://localhost:8081
  diginex-micro-service-1 --port=8080 --diginex_micro_sercvice_2_base_url=http://localhost:8081
```

* Run test cases and code coverage (go > 1.12)
```
go test --cover
```

* Build the binary or [Download from here](./diginex-micro-service-1)
```
go build
```

* Run the binary as a server
```
dignex-micro-service-1                  // default port is 8080
dignex-micro-service-1 --port=9090      // with custom port
```

#### Accessing the API:

```
curl -X POST \
  http://localhost:8080/api \
  -d '{
 "message":"abcdef"
}'
```
Output:
```
{
    "message": "fedcba",
    "rand": 8674665223082153551
}
```