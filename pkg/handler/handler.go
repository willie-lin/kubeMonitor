package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"log"
	"net/http"
)

func CreateDeployment(deploymentsClient v1.DeploymentInterface, deployment *appsv1.Deployment) echo.HandlerFunc {
	return func(c echo.Context) error {

		fmt.Println("Creating deployment......")
		result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}

		fmt.Println("Created deployment %q.\n", result.GetObjectMeta().GetName())

		return c.JSON(http.StatusOK, result.GetObjectMeta().GetName())
	}

}

func ListPods(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		return c.JSON(http.StatusOK, &podNames)
	}

}

func GetAllNamespaces(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
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
		return c.JSON(http.StatusOK, &namespaces)
	}
}

func WatchDeployment(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		//var deployment *v1beta1.Deployment
		//
		//clientset, err := GetClientSet()
		//if err != nil {
		//	panic(err.Error())
		//}

		watchInterface, err := clientSet.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(watchInterface)
		//for _, err := range watchInterface.

		return c.JSON(http.StatusOK, &watchInterface)
	}

}

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
