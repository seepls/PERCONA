##### Data retention :
Prometheus stores time-series data for 30 days, and QAN stores query data for 8 days 
Adjust data retention time as per need . " Via Settings dashboard"


##### Data Source Name: DSN 

database server attribute found on QAN page
" how PMM conects to selected database"


###### Grand total time:
% of time that database server spent running a specific query , compared to total time it ran for all query.

###### External Monitoring Service 
Monitoring service bound to Prometheus Exporter 
Asa service added , can set up metrics monitor to display the graphs.

###### Metrics 
Series of data visualized in PMM

###### Metrics Monitor
Component of PMM SERVER provides historical view of metrics critical to MySQL server instance 

###### Monitoring service
Service that collects information from database instance where PMM Client is installed .
pmm-admin add command .

###### PMM 
pmm - admin : chnages configurtion of pmm client . 
pmm-client : Collects MYSQL server metrics, general system metrics, and query analytics data for complete performance view.

Collected data sent to pmm-server.

pmm -server : aggregates data collected by pmm-client. presents in form of tables, graphs etc on the web interface.

##### QAN :
Quality analytics enables analysis of MySQL query performance.
