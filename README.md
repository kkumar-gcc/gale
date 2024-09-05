# Gale

Gale is an authentication addon for the Goravel framework, designed to simplify and enhance user authentication processes.

## Version

| kkumar-gcc/gale | goravel/framework |
|-----------------|-------------------|
| v1.0.x          | v1.14.x           |

## Install

### Global Installation

To install Gale globally, use the following command:

```bash
go install github.com/kkumar-gcc/gale@latest
```

### Project-Level Installation

Alternatively, you can install Gale directly within your Goravel project:

1. **Add Gale as a Dependency**\
Navigate to your Goravel project directory and run:
```bash
go get -u github.com/kkumar-gcc/gale
```

2. **Register Gale's Service Provider**\
To make Gale functional within your project, you need to register its service provider in your `config/app.go` file:
```go
import (
   gale "github.com/kkumar-gcc/gale/providers"
)

"providers": []foundation.ServiceProvider{
    // ...
    &gale.ServiceProvider{},
}
```

3. **Run Gale Installation Command**\
After adding Gale to your project, run the following command within your project directory to install the necessary files:
```bash
go run . artisan install

## global installation
gale install
```

**Options provided by the `install` command**:
- `--stack / -s`: Specifies the authentication stack to install. **Options**: `api`.
- `--driver / -m`: Specifies the database driver to use. **Options**: `mysql`, `postgres`, `sqlite`.


## Usage

To integrate Gale into your Goravel application, follow these steps:

1. **Add Authentication Routes**\
Add the `Auth` function from `routes/auth.go` to your `route_service_provider.go` file. This will register the authentication routes provided by Gale:

```go
func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
    // ...
    routes.Auth()
}
```

2. **Run Migrations**\
Run the following command to execute the migrations required for Gale to function:

```bash
go run . artisan migrate:fresh
```

## License

Gale is open-source software licensed under the [MIT license](LICENSE).