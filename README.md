# Go Luhn Credit Card Validator

A simple Go microservice for validating credit card numbers using the Luhn algorithm.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.21 or higher)

### Installing

1. Clone the repository to your local machine:
``` 
git clone https://github.com/Palaszontko/go-luhn-credit-validator
cd go-luhn-credit-validator
```

## Usage

1. Run the server:
```
go run cmd/server/main.go
```

2. Send a POST request to `http://localhost:8080/validate` with a JSON body containing the credit card number:

```json
{
    "cardNumber": "4111111111111111"
}
```

3. Response 
```json
{
    "isValid": true,
    "cardNetwork": "Visa"
}
```

## Running the tests

Use the following command in the project directory to run the tests:

``` bash 
go test ./...
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
