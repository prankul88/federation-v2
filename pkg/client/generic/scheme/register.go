/*
Copyright 2019 The Kubernetes Authors.

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

package scheme

import (
	fedapis "github.com/kubernetes-sigs/federation-v2/pkg/apis"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"
	crapis "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
)

var Scheme = runtime.NewScheme()
var localSchemeBuilder = runtime.SchemeBuilder{
	fedapis.AddToScheme,
	k8sscheme.AddToScheme,
	crapis.AddToScheme,
}

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}

var AddToScheme = localSchemeBuilder.AddToScheme
