# credentiald

credentiald manages credentials for cloud environments.

## Development

To run against a local Minikube:
```
./credentiald daemon \
    --service.kubernetes.incluster=false \
    --service.kubernetes.address=https://$(minikube ip):8443 \
    --service.kubernetes.tls.cafile=~/.minikube/ca.crt \
    --service.kubernetes.tls.crtfile=~/.minikube/client.crt \
    --service.kubernetes.tls.keyfile=~/.minikube/client.key
```

And to create a credential:
```
curl -Ss -v \
    http://localhost:8000/v4/organizations/foobar/credentials/ \
    -d '{
            "provider": "aws",
            "aws": {
                "roles": {
                    "admin": "arn...",
                    "awsoperator": "arn..."
                }
            }
        }' | jq
```
