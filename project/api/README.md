## gRPC guide
_(Almost entirely copied from [here](https://inst.eecs.berkeley.edu/~cs162/sp23/static/hw/lab-grpc-rs/))_

Machines in a distributed system must communicate with one another.
One method of doing this is using Remote Procedure Calls (RPCs), which make communication between machines look like ordinary function calls.
Effectively, RPCs abstract away the underlying serialization required for raw messaging and allow for an interaction like this:

- A client calls `remoteFileSystem->Read("rutabaga")`.
- The underlying RPC library serializes this function call as a request and sends it to the server.
- The RPC library deserializes the request on the server and runs `fileSys->Read("rutabaga")`, which returns a response.
- The RPC library serializes the response, sends it to the client, and deserializes it to give the client the return value for its function call.

Effectively, this allows the client to call a function on an entirely different machine as if it were just calling any other function.


In this lab, you will be familiarizing yourself with [gRPC](https://grpc.io/about/), an open-source remote procedure call (RPC) framework
that uses Google’s [protocol buffers](https://developers.google.com/protocol-buffers/docs/overview) serialization mechanism.
You will do so by implementing a key/value (KV) store database,
which allows clients to put keys in a server-side hash map and read them out later by making requests to the server.

For a brief rundown of how gRPC works, please read [this guide](https://grpc.io/docs/what-is-grpc/introduction/).

##### Serialization
Serialization is the process of converting structures into bytes that can be sent via communication protocols like TCP,
while deserialization recovers the original structure from the serialized bytes.
Serialization formats specify how serialization and deserialization should occur,
and allow communicating processes to understand the bytes being received from one another.
JSON, XML, and protobuf are some common serialization formats that often find use in configuration files.

### The project

#### What is available
The `.proto` file has already been implemented for you. This specifies the API between the client (our load testing program) and your server.

Reading the [protobuf tutorial](https://protobuf.dev/getting-started/gotutorial/),
you will find out how to compile protobuf `message`s to go-lang `structs` in `.pb.go`.
This file is machine-generated and should not be edited.

Reading the [gRPC tutorial](https://grpc.io/docs/languages/go/basics/),
you will find out how to compile protobuf `service`s and `rpc`s to go-lang `interfaces` and `func`s in `_grpc.pb.go`.
This file is also machine-generated and should not be edited.

The generated files are available at `server/internal/pb`