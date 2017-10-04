package api

import "github.com/Azure/acs-engine/pkg/api/agentPoolOnlyApi/v20170831"
import "github.com/Azure/acs-engine/pkg/api/agentPoolOnlyApi/v20171031"

///////////////////////////////////////////////////////////
// The converter exposes functions to convert the top level
// ContainerService resource
//
// All other functions are internal helper functions used
// for converting.
///////////////////////////////////////////////////////////

// ConvertContainerServiceToV20170831AgentPoolOnly converts an unversioned ContainerService to a v20170831 ContainerService
func ConvertContainerServiceToV20170831AgentPoolOnly(api *ContainerService) *v20170831.ManagedCluster {
	v20170831HCP := &v20170831.ManagedCluster{}
	v20170831HCP.ID = api.ID
	v20170831HCP.Location = api.Location
	v20170831HCP.Name = api.Name
	if api.Plan != nil {
		v20170831HCP.Plan = &v20170831.ResourcePurchasePlan{}
		convertResourcePurchasePlanToV20170831AgentPoolOnly(api.Plan, v20170831HCP.Plan)
	}
	v20170831HCP.Tags = map[string]string{}
	for k, v := range api.Tags {
		v20170831HCP.Tags[k] = v
	}
	v20170831HCP.Type = api.Type
	v20170831HCP.Properties = &v20170831.Properties{}
	convertPropertiesToV20170831AgentPoolOnly(api.Properties, v20170831HCP.Properties)
	return v20170831HCP
}

// convertResourcePurchasePlanToV20170831 converts a v20170831 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertResourcePurchasePlanToV20170831AgentPoolOnly(api *ResourcePurchasePlan, v20170831 *v20170831.ResourcePurchasePlan) {
	v20170831.Name = api.Name
	v20170831.Product = api.Product
	v20170831.PromotionCode = api.PromotionCode
	v20170831.Publisher = api.Publisher
}

func convertPropertiesToV20170831AgentPoolOnly(api *Properties, p *v20170831.Properties) {
	p.ProvisioningState = v20170831.ProvisioningState(api.ProvisioningState)
	if api.OrchestratorProfile != nil {
		if api.OrchestratorProfile.OrchestratorRelease != "" {
			p.KubernetesRelease = api.OrchestratorProfile.OrchestratorRelease
		}

		if api.OrchestratorProfile.OrchestratorVersion != "" {
			p.KubernetesVersion = api.OrchestratorProfile.OrchestratorVersion
		}
	}
	if api.HostedMasterProfile != nil {
		p.DNSPrefix = api.HostedMasterProfile.DNSPrefix
		p.FQDN = api.HostedMasterProfile.FQDN
	}
	p.AgentPoolProfiles = []*v20170831.AgentPoolProfile{}
	for _, apiProfile := range api.AgentPoolProfiles {
		v20170831Profile := &v20170831.AgentPoolProfile{}
		convertAgentPoolProfileToV20170831AgentPoolOnly(apiProfile, v20170831Profile)
		p.AgentPoolProfiles = append(p.AgentPoolProfiles, v20170831Profile)
	}
	if api.LinuxProfile != nil {
		p.LinuxProfile = &v20170831.LinuxProfile{}
		convertLinuxProfileToV20170831AgentPoolOnly(api.LinuxProfile, p.LinuxProfile)
	}
	if api.WindowsProfile != nil {
		p.WindowsProfile = &v20170831.WindowsProfile{}
		convertWindowsProfileToV20170831AgentPoolOnly(api.WindowsProfile, p.WindowsProfile)
	}
	if api.ServicePrincipalProfile != nil {
		p.ServicePrincipalProfile = &v20170831.ServicePrincipalProfile{}
		convertServicePrincipalProfileToV20170831AgentPoolOnly(api.ServicePrincipalProfile, p.ServicePrincipalProfile)
	}
}

func convertLinuxProfileToV20170831AgentPoolOnly(api *LinuxProfile, obj *v20170831.LinuxProfile) {
	obj.AdminUsername = api.AdminUsername
	obj.SSH.PublicKeys = []v20170831.PublicKey{}
	for _, d := range api.SSH.PublicKeys {
		obj.SSH.PublicKeys = append(obj.SSH.PublicKeys, v20170831.PublicKey{
			KeyData: d.KeyData,
		})
	}
}

func convertWindowsProfileToV20170831AgentPoolOnly(api *WindowsProfile, v20170831Profile *v20170831.WindowsProfile) {
	v20170831Profile.AdminUsername = api.AdminUsername
	v20170831Profile.AdminPassword = api.AdminPassword
}

func convertAgentPoolProfileToV20170831AgentPoolOnly(api *AgentPoolProfile, p *v20170831.AgentPoolProfile) {
	p.Name = api.Name
	p.Count = api.Count
	p.VMSize = api.VMSize
	p.OSType = v20170831.OSType(api.OSType)
	p.SetSubnet(api.Subnet)
	p.OSDiskSizeGB = api.OSDiskSizeGB
	p.StorageProfile = api.StorageProfile
	p.VnetSubnetID = api.VnetSubnetID
}

