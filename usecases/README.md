# Hypnogram Data Use Case

The `usecases` demonstrates a practical application of **SPADE** for hypnogram data.
The dataset comprises hypnogram data files and is utilized within a model involving
three key entities: Users, Server/Curator, and Analysts. To implement the client-server
model, we utilized standard gRPC and protocol buffer libraries.

## System Model

## Protocol Sequence Diagram

```mermaid
sequenceDiagram
    autonumber
    actor u as User(s)|id
    participant c as Curator/Server
    actor a as Analyst(s)
%% start the protocol 
    note over c: generate q, g <br> n=maximum number of users, <br> m=maximum number of elements
    c ->> c: msk, mpk = setup(n, m)
    c -->> u: (q, g, mpk)
    c -->> a: (q, g, mpk)
    u ->> u: generate random #alpha; <br> regKey = g^ #alpha; <br> ct = Enc(mpk, #alpha; , m)
    u -->> c: (id, regKey, ct)
    c ->> c: stores regKeys[id] = regKey <br> stores cts[id] = ct
%% analyst query
    a -->> c: (id, v)
    c ->> c: dkv = keyDer(id, v, msk, regKeys[id])
    c -->> a: (dkv, c[id])
%% end of protocol
```