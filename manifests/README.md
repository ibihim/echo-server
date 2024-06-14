# OpenShift Configuration

The provided OpenShift configuration is based on [image builds](https://docs.openshift.com/container-platform/4.15/cicd/builds/understanding-image-builds.html) and uses an [external build trigger like GitHub](https://docs.openshift.com/container-platform/4.15/cicd/builds/triggering-builds-build-hooks.html) to automate the build and deployment process.

- namespace.yaml: Creates a new namespace named echo-server for the Golang Echo Server application.
- build-config.yaml: Defines the BuildConfig resource that specifies the build strategy, source repository, and triggers for the build process.
- secret.yaml: Defines a Secret resource to store sensitive information required for the build process, such as authentication tokens or credentials.
- role-binding.yaml: Creates a RoleBinding in a specific namespace to grant the necessary permissions for the build and deployment process.
- cluster-role-binding.yaml: Creates a ClusterRoleBinding to grant the necessary permissions for the build and deployment process. This is a last resort if the user wants to avoid creating RoleBindings for all individual namespaces.

Please note that using a ClusterRoleBinding is a broad permission and should be used cautiously. It is recommended to create RoleBindings for individual namespaces whenever possible to follow the principle of least privilege.