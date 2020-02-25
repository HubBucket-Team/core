## Database migrations

Install `goose`:
```bash
go get -u github.com/pressly/goose/cmd/goose
```

Note: Up migration are automatically executed when the application is run.

```bash
goose -dir db create <name> sql                       # Create migration in db folder
goose -dir db postgres "${DB_DATASOURCE_NAME}" up     # Migrate the DB to the most recent version available
goose -dir db postgres "${DB_DATASOURCE_NAME}" down   # Roll back the version by 1
goose help                                            # See all available commands
```

## gRPC installation

Install gRPC:
```bash
go get -u google.golang.org/grpc
```

Download pre-compiled binaries for your platform(protoc-<version>-<platform>.zip) from here: https://github.com/google/protobuf/releases

On macOS or Linux:

- Unzip `protoc-<version>-<platform>.zip`
- Move `bin/protoc` to `/usr/local/bin/`
- Move `include/google` to `/usr/local/include`

Then use `go get -u` to download the following packages:

```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

This will place three binaries in your `$GOBIN`;

* `protoc-gen-grpc-gateway`
* `protoc-gen-swagger`
* `protoc-gen-go`

Make sure that your `$GOBIN` is in your `$PATH`.

## Code generation

Generate Go and Swagger APIs:
```bash
make api
```

Generate Python SDK:
```bash
make python-sdk
```

Note that for the Python SDK, you should setup a virtual environment.
Example:
```shell script
python3 -m pip install virtualenv
sudo apt install virtualenv #Ubuntu

virtualenv -p /usr/bin/python3 venv3
```
You may also need these packages:
```shell script
pip install six python-dateutil urllib3 certifi
```


To generate API and SDKs:
```bash
make all
```
