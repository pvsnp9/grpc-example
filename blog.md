## gRPC
gRPC, or Google Remote Procedure Call, is an open-source RPC (Remote Procedure Call) framework initially developed by Google. It is designed for building efficient and scalable distributed systems. gRPC uses HTTP/2 as the transport protocol and Protocol Buffers as the interface definition language.

#### Key Features of gRPC:

1. **Interface Definition Language (IDL):**
   - gRPC uses Protocol Buffers (protobuf) as its IDL. Protocol Buffers provide a language-agnostic way to define the structure of messages and services.

2. **HTTP/2 Transport:**
   - gRPC leverages the features of HTTP/2, including multiplexing, header compression, and bidirectional streaming, to optimize communication between clients and servers.

3. **Cross-Language Support:**
   - gRPC supports multiple programming languages, including but not limited to Go, Java, Python, C++, Ruby, JavaScript, and more. This enables interoperability between services written in different languages.

4. **Code Generation:**
   - gRPC generates client and server code from the protobuf service definition. This greatly simplifies the development process and helps ensure consistency between client and server implementations.

5. **Strong Typing:**
   - Protocol Buffers enforce strong typing, providing a clear contract between the client and server regarding the structure of messages and the available service methods.

#### Data Transfer Methods:

gRPC supports several data transfer methods, which are defined by the types of RPCs (Remote Procedure Calls) it supports:

1. **Unary RPC:**
   - In a unary RPC, the client sends a single request to the server and receives a single response. This is the simplest form of RPC, similar to traditional synchronous function calls.
    ```protobuf
    service Name{
        rpc myMethod(MessageRequest) returns (MessageResponse){}
    }
    ```

2. **Server Streaming RPC:**
   - In a server streaming RPC, the client sends a single request to the server and receives a stream of responses. The server sends data to the client as soon as it becomes available.
   ```protobuf
    service Name{
        rpc myMethod(MessageRequest) returns ( stream MessageResponse){}
    }
    ``` 

3. **Client Streaming RPC:**
   - In a client streaming RPC, the client sends a stream of requests to the server and receives a single response. This is useful for scenarios where the client needs to send a large amount of data.
   ```protobuf
    service Name{
        rpc myMethod(stream MessageRequest) returns (MessageResponse){}
    }
    ``` 

4. **Bidirectional Streaming RPC:**
   - In a bidirectional streaming RPC, both the client and server send a stream of messages to each other. This allows for full-duplex communication, enabling scenarios where the client and server need to send multiple messages asynchronously. For instance, financial trading applications. 
   ```protobuf
    service Name{
        rpc myMethod(stream MessageRequest) returns (stream MessageResponse){}
    }
    ``` 

gRPC's support for these different data transfer methods provides flexibility for building a wide range of distributed systems, from simple request-response scenarios to more complex streaming and interactive applications. 
```
 * Exactly one request and one response
 * Messaging order is guranteed for each call, for example - reques 1,2,3... response 1,2,3..
```
