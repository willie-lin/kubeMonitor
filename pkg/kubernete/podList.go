package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func podList() {

	restClient, err := GetRestClient()
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}
	if err := restClient.
		Get().
		Namespace("kube-system").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result); err != nil {
		panic(err)
	}

	fmt.Println("Print all listed pods.")

	// 打印所有获取到的pods资源，输出到标准输出
	for _, d := range result.Items {
		fmt.Printf("NAMESPACE: %v NAME: %v \t STATUS: %v \n", d.Namespace, d.Name, d.Status.Phase)
	}
}
