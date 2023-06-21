package demo

type MessageInfo struct {
	Name string
}

type ServiceMethodInfo struct {
	Name string

	Input  string
	Output string
}

type ServiceInfo struct {
	Name string

	Methods []*ServiceMethodInfo
}

type FileInfo struct {
	Package string
	Name    string

	Messages []*MessageInfo
	Services []*ServiceInfo
}
