# dSock Go Client

Go client for [dSock](https://github.com/Cretezy/dSock).

[GitHub](https://github.com/Cretezy/dSock-go)

## Installation

```bash
go get github.com/Cretezy/dSock-go
```

## Usage

```go
import "github.com/Cretezy/dSock-go/v0"

dSockClient := dsock.NewClient(dSockUrl, dSockToken)
```

### [Create claim](https://github.com/Cretezy/dSock#claims)

Create a claim for user authentication.

```go
claim, err := dSockClient.CreateClaim(CreateClaimOptions{
    User: "user",
    // optional
    Session:  "session",
    Id:       "id",
    Channels: []string{"channel"},
    // either or
    Duration:   30,         // in seconds
    Expiration: 1600000000, // in seconds since epoch
})
```

### [Send message](https://github.com/Cretezy/dSock#sending-message)

Send a message to a target (one or many clients).

```ts
err := dSockClient.SendMessage(SendMessageOptions{
    Type:    "text", // or `binary`
    Message: []byte("hello world"),
    // target (choose one or many)
    Target: Target{
        User:    "user",
        Session: "session",
        Id:      "id",
        Channel: "channel",
    },
})
```

### [Disconnecting](https://github.com/Cretezy/dSock#disconnecting)

Disconnect a target (one or many clients).

```ts
err := dSockClient.Disconnect(DisconnectOptions{
    // target (choose one or many)
    Target: Target{
        User:    "user",
        Session: "session",
        Id:      "id",
        Channel: "channel",
    },
})
```

### [Info](https://github.com/Cretezy/dSock#info)

Get claim and connection info from a target (one or many clients).

```ts
info, err := dSockClient.GetInfo(GetInfoOptions{
    // target (choose one or many)
    Target: Target{
        User:    "user",
        Session: "session",
        Id:      "id",
        Channel: "channel",
    },
})
```

### [Channels](https://github.com/Cretezy/dSock#channels)

Subscribe/unsubscribe a target to a channel (one or many clients).

```ts
err := dSockClient.SubscribeChannel(ChannelOptions{
    Channel: "new_channel",
    // Should claims also be subscribed?
    IgnoreClaims: true,
    // target (choose one or many)
    Target: Target{
        User:    "user",
        Session: "session",
        Id:      "id",
        Channel: "channel",
    },
})


err := dSockClient.UnsubscribeChannel(ChannelOptions{
    Channel: "remove_channel",
    // Should claims also be unsubscribed?
    IgnoreClaims: true,
    // target (choose one or many)
    Target: Target{
        User:    "user",
        Session: "session",
        Id:      "id",
        Channel: "channel",
    },
})
```

## License

[MIT](./LICENSE)
