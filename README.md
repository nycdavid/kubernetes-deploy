### Deploy the Kubernetes cluster with `kops`
1. Set the following environment variables
    * `KOPS_STATE_STORE=s3://foo-bar`
    * `NAME=myfirstcluster.k8s.local`
1. Run cluster creation command
    ```
    kops create cluster \
      --zones us-west-2a \
      --state s3://foo-bar \
      myfirstcluster.k8s.local
    ```

### Log into the cluster
1. Record the secret for login
    ```
    kops get secrets \
    --type secret admin \
    --state s3://foo-bar \
    -oplaintext
    ```
1. Proxy localhost to the remote cluster: `kubectl proxy`
1. Point browser to `localhost:8001`
