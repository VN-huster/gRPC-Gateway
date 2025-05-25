import grpc
from protogen.greeter import greeter_pb2, greeter_pb2_grpc
def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = greeter_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(greeter_pb2.HelloRequest(name='World'))
    print("Greeter client received: " + response.message)