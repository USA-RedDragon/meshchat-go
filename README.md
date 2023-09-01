# meshchat-go

A small Go implementation of MeshChat to be deployed anywhere.

## Commands

`meshchat sync` - run the sync daemon

`meshchat server` - run the http server

## Routes from Lua version

`/cgi-bin/meshchat?action=meshchat_nodes&zone_name=MeshChat-OK`

returns

```text
KI5VMF-MINI     8080
KD5TDK-5        8080
```
