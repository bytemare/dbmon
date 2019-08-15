# dbmon
A cluster monitoring tool.

> Fancy sparkling pew-pew badges
___
> What it does
___


> - Simplicity in use
>    - Most of the stuff runs in containers. Yay !
>    - When using existing connectors, simply fill yaml config file
>        - yaml config file holds parameters such as
>            - cluster names, ips, etc.
>            - requests to be sent, etc
>    - If connector for your db does not exist, it's easy to build one for your needs

## Installation / Prerequisites

> client
> - PS / shell
> - go
___
> server
> - docker
> - connector for your endpoint
___
> cluster
> - API

## Run

If you are starting from scratch, you first may want to have your cluster up and running. If you don't have one, there's a script for that:
> - give shell script to pop a cluster locally

Then, get your server running :
> - give it some info / configuration
> - start it locally or in the cloudz

Finally, start the client :
> - give it info to connect to server (ip, key)
> - start go client

## HLD

The Architecture can be represented like this

    
         Client  Client  Client
             \     |     /
              \    |    /
               \   |   /
                \--|--/------ gRPC
                 Server 
                   |--------- gRPC
                   |
                   | ---- Cache ?
                   |
                Collector
         __________|__________
         Agent   Agent   Agent
           |       |       |
           w       w       w  <-- Connectors  
           |       |       |
    CockroachDB  Istio   Whatever

You can monitor whatever you want, as long as it exposes an API and you have a connector for it.

Some connectors are already implemented (_for now CockroachDB only_).

A connector is a simple piece of code of go implementing the Connector interface, that runs with the Collector.

- Server and Collector are services that run in containers, so you can run them on the same machine or distribute them.

- The client is a go program, so you'll need a working go environment.

## Use / Adaptation

### Connecting dbmon to your cluster

step-by-step explanation.

### Create a new connector for your needs

step-by-step explanation.

## Todo

- well ... make it work, duh.
- authentication collector <-> cluster (no idea, yet)
- authentication server <-> collector (set up instance pki ?)
- authentication client <-> server (webauthn + macaroons ?)

## Ideas

v0 :
- [x] Agents are individual go routines that use a connector to communicate with a cluster ( agent 1:1 cluster )
- [x] A connector defines the requests to be send to a type of cluster, and implements the a client connection to it
- [ ] Connectors retrieve health status from clusters
- [x] When an agent gets a response, it hands it to the collector
- [x] The collector sends every result to the server as a report : A report is the response for a call to the cluster
- [x] Server holds a cache : a key:value map, that associates an identified cluster to a list of reports.
- [ ] Client operates unary RPC from to server and gets all the reports in a single response
- [ ] Server/Collector are bundled and Dockerised

v1 :
- [ ] Pub/Sub pattern in gRPC between client and server : the server streams continuously reports to the client
- [ ] Separate server and collector, and implement a gRPC interface between them
- [ ] Continuously stream reports from the collector to the server 
- [ ] Configure with yaml files

v2 :
- [ ] Operate a better cache system, maybe using a database
