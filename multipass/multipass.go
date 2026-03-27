package multipass

type Instance struct {
	Name        string
	State       string
	Snapshots   string
	IP          string
	Release     string
	Image       string
	ImageHash   string
	CPUs        string
	Load        string
	DiskUsage   string
	TotalDisk   string
	MemoryUsage string
	MemoryTotal string
	Mounts      string
}
