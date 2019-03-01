# How to use

## Install the operator framework sdk

For the beginning just follow [these
steps](https://github.com/operator-framework/operator-sdk#quick-start) of the
official doc.

Or execute `01_setup.sh`

## Initialize the a new operator

```shell
#!/usr/bin/env sh

# Create a new giphy-operator project
operator-sdk new giphy-operator
cd giphy-operator

# Add a new API for the custom resource AppGiphy
operator-sdk add api --api-version=app.zenika.com/v1alpha1 --kind=AppGiphy

# Add a new controller that watches for AppGiphy
operator-sdk add controller --api-version=app.zenika.com/v1alpha1 --kind=AppGiphy

# Build and push the app-operator image to a public registry
operator-sdk build pyaillet/giphy-operator
```

Or execute `02_init_operator.sh`

After this step you should have a functional build

## Modify the CRD to add your properties

Edit the file `pkg/apis/app/v1alpha1/appgiphy_types.go`
Find the lines:
```go
type AppGiphySpec struct {
  // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
```

Insert this line after the previous ones to specify the tag used for this
particular AppGiphy:
```go
  Tag string `json:"tag"`
```

As stated in the code comment, launch the command to generate the code
relative to this type in the `giphy-operator` dir:
```shell
operator-sdk generate k8s
```

## Modify the controller to adapt Pod creation

Edit the file `pkg/controller/appgiphy/appgiphy_controller.go`
Find the lines:
```go
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
```

- Replace the `busybox` name with `giphyserver`
- Replace the `busybox` image with `pyaillet/giphyserver:0.1`
- Remove the `Command` section

The content should then, be like:
```go
			Containers: []corev1.Container{
				{
					Name:    "giphyserver",
					Image:   "pyaillet/giphyserver:0.1",
					Env: []corev1.EnvVar{
					  {
						  Name: "TAG",
						  Value: cr.Spec.Tag,
            },
					},
				},
			},
```

## Rebuild and push the operator

```shell
operator-sdk build pyaillet/giphy-operator:0.1
docker push pyaillet/giphy-operator:0.1
```

## Deployment

- Replace the image name in the deployment descriptor `deploy/operator.yaml`
  with your image:

```yaml
containers:
  - name: giphy-operator
    # Replace this with the built image name
    image: pyaillet/giphy-operator:0.1
```

- Deploy the _Operator_ and other needed resources

```shell
kubectl apply -f deploy/crds/app_v1alpha1_appgiphy_crd.yaml
kubectl apply -f deploy/
```

- Verify that the _CRD_ has been created

```shell
$ kubectl get crd
NAME                        CREATED AT
appgiphies.app.zenika.com   2019-02-27T22:23:02Z
```

- Create a new Custom Resource

```shell
kubectl apply -f - <<EOF
apiVersion: app.zenika.com/v1alpha1
kind: AppGiphy
metadata:
  name: example-appgiphy
spec:
  tag: dog
```

- Verify that a new _Pod_ corresponding to your CR is being created:

```shell
$ kubectl get po -l app=example-appgiphy
NAME                   READY     STATUS    RESTARTS   AGE
example-appgiphy-pod   1/1       Running   0          100s
```
