load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
helm_repo('bitnami', 'https://charts.bitnami.com/bitnami')

# --------------------------------- Kafka --------------------------------- #

helm_resource(
  name='kafka',
  chart='bitnami/kafka',
  resource_deps=['bitnami'],
  flags=[
    '--values=./kafka/values.yaml',
    '--version=26.8.4',
  ]
)

# --------------------------------- Protobuf --------------------------------- #

# Compile command for the protobuf files
proto_compile_cmd = 'buf generate'

# Local resource to compile the protobuf files
local_resource(
  'proto_compile',
  proto_compile_cmd,
  deps=['./proto'],
)

# --------------------------------- Producer --------------------------------- #

# Compile command for the Go application
producer_compile_cmd = 'CGO_ENABLED=0 GOOS=linux go build -o ../build/producer .'

# Local resource to compile the Go application
local_resource(
  'producer_compile',
  producer_compile_cmd,
  deps=['./producer', './generated'],
  dir='./producer'
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

# --------------------------------- Consumer --------------------------------- #

# Compile command for the Go consumer application
consumer_compile_cmd = 'CGO_ENABLED=0 GOOS=linux go build -o ../build/consumer .'

# Local resource to compile the Go consumer application
local_resource(
  'consumer_compile',
  consumer_compile_cmd,
  deps=['./consumer', './generated'],
  dir='./consumer'
)

# Docker build with restart for the Go consumer application
docker_build_with_restart(
    'wcygan/kafka-on-kubernetes-consumer',
    '.',
    dockerfile='consumer/Dockerfile',
    entrypoint='/app/consumer',
    ignore=['./scripts', '.gitignore'],
    live_update=[
        sync('./', '/app'),
    ],
)

# --------------------------------- Resources --------------------------------- #

k8s_yaml([
  'producer/Deployment.yaml',
  'consumer/Deployment.yaml',
  'admin-dashboard.yaml'
])