# Gale

Gale is an authentication addon for the Goravel framework, designed to simplify and enhance user authentication processes.

## Installation

To install Gale, you can use the following command:

```bash
go install github.com/kkumar-gcc/gale@latest
```

## Usage

To integrate Gale into your Goravel application, follow these steps:

1. **Install Gale**\
   Run the following command in your terminal:
    ```bash
    gale install
    ```
   This command will install the necessary files and directories required for Gale to function in your Goravel application.
   
2. **Add Authentication Routes**\
   Next, you'll need to add the `Auth` function from `routes/auth.go` to your `route_service_provider.go` file. This will register the authentication routes provided by Gale:
   ```go
    func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
        // ...
        routes.Auth()
    }
    ```
   This will ensure that all the necessary authentication routes (e.g., login, registration, password reset) are available in your application.

## License

Gale is open-source software licensed under the [MIT license](LICENSE).
