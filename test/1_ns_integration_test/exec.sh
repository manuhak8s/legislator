#!/bin/sh

# prequireties to execute this test
# - connection to minikube cluster + cni pluging enabled & installed
# -- e.g. minikube start --network-plugin=cni --cni=false
# - kubectl & kubens 
# - helm

# K8S connection test
echo " "
echo "$(tput setaf 4) K8S CONNECTION CHECK $(tput setaf 7)"
echo "$(tput setaf 3) ... checking minikube connection ... $(tput setaf 7)"
k8s_ready=false
if kubectl -n kube-system wait pod --for=condition=Ready -l component=etcd,tier=control-plane
then 
    k8s_ready=true
fi

if $k8s_ready == true
then
    # creating test namespace
    echo " "
    echo "$(tput setaf 4) INITIALIZING K8S RESOURCES  $(tput setaf 7)"
    echo "$(tput setaf 3) ... creating namespace(s) ... $(tput setaf 7)"
    kubectl apply -f ./namespaces/legislator_test_namespace.yaml
    kubens legislator-test-namespace

    # create helm release
    echo " "
    echo "$(tput setaf 3) ... installing helm chart ... $(tput setaf 7)"
    helm install legislator-test-release curl_helm_chart/ --atomic

    # expose services
    echo " "
    echo "$(tput setaf 3) ... exposing services ... $(tput setaf 7)"
    declare -a pod_names=(
        "pod-1"
        "pod-2"
        "pod-3"
        "pod-4"
    )

    for name in "${pod_names[@]}"
    do
        kubectl expose pod $name --port=80
    done
    sleep 3

    # check connections before legislator config deployment
    echo " "
    echo "$(tput setaf 4) POD CONNECTION CHECKS $(tput setaf 7)"
    check1=()
    echo "$(tput setaf 3) ... checking connections WITHOUT network policies ... $(tput setaf 7)"
    for pod in "${pod_names[@]}"
    do
        src=$pod
        for pod in "${pod_names[@]}"
        do
            trgt=$pod
            echo "$(tput setaf 3) $src curling $trgt ... $(tput setaf 7)"
            kubectl exec -it $src -- /bin/bash -c "curl $trgt --max-time 1"
            if [ $? -eq 0 ]; then
                echo "$(tput setaf 2) OK: $src curling $trgt $(tput setaf 7)"
            else
                check1+=("$src:$trgt" )
            fi
        done 
    done 

    # deploy legislator config
    echo " "
    echo "$(tput setaf 5) ... deploying network policies ... $(tput setaf 7)"
    ./legislator apply --path=./constitution.yaml

    # check connections after legislator config deployment
    echo " "
    check2=()
    echo "$(tput setaf 3) ... checking connections WITH network policies ... $(tput setaf 7)"
    for pod in "${pod_names[@]}"
    do
        src=$pod
        for pod in "${pod_names[@]}"
        do
            trgt=$pod
            echo "$(tput setaf 3) $src curling $trgt ... $(tput setaf 7)"
            kubectl exec -it $src -- /bin/bash -c "curl $trgt --max-time 1"
            if [ $? -eq 0 ]; then
                echo "$(tput setaf 2) OK: $src curling $trgt $(tput setaf 7)"
            else
                check2+=("$src:$trgt" )
            fi
        done 
    done 

    # test summary
    echo " "
    echo "$(tput setaf 4) SUMMARY $(tput setaf 7)"
    # summary test 1 - before legislator deployment
    echo "$(tput setaf 3) ... connection check BEFORE legislator deployment ... $(tput setaf 7)"
    if [ -z "$check1" ]; then
        echo "$(tput setaf 2) connection check succeeded $(tput setaf 7)"
    else
        for opt in "${check1[@]}"
        do
            src=${opt%%:*}
            trgt=${opt#*:}
            echo "$(tput setaf 1) connection check failed: $src + $trgt $(tput setaf 7)"
        done
    fi

    # summary test 2 - after legislator deployment
    echo " "
    echo "$(tput setaf 3) ... connection check AFTER legislator deployment ... $(tput setaf 7)"
    if [ -z "$check2" ]; then
        echo "$(tput setaf 2) connection check succeeded $(tput setaf 7)"
    else
        for opt in "${check2[@]}"
        do
            src=${opt%%:*}
            trgt=${opt#*:}
            echo "$(tput setaf 1) connection check failed: $src ----> $trgt $(tput setaf 7)"
        done
    fi

    # removing all created instances from minikube cluster
    echo " "
    echo "$(tput setaf 4)  RESOURCE REMOVEMENT $(tput setaf 7)"
    echo "$(tput setaf 3) ... removing legislator netowrk policies ... $(tput setaf 7)"
    ./legislator destroy --path=./constitution.yaml
    echo " "
    echo "$(tput setaf 3) ... uninstalling helm release ... $(tput setaf 7)"
    helm uninstall legislator-test-release
    echo " "
    echo "$(tput setaf 3) ... deleting namespace(s) ... $(tput setaf 7)"
    kubectl delete namespace legislator-test-namespace
    echo " "
    echo "$(tput setaf 2) integration test finished :-) $(tput setaf 7)"
else
    echo "$(tput setaf 1) can not connect to minikube - please try again $(tput setaf 7)"
fi