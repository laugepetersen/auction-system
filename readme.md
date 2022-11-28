# Guide

## How to run program

Please note the following

- The primary manager should be set up first following the other managers. And the afterwards the clients
- The system is setup in a way that we can have three managers and three clients.
- A package.json has been setup to ease the process. The following shows how to start the program with and without the package.json.
- Each line should be used in a new terminal window/tab.
- All commands should be called from the root folder

### Managers

With package.json

```bash
npm run m1
npm run m2
npm run m3
```

Without package.json

```bash
go run manager/main.go 6000 6000 6001 6002 4000 4001 4002
go run manager/main.go 6000 6001 6002 6000 4000 4001 4002
go run manager/main.go 6000 6002 6000 6001 4000 4001 4002
```

The commands explained:

```bash
go run manager/main.go <primary manager port> <manager port 1> <manager port 2> <manager port 3> <client port 1> <client port 2> <client port 3>
```


### Clients

With package.json

```bash
npm run c1
npm run c2
npm run c3
```

Without package.json

```bash
go run frontend/main.go 6000 4000
go run frontend/main.go 6000 4001
go run frontend/main.go 6000 4002
```

The commands explained:

```bash
go run frontend/main.go <primary manager port> <own port>
```

## Auction

The following can be typed from the client, an info message is shown afterwards

### Bid

```bash
bid <int amount>
```

### Get result

```bash
result
```

## Simulate crash

To simulate a crash of a manger on mac do "ctrl+c" or simply force quit the one terminal
