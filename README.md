# go-xchange

Collects continuously (with time interval) currency data and transfers it to InfluxDB.

- On failed influx connection it saves fetch result to file for data recovery
- Configuration via yml file