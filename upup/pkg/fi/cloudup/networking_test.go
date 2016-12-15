/*
Copyright 2016 The Kubernetes Authors.

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

package cloudup

import (
	"testing"
	api "k8s.io/kops/pkg/apis/kops"
	"os"
)

func Test_FindCNIAssetFromEnvironmentVariable(t *testing.T) {

	desiredCNIVersion := "https://storage.googleapis.com/kubernetes-release/network-plugins/cni-TEST-VERSION.tar.gz"
	os.Setenv(CNI_VERSION_URL, desiredCNIVersion)
	defer func(){
		os.Unsetenv(CNI_VERSION_URL)
	}()

	cluster := &api.Cluster{}
	cniAsset, cniAssetHashString := findCNIAssets(cluster)

	if cniAsset != desiredCNIVersion {
		t.Errorf("Expected CNI version from Environment variable %v, but got %v instead\n", desiredCNIVersion, cniAsset)
	}

	if cniAssetHashString != "" {
		t.Errorf("Expected Empty CNI Version Hash String, but got %v instead\n", cniAssetHashString)
	}
}

func Test_FindCNIAssetDefaultValue(t *testing.T) {

	cluster := &api.Cluster{}
	cniAsset, cniAssetHashString := findCNIAssets(cluster)

	if cniAsset != defaultCNIAsset {
		 t.Errorf("Expected default CNI version %v and got %v\n", defaultCNIAsset, cniAsset)
	}

	if cniAssetHashString != defaultCNIAssetHashString {
		t.Errorf("Expected default CNI Version Hash String %v and got %v\n", defaultCNIAssetHashString, cniAssetHashString)
	}

}