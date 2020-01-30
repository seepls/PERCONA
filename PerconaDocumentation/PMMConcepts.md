###### ClientServer Architecture 

Client-server model
PMM Client : install on every database that I want to monitor.
{ collects data }

PMM Server : presents 

Customer application is connected to a ----> database ---> Install PMM client on DB 
client sends data to server and the user can infer from data presented over the web interface.

###### PMM-Client Package

pmm-admin :

pmm-agent :

node_exporter : 

mysql_exporter :

mongodb_exporter : 

postgres_exporter : 

proxysql_exporter : 



pmm-client to pmm-server secure  : all exporter use   SSL/TLS encrypted connections,
and to server protected by HTTP basic authentication.

###### PMM-SERVER :

PMM-server runs on central monitoring host.
Distributed as an Alliance via :
Docker Image
OVA : run in Virtual box or another hypervisor

AMI : run via Amazon web services

  
###### Pmm-server tools 
-QAN :
   -- QAN API :  backend for storing and accessing query data collected by QAN agent running on PMM Client 
   -- QAN web app : visualizing collected Query Analytics data.
   -- ClickHouse :
   
###### Prometheus is a third-party time-series database that connects to 
exporters running on a PMM Client and aggregates metrics collected by the exporters
###### ClickHouse is a third-party column-oriented database that facilitates the Query Analytics functionality
###### Grafana is a third-party dashboard and graph builder for visualizing data aggregated by Prometheus in an intuitive web interface
######
