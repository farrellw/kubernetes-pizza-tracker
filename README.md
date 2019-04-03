# pizza-tracker

Simple Client/API/Database that tracks when the last time we had pizza in the lab. Managed by Kubernetes

## Kubernetes moving pieces

### Client

Client consists of four pods across two deployments. The first deployment, shown in `client.yaml` deploys 3 pods of the blue version. `client-red.yaml` deploys 1 pod of the red version. Hitting the external service endpoint, configured in `service.yaml` will route you to one of the 4 pods, meaning you have a 1 in 4 chance of hitting the red version.
Client uses a reverse NGINX proxy to talk to the API. Requests to `/api` proxy to an internal kubernetes IP (a service that fronts the API), which is then forwarded to the API pods.

### API

One deployment, two pods defined in `api.yaml`. Fronted by the service in `service.yaml`.
Reads environment variables to contact the cosmosDB database, which is hosted outside the kubernetes framework.
