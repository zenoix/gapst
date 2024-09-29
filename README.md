# gapst

Gapst is a test repo for using [gRPC](https://grpc.io/) to communicate between a go API and a python ML model, without building a python REST API.

## Structure

The repo is structured into two main components:

- The go API (found in the api folder)
- The python ML model (found in the classifier folder)

## Requirements

To run the code in this repo, you will need the following:
- Go 1.22
- Python 3.11
- Make
- Sed
- Protobuf
- Protoc-gen-go
- Protoc-gen-go-grpc

and the following python packages:
- grpcio
- grpcio-tools
- numpy
- pandas
- scikit-learn

Other versions of go and python might work, but those are the versions I used.

These are all provided in the corresponding `flake.nix` and `flake.lock` files in the two folders. To use them, use the following command in either the api or classifier (depending on if you're running the go API or python server):

```sh
$ nix develop
```

> [!NOTE]
> If you are not using nix, you will have to set up the dependencies yourself, including setting up the python environment.

## Usage

### Python Server

To try it out, you'll want to first start the python ML model server. This can be done by using the following:

```sh
$ cd classifier

# If using nix, otherwise activate your python environment
$ nix develop

$ python server.py
```

### Go API Client

Then start the go API with the following in a separate terminal:
```sh
$ cd api

# If using nix
$ nix develop

$ go run main.go
```

### Using the Go API as the Client

You can then make a POST request to `http://localhost:8080/predict` to make predictions using a model trained on the [iris dataset](https://www.kaggle.com/datasets/uciml/iris/data) using curl, postman, or any other tool:

```sh
# Example using curl
$ curl -X POST http://localhost:8080/predict -d "{\"SepalLengthCm\": 2.4, \"SepalWidthCm\": 1.2, \"PetalLengthCm\": 1.1, \"PetalWidthCm\": 1.0}"
```

## Protobuf Modifications
If you make any protobuf modifications, you use use `make generate` to re-generate the gRPC code.
