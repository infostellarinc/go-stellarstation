{
  sources ? import ./nix/sources.nix,
  pkgs ? import sources.nixpkgs {}
}:

pkgs.mkShell {
  buildInputs = [
    # golang
    pkgs.go_1_21
    # protobuf
    pkgs.buf
    # protobuf: golang
    pkgs.protoc-gen-go
    pkgs.protoc-gen-go-grpc
    pkgs.protoc-gen-connect-go
    # typescript
    pkgs.protoc-gen-grpc-web
    # accessory
    pkgs.wget
    pkgs.unzip
    pkgs.just
  ];
}
