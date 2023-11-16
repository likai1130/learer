package main

import (
	"context"
	"fmt"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/core"
	"github.com/ipfs/kubo/core/coreapi"
	"github.com/ipfs/kubo/core/node"
	"os"
)

func main() {
	cid, err := preGeneraForFile("/Users/likai/hisun/gospace/src/learner/pkg/i18n/learni18n.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("file cid is ", cid)

	data := []byte("hello world")
	cid, err = preGeneraFor(data)
	if err != nil {
		panic(err)
	}
	fmt.Println("data cid is ", cid)

}

func preGeneraFor(data []byte) (string, error) {
	ipfsNode, err := core.NewNode(context.Background(), &node.BuildCfg{Online: false})
	if err != nil {
		return "", nil
	}

	api, err := coreapi.NewCoreAPI(ipfsNode)
	if err != nil {
		return "", nil
	}
	f := files.NewBytesFile(data)

	cid, err := api.Unixfs().Add(context.Background(), f, options.Unixfs.CidVersion(1))
	if err != nil {
		return "", nil
	}
	return cid.Root().String(), nil
}

func preGeneraForFile(filePath string) (string, error) {
	// Create a new IPFS node
	ipfsNode, err := core.NewNode(context.Background(), &node.BuildCfg{Online: false})
	if err != nil {
		return "", nil
	}

	// Get the core API
	api, err := coreapi.NewCoreAPI(ipfsNode)
	if err != nil {
		return "", nil
	}

	// Read the file into memory
	file, err := os.Open(filePath)
	if err != nil {
		return "", nil
	}

	f := files.NewReaderFile(file)
	// Add the file to IPFS
	cid, err := api.Unixfs().Add(context.Background(), f, options.Unixfs.CidVersion(1))
	if err != nil {
		return "", nil
	}

	return cid.Root().String(), nil
}
