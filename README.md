# legislator CLI

**legislator** is an CLI for managing kubernetes network policies from a single config file. Like a member
of a legislative body, the legislator CLI can create rules or laws for network traffic in a kubernetes cluster.
The added value of this tool is a less complex level of defining multiple network policies for kubernetes and
an aspect of less manual configuration. 

Various use cases could be made more pleasant:
* general network segmentation in kubernetes
* pod or namespace access restrictons 
* user access restrictions
* creating development environments
* network security apsetcs for layered architectures
* public/private networking

## legislator config
The base of the legislator CLI is an individual config with its own fields and syntax. 
### config fields
A legislator config file has to contain some mandatory fields for a successful creation and deployment of network policies.
The following table contains all actual fields of the current release with an additional description.

Field  | Descritopn
------------- | -------------
connectedSets  | can be described as the opening sequence and list representation of connected amounts. 
name  | represents the identification of a connected set and has to be unique within a config file.
targetNamespaces  | contains any information from namespaces that are to be adressed 
podSelector  | contains any information from pods that are to be adressed 
matchLabels  | key-value representations of the labeling fields from namespace or pod instances
### example configuration
```yaml
connectedSets:
  - name: first-layer-set
    targetNamespaces:
      matchLabels:
        project: dev
    podSelector:
      matchLabels:
        set: first
  - name: second-layer-set
    targetNamespaces:
      matchLabels:
        project: prod
    podSelector:
      matchLabels:
        set: second
```
connectedSets represents a list of connected amounts that can communicate to each others.
In this example two sets are defined as **first-layer-set** and **second-layer-set**. 
* The first-layer-set adresses all namespaces with the labeling **project: dev** and creates network policies with associated ingress rules for all pods that are containing the labeling **set: first**.
* The second-layer-set adresses all namespaces with the labeling **project: prod** and creates network policies with associated ingress rules for all pods that are containing the labeling **set: second**.
**-->** That means that two isolated quantities of pods were created after an executed legislator deployment

## legislator binary download
[legislator](https://github.com/manuhak8s/legislator/blob/validate-config/legislator)