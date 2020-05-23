package marathon

//=== Marathon Server Info definitions

// Marathon server info vía API endpoint /v2/info
type Info struct {
	Name            string     `json:"name"`
	Version         string     `json:"version"`
	Buildref        string     `json:"buildref"`
	Elected         bool       `json:"elected"`
	Leader          string     `json:"leader"`
	FrameworkID     string     `json:"frameworkId"`
	MarathonConfig  Config     `json:"marathon_config"`
	ZookeeperConfig ZkConfig   `json:"zookeeper_config"`
	HTTPConf        HTTPConfig `json:"http_config"`
}

// Marathon Server Config definitions
type Config struct {
	AccessControlAllowOrigin       []string      `json:"access_control_allow_origin"`
	Checkpoint                     bool          `json:"checkpoint"`
	DeclineOfferDuration           int           `json:"decline_offer_duration"`
	DefaultNetworkName             string        `json:"default_network_name"`
	EnvVarsPrefix                  string        `json:"env_vars_prefix"`
	Executor                       string        `json:"executor"`
	FailoverTimeout                int           `json:"failover_timeout"`
	Features                       []interface{} `json:"features"`
	FrameworkName                  string        `json:"framework_name"`
	Ha                             bool          `json:"ha"`
	Hostname                       string        `json:"hostname"`
	LaunchToken                    int           `json:"launch_token"`
	LaunchTokenRefreshInterval     int           `json:"launch_token_refresh_interval"`
	LeaderProxyConnectionTimeoutMs int           `json:"leader_proxy_connection_timeout_ms"`
	LeaderProxyReadTimeoutMs       int           `json:"leader_proxy_read_timeout_ms"`
	LocalPortMax                   int           `json:"local_port_max"`
	LocalPortMin                   int           `json:"local_port_min"`
	Master                         string        `json:"master"`
	MaxInstancesPerOffer           int           `json:"max_instances_per_offer"`
	MesosBridgeName                string        `json:"mesos_bridge_name"`
	MesosHeartbeatFailureThreshold int           `json:"mesos_heartbeat_failure_threshold"`
	MesosHeartbeatInterval         int           `json:"mesos_heartbeat_interval"`
	MesosLeaderUIURL               string        `json:"mesos_leader_ui_url"`
	MesosRole                      string        `json:"mesos_role"`
	MesosUser                      string        `json:"mesos_user"`
	MinReviveOffersInterval        int           `json:"min_revive_offers_interval"`
	OfferMatchingTimeout           int           `json:"offer_matching_timeout"`
	OnElectedPrepareTimeout        int           `json:"on_elected_prepare_timeout"`
	ReconciliationInitialDelay     int           `json:"reconciliation_initial_delay"`
	ReconciliationInterval         int           `json:"reconciliation_interval"`
	ReviveOffersForNewApps         bool          `json:"revive_offers_for_new_apps"`
	ReviveOffersRepetitions        int           `json:"revive_offers_repetitions"`
	ScaleAppsInitialDelay          int           `json:"scale_apps_initial_delay"`
	ScaleAppsInterval              int           `json:"scale_apps_interval"`
	StoreCache                     bool          `json:"store_cache"`
	TaskLaunchConfirmTimeout       int           `json:"task_launch_confirm_timeout"`
	TaskLaunchTimeout              int           `json:"task_launch_timeout"`
	TaskLostExpungeInitialDelay    int           `json:"task_lost_expunge_initial_delay"`
	TaskLostExpungeInterval        int           `json:"task_lost_expunge_interval"`
	TaskReservationTimeout         int           `json:"task_reservation_timeout"`
	WebuiURL                       string        `json:"webui_url"`
}

type ZkConfig struct {
	Zk                     string `json:"zk"`
	ZkCompression          bool   `json:"zk_compression"`
	ZkCompressionThreshold int    `json:"zk_compression_threshold"`
	ZkConnectionTimeout    int    `json:"zk_connection_timeout"`
	ZkMaxNodeSize          int    `json:"zk_max_node_size"`
	ZkMaxVersions          int    `json:"zk_max_versions"`
	ZkSessionTimeout       int    `json:"zk_session_timeout"`
	ZkTimeout              int    `json:"zk_timeout"`
}

type HTTPConfig struct {
	Port       int `json:"http_port"`
	SecurePort int `json:"https_port"`
}