func convertServicePrincipalProfileToV20170831AgentPoolOnly(api *ServicePrincipalProfile, v20170831 *v20170831.ServicePrincipalProfile) {
	v20170831.ClientID = api.ClientID
	v20170831.Secret = api.Secret
	// v20170831.KeyvaultSecretRef = api.KeyvaultSecretRef
}

// ConvertContainerServiceTov20171031AgentPoolOnly converts an unversioned ContainerService to a v20171031 ContainerService
func ConvertContainerServiceTov20171031AgentPoolOnly(api *ContainerService) *v20171031.ManagedCluster {
	v20171031HCP := &v20171031.ManagedCluster{}
	v20171031HCP.ID = api.ID
	v20171031HCP.Location = api.Location
	v20171031HCP.Name = api.Name
	if api.Plan != nil {
		v20171031HCP.Plan = &v20171031.ResourcePurchasePlan{}
		convertResourcePurchasePlanTov20171031AgentPoolOnly(api.Plan, v20171031HCP.Plan)
	}
	v20171031HCP.Tags = map[string]string{}
	for k, v := range api.Tags {
		v20171031HCP.Tags[k] = v
	}
	v20171031HCP.Type = api.Type
	v20171031HCP.Properties = &v20171031.Properties{}
	convertPropertiesTov20171031AgentPoolOnly(api.Properties, v20171031HCP.Properties)
	return v20171031HCP
}

// convertResourcePurchasePlanTov20171031 converts a v20171031 ResourcePurchasePlan to an unversioned ResourcePurchasePlan
func convertResourcePurchasePlanTov20171031AgentPoolOnly(api *ResourcePurchasePlan, v20171031 *v20171031.ResourcePurchasePlan) {
	v20171031.Name = api.Name
	v20171031.Product = api.Product
	v20171031.PromotionCode = api.PromotionCode
	v20171031.Publisher = api.Publisher
}

func convertPropertiesTov20171031AgentPoolOnly(api *Properties, p *v20171031.Properties) {
	p.ProvisioningState = v20171031.ProvisioningState(api.ProvisioningState)
	if api.OrchestratorProfile != nil {
		if api.OrchestratorProfile.OrchestratorVersion != "" {
			p.KubernetesVersion = api.OrchestratorProfile.OrchestratorVersion
		}
	}
	if api.HostedMasterProfile != nil {
		p.DNSPrefix = api.HostedMasterProfile.DNSPrefix
		p.FQDN = api.HostedMasterProfile.FQDN
	}
	p.AgentPoolProfiles = []*v20171031.AgentPoolProfile{}
	for _, apiProfile := range api.AgentPoolProfiles {
		v20171031Profile := &v20171031.AgentPoolProfile{}
		convertAgentPoolProfileTov20171031AgentPoolOnly(apiProfile, v20171031Profile)
		p.AgentPoolProfiles = append(p.AgentPoolProfiles, v20171031Profile)
	}
	if api.LinuxProfile != nil {
		p.LinuxProfile = &v20171031.LinuxProfile{}
		convertLinuxProfileTov20171031AgentPoolOnly(api.LinuxProfile, p.LinuxProfile)
	}
	if api.WindowsProfile != nil {
		p.WindowsProfile = &v20171031.WindowsProfile{}
		convertWindowsProfileTov20171031AgentPoolOnly(api.WindowsProfile, p.WindowsProfile)
	}
	if api.ServicePrincipalProfile != nil {
		p.ServicePrincipalProfile = &v20171031.ServicePrincipalProfile{}
		convertServicePrincipalProfileTov20171031AgentPoolOnly(api.ServicePrincipalProfile, p.ServicePrincipalProfile)
	}
}

func convertLinuxProfileTov20171031AgentPoolOnly(api *LinuxProfile, obj *v20171031.LinuxProfile) {
	obj.AdminUsername = api.AdminUsername
	obj.SSH.PublicKeys = []v20171031.PublicKey{}
	for _, d := range api.SSH.PublicKeys {
		obj.SSH.PublicKeys = append(obj.SSH.PublicKeys, v20171031.PublicKey{
			KeyData: d.KeyData,
		})
	}
}

func convertWindowsProfileTov20171031AgentPoolOnly(api *WindowsProfile, v20171031Profile *v20171031.WindowsProfile) {
	v20171031Profile.AdminUsername = api.AdminUsername
	v20171031Profile.AdminPassword = api.AdminPassword
}

func convertAgentPoolProfileTov20171031AgentPoolOnly(api *AgentPoolProfile, p *v20171031.AgentPoolProfile) {
	p.Name = api.Name
	p.Count = api.Count
	p.VMSize = api.VMSize
	p.OSType = v20171031.OSType(api.OSType)
	p.SetSubnet(api.Subnet)
	p.OSDiskSizeGB = api.OSDiskSizeGB
	p.StorageProfile = api.StorageProfile
	p.VnetSubnetID = api.VnetSubnetID
}

func convertServicePrincipalProfileTov20171031AgentPoolOnly(api *ServicePrincipalProfile, v20171031 *v20171031.ServicePrincipalProfile) {
	v20171031.ClientID = api.ClientID
	v20171031.Secret = api.Secret
	// v20171031.KeyvaultSecretRef = api.KeyvaultSecretRef
}
