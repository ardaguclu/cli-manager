package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PluginSpec defines the desired state of Plugin
type PluginSpec struct {
	// ShortDescription of the plugin.
	// +required
	ShortDescription string `json:"shortDescription"`

	// Description of the plugin.
	// +optional
	Description string `json:"description,omitempty"`

	// Caveats of using the plugin.
	// +optional
	Caveats string `json:"caveats,omitempty"`

	// Homepage of the plugin.
	// +optional
	Homepage string `json:"homepage,omitempty"`

	// Version of the plugin.
	// +required
	Version string `json:"version"`

	// Platforms the plugin supports.
	// +required
	Platforms []PluginPlatform `json:"platforms"`
}

// PluginPlatform defines per-OS and per-Arch binaries for the given plugin.
type PluginPlatform struct {
	// Platform for the given binary (i.e. linux/amd64, darwin/amd64, windows/amd64).
	// +required
	Platform string `json:"platform"`

	// Image containing plugin.
	// +required
	Image string `json:"image"`

	// ImagePullSecret to use when connecting to an image registry that requires authentication.
	// +optional
	ImagePullSecret string `json:"imagePullSecret,omitempty"`

	// Files is a list of file locations within the image that need to be extracted.
	// +required
	Files []FileLocation `json:"files"`

	// Bin specifies the path to the plugin executable.
	// The path is relative to the root of the installation folder.
	// The binary will be linked after all FileOperations are executed.
	// If not specified, the first From item in the Files is set.
	// +optional
	Bin string `json:"bin"`
}

// FileLocation specifies a file copying operation from plugin archive to the
// installation directory.
type FileLocation struct {
	// From is the absolute file path within the image to copy from.
	// Directories and wildcards are not currently supported.
	// +required
	From string `json:"from"`

	// To is the relative path within the root of the installation folder to place the file.
	// Default is set to "." where points the default Krew directory.
	// +required
	// +kubebuilder:default:="."
	To string `json:"to"`
}

// PluginStatus defines the observed state of Plugin.
type PluginStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=Plugins,scope=Cluster

// Plugin is the Schema for the plugins API
type Plugin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PluginSpec   `json:"spec,omitempty"`
	Status PluginStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PluginList contains a list of Plugin
type PluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plugin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Plugin{}, &PluginList{})
}
