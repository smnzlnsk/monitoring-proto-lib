import sys
from google.protobuf.json_format import MessageToDict

try:
    from monitoring_proto_lib.monitoring.v1 import delete_pb2
    from monitoring_proto_lib.monitoring.v1 import deploy_pb2
    from monitoring_proto_lib.monitoring.v1 import monitoring_pb2_grpc
    
    # Test creating a deployment request
    deploy_request = deploy_pb2.NotifyDeploymentRequest(
        job_name="test-job",
        job_hash="abc123",
        instance_number=1,
        resource=deploy_pb2.ResourceInfo(
            cpu="2",
            memory="4Gi",
            gpu="1",
            disk="100Gi",
            network=deploy_pb2.NetworkInfo(
                bandwidth_in="1Gbps",
                bandwidth_out="1Gbps"
            )
        ),
        calculation_requests=[
            deploy_pb2.CalculationRequest(
                metric_name="test_metric",
                formula="sum(rate(http_requests_total[5m]))",
                description="Test metric description",
                states=["warning", "critical"],
                goal="less",
                unit="requests/second"
            )
        ]
    )
    
    # Test creating a deletion request
    delete_request = delete_pb2.NotifyDeletionRequest(
        job_name="test-job",
        instance_number=1
    )
    
    print("✅ Imports successful!")
    print("\nDeployment Request:")
    print(MessageToDict(deploy_request))
    print("\nDeletion Request:")
    print(MessageToDict(delete_request))
    
except ImportError as e:
    print(f"❌ Import error: {e}")
    print(f"Python path: {sys.path}")
    try:
        import monitoring_proto_lib
        print(f"monitoring_proto_lib location: {monitoring_proto_lib.__file__}")
    except ImportError as e2:
        print(f"Cannot import monitoring_proto_lib: {e2}")