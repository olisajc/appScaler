

# ========
path="./k8/crd.yaml"
{ kubectl apply -f $path || echo "failed to create resource" ; }


# ========
{ kubectl get pscaler || echo "failed to get resource" ; }