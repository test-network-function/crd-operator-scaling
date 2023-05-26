package controllers

import (
	"context"
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
	tutorialv1 "my.domain/tutorial/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// FooReconciler reconciles a Foo object
type FooReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// RBAC permissions to monitor foo custom resources
//+kubebuilder:rbac:groups=tutorial.my.domain,resources=foos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tutorial.my.domain,resources=foos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tutorial.my.domain,resources=foos/finalizers,verbs=update

// RBAC permissions to monitor pods
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *FooReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("reconciling foo custom resource")
	// Get the Foo resource that triggered the reconciliation request
	var foo tutorialv1.Foo
	if err := r.Get(ctx, req.NamespacedName, &foo); err != nil {
		log.Error(err, "unable to fetch Foo")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("create deployment")
	size := foo.Spec.Replicas
	faleb := false
	tim := int64(30)
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "jack",
			Namespace: "tnf",
			Labels: map[string]string{
				"app": "jack",
			},
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion:         "tutorial.my.domain/v1",
				Kind:               "Foo",
				Name:               "foo-sample",
				UID:                foo.GetUID(),
				BlockOwnerDeletion: pointer.Bool(true),
				Controller:         pointer.Bool(true),
			},
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "jack",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":                                 "jack",
						"test-network-function.com/generic":   "target",
						"test-network-function.com/container": "target",
					},
				},
				Spec: corev1.PodSpec{
					TerminationGracePeriodSeconds: &tim,
					AutomountServiceAccountToken:  &faleb,
					Containers: []corev1.Container{{
						Command:                  []string{"./bin/app"},
						TerminationMessagePolicy: "FallbackToLogsOnError",
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								corev1.ResourceMemory: resource.MustParse("20Mi"),
								corev1.ResourceCPU:    resource.MustParse("10m"),
							},
							Requests: corev1.ResourceList{
								corev1.ResourceCPU: resource.MustParse("10m"),
							},
						},
						Image:           "quay.io/testnetworkfunction/cnf-test-partner:latest",
						ImagePullPolicy: corev1.PullIfNotPresent,
						Name:            "jack",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8080,
							Name:          "http-jack",
						}},
						Lifecycle: &corev1.Lifecycle{
							PostStart: &corev1.LifecycleHandler{
								Exec: &corev1.ExecAction{
									Command: []string{"/bin/sh", "-c", "echo Hello from the postStart handler > /tmp/message"},
								},
							},
							PreStop: &corev1.LifecycleHandler{
								Exec: &corev1.ExecAction{
									Command: []string{"/bin/sh", "-c", "killall -0 tail"},
								},
							},
						},
						LivenessProbe: &corev1.Probe{
							ProbeHandler: corev1.ProbeHandler{
								HTTPGet: &corev1.HTTPGetAction{
									Port: intstr.IntOrString{
										IntVal: 8080,
									},
									Path: "/health",
									HTTPHeaders: []corev1.HTTPHeader{{
										Name:  "health-check",
										Value: "liveness",
									},
									},
								},
							},
						},
						ReadinessProbe: &corev1.Probe{
							ProbeHandler: corev1.ProbeHandler{
								HTTPGet: &corev1.HTTPGetAction{
									Port: intstr.IntOrString{
										IntVal: 8080,
									},
									Path: "/ready",
									HTTPHeaders: []corev1.HTTPHeader{{
										Name:  "health-check",
										Value: "readiness",
									},
									},
								},
							},
						},
						StartupProbe: &corev1.Probe{
							ProbeHandler: corev1.ProbeHandler{
								HTTPGet: &corev1.HTTPGetAction{
									Port: intstr.IntOrString{
										IntVal: 8080,
									},
									Path: "/ready",
									HTTPHeaders: []corev1.HTTPHeader{{
										Name:  "health-check",
										Value: "startup",
									},
									},
								},
							},
						},
					}},
					Affinity: &corev1.Affinity{
						PodAntiAffinity: &corev1.PodAntiAffinity{},
					},
				},
			},
		},
	}
	log.Info("after create deployment")
	found := &appsv1.Deployment{}
	errf := r.Get(context.TODO(), types.NamespacedName{
		Name:      "jack",
		Namespace: "tnf",
	}, found)
	if errf != nil {
		log.Error(errf, "unable to list pods")

	}
	if errf != nil && errors.IsNotFound(errf) {
		errdep := r.Create(context.TODO(), dep)
		found = dep
		if errdep != nil {
			log.Error(errdep, "unable to list pods")
		}
	}
	selector, err := metav1.LabelSelectorAsSelector(found.Spec.Selector)
	if err != nil {
		log.Error(err, "Error retrieving Deployment labels")
		return reconcile.Result{}, err
	}
	// 3. Retrieve the current number of replicas from the
	replicas := *found.Spec.Replicas
	// 4. Retrieve and update the CR
	foo.Status.Selector = selector.String()
	foo.Status.Replicas = replicas
	log.Info(fmt.Sprintf("%d", int(replicas)))
	// Get pods with the same name as Foo's friend
	if err := r.Status().Update(ctx, &foo); err != nil {
		return ctrl.Result{}, err
	}
	time.Sleep(3 * time.Second)
	var newfoo tutorialv1.Foo
	if err := r.Get(ctx, req.NamespacedName, &newfoo); err != nil {
		log.Error(err, "unable to fetch Foo")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	log.Info(fmt.Sprintf("newfoo replicas is %d", int(newfoo.Spec.Replicas)))

	if replicas != newfoo.Spec.Replicas {
		log.Info(fmt.Sprintf("foos replicas is %d and dep replica is %d ", int(newfoo.Spec.Replicas), int(replicas)))
		found.Spec.Replicas = &newfoo.Spec.Replicas
		err = r.Update(context.TODO(), found)
		if err != nil {
			log.Error(err, "Error retrieving Deployment labels")
		} else {
			newfoo.Status.Replicas = *found.Spec.Replicas
			if err := r.Status().Update(ctx, &newfoo); err != nil {
				return ctrl.Result{}, err
			}
			time.Sleep(3 * time.Second)
		}
	}
	log.Info("foo custom resource reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FooReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&tutorialv1.Foo{}).
		Complete(r)
}

func (r *FooReconciler) mapPodsReqToFooReq(obj client.Object) []reconcile.Request {
	ctx := context.Background()
	log := log.FromContext(ctx)

	// List all the Foo custom resource
	req := []reconcile.Request{}
	var list tutorialv1.FooList
	if err := r.Client.List(context.TODO(), &list); err != nil {
		log.Error(err, "unable to list foo custom resources")
	} else {
		// Only keep Foo custom resources related to the Pod that triggered the reconciliation request
		for _, item := range list.Items {
			if item.Spec.Name == obj.GetName() {
				req = append(req, reconcile.Request{
					NamespacedName: types.NamespacedName{Name: item.Name, Namespace: item.Namespace},
				})
				log.Info("pod linked to a foo custom resource issued an event", "name", obj.GetName())
			}
		}
	}
	return req
}
