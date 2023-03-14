<div style="text-align: center;">
  <h1 align="center">MongoDB Transactions POC</h1>
  <p align="center">
    <img src=".images/cluster-mongodb.png"  width="764" />
  </p> 
  <h1 align="center">Using Docker to Deploy a MongoDB Cluster/Replica Set</h1>  
</div>
<br />

### Create a Docker Network

```
$ docker network create mongoCluster
```

### Start MongoDB Instances and Initiate the Replica Set

```
$ ./startdb.sh
```

### Verify the Replica Set

```
$ docker exec -it mongo1 mongo --eval "rs.status()"
```

### Running

```
$ go run main.go
```