package e2e

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("operand deletion", func() {
	// a ClusterServiceVersions owns a cluster-scoped CustomResourceDefinition
	//   a second ClusterServiceVersion owns the same CustomResourceDefinition
	//   a second ClusterServiceVersion requires the same CustomResourceDefinition
	// a ClusterServiceVersions owns a namespace-scoped CustomResourceDefinition
	//   no associated OperatorGroup exists
	//   multiple associated OperatorGroups exist
	//   an incorrect target namespaces annotation is set on the ClusterServiceVersion
	//   a second ClusterServiceVersion owns the same CustomResourceDefinition (both in and out of target namespace set)
	//   a second ClusterServiceVersion requires the same CustomResourceDefinition (both in and out of target namespace set)

	PIt("has end-to-end tests", func() {
		Fail("not implemented")
	})
})
