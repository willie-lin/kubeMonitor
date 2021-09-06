package handler

import (
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

func Hello(clientSet *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {

		return c.JSON(http.StatusOK, "hello, welcome to kube-Monitor!!!")
	}
}
