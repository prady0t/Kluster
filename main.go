package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"time"
)

func main() {
	//kubeconfig := flag.String("kubeconfig", "/Users/pradyotranjan/.kube/config", "Location of kubeconfig")
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	fmt.Println("Error fetching kubeconfig")
	//}
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Printf("Error while InClusterConfig(): %s", err.Error())
	}
	//runtime.Object()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %s", err.Error())
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing Pods: %s", err.Error())
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)

	}
	depls, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{
		TypeMeta:             metav1.TypeMeta{},
		LabelSelector:        "",
		FieldSelector:        "",
		Watch:                false,
		AllowWatchBookmarks:  false,
		ResourceVersion:      "",
		ResourceVersionMatch: "",
		TimeoutSeconds:       nil,
		Limit:                0,
		Continue:             "",
		SendInitialEvents:    nil,
	})
	informerfactory := informers.NewSharedInformerFactory(clientset, 10*time.Minute)

	podInformer := informerfactory.Core().V1().Pods().Informer()
	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(new interface{}) {
			fmt.Println("New was called!")
		},
		UpdateFunc: func(oldObj, newObj interface{}) {

		}})
	//Deployments are part of appsv1
	if err != nil {
		fmt.Printf("Error listing Deployments: %s", err.Error())
	}
	for _, d := range depls.Items {
		fmt.Printf("Deployment - %s\n", d.GetName())
	}
	//Pods are from api core.ve so after clientset we use core.v1 you will have to know about
	//which resource belongs to which api version herw.
}
