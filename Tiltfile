load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')

helm_resource(
  name='kafka',
  chart='bitnami/kafka',
  resource_deps=['bitnami'],
  flags=[
    '--values=./kafka/values.yaml',
    '--version=26.8.4',
  ]
)

# Compile command for the Go application
compile_cmd = 'CGO_ENABLED=0 GOOS=linux go build -o ../build/producer .'

# Compile command for the protobuf files
proto_compile_cmd = 'buf generate'

# Local resource to compile the Go application
local_resource(
  'compile',
  compile_cmd,
  deps=['./producer', './generated'],
  dir='./producer'
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
    entrypoint='/app/producer',
    ignore=['./scripts', '.gitignore'],
    live_update=[
        sync('./', '/app'),
    ],
)

# Kubernetes resources
k8s_yaml(['producer/Deployment.yaml', 'admin-dashboard.yaml'])