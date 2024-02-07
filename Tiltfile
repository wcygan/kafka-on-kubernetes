load('ext://restart_process', 'docker_build_with_restart')

# Compile command for the Go application
compile_cmd = 'CGO_ENABLED=0 GOOS=linux go build -o ../build/producer .'

# Compile command for the protobuf files
proto_compile_cmd = 'buf generate'

# Local resource to compile the Go application
local_resource(
  'compile',
  compile_cmd,
  deps=['./producer', './generated'],
  dir='./producer'  # Set the working directory for the command
)

# Local resource to compile the protobuf files
local_resource(
  'proto_compile',
  proto_compile_cmd,
  deps=['./proto'],
)

# Docker build with restart for the Go application
docker_build_with_restart(
    'wcygan/kafka-on-kubernetes-producer',
    '.',
    dockerfile='producer/Dockerfile',
    entrypoint='/app/build/producer',
    ignore=['./scripts', '.gitignore'],
    live_update=[
        sync('./', '/app'),
    ],
)

# Kubernetes resources
k8s_yaml(['producer/Deployment.yaml', 'kafka/deployment.yaml', 'kafka/topics/packet.yaml'])