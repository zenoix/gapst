generate:
	python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. --pyi_out=. model.proto
	sed -i -e 's/import model_pb2 as model__pb2/from . import model_pb2 as model__pb2/' model_pb2_grpc.py
