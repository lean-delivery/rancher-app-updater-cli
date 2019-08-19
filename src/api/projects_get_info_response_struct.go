package api

type ProjectsGetInfoResponse struct {
	Actions     struct{} `json:"actions"`
	CreateTypes struct {
		Cluster string `json:"cluster"`
	} `json:"createTypes"`
	Data []projectsGetInfoResponse_Project `json:"data"`
	Filters struct {
		AgentImage                          interface{} `json:"agentImage"`
		APIEndpoint                         interface{} `json:"apiEndpoint"`
		AppliedEnableNetworkPolicy          interface{} `json:"appliedEnableNetworkPolicy"`
		AppliedPodSecurityPolicyTemplateID  interface{} `json:"appliedPodSecurityPolicyTemplateId"`
		AuthImage                           interface{} `json:"authImage"`
		CaCert                              interface{} `json:"caCert"`
		Created                             interface{} `json:"created"`
		CreatorID                           interface{} `json:"creatorId"`
		DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
		DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
		Description                         interface{} `json:"description"`
		DesiredAgentImage                   interface{} `json:"desiredAgentImage"`
		DesiredAuthImage                    interface{} `json:"desiredAuthImage"`
		DockerRootDir                       interface{} `json:"dockerRootDir"`
		Driver                              interface{} `json:"driver"`
		EnableClusterAlerting               interface{} `json:"enableClusterAlerting"`
		EnableClusterMonitoring             interface{} `json:"enableClusterMonitoring"`
		EnableNetworkPolicy                 interface{} `json:"enableNetworkPolicy"`
		ID                                  interface{} `json:"id"`
		Internal                            interface{} `json:"internal"`
		Name                                []struct {
			Modifier string `json:"modifier"`
			Value    string `json:"value"`
		} `json:"name"`
		Removed              interface{} `json:"removed"`
		State                interface{} `json:"state"`
		Transitioning        interface{} `json:"transitioning"`
		TransitioningMessage interface{} `json:"transitioningMessage"`
		UUID                 interface{} `json:"uuid"`
	} `json:"filters"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Pagination struct {
		Limit int `json:"limit"`
		Total int `json:"total"`
	} `json:"pagination"`
	ResourceType string `json:"resourceType"`
	Sort         struct {
		Links struct {
			AgentImage                         string `json:"agentImage"`
			APIEndpoint                        string `json:"apiEndpoint"`
			AppliedPodSecurityPolicyTemplateID string `json:"appliedPodSecurityPolicyTemplateId"`
			AuthImage                          string `json:"authImage"`
			CaCert                             string `json:"caCert"`
			Description                        string `json:"description"`
			DesiredAgentImage                  string `json:"desiredAgentImage"`
			DesiredAuthImage                   string `json:"desiredAuthImage"`
			DockerRootDir                      string `json:"dockerRootDir"`
			Driver                             string `json:"driver"`
			State                              string `json:"state"`
			Transitioning                      string `json:"transitioning"`
			TransitioningMessage               string `json:"transitioningMessage"`
			UUID                               string `json:"uuid"`
		} `json:"links"`
		Order   string `json:"order"`
		Reverse string `json:"reverse"`
	} `json:"sort"`
	Type string `json:"type"`
}

type projectsGetInfoResponse_Project struct {
	Actions struct {
		BackupEtcd            string `json:"backupEtcd"`
		EnableMonitoring      string `json:"enableMonitoring"`
		ExportYaml            string `json:"exportYaml"`
		GenerateKubeconfig    string `json:"generateKubeconfig"`
		ImportYaml            string `json:"importYaml"`
		RestoreFromEtcdBackup string `json:"restoreFromEtcdBackup"`
		RotateCertificates    string `json:"rotateCertificates"`
	} `json:"actions"`
	AgentImage  string `json:"agentImage"`
	Allocatable struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
		Pods   string `json:"pods"`
	} `json:"allocatable"`
	Annotations struct {
		Authz_management_cattle_io_creator_role_bindings            string `json:"authz.management.cattle.io/creator-role-bindings"`
		Lifecycle_cattle_io_create_cluster_agent_controller_cleanup string `json:"lifecycle.cattle.io/create.cluster-agent-controller-cleanup"`
		Lifecycle_cattle_io_create_cluster_provisioner_controller   string `json:"lifecycle.cattle.io/create.cluster-provisioner-controller"`
		Lifecycle_cattle_io_create_cluster_scoped_gc                string `json:"lifecycle.cattle.io/create.cluster-scoped-gc"`
		Lifecycle_cattle_io_create_mgmt_cluster_rbac_remove         string `json:"lifecycle.cattle.io/create.mgmt-cluster-rbac-remove"`
		Provisioner_cattle_io_ke_driver_update                      string `json:"provisioner.cattle.io/ke-driver-update"`
	} `json:"annotations"`
	APIEndpoint                        string `json:"apiEndpoint"`
	AppliedEnableNetworkPolicy         bool   `json:"appliedEnableNetworkPolicy"`
	AppliedPodSecurityPolicyTemplateID string `json:"appliedPodSecurityPolicyTemplateId"`
	AppliedSpec                        struct {
		DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
		DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
		Description                         string      `json:"description"`
		DesiredAgentImage                   string      `json:"desiredAgentImage"`
		DesiredAuthImage                    string      `json:"desiredAuthImage"`
		DisplayName                         string      `json:"displayName"`
		DockerRootDir                       string      `json:"dockerRootDir"`
		EnableClusterAlerting               bool        `json:"enableClusterAlerting"`
		EnableClusterMonitoring             bool        `json:"enableClusterMonitoring"`
		EnableNetworkPolicy                 bool        `json:"enableNetworkPolicy"`
		Internal                            bool        `json:"internal"`
		LocalClusterAuthEndpoint            struct {
			Enabled bool   `json:"enabled"`
			Type    string `json:"type"`
		} `json:"localClusterAuthEndpoint"`
		RancherKubernetesEngineConfig struct {
			AddonJobTimeout int `json:"addonJobTimeout"`
			Authentication  struct {
				Strategy string `json:"strategy"`
				Type     string `json:"type"`
			} `json:"authentication"`
			Authorization struct {
				Type string `json:"type"`
			} `json:"authorization"`
			BastionHost struct {
				SSHAgentAuth bool   `json:"sshAgentAuth"`
				Type         string `json:"type"`
			} `json:"bastionHost"`
			CloudProvider struct {
				AwsCloudProvider struct {
					Global struct {
						Disable_security_group_ingress bool   `json:"disable-security-group-ingress"`
						Disable_strict_zone_check      bool   `json:"disable-strict-zone-check"`
						Elb_security_group             string `json:"elb-security-group"`
						Kubernetes_cluster_id          string `json:"kubernetes-cluster-id"`
						Kubernetes_cluster_tag         string `json:"kubernetes-cluster-tag"`
						Role_arn                       string `json:"role-arn"`
						Routetable_id                  string `json:"routetable-id"`
						Subnet_id                      string `json:"subnet-id"`
						Type                           string `json:"type"`
						Vpc                            string `json:"vpc"`
						Zone                           string `json:"zone"`
					} `json:"global"`
					Type string `json:"type"`
				} `json:"awsCloudProvider"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"cloudProvider"`
			IgnoreDockerVersion bool `json:"ignoreDockerVersion"`
			Ingress             struct {
				Provider string `json:"provider"`
				Type     string `json:"type"`
			} `json:"ingress"`
			KubernetesVersion string `json:"kubernetesVersion"`
			Monitoring        struct {
				Provider string `json:"provider"`
				Type     string `json:"type"`
			} `json:"monitoring"`
			Network struct {
				Options struct {
					FlannelBackendType string `json:"flannel_backend_type"`
				} `json:"options"`
				Plugin string `json:"plugin"`
				Type   string `json:"type"`
			} `json:"network"`
			Nodes []struct {
				Address          string `json:"address"`
				HostnameOverride string `json:"hostnameOverride"`
				InternalAddress  string `json:"internalAddress"`
				Labels           struct {
					Cattle_io_creator string `json:"cattle.io/creator"`
				} `json:"labels"`
				NodeID       string   `json:"nodeId"`
				Port         string   `json:"port"`
				Role         []string `json:"role"`
				SSHAgentAuth bool     `json:"sshAgentAuth"`
				Type         string   `json:"type"`
				User         string   `json:"user"`
			} `json:"nodes"`
			Restore struct {
				Restore bool   `json:"restore"`
				Type    string `json:"type"`
			} `json:"restore"`
			Services struct {
				Etcd struct {
					BackupConfig struct {
						Enabled        bool        `json:"enabled"`
						IntervalHours  int         `json:"intervalHours"`
						Retention      int         `json:"retention"`
						S3BackupConfig interface{} `json:"s3BackupConfig"`
						Type           string      `json:"type"`
					} `json:"backupConfig"`
					Creation  string `json:"creation"`
					ExtraArgs struct {
						Election_timeout   string `json:"election-timeout"`
						Heartbeat_interval string `json:"heartbeat-interval"`
					} `json:"extraArgs"`
					Retention string `json:"retention"`
					Snapshot  bool   `json:"snapshot"`
					Type      string `json:"type"`
				} `json:"etcd"`
				KubeAPI struct {
					AlwaysPullImages     bool   `json:"alwaysPullImages"`
					PodSecurityPolicy    bool   `json:"podSecurityPolicy"`
					ServiceNodePortRange string `json:"serviceNodePortRange"`
					Type                 string `json:"type"`
				} `json:"kubeApi"`
				KubeController struct {
					Type string `json:"type"`
				} `json:"kubeController"`
				Kubelet struct {
					FailSwapOn bool   `json:"failSwapOn"`
					Type       string `json:"type"`
				} `json:"kubelet"`
				Kubeproxy struct {
					Type string `json:"type"`
				} `json:"kubeproxy"`
				Scheduler struct {
					Type string `json:"type"`
				} `json:"scheduler"`
				Type string `json:"type"`
			} `json:"services"`
			SSHAgentAuth bool   `json:"sshAgentAuth"`
			Type         string `json:"type"`
		} `json:"rancherKubernetesEngineConfig"`
		Type string `json:"type"`
	} `json:"appliedSpec"`
	AuthImage    string `json:"authImage"`
	BaseType     string `json:"baseType"`
	CaCert       string `json:"caCert"`
	Capabilities struct {
		IngressCapabilities []struct {
			CustomDefaultBackend bool   `json:"customDefaultBackend"`
			IngressProvider      string `json:"ingressProvider"`
			Type                 string `json:"type"`
		} `json:"ingressCapabilities"`
		LoadBalancerCapabilities struct {
			Enabled              bool     `json:"enabled"`
			HealthCheckSupported bool     `json:"healthCheckSupported"`
			ProtocolsSupported   []string `json:"protocolsSupported"`
			Provider             string   `json:"provider"`
			Type                 string   `json:"type"`
		} `json:"loadBalancerCapabilities"`
		NodePoolScalingSupported bool   `json:"nodePoolScalingSupported"`
		NodePortRange            string `json:"nodePortRange"`
		Type                     string `json:"type"`
	} `json:"capabilities"`
	Capacity struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
		Pods   string `json:"pods"`
	} `json:"capacity"`
	ComponentStatuses []struct {
		Conditions []struct {
			Message string `json:"message"`
			Status  string `json:"status"`
			Type    string `json:"type"`
		} `json:"conditions"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"componentStatuses"`
	Conditions []struct {
		Status string `json:"status"`
		Type   string `json:"type"`
	} `json:"conditions"`
	Created                             string      `json:"created"`
	CreatedTS                           int         `json:"createdTS"`
	CreatorID                           string      `json:"creatorId"`
	DefaultClusterRoleForProjectMembers interface{} `json:"defaultClusterRoleForProjectMembers"`
	DefaultPodSecurityPolicyTemplateID  interface{} `json:"defaultPodSecurityPolicyTemplateId"`
	Description                         string      `json:"description"`
	DesiredAgentImage                   string      `json:"desiredAgentImage"`
	DesiredAuthImage                    string      `json:"desiredAuthImage"`
	DockerRootDir                       string      `json:"dockerRootDir"`
	Driver                              string      `json:"driver"`
	EnableClusterAlerting               bool        `json:"enableClusterAlerting"`
	EnableClusterMonitoring             bool        `json:"enableClusterMonitoring"`
	EnableNetworkPolicy                 bool        `json:"enableNetworkPolicy"`
	ID                                  string      `json:"id"`
	Internal                            bool        `json:"internal"`
	Labels                              struct {
		Cattle_io_creator string `json:"cattle.io/creator"`
	} `json:"labels"`
	Limits struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
		Pods   string `json:"pods"`
	} `json:"limits"`
	Links struct {
		ClusterAlertGroups          string `json:"clusterAlertGroups"`
		ClusterAlertRules           string `json:"clusterAlertRules"`
		ClusterAlerts               string `json:"clusterAlerts"`
		ClusterCatalogs             string `json:"clusterCatalogs"`
		ClusterLoggings             string `json:"clusterLoggings"`
		ClusterMonitorGraphs        string `json:"clusterMonitorGraphs"`
		ClusterRegistrationTokens   string `json:"clusterRegistrationTokens"`
		ClusterRoleTemplateBindings string `json:"clusterRoleTemplateBindings"`
		EtcdBackups                 string `json:"etcdBackups"`
		Namespaces                  string `json:"namespaces"`
		NodePools                   string `json:"nodePools"`
		Nodes                       string `json:"nodes"`
		Notifiers                   string `json:"notifiers"`
		PersistentVolumes           string `json:"persistentVolumes"`
		Projects                    string `json:"projects"`
		Remove                      string `json:"remove"`
		Self                        string `json:"self"`
		Shell                       string `json:"shell"`
		StorageClasses              string `json:"storageClasses"`
		Subscribe                   string `json:"subscribe"`
		Templates                   string `json:"templates"`
		Tokens                      string `json:"tokens"`
		Update                      string `json:"update"`
	} `json:"links"`
	LocalClusterAuthEndpoint struct {
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
	} `json:"localClusterAuthEndpoint"`
	Name                          string `json:"name"`
	RancherKubernetesEngineConfig struct {
		AddonJobTimeout int `json:"addonJobTimeout"`
		Authentication  struct {
			Strategy string `json:"strategy"`
			Type     string `json:"type"`
		} `json:"authentication"`
		Authorization struct {
			Type string `json:"type"`
		} `json:"authorization"`
		BastionHost struct {
			SSHAgentAuth bool   `json:"sshAgentAuth"`
			Type         string `json:"type"`
		} `json:"bastionHost"`
		CloudProvider struct {
			AwsCloudProvider struct {
				Global struct {
					Disable_security_group_ingress bool   `json:"disable-security-group-ingress"`
					Disable_strict_zone_check      bool   `json:"disable-strict-zone-check"`
					Elb_security_group             string `json:"elb-security-group"`
					Kubernetes_cluster_id          string `json:"kubernetes-cluster-id"`
					Kubernetes_cluster_tag         string `json:"kubernetes-cluster-tag"`
					Role_arn                       string `json:"role-arn"`
					Routetable_id                  string `json:"routetable-id"`
					Subnet_id                      string `json:"subnet-id"`
					Type                           string `json:"type"`
					Vpc                            string `json:"vpc"`
					Zone                           string `json:"zone"`
				} `json:"global"`
				Type string `json:"type"`
			} `json:"awsCloudProvider"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"cloudProvider"`
		IgnoreDockerVersion bool `json:"ignoreDockerVersion"`
		Ingress             struct {
			Provider string `json:"provider"`
			Type     string `json:"type"`
		} `json:"ingress"`
		KubernetesVersion string `json:"kubernetesVersion"`
		Monitoring        struct {
			Provider string `json:"provider"`
			Type     string `json:"type"`
		} `json:"monitoring"`
		Network struct {
			Options struct {
				FlannelBackendType string `json:"flannel_backend_type"`
			} `json:"options"`
			Plugin string `json:"plugin"`
			Type   string `json:"type"`
		} `json:"network"`
		Restore struct {
			Restore bool   `json:"restore"`
			Type    string `json:"type"`
		} `json:"restore"`
		Services struct {
			Etcd struct {
				BackupConfig struct {
					Enabled        bool        `json:"enabled"`
					IntervalHours  int         `json:"intervalHours"`
					Retention      int         `json:"retention"`
					S3BackupConfig interface{} `json:"s3BackupConfig"`
					Type           string      `json:"type"`
				} `json:"backupConfig"`
				Creation  string `json:"creation"`
				ExtraArgs struct {
					Election_timeout   string `json:"election-timeout"`
					Heartbeat_interval string `json:"heartbeat-interval"`
				} `json:"extraArgs"`
				Retention string `json:"retention"`
				Snapshot  bool   `json:"snapshot"`
				Type      string `json:"type"`
			} `json:"etcd"`
			KubeAPI struct {
				AlwaysPullImages     bool   `json:"alwaysPullImages"`
				PodSecurityPolicy    bool   `json:"podSecurityPolicy"`
				ServiceNodePortRange string `json:"serviceNodePortRange"`
				Type                 string `json:"type"`
			} `json:"kubeApi"`
			KubeController struct {
				Type string `json:"type"`
			} `json:"kubeController"`
			Kubelet struct {
				FailSwapOn bool   `json:"failSwapOn"`
				Type       string `json:"type"`
			} `json:"kubelet"`
			Kubeproxy struct {
				Type string `json:"type"`
			} `json:"kubeproxy"`
			Scheduler struct {
				Type string `json:"type"`
			} `json:"scheduler"`
			Type string `json:"type"`
		} `json:"services"`
		SSHAgentAuth bool   `json:"sshAgentAuth"`
		Type         string `json:"type"`
	} `json:"rancherKubernetesEngineConfig"`
	Requested struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
		Pods   string `json:"pods"`
	} `json:"requested"`
	State                string `json:"state"`
	Transitioning        string `json:"transitioning"`
	TransitioningMessage string `json:"transitioningMessage"`
	Type                 string `json:"type"`
	UUID                 string `json:"uuid"`
	Version              struct {
		BuildDate    string `json:"buildDate"`
		Compiler     string `json:"compiler"`
		GitCommit    string `json:"gitCommit"`
		GitTreeState string `json:"gitTreeState"`
		GitVersion   string `json:"gitVersion"`
		GoVersion    string `json:"goVersion"`
		Major        string `json:"major"`
		Minor        string `json:"minor"`
		Platform     string `json:"platform"`
		Type         string `json:"type"`
	} `json:"version"`
}