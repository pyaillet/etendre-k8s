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

Find the line:
```go
type AppGiphySpec struct {
  // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
```

Insert this line after the previous ones to specify the tag used for this
particular AppGiphy:
```go
  Tag string
```

As stated in the code comment, launch the command to generate the code
relative to this type in the `giphy-operator` dir:
```shell
operator-sdk generate k8s
```

## Modify the controller to adapt Pod creation

TODO


