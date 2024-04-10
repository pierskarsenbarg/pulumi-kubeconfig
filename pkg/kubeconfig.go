package pkg

type KubeConfig struct {
	APIVersion     string      `json:"apiVersion"`
	Clusters       []Clusters  `json:"clusters"`
	Contexts       []Contexts  `json:"contexts"`
	CurrentContext string      `json:"current-context"`
	Kind           string      `json:"kind"`
	Preferences    Preferences `json:"preferences"`
	Users          []Users     `json:"users"`
}
type Cluster struct {
	Server                   string `json:"server"`
	CertificateAuthorityData string `json:"certificate-authority-data"`
}
type Clusters struct {
	Cluster Cluster `json:"cluster"`
	Name    string  `json:"name"`
}
type Context struct {
	Cluster string `json:"cluster"`
	User    string `json:"user"`
}
type Contexts struct {
	Context Context `json:"context"`
	Name    string  `json:"name"`
}
type Preferences struct {
}
type Exec struct {
	APIVersion        string   `json:"apiVersion"`
	Command           string   `json:"command"`
	Args              []string `json:"args,omitempty"`
	Env               []Env    `json:"env,omitempty"`
	ProvideCluserInfo bool     `json:"providerClusterInfo,omitempty"`
	InstallHint       string   `json:"installHint,omitempty"`
}
type Env struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type User struct {
	Exec Exec `json:"exec"`
}
type Users struct {
	Name string `json:"name"`
	User User   `json:"user"`
}
