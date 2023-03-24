package policies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bastionzero/bastionzero-sdk-go/bastionzero/service/policies/policytype"
	"github.com/bastionzero/bastionzero-sdk-go/internal"
	"github.com/bastionzero/bastionzero-sdk-go/internal/client"
)

const (
	kubernetesBasePath   = policiesBasePath + "/kubernetes"
	kubernetesSinglePath = kubernetesBasePath + "/%s"
)

// ClusterUser refers to the Kubernetes subject user that a Kubernetes policy
// applies to
type ClusterUser struct {
	Name string `json:"name"`
}

// ClusterGroup refers to the Kubernetes subject group that a Kubernetes policy
// applies to
type ClusterGroup struct {
	Name string `json:"name"`
}

// Cluster refers to the BastionZero Cluster target that a Kubernetes policy
// applies to
type Cluster struct {
	ID string `json:"id"`
}

// KubernetesPolicy represents a Kubernetes policy. Kubernetes policies provide
// access to Cluster targets.
type KubernetesPolicy struct {
	*Policy

	Environments  *[]Environment  `json:"environments,omitempty"`
	Clusters      *[]Cluster      `json:"clusters,omitempty"`
	ClusterUsers  *[]ClusterUser  `json:"clusterUsers,omitempty"`
	ClusterGroups *[]ClusterGroup `json:"clusterGroups,omitempty"`
}

// ListKubernetesPolicies lists all Kubernetes policies
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/kubernetes
func (s *PoliciesService) ListKubernetesPolicies(ctx context.Context, opts *ListPolicyOptions) ([]KubernetesPolicy, *http.Response, error) {
	u := kubernetesBasePath
	u, err := client.AddOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policyList := new([]KubernetesPolicy)
	resp, err := s.Client.Do(req, policyList)
	if err != nil {
		return nil, resp, err
	}

	return *policyList, resp, nil
}

// CreateKubernetesPolicy creates a new Kubernetes policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#post-/api/v2/policies/kubernetes
func (s *PoliciesService) CreateKubernetesPolicy(ctx context.Context, policy *KubernetesPolicy) (*KubernetesPolicy, *http.Response, error) {
	u := kubernetesBasePath
	req, err := s.Client.NewRequest(ctx, http.MethodPost, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(KubernetesPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// GetKubernetesPolicy fetches the specified Kubernetes policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#get-/api/v2/policies/kubernetes/-id-
func (s *PoliciesService) GetKubernetesPolicy(ctx context.Context, policyID string) (*KubernetesPolicy, *http.Response, error) {
	u := fmt.Sprintf(kubernetesSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	policy := new(KubernetesPolicy)
	resp, err := s.Client.Do(req, policy)
	if err != nil {
		return nil, resp, err
	}

	return policy, resp, nil
}

// DeleteKubernetesPolicy deletes the specified Kubernetes policy
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#delete-/api/v2/policies/kubernetes/-id-
func (s *PoliciesService) DeleteKubernetesPolicy(ctx context.Context, policyID string) (*http.Response, error) {
	u := fmt.Sprintf(kubernetesSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ModifyKubernetesPolicy updates a Kubernetes policy. All user populated
// fields are mutable except for policy.TimeExpires.
//
// BastionZero API docs: https://cloud.bastionzero.com/api/#patch-/api/v2/policies/kubernetes/-id-
func (s *PoliciesService) ModifyKubernetesPolicy(ctx context.Context, policyID string, policy *KubernetesPolicy) (*KubernetesPolicy, *http.Response, error) {
	u := fmt.Sprintf(kubernetesSinglePath, policyID)
	req, err := s.Client.NewRequest(ctx, http.MethodPatch, u, policy)
	if err != nil {
		return nil, nil, err
	}

	p := new(KubernetesPolicy)
	resp, err := s.Client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}

// Ensure KubernetesPolicy implementation satisfies the expected interfaces.
var (
	// KubernetesPolicy implements PolicyInterface
	_ PolicyInterface = &KubernetesPolicy{}
)

func (p *KubernetesPolicy) GetPolicyType() policytype.PolicyType { return policytype.Kubernetes }

func (p *KubernetesPolicy) GetEnvironments() []Environment {
	if p.Environments == nil {
		return []Environment{}
	}
	return *p.Environments
}
func (p *KubernetesPolicy) GetClusters() []Cluster {
	if p.Clusters == nil {
		return []Cluster{}
	}
	return *p.Clusters
}
func (p *KubernetesPolicy) GetClusterUsers() []ClusterUser {
	if p.ClusterUsers == nil {
		return []ClusterUser{}
	}
	return *p.ClusterUsers
}
func (p *KubernetesPolicy) GetClusterGroups() []ClusterGroup {
	if p.ClusterGroups == nil {
		return []ClusterGroup{}
	}
	return *p.ClusterGroups
}

func (p *KubernetesPolicy) GetEnvironmentsAsStringList() []string {
	return internal.MapSlice(p.GetEnvironments(), func(e Environment) string { return e.ID })
}
func (p *KubernetesPolicy) GetClustersAsStringList() []string {
	return internal.MapSlice(p.GetClusters(), func(e Cluster) string { return e.ID })
}
func (p *KubernetesPolicy) GetClusterUsersAsStringList() []string {
	return internal.MapSlice(p.GetClusterUsers(), func(e ClusterUser) string { return e.Name })
}
func (p *KubernetesPolicy) GetClusterGroupsAsStringList() []string {
	return internal.MapSlice(p.GetClusterGroups(), func(e ClusterGroup) string { return e.Name })
}
