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

# Build and push the app-operator image to a public registry such as quay.io
operator-sdk build pyaillet/giphy-operator
```

Or execute `02_init_operator.sh`

After this step you should have a functional build

## Modify the CRD to add your properties

TODO

## Modify the controller to adapt Pod creation

TODO


