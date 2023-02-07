# State of Open Con :: 2023 Average Function

Function using Dapr

Steps to connect with cloud accounts: 

- `docker login`
- `gcloud init`
- `gcloud connect ...`

Check that you are connected to the correct enviornment: 
- `kubectl get nodes`
- `dapr components -k`

Create a function and deploy it: 

- `func create -l go -t dapr -r https://github.com/salaboy/func`
- Change `statestore` to `statestore` in handle.go
- Add annotations
```
  annotations:
    dapr.io/app-id: avg
    dapr.io/app-port: "8080"
    dapr.io/enabled: "true"
    dapr.io/metrics-port: "9099"
```
- `func deploy -v -r docker.io/salaboy`

