import sys

try:
    from monitoring_proto_lib.monitoring.v1 import delete_pb2
    from monitoring_proto_lib.monitoring.v1 import deploy_pb2
    print("✅ Imports successful!")
except ImportError as e:
    print(f"❌ Import error: {e}")
    print(f"Python path: {sys.path}")
    try:
        import monitoring_proto_lib
        print(f"monitoring_proto_lib location: {monitoring_proto_lib.__file__}")
    except ImportError as e2:
        print(f"Cannot import monitoring_proto_lib: {e2}")