generate:
	mkdir -p pb
	touch pb/__init__.py
	python -m grpc_tools.protoc -I. --proto_path=proto --python_out=./pb --grpc_python_out=./pb --pyi_out=./pb model.proto
	sed -i -e 's/import model_pb2 as model__pb2/from . import model_pb2 as model__pb2/' ./pb/model_pb2_grpc.py
