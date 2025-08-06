# Terraform Hex Provider

A Terraform provider that converts strings to hexadecimal representation.

## Features

- **hex_string resource**: Converts input data to hexadecimal encoding
- Built using the modern Terraform Plugin Framework
- Lightweight and efficient

## Example Usage

```hcl
resource "hex_string" "example" {
  data = "hello world"
}

output "hex_result" {
  value = hex_string.example.result  # Outputs: "68656c6c6f20776f726c64"
}
```

### Use Case: imgproxy Configuration

`imgproxy` requires keys and salts to be hex-encoded:

```hcl
resource "random_password" "imgproxy_key" {
  length = 32
}

resource "hex_string" "imgproxy_key_hex" {
  data = random_password.imgproxy_key.result
}

resource "helm_release" "imgproxy" {
  name       = "imgproxy"
  chart      = "imgproxy"
  version    = "1.0.0"

  set_sensitive {
    name  = "key"
    value = hex_string.imgproxy_key_hex.result
  }
}
```

## Schema

### Resource: hex_string

#### Arguments

- `data` (String, Required) - The input data to convert to hexadecimal

#### Attributes

- `id` (String) - Unique identifier for the resource (same as result)
- `result` (String) - The hexadecimal representation of the input data

## Installation

### Terraform Registry

```hcl
terraform {
  required_providers {
    hex = {
      source  = "AnticliMaxtic/hex"
      version = "~> 1.0"
    }
  }
}
```

### Local Development

1. Clone this repository
2. Build: `go build`
3. Install locally using Terraform's development overrides

## Development

### Building

```bash
go build
```

### Testing

```bash
go test ./internal/provider/
```

### Acceptance Testing

```bash
TF_ACC=1 go test ./internal/provider/ -v
```

## License

This project is licensed under the MPL 2.0 License - see the [LICENSE](LICENSE) file for details.
