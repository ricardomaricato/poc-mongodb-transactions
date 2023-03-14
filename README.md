<div style="text-align: center;">
  <h1 align="center">MongoDB Transactions POC</h1>
  <p align="center">
    <img src=".images/cluster-mongodb.png"  width="764" />
  </p>  
</div>
<br />

### Using Docker to Deploy a MongoDB Cluster/Replica Set
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