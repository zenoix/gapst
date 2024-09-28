from typing import Protocol

import numpy as np
import pandas as pd
from sklearn.datasets import load_iris
from sklearn.ensemble import RandomForestClassifier


class PredictiveModel(Protocol):
    def predict(self, X: pd.DataFrame) -> np.ndarray: ...  # noqa: N803


def create_demo_rf_model() -> RandomForestClassifier:
    iris_data = load_iris()

    iris_df = pd.DataFrame(data=iris_data["data"], columns=iris_data["feature_names"])

    iris_df = iris_df.rename(
        columns={
            "sepal length (cm)": "SepalLengthCm",
            "sepal width (cm)": "SepalWidthCm",
            "petal length (cm)": "PetalLengthCm",
            "petal width (cm)": "PetalWidthCm",
        }
    )

    iris_df["Species"] = [
        "sentosa" if x == 0 else ("versicolor" if x == 1 else "virginica")
        for x in iris_data["target"].tolist()
    ]

    X = iris_df[["SepalLengthCm", "SepalWidthCm", "PetalLengthCm", "PetalWidthCm"]]  # noqa: N806
    y = iris_df["Species"]

    rf_model = RandomForestClassifier(max_depth=5)

    rf_model.fit(X, y)

    return rf_model
