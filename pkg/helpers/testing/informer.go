package testing

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	clusterclient "open-cluster-management.io/api/client/cluster/clientset/versioned"
	clusterinformers "open-cluster-management.io/api/client/cluster/informers/externalversions"
	clusterapiv1 "open-cluster-management.io/api/cluster/v1"
	clusterapiv1alpha1 "open-cluster-management.io/api/cluster/v1alpha1"
	clusterapiv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
)

func NewClusterInformerFactory(clusterClient clusterclient.Interface, objects ...runtime.Object) clusterinformers.SharedInformerFactory {
	clusterInformerFactory := clusterinformers.NewSharedInformerFactory(clusterClient, time.Minute*10)
	clusterStore := clusterInformerFactory.Cluster().V1().ManagedClusters().Informer().GetStore()
	clusterSetStore := clusterInformerFactory.Cluster().V1beta1().ManagedClusterSets().Informer().GetStore()
	clusterSetBindingStore := clusterInformerFactory.Cluster().V1beta1().ManagedClusterSetBindings().Informer().GetStore()
	placementStore := clusterInformerFactory.Cluster().V1alpha1().Placements().Informer().GetStore()
	placementDecisionStore := clusterInformerFactory.Cluster().V1alpha1().PlacementDecisions().Informer().GetStore()
	addOnPlacementStore := clusterInformerFactory.Cluster().V1alpha1().AddOnPlacementScores().Informer().GetStore()

	for _, obj := range objects {
		switch obj.(type) {
		case *clusterapiv1.ManagedCluster:
			clusterStore.Add(obj)
		case *clusterapiv1beta1.ManagedClusterSet:
			clusterSetStore.Add(obj)
		case *clusterapiv1beta1.ManagedClusterSetBinding:
			clusterSetBindingStore.Add(obj)
		case *clusterapiv1alpha1.Placement:
			placementStore.Add(obj)
		case *clusterapiv1alpha1.PlacementDecision:
			placementDecisionStore.Add(obj)
		case *clusterapiv1alpha1.AddOnPlacementScore:
			addOnPlacementStore.Add(obj)
		}
	}

	return clusterInformerFactory
}
