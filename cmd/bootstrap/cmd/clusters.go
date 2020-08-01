package cmd

import (
	"fmt"

	"github.com/dapperlabs/flow-go/cmd/bootstrap/run"
	hotstuff "github.com/dapperlabs/flow-go/consensus/hotstuff/model"
	model "github.com/dapperlabs/flow-go/model/bootstrap"
	"github.com/dapperlabs/flow-go/model/cluster"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/flow/filter"
)

// Construct cluster assignment with internal and partner nodes uniformly
// distributed across clusters
func constructClusterAssignment(partnerNodes, internalNodes []model.NodeInfo) (flow.AssignmentList, flow.ClusterList) {

	partners := model.ToIdentityList(partnerNodes).Filter(filter.HasRole(flow.RoleCollection))
	internals := model.ToIdentityList(internalNodes).Filter(filter.HasRole(flow.RoleCollection))

	nClusters := flagCollectionClusters
	assignments := make(flow.AssignmentList, nClusters)

	// first, round-robin internal nodes into each cluster
	for i, node := range internals {
		assignments[i%len(assignments)] = append(assignments[i%len(assignments)], node.NodeID)
	}

	// next, round-robin partner nodes into each cluster
	for i, node := range partners {
		assignments[i%len(assignments)] = append(assignments[i%len(assignments)], node.NodeID)
	}

	collectors := append(partners, internals...)
	clusters, err := flow.NewClusterList(assignments, collectors)
	if err != nil {
		log.Fatal().Err(err).Msg("could not create cluster list")
	}

	return assignments, clusters
}

// TODO this should be defined in protocol state
func constructRootBlocksForClusters(clusters flow.ClusterList) []*cluster.Block {
	clusterBlocks := run.GenerateRootClusterBlocks(clusters)

	for _, clusterBlock := range clusterBlocks {
		// cluster ID is equivalent to chain ID
		clusterID := clusterBlock.Header.ChainID
		// TODO remove
		writeJSON(fmt.Sprintf(model.PathRootClusterBlock, clusterID), clusterBlock)
	}

	return clusterBlocks
}

func constructRootQCsForClusters(clusterList flow.ClusterList, nodeInfos []model.NodeInfo, block *flow.Block, clusterBlocks []*cluster.Block) []*hotstuff.QuorumCertificate {

	if len(clusterBlocks) != len(clusterList) {
		log.Fatal().Int("len(clusterBlocks)", len(clusterBlocks)).Int("len(clusterList)", len(clusterList)).
			Msg("number of clusters needs to equal number of cluster blocks")
	}

	for i, cluster := range clusterList {
		signers := filterClusterSigners(cluster, nodeInfos)

		qc, err := run.GenerateClusterRootQC(signers, block, clusterBlocks[i])
		if err != nil {
			log.Fatal().Err(err).Int("cluster index", i).Msg("generating collector cluster root QC failed")
		}

		// cluster ID is equivalent to chain ID
		clusterID := clusterBlocks[i].Header.ChainID
		// TODO remove
		writeJSON(fmt.Sprintf(model.PathRootClusterQC, clusterID), qc)
	}
}

// Filters a list of nodes to include only nodes that will sign the QC for the
// given cluster. The resulting list of nodes is only nodes that are in the
// given cluster AND are not partner nodes (ie. we have the private keys).
func filterClusterSigners(cluster flow.IdentityList, nodeInfos []model.NodeInfo) []model.NodeInfo {

	var filtered []model.NodeInfo
	for _, node := range nodeInfos {
		_, isInCluster := cluster.ByNodeID(node.NodeID)
		isNotPartner := node.Type() == model.NodeInfoTypePrivate

		if isInCluster && isNotPartner {
			filtered = append(filtered, node)
		}
	}

	return filtered
}
