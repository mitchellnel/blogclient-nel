package main

import (
	"context"
	"fmt"
	"log"

	// general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// types from our blockchain blog
	"blockchainblognel/x/blockchainblognel/types"
)

func main() {
	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithHome("/Users/user/.blockchain-blog-nel"),
	)
	if err != nil {
		log.Fatalf("Error creating cosmosclient instance: %v", err)
	}

	fmt.Println(cosmos)

	accounts, err := cosmos.AccountRegistry.List()
	if err != nil {
		log.Fatalf("Error getting accounts: %v", err)
	}
	fmt.Println(accounts)
	fmt.Println("\n\n")

	// account "alice" already exists -- created during first instance of `ignite chains serve`
	accountName := "alice"

	// get account from the keyring by account name, and return a bech32 address
	accountAddress, err := cosmos.Address(accountName)
	if err != nil {
		log.Fatalf("Error getting address for account with name %v: %v", accountName, err)
	}

	// define a message to create a post
	msg := &types.MsgCreatePost{
		Creator: accountAddress.String(),
		Title:   "Hello!",
		Body:    "This is the first post...",
	}

	// broadcast a transaction from the account "alice" with the message to create a Post
	// store the response in txResp
	txResp, err := cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		log.Fatalf("Error broadcasting transaction to create post: %v", err)
	}

	// print response from tx broadcast
	fmt.Print("MsgCreatePost:\n\n")
	fmt.Println(txResp)

	// instantiate a query client for our `blockchainblognel`
	queryClient := types.NewQueryClient(cosmos.Context())

	// query the blockchain using the `posts` query to get all the Posts
	// store the posts in queryResp
	queryResp, err := queryClient.Posts(context.Background(), &types.QueryPostsRequest{})
	if err != nil {
		log.Fatalf("Error querying posts from blockchain: %v", err)
	}

	// print the response from querying the posts
	fmt.Print("\n\nAll Posts:\n\n")
	fmt.Println(queryResp)
}
