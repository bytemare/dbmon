
# dbmon

A cluster monitoring tool.

[![Build Status](https://travis-ci.com/bytemare/dbmon.svg?branch=master)](https://travis-ci.com/bytemare/dbmon)
[![Go Report Card](https://goreportcard.com/badge/github.com/bytemare/dbmon)](https://goreportcard.com/report/github.com/bytemare/dbmon)
[![codebeat badge](https://codebeat.co/badges/7e86ba65-e7b9-4982-9996-6b42c0eb763e)](https://codebeat.co/projects/github-com-bytemare-dbmon-master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/2f906e2104b24db88596a918c25e59e6)](https://www.codacy.com/app/bytemare/dbmon?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=bytemare/dbmon&amp;utm_campaign=Badge_Grade)
[![GolangCI](https://golangci.com/badges/github.com/bytemare/dbmon.svg)](https://golangci.com/r/github.com/bytemare/dbmon)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bytemare_dbmon&metric=security_rating)](https://sonarcloud.io/dashboard?id=bytemare_dbmon)
[![Docker](https://images.microbadger.com/badges/image/bytemare/dbmon:dbmon.server.v0.svg)](https://microbadger.com/images/bytemare/dbmon:dbmon.server.v0)
[![GoDoc](https://godoc.org/github.com/bytemare/dbmon?status.svg)](https://godoc.org/github.com/bytemare/dbmon)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fbytemare%2Fdbmon.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fbytemare%2Fdbmon)

## Testing

The platform is _functional_. To see it live, you must first run a cluster, see instructions [here](https://www.cockroachlabs.com/docs/stable/start-a-local-cluster-in-docker.html). (_NB: If you changed the port number used in the instructions, please adapt it in the server test program : app/dbmon.go_)

Run the server in a docker container or compile and run it yourself :

    $ docker run bytemare/dbmon:dbmon.server.v0
    
    or
    
    $ go run Examples/server/dbmon.go

Then, in a different terminal

    $ go run Tests/client/dbmon-client.go

These are only showcase programs, and don't yet implement all desired functionalities (See Roadmap below).

To stop the server gracefully, interrupt with ctrl+c.  

> Fancy sparkling pew-pew badges
___
> What it does
___

> Explanation draft
>
> - Simplicity in use
>   - Most of the stuff runs in containers. Yay !
>   - When using existing connectors, simply fill yaml config file
>     - yaml config file holds parameters such as
>       - cluster names, ips, etc.
>       - requests to be sent, etc
>   - If connector for your db does not exist, it's easy to build one for your needs

## Installation / Prerequisites

> client
>
> - shell / ps
> - go
>
___
> server
>
> - docker
> - connector for your endpoint
>
___
> cluster
>
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

Think of the platform as a cache, or a broker. It continuously fetches data from your cluster(s), holds them warm, and relays it back to you.
The only thing you need is a compatible connector, the connection plug that links the platform to your cluster.

The Architecture can be represented like this

         Client  Client  Client             
             \     |     /                  
              \    |    /                   
               \   |   /                    
                \--|--/------ gRPC          
                 Server                     
                   |          
                   |                        
                   |           
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

- The client is a piece of software your machine, or your monitoring tool.

## Use / Adaptation

### Connecting dbmon to your cluster

> step-by-step explanation.

### Create a new connector for your needs

> step-by-step explanation.

## Roadmap

v0 :

- [x] Agents are individual go routines that use a connector to communicate with a cluster ( agent 1:1 cluster )
- [x] A connector defines the requests to be send to a type of cluster, and implements the a client connection to it
- [x] Connectors retrieve health status from clusters
- [x] When an agent gets a response, it hands it to the collector
- [x] The collector sends every result to the server as a report : A report is the response for a call to the cluster
- [x] Server holds a cache : a key:value map, that associates an identified cluster to a list of reports.
- [x] Client operates unary RPC from to server and gets all the reports in a single response
- [ ] Purge data to only keep desired ones ( in the connector ?) : CockroachDB
  - [x] Cluster's status
  - [ ] Useful information about health
    - [ ] Ranges / Replicas / Nodes
      - [x] Number of ranges
      - [x] Number of nodes that are up
      - [ ] decommissioned or shutting down
        - [x] Capacity usage ( in % )
        - [ ] Memory usage ( in % )
        - [ ] Unavailable ranges
        - [x] Queries per second
        - [ ] Heartbeat Latency: 99th percentile
        - [x] Mean clock offset with other nodes in nanoseconds

        More ? :
        - [ ] Round-trip time to reach the cluster from the platform
        - [ ] Number of pending tasks
        - [ ] Oldest task (time the earliest initiated, still pending, task has been waiting for)
- [x] Server/Collector are bundled and Dockerised
- [x] Testing
- [ ] Coverage
- [ ] Better logging, and to a file (but not for docker)
- [ ] Better documentation

v0.1 : Separate units and configuration independent boot

- [ ] Add a cluster through client (client tells monitor which cluster/s to connect to)
- [ ] Add a lifetime to cache : time based or buffer size (last n probes)
- [ ] Make server/collector as agnostic as possible on launch (only do specific things as demanded by client)
- [ ] Add parameters as command line arguments (ports, etc.)

v0.2 : Further separation and make use of streams

- [ ] Separate server and collector, and implement a gRPC interface between them
- [ ] Continuously stream reports from the collector to the server
- [ ] Pub/Sub pattern in gRPC between client and server : the server streams continuously reports to the client

v0.3 : Performance and extendability

- [ ] Operate a better cache system, maybe using a database
- [ ] Add more connectors

v0.4 : Authentication

- [ ] authentication client <-> server (mTLS ? WebAuthN + macaroons ?)
- [ ] authentication server <-> collector (set up instance pki ?)
- [ ] Authentication collector <-> cluster (no idea, yet)

v0.5 : Client enhancement

- [ ] Clients can send commands, like 'health' to get a healthcheck, 'info' for general info, and more ad-hoc commands
    ( a command must then be registered in the connector )
- [ ] Configure client with yaml file ?
