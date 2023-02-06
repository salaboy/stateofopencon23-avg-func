# stateofopencon23-avg-func
Function using Dapr

Steps: 

- `docker login`
- `gcloud init`
- `gcloud connect`
- `dapr components -k`
- `func create -l go -t dapr -r https://github.com/salaboy/func`
- Change `statestore` to `statestore-redis` in handle.go
- Add annotations
```
  annotations:
    dapr.io/app-id: avg
    dapr.io/app-port: "8080"
    dapr.io/enabled: "true"
    dapr.io/metrics-port: "9099"
```
- `func deploy -v -r docker.io/salaboy`
