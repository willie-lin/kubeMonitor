package main

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

// GetAllNamespaces func GetAllNamespaces(clientSet *kubernetes.Clientset) ([]string, error ){
func GetAllNamespaces(clientSet *kubernetes.Clientset) []string {
	var namespaces []string
	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	} else {
		for _, nsList := range namespaceList.Items {
			namespaces = append(namespaces, nsList.Name)
		}
	}
	//return namespaces, err
	return namespaces
}
