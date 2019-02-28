#!/usr/bin/env sh

set -e

# Create a new giphy-operator project
operator-sdk new giphy-operator
cd giphy-operator

# Add a new API for the custom resource AppGiphy
operator-sdk add api --api-version=app.zenika.com/v1alpha1 --kind=AppGiphy

# Add a new controller that watches for AppGiphy
operator-sdk add controller --api-version=app.zenika.com/v1alpha1 --kind=AppGiphy

# Build and push the app-operator image to a public registry
operator-sdk build pyaillet/giphy-operator
