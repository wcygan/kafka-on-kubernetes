module github.com/wcygan/kafka-on-kubernetes/consumer-2

go 1.21.6

require (
	github.com/segmentio/kafka-go v0.4.47
	github.com/wcygan/kafka-on-kubernetes/generated/go v0.0.0
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)

replace github.com/wcygan/kafka-on-kubernetes/generated/go => ../generated/go
