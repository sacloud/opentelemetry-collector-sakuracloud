package sacloud

import (
	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/iaas-api-go"
)

type Client struct {
	caller *iaas.Client

	authStatus    authStatusClient
	AutoBackup    AutoBackupClient
	Bill          BillClient
	Coupon        CouponClient
	Database      DatabaseClient
	ESME          ESMEClient
	Internet      InternetClient
	LoadBalancer  LoadBalancerClient
	LocalRouter   LocalRouterClient
	MobileGateway MobileGatewayClient
	NFS           NFSClient
	ProxyLB       ProxyLBClient
	Server        ServerClient
	SIM           SIMClient
	VPCRouter     VPCRouterClient
	Zone          ZoneClient

	WebAccel WebAccelClient
}

func NewClient(endpoint string, zones []string, options *client.Options) (*Client, error) {
	defaults, err := client.DefaultOption()
	if err != nil {
		return nil, err
	}

	caller := iaas.NewClientWithOptions(client.MergeOptions(defaults, options))

	iaas.SakuraCloudAPIRoot = endpoint
	iaas.SakuraCloudZones = zones

	return &Client{
		authStatus:    getAuthStatusClient(caller),
		AutoBackup:    getAutoBackupClient(caller, zones),
		Bill:          getBillClient(caller),
		Coupon:        getCouponClient(caller),
		Database:      getDatabaseClient(caller, zones),
		ESME:          getESMEClient(caller),
		Internet:      getInternetClient(caller, zones),
		LoadBalancer:  getLoadBalancerClient(caller, zones),
		LocalRouter:   getLocalRouterClient(caller),
		MobileGateway: getMobileGatewayClient(caller, zones),
		NFS:           getNFSClient(caller, zones),
		ProxyLB:       getProxyLBClient(caller),
		Server:        getServerClient(caller, zones),
		SIM:           getSIMClient(caller),
		VPCRouter:     getVPCRouterClient(caller, zones),
		Zone:          getZoneClient(caller),

		//WebAccel: getWebAccelClient(webaccelCaller),
	}, nil
}
