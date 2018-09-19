# credentiald

credentiald manages credentials for cloud environments.

## Development

To run against a local Minikube:

```
$ kubectl create namespace giantswarm
$ ./credentiald daemon \
    --service.kubernetes.incluster=false \
    --service.kubernetes.address=https://$(minikube ip):8443 \
    --service.kubernetes.tls.cafile=~/.minikube/ca.crt \
    --service.kubernetes.tls.crtfile=~/.minikube/client.crt \
    --service.kubernetes.tls.keyfile=~/.minikube/client.key
```

And to create a credential:
```
curl -Ss -i \
    http://localhost:8000/v4/organizations/acme/credentials/ \
    -d '{
            "provider": "aws",
            "aws": {
                "roles": {
                    "admin": "this-is-a-fake-admin-arn",
                    "awsoperator": "this-is-a-fake-awsoperator-arn"
                }
            }
        }'
```

Retrieve credentials for an org:

```
curl -s -i http://localhost:8000/v4/organizations/acme/credentials/
```
