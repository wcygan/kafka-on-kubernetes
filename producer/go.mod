module github.com/wcygan/kafka-on-kubernetes/producer

go 1.21.6

require (
	github.com/segmentio/kafka-go v0.4.47
	github.com/wcygan/kafka-on-kubernetes/generated/go v0.0.0
)

require (
	github.com/klauspost/compress v1.17.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/wcygan/kafka-on-kubernetes/generated/go => ../generated/go
