# Private/Whitelisted Content on IPFS

Shows a demo for:

- https://github.com/ipfs/go-ipfs/issues/8763

Relies on branches:

- https://github.com/ipfs/go-bitswap/pull/549
- https://github.com/ipfs/go-ipfs/pull/8772

Follow up from our js poc @ https://github.com/kylehuntsman/private-ipfs-content

## What is this?

This takes the [go-ipfs example](https://github.com/ipfs/go-ipfs/tree/master/docs/examples/go-ipfs-as-a-library)
and shows how to pass a "peer block filter" options.

When you start the node it will add files from the `./my-test-folder` folder.

The node also provides an API so you can try to protect your content & allow access to specific nodes:

```
# this will "deny" every incoming request for cid: Qmaz..Qb
curl http://localhost:4444/add?deny=true&cid=QmazWyBg8HQ1agguv6XM9XGCUWH8JsrmMctBZtkromnuQb

# this will "allow" every incoming request for cid: Qmaz..Qb & peer: QmNm..8u
curl http://localhost:4444/add?deny=true&cid=QmazWyBg8HQ1agguv6XM9XGCUWH8JsrmMctBZtkromnuQb&peer=QmNmFPULodfEnSE1jgfx5g1rRjMMaUJWVYtQGEuRhbHS8u

# show all your active rules
http://localhost:4444/list

# remove a rule (id is it's index in the list)
http://localhost:4444/remove?id=42
```

## Setup

Clone the wip implementations for bitswap and ipfs at:

- github.com/laurentsenta/go-bitswap/commits/44432bc3448a50ff3a013b2b367513c6f1203a18
- github.com/laurentsenta/go-ipfs/commis/fb9626878058b5acbd968812e0ce71379a0d28bf

Then make sure you rewrite both modules with your local local, use something like:

```sh
go mod edit -replace github.com/ipfs/go-bitswap=../go-bitswap             
go mod edit -replace github.com/ipfs/go-ipfs=../go-ipfs                               
```

Then run `make`, this will start your node, print a few notes on the CLI and start the API.

## Todo

- [ ] Merge this feature into master
- [ ] make this demo thread-safe