import logging
from concurrent import futures

import grpc
import pandas as pd

from model import PredictiveModel
from pb import model_pb2, model_pb2_grpc

PORT = 50051

logger = logging.getLogger(__name__)
logging.basicConfig(level=logging.INFO)


class ModelServicer(model_pb2_grpc.ModelServicer):
    def __init__(self, model: PredictiveModel) -> None:
        self._model = model

    def Predict(self, request, context):  # noqa: ANN001, ANN201, ARG002, N802
        result = self._model.predict(
            pd.DataFrame(
                {
                    "SepalLengthCm": [request.SepalLengthCm],
                    "SepalWidthCm": [request.SepalWidthCm],
                    "PetalLengthCm": [request.PetalLengthCm],
                    "PetalWidthCm": [request.PetalWidthCm],
                }
            )
        )[0]

        logger.info(f"Predicted value: {result}")

        return model_pb2.PredictResponse(Species=result)


def serve(model: PredictiveModel) -> None:
    server: grpc.Server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    model_pb2_grpc.add_ModelServicer_to_server(ModelServicer(model), server)
    _ = server.add_insecure_port(f"[::]:{PORT}")
    server.start()
    logger.info(f"Server started, listening on {PORT}")
    _ = server.wait_for_termination()


if __name__ == "__main__":
    from model import create_demo_rf_model

    rf_model = create_demo_rf_model()

    serve(rf_model)
