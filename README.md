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
### config anatomy
```yaml
connectedSets:
  - name: <set-name>
    targetNamespaces:
      matchLabels:
        <key>: <value>
    podSelector:
      matchLabels:
        <key>: <value>
```
connectedSets represents a list of connected amounts that can communicate to each others. Every set can be desribed as a single list object encapsulated from the others. Every set has to contain a name for identification. At least one key-value pair has to be configured in a config to adress a namespace fornetwork policy deployment. The same applies to the podSelector-field.
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
In this example two sets are defined as **first-layer-set** and **second-layer-set**. 
* The first-layer-set adresses all namespaces with the labeling **project: dev** and creates network policies with associated ingress rules for all pods that are containing the labeling **set: first**.
* The second-layer-set adresses all namespaces with the labeling **project: prod** and creates network policies with associated ingress rules for all pods that are containing the labeling **set: second**.

**-->** That means that two isolated quantities of pods were created after an executed legislator deployment

## commands and flags 
Following commands and flags are executable by using the current release of legislator CLI:
command/flag | example | description
apply  | legislator apply --path=<config path> | Creates network policies based on the given config path to current kubecontext
destroy  | legislator destroy --path=<config path> | Removes network policies based on the given config path from current kubecontext
--path  | legislator destroy --path=/path/to/config.yaml | Flag that accepts a valid path to an existant config file.

## notes - FAQs
* every network policy is associated to its config file - that means by executing the destroy command, every network policy based on the config will be deleted from the current kubecontext
* legislator CLI has no update functionality yet so that it is recommended to execute destroy and apply subsequently the updated config file
* this tool is a result of a bachelor thesis project and possibly will not be followed up 

# installation
## legislator binary download link
[legislator](https://github.com/manuhak8s/legislator/blob/validate-config/legislator)
## git clone and go build
```bash
git clone https://github.com/manuhak8s/legislator.git
cd legislator
go build .
```