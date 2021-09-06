package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/willie-lin/kubeMonitor/pkg/handler"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	//log, _ := zap.NewDevelopment()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	clientSet, err := handler.GetClientSet()
	if err != nil {
		panic(err.Error())
	}

	deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-nginx",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	//createDeployment(err, deploymentsClient, deployment)

	e.GET("/", handler.Hello(clientSet))
	e.POST("/deployment", handler.CreateDeployment(deploymentsClient, deployment))
	e.GET("/pods", handler.ListPods(clientSet))
	e.GET("/namespaces", handler.GetAllNamespaces(clientSet))
	e.GET("/watchInterface", handler.WatchDeployment(clientSet))
	e.GET("/nodes", handler.ListNode(clientSet))
	e.GET("/configmap", handler.ListConfigMap(clientSet))
	e.GET("/services", handler.ListServices(clientSet))
	e.GET("/secrets", handler.ListSecrets(clientSet))

	//client, err := NewClient()
	//if err != nil {
	//	log.Fatal("opening ent client", zap.Error(err))
	//	return
	//}

	fmt.Println("eeee")

	//defer client.Close()
	//ctx := context.Background()

	//autoMigration := database.AutoMigration
	//autoMigration := AutoMigration
	//autoMigration(client, ctx)

	//debugMode := database.DebugMode
	//debugMode := DebugMode
	//
	//debugMode(err, client, ctx)

	//e.GET("/", Hello(client))

	e.Logger.Fatal(e.Start(":2021"))

}

func int32Ptr(i int32) *int32 {
	return &i

}
