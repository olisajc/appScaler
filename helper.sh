# ===========
path="./k8/crd.yaml"
{ kubectl apply -f $path || echo "failed to create resource" ; }


# =================================
{ kubectl get pscaler || echo "failed to get resource" ; }

# =================================
kubectl get role -n workspace


# =================================
kubectl get rolebinding -n workspace

# =================================
kubectl get pscaler -n workspace


# =================================
kubectl apply -f k8/policyScaler.yaml -n workspace


# =================================
kubectl create clusterrolebinding policyscaler-controller --clusterrole=cluster-admin  --serviceaccount=workspace:policy-scaler


# ==================================
kubectl run test-pod --image=debian:latest --command -- sleep infinity