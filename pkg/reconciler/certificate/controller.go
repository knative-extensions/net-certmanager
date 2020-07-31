/*
Copyright 2020 The Knative Authors

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

package certificate

import (
	"context"

	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"

	cmclient "knative.dev/net-certmanager/pkg/client/certmanager/injection/client"
	cmchallengeinformer "knative.dev/net-certmanager/pkg/client/certmanager/injection/informers/acme/v1alpha2/challenge"
	cmcertinformer "knative.dev/net-certmanager/pkg/client/certmanager/injection/informers/certmanager/v1alpha2/certificate"
	clusterinformer "knative.dev/net-certmanager/pkg/client/certmanager/injection/informers/certmanager/v1alpha2/clusterissuer"
	"knative.dev/net-certmanager/pkg/reconciler/certificate/config"
	network "knative.dev/networking/pkg"
	"knative.dev/networking/pkg/apis/networking"
	kcertinformer "knative.dev/networking/pkg/client/injection/informers/networking/v1alpha1/certificate"
	certreconciler "knative.dev/networking/pkg/client/injection/reconciler/networking/v1alpha1/certificate"
	serviceinformer "knative.dev/pkg/client/injection/kube/informers/core/v1/service"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/logging/logkey"
	pkgreconciler "knative.dev/pkg/reconciler"
	"knative.dev/pkg/tracker"
)

const controllerAgentName = "certificate-controller"

// AnnotateLoggerWithName names the logger in the context with the supplied name
//
// This is a stop gap until the generated reconcilers can do this
// automatically for you
func AnnotateLoggerWithName(ctx context.Context, name string) context.Context {
	logger := logging.FromContext(ctx).
		Named(name).
		With(zap.String(logkey.ControllerType, name))

	return logging.WithLogger(ctx, logger)

}

// NewController initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	ctx = AnnotateLoggerWithName(ctx, controllerAgentName)
	logger := logging.FromContext(ctx)
	knCertificateInformer := kcertinformer.Get(ctx)
	cmCertificateInformer := cmcertinformer.Get(ctx)
	cmChallengeInformer := cmchallengeinformer.Get(ctx)
	clusterIssuerInformer := clusterinformer.Get(ctx)
	svcInformer := serviceinformer.Get(ctx)

	c := &Reconciler{
		cmCertificateLister: cmCertificateInformer.Lister(),
		cmChallengeLister:   cmChallengeInformer.Lister(),
		cmIssuerLister:      clusterIssuerInformer.Lister(),
		svcLister:           svcInformer.Lister(),
		certManagerClient:   cmclient.Get(ctx),
	}

	impl := certreconciler.NewImpl(ctx, c, network.CertManagerCertificateClassName,
		func(impl *controller.Impl) controller.Options {
			logger.Info("Setting up ConfigMap receivers")
			resyncCertOnCertManagerconfigChange := configmap.TypeFilter(&config.CertManagerConfig{})(func(string, interface{}) {
				impl.GlobalResync(knCertificateInformer.Informer())
			})
			configStore := config.NewStore(logger.Named("config-store"), resyncCertOnCertManagerconfigChange)
			configStore.WatchConfigs(cmw)
			return controller.Options{ConfigStore: configStore}
		})

	logger.Info("Setting up event handlers")
	classFilterFunc := pkgreconciler.AnnotationFilterFunc(networking.CertificateClassAnnotationKey, network.CertManagerCertificateClassName, true)
	certHandler := cache.FilteringResourceEventHandler{
		FilterFunc: classFilterFunc,
		Handler:    controller.HandleAll(impl.Enqueue),
	}
	knCertificateInformer.Informer().AddEventHandler(certHandler)

	cmCertificateInformer.Informer().AddEventHandler(controller.HandleAll(impl.EnqueueControllerOf))

	c.tracker = tracker.New(impl.EnqueueKey, controller.GetTrackerLease(ctx))

	// Make sure trackers are deleted once the observers are removed.
	knCertificateInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.tracker.OnDeletedObserver,
	})

	svcInformer.Informer().AddEventHandler(controller.HandleAll(
		controller.EnsureTypeMeta(
			c.tracker.OnChanged,
			corev1.SchemeGroupVersion.WithKind("Service"),
		),
	))

	return impl
}
