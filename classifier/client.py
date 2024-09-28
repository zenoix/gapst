import grpc

from proto import model_pb2, model_pb2_grpc


def run(
    sepal_length_cm: float,
    sepal_width_cm: float,
    petal_length_cm: float,
    petal_width_cm: float,
) -> None:
    with grpc.insecure_channel("localhost:50051") as channel:
        stub = model_pb2_grpc.ModelStub(channel)
        response = stub.Predict(
            model_pb2.PredictRequest(
                SepalLengthCm=sepal_length_cm,
                SepalWidthCm=sepal_width_cm,
                PetalLengthCm=petal_length_cm,
                PetalWidthCm=petal_width_cm,
            )
        )
    print(f"Result: {response.Species}")


if __name__ == "__main__":
    sepal_length_cm = 0.2
    sepal_width_cm = 1.24
    petal_length_cm = 2.3
    petal_width_cm = 3.5
    run(sepal_length_cm, sepal_width_cm, petal_length_cm, petal_width_cm)
