CREATE KEYSPACE epicfda WITH replication = {
	'class' : 'NetworkTopologyStrategy',
	'datacenter1' : 1,
	'datacenter2' : 1
};