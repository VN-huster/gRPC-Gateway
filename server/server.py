from concurrent import futures
import grpc
import time
from protogen.greeter import greeter_pb2, greeter_pb2_grpc


class Greeter(greeter_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        return greeter_pb2.HelloReply(message=f"Hello, {request.name}!")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    greeter_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port("[::]:8080")
    server.start()
    print("Server started on port 8080")
    server.wait_for_termination()
    # try:
    #     while True:
    #         time.sleep(86400)  # Keep the server running for 1 day
    # except KeyboardInterrupt:
    #     server.stop(0)


# if __name__ == "__main__":
#     serve()
