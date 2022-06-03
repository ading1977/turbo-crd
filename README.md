# Turbonomic Custom Resource
This repository hosts all Turbonomic custom resource API definitions and manifests.

## Description
The following Turbonomic CRDs are defined:

### Turbonomic Policy
* `policy.turbonomic.io/v1alpha1/slohorizontalscales`
* `policy.turbonomic.io/v1alpha1/policybindings`

## Getting Started
### Install CRDs
To install the CRDs from the cluster:

```sh
make install
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Modifying the API definitions
If you are editing the API definitions, generate the manifests such as CRs or CRDs using:

```sh
make manifests
```

**NOTE:** Run `make --help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

