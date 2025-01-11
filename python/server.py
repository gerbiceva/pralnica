# GRPC
import asyncio
import uvicorn
from grpc_server.service_pb2 import Weather, Void
from grpc_server.service_pb2_grpc import WeatherServiceServicer, add_MyServiceServicer_to_server
from sonora.asgi import grpcASGI
from starlette.middleware.cors import CORSMiddleware
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.types import ASGIApp



# class for handling actual communication with the client
class MyService(WeatherServiceServicer):
    async def GetWeather(self, request, context):
        """get the weather string description
        """
        # TODO: ENEEEII
        # dobi nek weather skif api in poslji gor nek string tipo : "sunny"
        pass


class GRPCWebMiddleware(BaseHTTPMiddleware):
    def __init__(self, app: ASGIApp):
        super().__init__(app)

    async def dispatch(self, request, call_next):
        response = await call_next(request)
        response.headers["Content-Type"] = "application/grpc-web+proto"

        # modifying raw headers is necesarry
        del response.raw_headers[0]
        return response
    
async def grpc():
    application = grpcASGI(uvicorn, False)
    # Attach your gRPC server implementation.
    add_MyServiceServicer_to_server(MyService(), application)

     # Add CORS middleware
    application = CORSMiddleware(
        application,
        allow_origins=["*"],  # Adjust this to your allowed origins
        allow_credentials=True,
        allow_methods=["*"],
        allow_headers=["*"],
    )

    # Wrap the app with GRPCWebMiddleware
    application = GRPCWebMiddleware(application)

    config = uvicorn.Config(application, port=8080, log_level="info")
    server = uvicorn.Server(config)
    await server.serve()

if __name__=="__main__":
    asyncio.run(grpc())