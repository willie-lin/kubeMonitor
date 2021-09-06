package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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

func ListNode(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		var nodes []string
		//var node1 map[string]string

		nodeList, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		fmt.Println(nodeList)
		for _, node := range nodeList.Items {
			fmt.Printf("nodeName: v%, status: %v", node.GetName(), node.GetCreationTimestamp())
			//node1[node.ClusterName] = node.Name

			//nodes["name"] = node.GetName()
			//nodes["namespaces"] = node.Namespace
			//nodes["clusterName"] = node.GetClusterName()
			//nodes["ResourceVersion"] = node.GetResourceVersion()
			//nodes["APIVersion"] = node.APIVersion
			//nodes["Kind"] = node.Kind
			nodes = append(nodes, node.GetName())

		}

		return c.JSON(http.StatusOK, &nodes)
	}
}

func ListConfigMap(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		var configMaps []string

		configMapList, err := clientSet.CoreV1().ConfigMaps("").List(context.TODO(), metav1.ListOptions{})
		//configMapList, err := clientSet.CoreV1().ConfigMap().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		fmt.Println(configMapList)
		for _, configMap := range configMapList.Items {
			fmt.Printf("nodeName: v%, status: %v", configMap.GetName(), configMap.GetCreationTimestamp())

			//nodes["name"] = node.GetName()
			//nodes["namespaces"] = node.Namespace
			//nodes["clusterName"] = node.GetClusterName()
			//nodes["ResourceVersion"] = node.GetResourceVersion()
			//nodes["APIVersion"] = node.APIVersion
			//nodes["Kind"] = node.Kind
			configMaps = append(configMaps, configMap.GetName())

		}

		return c.JSON(http.StatusOK, &configMaps)
	}
}

func ListServices(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		var services []string

		servicesList, err := clientSet.CoreV1().Services("").List(context.TODO(), metav1.ListOptions{})
		//configMapList, err := clientSet.CoreV1().ConfigMap().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		fmt.Println(servicesList)
		for _, service := range servicesList.Items {
			fmt.Printf("nodeName: v%, status: %v", service.GetName(), service.GetCreationTimestamp())

			services = append(services, service.GetName())

		}

		return c.JSON(http.StatusOK, &services)
	}
}

func ListSecrets(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		var secretses []string

		secretsList, err := clientSet.CoreV1().Secrets("").List(context.TODO(), metav1.ListOptions{})
		//configMapList, err := clientSet.CoreV1().ConfigMap().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		fmt.Println(secretsList)
		for _, secrets := range secretsList.Items {
			fmt.Printf("nodeName: v%, status: %v", secrets.GetName(), secrets.GetCreationTimestamp())

			secretses = append(secretses, secrets.GetName())

		}

		return c.JSON(http.StatusOK, &secretses)
	}
}
