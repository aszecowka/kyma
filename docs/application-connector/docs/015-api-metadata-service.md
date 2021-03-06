---
title: Metadata Service
type: API
---

You can get the API specification of the Metadata Service for a given version of the service using this command:
```
curl https://gateway.{CLUSTER_DOMAIN}/{RE_NAME}/v1/metadata/api.yaml
```

To access the API specification of the Metadata Service locally, provide the NodePort of the `core-nginx-ingress-controller`.

To get the NodePort, run this command:

```
kubectl -n kyma-system get svc core-nginx-ingress-controller -o 'jsonpath={.spec.ports[?(@.port==443)].nodePort}'
```

To access the specification, run:

```
curl https://gateway.kyma.local:{NODE_PORT}/{RE_NAME}/v1/metadata/api.yaml
```
