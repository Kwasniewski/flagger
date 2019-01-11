/*
Copyright The Flagger Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	flaggerv1alpha2 "github.com/stefanprodan/flagger/pkg/apis/flagger/v1alpha2"
	versioned "github.com/stefanprodan/flagger/pkg/client/clientset/versioned"
	internalinterfaces "github.com/stefanprodan/flagger/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/stefanprodan/flagger/pkg/client/listers/flagger/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CanaryInformer provides access to a shared informer and lister for
// Canaries.
type CanaryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.CanaryLister
}

type canaryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCanaryInformer constructs a new informer for Canary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCanaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCanaryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCanaryInformer constructs a new informer for Canary type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCanaryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlaggerV1alpha2().Canaries(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlaggerV1alpha2().Canaries(namespace).Watch(options)
			},
		},
		&flaggerv1alpha2.Canary{},
		resyncPeriod,
		indexers,
	)
}

func (f *canaryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCanaryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *canaryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flaggerv1alpha2.Canary{}, f.defaultInformer)
}

func (f *canaryInformer) Lister() v1alpha2.CanaryLister {
	return v1alpha2.NewCanaryLister(f.Informer().GetIndexer())
}