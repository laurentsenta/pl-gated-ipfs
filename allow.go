package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ipfs/go-bitswap"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

var alf AllowListFile

// An AllowListFile defines the rule for the bitswap replication,
//
// Example:
// {
// 	"items": [
// 	  {
// 		"cid": {
// 		  "/": "bafybeiezp42osao4gtmytbi7k6xsltuiywozrwumelllwldcv35twuhdfe"
// 		},
// 		"peerID": "12D3KooWJS1mDHAVcEi9wkyiBcUGcbc8JVJf2ojN3YQdfpXSWsay",
// 		"allow": true
// 	  },
// 	  {
// 		"peerID": "12D3KooWJS1mDHAVcEi9wkyiBcUGcbc8JVJf2ojN3YQdfpXSWsay",
// 		"allow": true
// 	  },
// 	  {
// 		"cid": {
// 		  "/": "bafybeihcyruaeza7uyjd6ugicbcrqumejf6uf353e5etdkhotqffwtguva"
// 		},
// 		"deny": true
// 	  }
// 	]
// }
type AllowListFile struct {
	Items []AllowListItem `json:"items"`
}

type AllowListItem struct {
	Cid    *cid.Cid `json:"cid,omitempty"`
	PeerID *peer.ID `json:"peerID,omitempty"`
	Allow  *bool    `json:"allow,omitempty"`
	Deny   *bool    `json:"deny,omitempty"`
}

func (a *AllowListItem) String() string {
	result := make([]string, 0, 4)

	if a.Cid != nil {
		result = append(result, fmt.Sprintf("cid: %s", *a.Cid))
	}
	if a.PeerID != nil {
		result = append(result, fmt.Sprintf("peerID: %s", *a.PeerID))
	}
	if a.Allow != nil {
		result = append(result, fmt.Sprintf("allow: %s", *a.Allow))
	}
	if a.Deny != nil {
		result = append(result, fmt.Sprintf("deny: %s", *a.Deny))
	}

	return fmt.Sprintf("{%s}", strings.Join(result, ","))
}

func loadAllowListJSON(path string) (AllowListFile, error) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	var alf AllowListFile
	err = json.Unmarshal(byteValue, &alf)
	if err != nil {
		panic(err)
	}

	return alf, nil
}

func loadPeerBlockRequestFilter(path string) bitswap.PeerBlockRequestFilter {
	alf, err := loadAllowListJSON(path)
	if err != nil {
		panic(err)
	}

	// fmt.Println("alf:", alf)

	return func(p peer.ID, c cid.Cid) bool {
		matched := make([]*AllowListItem, 3)

		// log.Printf("(peer: %v, cid: %v) Evaluating filter with %d rules.\n", p, c, len(alf.Items))

		for i, rule := range alf.Items {
			// log.Printf("(peer: %v, cid: %v) Evaluating rule: %s.\n", p, c, &rule)
			precision := 0

			if rule.Cid != nil {
				if !rule.Cid.Equals(c) {
					continue
				}
				// log.Printf("(peer: %v, cid: %v) cid matched.\n", p, c)
				precision++
			}

			if rule.PeerID != nil {
				if *rule.PeerID != p {
					continue
				}
				// log.Printf("(peer: %v, cid: %v) peerID matched.\n", p, c)
				precision++
			}

			if matched[precision] == nil && precision > 0 {
				// Note: careful here, if you use &rule it'll point to the iterator and break.
				matched[precision] = &alf.Items[i]
			}
		}

		for i := len(matched) - 1; i >= 0; i-- {
			if matched[i] != nil {
				log.Printf("(peer: %v, cid: %v) Evaluating filter with %d rules.\n", p, c, len(alf.Items))
				r := isAllow(*matched[i])
				log.Printf("(peer: %v, cid: %v) rule `%s' matched %d parameters, access allowed: %v.\n", p, c, matched[i], i, r)
				return r
			}
		}

		return true
	}
}

func isAllow(rule AllowListItem) bool {
	hasAllow, hasDeny := rule.Allow != nil, rule.Deny != nil

	if hasAllow == hasDeny {
		panic(fmt.Sprintf("invalid rule: `%v' it should define at most one allow or deny.", rule))
	}

	if rule.Allow != nil {
		return *rule.Allow
	}

	return !*rule.Deny
}
