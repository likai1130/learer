package goclient

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"testing"
)

func TestK8s(t *testing.T) {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// 创建kube config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()

	cronJob := &v1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hello",
			Namespace: "default",
		},
		Spec: v1.CronJobSpec{
			Schedule: "*/1 * * * *",
			JobTemplate: v1.JobTemplateSpec{
				Spec: v1.JobSpec{
					Template: corev1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{},
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								corev1.Container{
									Name:  "hello",
									Image: "busybox",
									Args:  []string{"/bin/sh", "-c", "date"},
								},
							},
							RestartPolicy: "OnFailure",
						},
					},
				},
			},
		},
	}

	cj, err := clientset.BatchV1().CronJobs("default").Create(ctx, cronJob, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("name ->", cj.GetName())
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
