//
// Copyright (c) 2019-2020 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package apis

import (
	"github.com/che-incubator/che-workspace-operator/internal/cluster"
	apis "github.com/che-incubator/che-workspace-operator/pkg/apis/workspace/v1alpha1"
	routeV1 "github.com/openshift/api/route/v1"
	templateV1 "github.com/openshift/api/template/v1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		apis.SchemeBuilder.AddToScheme,
	)
	if isOS, err := cluster.IsOpenShift(); isOS && err == nil {
		AddToSchemes = append(AddToSchemes,
			routeV1.AddToScheme,
			templateV1.AddToScheme,
		)
	}
}
