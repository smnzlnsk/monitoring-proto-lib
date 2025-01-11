from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class NotifyDeletionRequest(_message.Message):
    __slots__ = ("job_name", "instance_number")
    JOB_NAME_FIELD_NUMBER: _ClassVar[int]
    INSTANCE_NUMBER_FIELD_NUMBER: _ClassVar[int]
    job_name: str
    instance_number: int
    def __init__(self, job_name: _Optional[str] = ..., instance_number: _Optional[int] = ...) -> None: ...

class NotifyDeletionResponse(_message.Message):
    __slots__ = ("acknowledged", "message")
    ACKNOWLEDGED_FIELD_NUMBER: _ClassVar[int]
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    acknowledged: bool
    message: str
    def __init__(self, acknowledged: bool = ..., message: _Optional[str] = ...) -> None: ...
