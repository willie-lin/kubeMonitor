/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {

	clientSet, err := GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	//deploymentsClient, _ := clientSet.AppsV1().Deployments().List(context.TODO(), metav1.ListOptions{})

	fmt.Println(clientSet)

	//var err error
	//var nss []string
	//nss, err := GetAllNamespaces(clientSet)
	nss := GetAllNamespaces(clientSet)

	fmt.Println(len(nss))

	fmt.Println(nss[1:3])

	discoveryClient, err := GetDiscoveryClient()
	//discoveryClient, err := discovery.NewDiscoveryClientForConfig(getCondig())
	if err != nil {
		panic(err.Error())
	}

	//grv := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	//
	//unstructObj, err := dynamicClient.Resource(grv).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
	//if err != nil {
	//	panic(err.Error())
	//}

	//podList := &apiv1.NewFirestoreAdminClient()

	APIResourceListSlice, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		panic(err.Error())
	}

	//fmt.Printf("APIGROUP: \n\n %v\n\n\n\n\n", APIGroup)

	for _, singleAPIResourceList := range APIResourceListSlice {

		groupVersionStr := singleAPIResourceList.GroupVersion
		gv, err := schema.ParseGroupVersion(groupVersionStr)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println("-------------------------------")

		fmt.Printf("GV string [%v]\nGV", groupVersionStr, gv)

		//for _, singleAPIResource := range APIRes

	}

	//listDeployment()

	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kubeconfig", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//
	//fmt.Println(kubeconfig)
	//
	//// use the current context in kubeconfig
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//// create the clientset
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err.Error())
	//}
	//for {
	//	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//	// Examples for error handling:
	//	// - Use helper functions like e.g. errors.IsNotFound()
	//	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//	namespace := "default"
	//	pod := "example-xxxxx"
	//	_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	//	if errors.IsNotFound(err) {
	//		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	//	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//			pod, namespace, statusError.ErrStatus.Message)
	//	} else if err != nil {
	//		panic(err.Error())
	//	} else {
	//		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	//	}
	//
	//	time.Sleep(10 * time.Second)
	//}

}
