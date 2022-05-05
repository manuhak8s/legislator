package main

import (
	_"fmt"

	_ "github.com/k0kubun/pp"

	_ "github.com/manuhak8s/legislator/cmd"
	_ "github.com/manuhak8s/legislator/pkg/config"
	_ "github.com/manuhak8s/legislator/pkg/k8s"
	"github.com/manuhak8s/legislator/pkg/luther"
	//v1 "k8s.io/api/networking/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	
	//luther.DestroyAllNetworkPolicies()
	
	//luther.ExecuteLegislation("test_data/configs/v3_data.yaml")
	//luther.ExecuteLegislation("2_test.constitution.yaml")

	//luther.ExecuteDestruction("test_data/configs/v3_data.yaml")
	luther.ExecuteDestruction("2_test.constitution.yaml")
	
	//cmd.Execute()
 }