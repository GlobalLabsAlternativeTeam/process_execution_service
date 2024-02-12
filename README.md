# [process_execution_service](https://versed-erica-20d.notion.site/process_execution_service-2185163da6be429caec9662e1f0dd46f)

This service is accountable for execution of the schemas

## Dependencies

In order to generate the interfaces using `proto`, you will need to install some previous dependencies:

```bash
TODO
```

## Usage

### Generating proto interfaces

In order to run the service, we first need to **generate the interfaces** using protoc.

To do this, run the following command from the root directory:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/process_execution_service.proto
```

### Running the service

You can now launch the service using:

```bash
TODO
```
