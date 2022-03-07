module github.com/laurentsenta/pl-gated-ipfs

require (
	github.com/ipfs/go-bitswap v0.5.1
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-ipfs v0.11.0
	github.com/ipfs/go-ipfs-files v0.0.9
	github.com/ipfs/interface-go-ipfs-core v0.5.2
	github.com/libp2p/go-libp2p-core v0.11.0
	github.com/multiformats/go-multiaddr v0.4.1
)

replace github.com/ipfs/go-ipfs => ./../../ipfs/go-ipfs
replace github.com/ipfs/go-bitswap => ./../../ipfs/go-bitswap

go 1.17