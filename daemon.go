package main

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type IdentifyOutput struct {
}

// Runs the full check for a cid and content id.
func runAdd(writer http.ResponseWriter, uristr string) error {
	u, err := url.ParseRequestURI(uristr)
	if err != nil {
		return err
	}

	var newItem AllowListItem
	validParams, validRule := false, false

	if u.Query().Has("peer") {
		x, err := peer.IDFromString(u.Query().Get("peer"))
		if err != nil {
			return err
		}
		newItem.PeerID = &x
		validParams = true
	}

	if u.Query().Has("cid") {
		x, err := cid.Decode(u.Query().Get("cid"))
		if err != nil {
			return err
		}
		newItem.Cid = &x
		validParams = true
	}

	if u.Query().Has("allow") {
		x, err := peer.IDFromString(u.Query().Get("peer"))
		if err != nil {
			return err
		}
		newItem.PeerID = &x
		validParams = true
	}

	if u.Query().Has("deny") {
		b := u.Query().Get("deny") == "true"
		newItem.Deny = &b
		validRule = !validRule
	}
	if u.Query().Has("allow") {
		b := u.Query().Get("allow") == "true"
		newItem.Allow = &b
		validRule = !validRule
	}

	if !validRule || !validParams {
		return errors.New("invalid input")
	}

	alf.Items = append(alf.Items, newItem)

	return nil
}

func runRemove(writer http.ResponseWriter, uristr string) error {
	u, err := url.ParseRequestURI(uristr)
	if err != nil {
		return err
	}

	var newItem AllowListItem
	validParams, validRule := false, false

	if u.Query().Has("peer") {
		x, err := peer.IDFromString(u.Query().Get("peer"))
		if err != nil {
			return err
		}
		newItem.PeerID = &x
		validParams = true
	}

	if u.Query().Has("cid") {
		x, err := cid.Decode(u.Query().Get("cid"))
		if err != nil {
			return err
		}
		newItem.Cid = &x
		validParams = true
	}

	if u.Query().Has("allow") {
		x, err := peer.IDFromString(u.Query().Get("peer"))
		if err != nil {
			return err
		}
		newItem.PeerID = &x
		validParams = true
	}

	if u.Query().Has("deny") {
		b := u.Query().Get("deny") == "true"
		newItem.Deny = &b
		validRule = !validRule
	}
	if u.Query().Has("allow") {
		b := u.Query().Get("allow") == "true"
		newItem.Allow = &b
		validRule = !validRule
	}

	if !validRule || !validParams {
		return errors.New("invalid input")
	}

	newItems := make([]AllowListItem, 0, len(alf.Items))
	found := false

	for _, item := range alf.Items {
		if item.Cid == newItem.Cid && item.PeerID == newItem.PeerID && item.Allow == newItem.Allow && item.Deny == newItem.Deny {
			found = true
			continue
		}
		newItems = append(newItems, item)
	}

	if !found {
		return errors.New("no rule matched")
	}

	alf.Items = newItems
	return nil
}

func runList(writer http.ResponseWriter, uristr string) (interface{}, error) {
	return alf, nil
}
