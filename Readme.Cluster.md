  ## DOCKER Cassandra Cluster Configuration

  Cluster is based on Apache Cassandra Cluster running on Docker, simulating a Datacenter segregation and sharding

  ```bash
  
  docker run --name cas1 -p 9042:9042 -e CASSANDRA_CLUSTER_NAME=MyCluster -e CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch -e CASSANDRA_DC=datacenter1 -d cassandra
  docker inspect --format='{{ .NetworkSettings.IPAddress }}' cas1
  ```

  We run cassandra 2 and Cassandra 3 nodes 

  ```bash
  docker run --name cas2 -e CASSANDRA_SEEDS="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' cas1)" -e CASSANDRA_CLUSTER_NAME=MyCluster -e CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch -e CASSANDRA_DC=datacenter1 -d cassandra
  
  
  docker run --name cas3 -e CASSANDRA_SEEDS="$(docker inspect --format='{{ .NetworkSettings.IPAddress }}' cas1)" -e CASSANDRA_CLUSTER_NAME=MyCluster -e CASSANDRA_ENDPOINT_SNITCH=GossipingPropertyFileSnitch -e CASSANDRA_DC=datacenter2 -d cassandra

  ```
  

  We can check the nodes by 
  
  
  ``` bash
  docker exec -ti cas1 nodetool status

  ```
  
We can access CSQL terminal by executing 

``` bash 

  docker exec -ti cas1 cqlsh

```


Check folder extras/database for SQL scripts for creating schema and tables needed.

