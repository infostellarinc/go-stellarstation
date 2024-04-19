project := justfile_directory()
tmp := project / "tmp"

protos-path := ".protos"
# where the proto/stellarstation folder is located in the stellarstation-api repo src
# should land you in front of stellarstation/api/v1 folders
stellarapi-proto-path := "api/src/main/proto"

@help:
    just --list

default:
    @help


# Clean generate bindings for all languages
buf version: buf-api-clean buf-clean-dep
    just buf-setup-dep {{version}}
    just pb
    just move-generated

# Delete all generated
buf-api-clean:
    cd {{project}} && rm -rf {{project}}/api

# pull dependency
buf-clean-dep:
    cd {{project}} && rm -rf {{protos-path}}
    cd {{project}} && rm -rf {{tmp}}

# Create a directory to create symlinks in
buf-setup-dep version:
    cd {{project}} && mkdir -p {{protos-path}}
    cd {{project}} && mkdir -p {{tmp}}
    wget https://github.com/infostellarinc/stellarstation-api/archive/refs/tags/{{version}}.zip -O {{tmp}}/api.zip
    unzip {{tmp}}/api.zip -d {{tmp}}/api-out
    mv {{tmp}}/api-out/stellarstation-api-{{version}}/{{stellarapi-proto-path}} {{tmp}}/.protos

pb:
    cd {{project}} && buf generate {{tmp}}/.protos --template buf.gen.yaml

# move the pb generated protos into the ready to use /api directory
move-generated:
    cd {{project}} && mv {{tmp}}/gen/api/stellarstation/api {{project}}/api
