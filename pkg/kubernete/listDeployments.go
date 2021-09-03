package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func createDeployment(err error, deploymentsClient v1.DeploymentInterface, deployment *appsv1.Deployment) {
	fmt.Println("Creating deployment......")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

func listDeployment() {

	clientSet, err := GetClientSet()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(clientSet)

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

func watchDeployment(clientSet *kubernetes.Clientset) {
	//var deployment *v1beta1.Deployment

	clientset, err := GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	watchInterface, err := clientset.AppsV1().Deployments("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(watchInterface)
}

func ListPods(clientSet *kubernetes.Clientset) []string {
	//clientSet := GetClientSet()
	var podNames []string
	//var label  map[string]string
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	} else {
		for _, podList := range pods.Items {
			//fmt.Println(podList)
			podNames = append(podNames, podList.GetName())
			//label = append(podNames, podList.Labels)
		}
	}

	//fmt.Println(po)
	//fmt.Println(pods)
	return podNames
}
