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

	podNames := ListPods(clientSet)

	fmt.Println(len(podNames))
	fmt.Println(podNames)
	//listDeployment()

	podList()

	//watchDeployment(clientSet)
}

//	deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
//
//	deployment := &appsv1.Deployment{
//		ObjectMeta: metav1.ObjectMeta{
//			Name: "demo-nginx",
//		},
//		Spec: appsv1.DeploymentSpec{
//			Replicas: int32Ptr(2),
//			Selector: &metav1.LabelSelector{
//				MatchLabels: map[string]string{
//					"app": "demo",
//				},
//			},
//			Template: apiv1.PodTemplateSpec{
//				ObjectMeta: metav1.ObjectMeta{
//					Labels: map[string]string{
//						"app": "demo",
//					},
//				},
//				Spec: apiv1.PodSpec{
//					Containers: []apiv1.Container{
//						{
//							Name: "web",
//							Image: "nginx",
//							Ports: []apiv1.ContainerPort{
//								{
//									Name: "http",
//									Protocol: apiv1.ProtocolTCP,
//									ContainerPort: 80,
//								},
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	createDeployment(err, deploymentsClient, deployment)
//
//
//	// Update Deployment
//	prompt()
//	fmt.Println("Updating deployment...")
//	//    You have two options to Update() this Deployment:
//	//
//	//    1. Modify the "deployment" variable and call: Update(deployment).
//	//       This works like the "kubectl replace" command and it overwrites/loses changes
//	//       made by other clients between you Create() and Update() the object.
//	//    2. Modify the "result" returned by Get() and retry Update(result) until
//	//       you no longer get a conflict error. This way, you can preserve changes made
//	//       by other clients between Create() and Update(). This is implemented below
//	//			 using the retry utility package included with client-go. (RECOMMENDED)
//	//
//	// More Info:
//	// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency
//	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
//		// Retrieve the latest version of Deployment before attempting update
//		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
//		result, getErr := deploymentsClient.Get(context.TODO(), "demo-nginx", metav1.GetOptions{})
//		if getErr != nil {
//			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
//		}
//
//		result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
//		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
//		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
//		return updateErr
//	})
//	if retryErr != nil {
//		panic(fmt.Errorf("Update failed: %v", retryErr))
//	}
//	fmt.Println("Updated deployment...")
//
//	// List Deployments
//	prompt()
//	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
//	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
//	if err != nil {
//		panic(err)
//	}
//	for _, d := range list.Items {
//		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
//	}
//
//	// Delete Deployment
//	prompt()
//	fmt.Println("Deleting deployment...")
//	deletePolicy := metav1.DeletePropagationForeground
//	if err := deploymentsClient.Delete(context.TODO(), "demo-nginx", metav1.DeleteOptions{
//		PropagationPolicy: &deletePolicy,
//	}); err != nil {
//		panic(err)
//	}
//	fmt.Println("Deleted deployment.")
//}
//
//
//
//func prompt() {
//	fmt.Printf("-> Press Return key to continue.")
//	scanner := bufio.NewScanner(os.Stdin)
//	for scanner.Scan() {
//		break
//	}
//	if err := scanner.Err(); err != nil {
//		panic(err)
//	}
//	fmt.Println()
//}

//client, err := GetDiscoveryClient()
//fmt.Println(discoveryClient)
////discoveryClient, err := discovery.NewDiscoveryClientForConfig(getCondig())
//if err != nil {
//	panic(err.Error())
//}

//grv := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
//
//unstructObj, err := dynamicClient.Resource(grv).Namespace("kube-system").List(context.TODO(), metav1.ListOptions{Limit: 100})
//if err != nil {
//	panic(err.Error())
//}

//podList := &apiv1.NewFirestoreAdminClient()

//APIResourceListSlice, err := discoveryClient.ServerPreferredResources()

//fmt.Printf("APIGROUP: \n\n %v\n\n\n\n\n", APIGroup)

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
//
//
//func int32Ptr(i int32) *int32 {
//	return &i
//}
