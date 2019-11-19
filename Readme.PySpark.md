## PySpark configuration

In order to process the information being loaded into Cassandra, we choose spark to access it.

A few steps are needed in order to get Jupyter notebooks and Pyspark with Cassandra database working.


***DEFAULT JDK SHOULD BE 1.8***

Download Spark 2.4.4

```bash

 wget http://apache.dattatec.com/spark/spark-2.4.4/spark-2.4.4-bin-hadoop2.7.tgz

 tar zxvf spark-2.4.4-bin-hadoop2.7.tgz

 cd spark-2.4.4-bin-hadoop2.7/
 
```

We need to change default configuration of Spark in order to work with Cassandra Driver


```bash
    vi conf/spark-defaults.conf.template

```

The content should look like this

```bash
#
# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Default system properties included when running spark-submit.
# This is useful for setting default environmental settings.

# Example:
# spark.master                     spark://master:7077
# spark.eventLog.enabled           true
# spark.eventLog.dir               hdfs://namenode:8021/directory
# spark.serializer                 org.apache.spark.serializer.KryoSerializer
# spark.driver.memory              5g
# spark.executor.extraJavaOptions  -XX:+PrintGCDetails -Dkey=value -Dnumbers="one two three"
spark.jars.packages     com.datastax.spark:spark-cassandra-connector_2.11:2.4.1 

```

The last line allows to add the Cassandra connector.

With this configuration we can start the single node spark as master

```bash

spark-2.4.4-bin-hadoop2.7/sbin$ ./start-master.sh 

starting org.apache.spark.deploy.master.Master, logging to spark-2.4.4-bin-hadoop2.7//logs/spark-users-org.apache.spark.deploy.master.Master-1-osboxes.out 
```

with pip3 installed we can install and start Jupyter notebook

```bash

pip3 install jupyter

jupyter notebook

```

This configuration on a Ubuntu 19.04 should work and we should get a good Development enviroment.


Note: In addition the binary format the loader and REST API makes it easier to get the data enviroment up and running with minimal configuration.

