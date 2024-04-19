# Go generated API subs

This repository contains the generated Go packages for gRPC stubs to the StellarStation API.
The definitions come from [infostellarinc/stellarstation-api](https://github.com/infostellarinc/stellarstation-api).

```go
import stellarstation "github.com/infostellarinc/go-stellarstation/api/v1"
```

## Building

Utilizes nix-shell, you can run `just buf 0.14.0` and it'll that tagged version of `stellarstation-api`, build the stubs, and put them in the `/api` folder
