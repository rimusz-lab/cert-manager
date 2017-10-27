/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1beta1

import (
	internalinterfaces "github.com/jetstack-experimental/cert-manager/third_party/informers/internalinterfaces"
	extensions_v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/extensions/v1beta1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// PodSecurityPolicyInformer provides access to a shared informer and lister for
// PodSecurityPolicies.
type PodSecurityPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PodSecurityPolicyLister
}

type podSecurityPolicyInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewPodSecurityPolicyInformer constructs a new informer for PodSecurityPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ExtensionsV1beta1().PodSecurityPolicies().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ExtensionsV1beta1().PodSecurityPolicies().Watch(options)
			},
		},
		&extensions_v1beta1.PodSecurityPolicy{},
		resyncPeriod,
		indexers,
	)
}

func defaultPodSecurityPolicyInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewPodSecurityPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *podSecurityPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&extensions_v1beta1.PodSecurityPolicy{}, defaultPodSecurityPolicyInformer)
}

func (f *podSecurityPolicyInformer) Lister() v1beta1.PodSecurityPolicyLister {
	return v1beta1.NewPodSecurityPolicyLister(f.Informer().GetIndexer())
}
