module github.com/wcygan/kafka-on-kubernetes/consumer

go 1.21.6

require github.com/wcygan/kafka-on-kubernetes/generated/go v0.0.0

require google.golang.org/protobuf v1.32.0 // indirect

replace github.com/wcygan/kafka-on-kubernetes/generated/go => ../generated/go
