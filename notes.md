# installation

## Protocol  Buffers

https://gist.github.com/diegopacheco/cd795d36e6ebcd2537cd18174865887b

go get -u google.golang.org/grpc 

protoc -I=/home/dave/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I=/home/dave/Documents/spacemesh/test --go_out=pb api.proto