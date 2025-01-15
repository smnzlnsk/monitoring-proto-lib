# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from monitoring_proto_lib.monitoring.v1 import delete_pb2 as monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2
from monitoring_proto_lib.monitoring.v1 import deploy_pb2 as monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2


class MonitoringServiceStub(object):
    """Service definition
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.NotifyDeployment = channel.unary_unary(
                '/monitoring.v1.MonitoringService/NotifyDeployment',
                request_serializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentRequest.SerializeToString,
                response_deserializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentResponse.FromString,
                )
        self.NotifyDeletion = channel.unary_unary(
                '/monitoring.v1.MonitoringService/NotifyDeletion',
                request_serializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionRequest.SerializeToString,
                response_deserializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionResponse.FromString,
                )


class MonitoringServiceServicer(object):
    """Service definition
    """

    def NotifyDeployment(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def NotifyDeletion(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MonitoringServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'NotifyDeployment': grpc.unary_unary_rpc_method_handler(
                    servicer.NotifyDeployment,
                    request_deserializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentRequest.FromString,
                    response_serializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentResponse.SerializeToString,
            ),
            'NotifyDeletion': grpc.unary_unary_rpc_method_handler(
                    servicer.NotifyDeletion,
                    request_deserializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionRequest.FromString,
                    response_serializer=monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'monitoring.v1.MonitoringService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MonitoringService(object):
    """Service definition
    """

    @staticmethod
    def NotifyDeployment(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/monitoring.v1.MonitoringService/NotifyDeployment',
            monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentRequest.SerializeToString,
            monitoring__proto__lib_dot_monitoring_dot_v1_dot_deploy__pb2.NotifyDeploymentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def NotifyDeletion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/monitoring.v1.MonitoringService/NotifyDeletion',
            monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionRequest.SerializeToString,
            monitoring__proto__lib_dot_monitoring_dot_v1_dot_delete__pb2.NotifyDeletionResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
