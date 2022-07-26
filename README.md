# blogclient-nel

This is a Go client designed to be used to interact with the blockchain blog app found [here](https://github.com/mitchellnel/blockchain-blog-nel).

## Getting started

First, clone this repository:

```
git clone https://github.com/mitchellnel/blogclient-nel
```

Then, get the necessary Go packages by invoking:

```
go mod tidy
```

After that, follow the getting started steps found [here](https://github.com/mitchellnel/blockchain-blog-nel). Specifically, the blockchain app have been started up and running in a separate terminal window.

Once this is done, simply invoke:

```
go run main.go
```

This will create a post and query the blockchain to display existing posts. The results of the query will be output to the user.
