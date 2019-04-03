# pizza-tracker

Simple Client/API/Database that tracks when the last time we had pizza in the lab. Managed by Kubernetes

## Client

Client consists of 4 deployed applications. 3 of the applications are version blue. 1 of the applications are version red. Hitting the external service endpoint will route you within all 4 services.
Client talks a service which fronts the API layer. It does so via NGINX reverse proxy from `/api` to the services internal kubernetes IP

## API

Two kubernetes deployments.
Reads environment variables to contact the cosmosDB database, which is hosted outside the kubernetes framework.
