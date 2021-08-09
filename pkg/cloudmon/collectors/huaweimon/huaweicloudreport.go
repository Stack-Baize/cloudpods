package huaweimon

import (
	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/apis/compute"
	"yunion.io/x/onecloud/pkg/cloudmon/collectors/common"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
)

func init() {
	factory := SHwCloudReportFactory{}
	common.RegisterFactory(&factory)
}

type SHwCloudReportFactory struct {
}

func (self *SHwCloudReportFactory) NewCloudReport(provider *common.SProvider, session *mcclient.ClientSession,
	args *common.ReportOptions, operatorType string) common.ICloudReport {
	return &SHwCloudReport{
		common.CloudReportBase{
			SProvider: provider,
			Session:   session,
			Args:      args,
			Operator:  operatorType,
		},
	}
}

func (self *SHwCloudReportFactory) GetId() string {
	return compute.CLOUD_PROVIDER_HUAWEI
}

type SHwCloudReport struct {
	common.CloudReportBase
}

func (self *SHwCloudReport) Report() error {
	var servers []jsonutils.JSONObject
	var err error
	switch self.Operator {
	case "redis":
		servers, err = self.GetAllserverOfThisProvider(&modules.ElasticCache)
	case "rds":
		servers, err = self.GetAllserverOfThisProvider(&modules.DBInstance)
	case "oss":
		servers, err = self.GetAllserverOfThisProvider(&modules.Buckets)
	case "elb":
		servers, err = self.GetAllserverOfThisProvider(&modules.Loadbalancers)
	default:
		servers, err = self.GetAllserverOfThisProvider(&modules.Servers)
	}
	if err != nil {
		return err
	}
	providerInstance, err := self.InitProviderInstance()
	if err != nil {
		return err
	}
	regionList, regionServerMap, err := self.GetAllRegionOfServers(servers, providerInstance)
	if err != nil {
		return err
	}
	for _, region := range regionList {
		servers := regionServerMap[region.GetGlobalId()]
		switch self.Operator {
		case "server":
			err = self.collectRegionMetricOfHost(region, servers)
		case "redis":
			err = self.collectRegionMetricOfRedis(region, servers)
		case "rds":
			err = self.collectRegionMetricOfRds(region, servers)
		case "oss":
			err = self.collectRegionMetricOfOss(region, servers)
		case "elb":
			err = self.collectRegionMetricOfElb(region, servers)
		}
		if err != nil {
			return err
		}
	}
	return nil
}