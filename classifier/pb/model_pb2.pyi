from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class PredictRequest(_message.Message):
    __slots__ = ("SepalLengthCm", "SepalWidthCm", "PetalLengthCm", "PetalWidthCm")
    SEPALLENGTHCM_FIELD_NUMBER: _ClassVar[int]
    SEPALWIDTHCM_FIELD_NUMBER: _ClassVar[int]
    PETALLENGTHCM_FIELD_NUMBER: _ClassVar[int]
    PETALWIDTHCM_FIELD_NUMBER: _ClassVar[int]
    SepalLengthCm: float
    SepalWidthCm: float
    PetalLengthCm: float
    PetalWidthCm: float
    def __init__(self, SepalLengthCm: _Optional[float] = ..., SepalWidthCm: _Optional[float] = ..., PetalLengthCm: _Optional[float] = ..., PetalWidthCm: _Optional[float] = ...) -> None: ...

class PredictResponse(_message.Message):
    __slots__ = ("Species",)
    SPECIES_FIELD_NUMBER: _ClassVar[int]
    Species: str
    def __init__(self, Species: _Optional[str] = ...) -> None: ...
