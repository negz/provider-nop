/*
Copyright 2019 The Crossplane Authors.

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

package controller

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/negz/provider-nop/apis/v1alpha1"
)

// SetupNopResource adds a controller that reconciles NopResources.
func SetupNopResource(mgr ctrl.Manager, l logging.Logger) error {
	name := managed.ControllerName(v1alpha1.NopResourceGroupKind)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		For(&v1alpha1.NopResource{}).
		Complete(managed.NewReconciler(mgr,
			resource.ManagedKind(v1alpha1.NopResourceGroupVersionKind),
			managed.WithExternalConnecter(&connecter{client: mgr.GetClient()}),
			managed.WithLogger(l.WithValues("controller", name)),
			managed.WithInitializers(managed.NewNameAsExternalName(mgr.GetClient())),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}

type connecter struct {
	client client.Client
}

func (c *connecter) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	fns := managed.ExternalClientFns{
		ObserveFn: func(_ context.Context, _ resource.Managed) (managed.ExternalObservation, error) {
			// If our managed resource has been deleted we need to report that
			// our pretend external resource is gone in order for the delete
			// process to complete. This means we'll never call the DeleteFn.
			if meta.WasDeleted(mg) {
				return managed.ExternalObservation{ResourceExists: false}, nil
			}

			// If our managed resource has not been deleted we report that our
			// pretend external resource exists and is up-to-date. This means
			// we'll never call the CreateFn or UpdateFn.
			mg.SetConditions(xpv1.Available())
			return managed.ExternalObservation{ResourceExists: true, ResourceUpToDate: true}, nil
		},
	}

	return fns, nil
}
