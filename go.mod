module emilianobonassi/bsc-hack-analysis

go 1.15

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/iavl v0.12.0
	github.com/tendermint/tendermint v0.31.11
)

replace (
	github.com/etcd-io/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
	github.com/gogo/protobuf v1.1.1 => github.com/gogo/protobuf v1.3.2
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2
)
