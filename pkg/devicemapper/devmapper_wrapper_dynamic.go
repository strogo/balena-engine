// +build linux,cgo,!static_build
// +build !no_devicemapper

package devicemapper // import "github.com/docker/docker/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
