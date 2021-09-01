package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listDeployment() {

	clientSet, err := GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	deployments, err := clientSet.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for i, deployment := range deployments.Items {
		fmt.Printf("[%d] %s\n", i, deployment.GetName())
		fmt.Printf("[%d] %s\n", i, deployment.GetCreationTimestamp())
		fmt.Printf("[%d] %s\n", i, deployment.GetNamespace())
		fmt.Printf("[%d] %s\n", i, deployment.GetLabels())
		fmt.Printf("[%d] %s\n", i, deployment.GetUID())
		fmt.Printf("[%d] %s\n", i, deployment.GetSelfLink())
		fmt.Println("-----------------")

	}

}
