package api

import (
	"net/http"

	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/disk"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/firewall"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/floatingip"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/loadbalancer"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/network"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/objectstorage"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/user"
	"github.com/teliti-dev/tf-idcloudhost-client/idcloudhost/vm"
)

type APIClient struct {
	VM            *vm.VirtualMachineAPI
	Disk          *disk.DiskAPI
	FloatingIP    *floatingip.FloatingIPAPI
	User          *user.UserAPI
	Network       *network.NetworkAPI
	Firewall      *firewall.FirewallAPI
	ObjectStorage *objectstorage.ObjectStorageAPI
	LoadBalancer  *loadbalancer.LoadBalancerAPI
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewClient(authToken string, loc string) (*APIClient, error) {
	c := http.Client{}
	var ApiClient = APIClient{
		VM:            &vm.VirtualMachineAPI{},
		Disk:          &disk.DiskAPI{},
		FloatingIP:    &floatingip.FloatingIPAPI{},
		User:          &user.UserAPI{},
		Network:       &network.NetworkAPI{},
		Firewall:      &firewall.FirewallAPI{},
		ObjectStorage: &objectstorage.ObjectStorageAPI{},
		LoadBalancer:  &loadbalancer.LoadBalancerAPI{},
	}

	if err := ApiClient.VM.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.Disk.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.FloatingIP.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.User.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.Network.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.Firewall.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.ObjectStorage.Init(&c, authToken, loc); err != nil {
		return nil, err
	}
	if err := ApiClient.LoadBalancer.Init(&c, authToken, loc); err != nil {
		return nil, err
	}

	return &ApiClient, nil
}
