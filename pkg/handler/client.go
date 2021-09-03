package handler

import (
	"flag"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetClientSet() (*kubernetes.Clientset, error) {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kubeconfig", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	fmt.Println(kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet, err
}

func GetRestClient() (*rest.RESTClient, error) {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kubeconfig", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	fmt.Println(kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	fmt.Println("Init RESTClient.")

	// 定义RestClient，用于与k8s API server进行交互
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	return restClient, err
}

//func GetDiscoveryClient() (*kubernetes.Clientset, error) {
//
//	var kubeconfig *string
//	if home := homedir.HomeDir(); home != "" {
//		kubeconfig = flag.String("kubeconfig", filepath.Join(home, "kubeconfig", "config"), "(optional) absolute path to the kubeconfig file")
//	} else {
//		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
//	}
//	flag.Parse()
//
//	fmt.Println(kubeconfig)
//
//	// use the current context in kubeconfig
//	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// create the client
//	//discoveryClient, err := clientset.DiscoveryClient.ServerGroupsAndResources()
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	return clientset, nil
//}
