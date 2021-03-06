package docker

import "time"

type ChangeType int

const (
	ChangeModify = iota
	ChangeAdd
	ChangeDelete
)

type ContainerConfig struct {
	Hostname        string
	Domainname      string
	User            string
	Memory          int64
	MemorySwap      int64
	CpuShares       int64
	Cpuset          string
	AttachStdin     bool
	AttachStdout    bool
	AttachStderr    bool
	PortSpecs       []string
	ExposedPorts    map[string]struct{}
	Tty             bool
	OpenStdin       bool
	StdinOnce       bool
	Env             []string
	Cmd             []string
	Image           string
	Volumes         map[string]struct{}
	WorkingDir      string
	Entrypoint      []string
	NetworkDisabled bool
	OnBuild         []string

	// This is used only by the create command
	HostConfig HostConfig
}

type HostConfig struct {
	Binds           []string
	ContainerIDFile string
	LxcConf         []map[string]string
	Privileged      bool
	PortBindings    map[string][]PortBinding
	Links           []string
	PublishAllPorts bool
	Dns             []string
	DnsSearch       []string
	VolumesFrom     []string
	NetworkMode     string
	RestartPolicy   RestartPolicy
	Memory          int64
	MemorySwap      int64
	CpuShares       int64
	CpuPeriod       int64
	CpusetCpus      string
	CpusetMems      string
	LogConfig       LogConfig
}

type LogConfig struct {
	Type   string
	Config map[string]string
}

type ExecConfig struct {
	AttachStdin  bool
	AttachStdout bool
	AttachStderr bool
	Tty          bool
	Cmd          []string
	Container    string
	Detach       bool
}

type LogOptions struct {
	Follow     bool
	Stdout     bool
	Stderr     bool
	Timestamps bool
	Tail       int64
}

type RestartPolicy struct {
	Name              string
	MaximumRetryCount int64
}

type PortBinding struct {
	HostIp   string
	HostPort string
}

type ContainerInfo struct {
	Id      string
	Created string
	Path    string
	Name    string
	Args    []string
	ExecIDs []string
	Config  *ContainerConfig
	State   struct {
		Running    bool
		Paused     bool
		Restarting bool
		Pid        int
		ExitCode   int
		StartedAt  time.Time
		FinishedAt time.Time
		Ghost      bool
	}
	Image           string
	NetworkSettings struct {
		IpAddress   string
		IpPrefixLen int
		Gateway     string
		Bridge      string
		Ports       map[string][]PortBinding
	}
	SysInitPath    string
	ResolvConfPath string
	Volumes        map[string]string
	HostConfig     *HostConfig
}

type Port struct {
	IP          string
	PrivatePort int
	PublicPort  int
	Type        string
}

type Container struct {
	Id         string
	Names      []string
	Image      string
	Command    string
	Created    int64
	Status     string
	Ports      []Port
	SizeRw     int64
	SizeRootFs int64
}

type Event struct {
	Id     string
	Status string
	From   string
	Time   int64
}

type Version struct {
	Version   string
	GitCommit string
	GoVersion string
}

type RespContainersCreate struct {
	Id       string
	Warnings []string
}

type Image struct {
	Created     int64
	Id          string
	ParentId    string
	RepoTags    []string
	Size        int64
	VirtualSize int64
}

//returned by History
type ImageLayer struct {
	Id        string
	Created   int64
	CreatedBy string
	Comment   string
	Size      int64
	Tags      []string
}

// returned by InspectImage
type ImageInfo struct {
	Created   string
	Container string
	Id        string
	Parent    string
	Size      int
}

type Info struct {
	ID              string
	Containers      int64
	DockerRootDir   string
	Driver          string
	DriverStatus    [][]string
	ExecutionDriver string
	Images          int64
	KernelVersion   string
	OperatingSystem string
	NCPU            int64
	MemTotal        int64
	Name            string
	Labels          []string
}

type ImageDelete struct {
	Deleted  string
	Untagged string
}

type AttachOptions struct {
	Logs   bool
	Stream bool
	Stdin  bool
	Stdout bool
	Stderr bool
}

type CommitOptions struct {
	Container string
	Repo      string
	Tag       string
	Comment   string
	Author    string
}

type ContainerChange struct {
	Kind int
	Path string
}

type TagOptions struct {
	Repo  string
	Force bool
	Tag   string
}
