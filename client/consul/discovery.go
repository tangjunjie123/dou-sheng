package consul

import (
	"unsafe"
)

type DiscoveryClient interface {

	/**
	 * Gets all ServiceInstances associated with a particular serviceId.
	 * @param serviceId The serviceId to query.
	 * @return A List of ServiceInstance.
	 */
	GetInstances(serviceId string) ([]ServiceInstance, error)

	/**
	 * @return All known service IDs.
	 */
	GetServices() ([]string, error)
}

func (c consulServiceRegistry) GetInstances(serviceId string) ([]ServiceInstance, error) {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	if len(catalogService) > 0 {
		result := make([]ServiceInstance, len(catalogService))
		for index, sever := range catalogService {
			s := DefaultServiceInstance{
				InstanceId: sever.ServiceID,
				ServiceId:  sever.ServiceName,
				Host:       sever.Address,
				Port:       sever.ServicePort,
				Metadata:   sever.ServiceMeta,
			}
			result[index] = s
		}
		return result, nil
	}
	return nil, nil
}

func (c consulServiceRegistry) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	result := make([]string, unsafe.Sizeof(services))
	index := 0
	for serviceName, _ := range services {
		result[index] = serviceName
		index++
	}
	return result, nil
}

// new a consulServiceRegistry instance
// common is optional

func Discovery(consulIp, token string, consulPort int, serviceId string) ([]ServiceInstance, error) {
	registryDiscoveryClient, err := NewConsulServiceRegistry(consulIp, consulPort, token)
	if err != nil {
		panic(err)
	}
	return registryDiscoveryClient.GetInstances(serviceId)
}
