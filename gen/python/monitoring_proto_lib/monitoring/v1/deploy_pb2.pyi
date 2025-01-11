from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class NotifyDeploymentRequest(_message.Message):
    __slots__ = ("job_name", "job_hash", "instance_number", "resource", "calculation_requests")
    JOB_NAME_FIELD_NUMBER: _ClassVar[int]
    JOB_HASH_FIELD_NUMBER: _ClassVar[int]
    INSTANCE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    RESOURCE_FIELD_NUMBER: _ClassVar[int]
    CALCULATION_REQUESTS_FIELD_NUMBER: _ClassVar[int]
    job_name: str
    job_hash: str
    instance_number: int
    resource: ResourceInfo
    calculation_requests: _containers.RepeatedCompositeFieldContainer[CalculationRequest]
    def __init__(self, job_name: _Optional[str] = ..., job_hash: _Optional[str] = ..., instance_number: _Optional[int] = ..., resource: _Optional[_Union[ResourceInfo, _Mapping]] = ..., calculation_requests: _Optional[_Iterable[_Union[CalculationRequest, _Mapping]]] = ...) -> None: ...

class ResourceInfo(_message.Message):
    __slots__ = ("cpu", "memory", "gpu", "disk", "network")
    CPU_FIELD_NUMBER: _ClassVar[int]
    MEMORY_FIELD_NUMBER: _ClassVar[int]
    GPU_FIELD_NUMBER: _ClassVar[int]
    DISK_FIELD_NUMBER: _ClassVar[int]
    NETWORK_FIELD_NUMBER: _ClassVar[int]
    cpu: str
    memory: str
    gpu: str
    disk: str
    network: NetworkInfo
    def __init__(self, cpu: _Optional[str] = ..., memory: _Optional[str] = ..., gpu: _Optional[str] = ..., disk: _Optional[str] = ..., network: _Optional[_Union[NetworkInfo, _Mapping]] = ...) -> None: ...

class NetworkInfo(_message.Message):
    __slots__ = ("bandwidth_in", "bandwidth_out")
    BANDWIDTH_IN_FIELD_NUMBER: _ClassVar[int]
    BANDWIDTH_OUT_FIELD_NUMBER: _ClassVar[int]
    bandwidth_in: str
    bandwidth_out: str
    def __init__(self, bandwidth_in: _Optional[str] = ..., bandwidth_out: _Optional[str] = ...) -> None: ...

class CalculationRequest(_message.Message):
    __slots__ = ("metric_name", "formula", "description", "states", "goal", "unit")
    METRIC_NAME_FIELD_NUMBER: _ClassVar[int]
    FORMULA_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    STATES_FIELD_NUMBER: _ClassVar[int]
    GOAL_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    metric_name: str
    formula: str
    description: str
    states: _containers.RepeatedScalarFieldContainer[str]
    goal: str
    unit: str
    def __init__(self, metric_name: _Optional[str] = ..., formula: _Optional[str] = ..., description: _Optional[str] = ..., states: _Optional[_Iterable[str]] = ..., goal: _Optional[str] = ..., unit: _Optional[str] = ...) -> None: ...

class NotifyDeploymentResponse(_message.Message):
    __slots__ = ("acknowledged", "message")
    ACKNOWLEDGED_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    acknowledged: bool
    message: str
    def __init__(self, acknowledged: bool = ..., message: _Optional[str] = ...) -> None: ...
