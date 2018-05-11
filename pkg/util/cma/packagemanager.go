package cma

import(
	sdsapi "github.com/samsung-cnct/cluster-manager-api/pkg/apis/cma/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/samsung-cnct/cluster-manager-api/pkg/util/k8sutil"
	"k8s.io/client-go/rest"

)

type SDSPackageManagerOptions struct {
	Name string
	Namespace string
	Version string
	ClusterWide bool
	AdminNamespaces []string
}

func GenerateSDSPackageManager(options SDSPackageManagerOptions) sdsapi.SDSPackageManager {
	return sdsapi.SDSPackageManager{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.Name,
		},
		Spec: sdsapi.SDSPackageManagerSpec{
			Name: options.Name,
			Namespace: options.Namespace,
			Version: options.Version,
			Permissions: sdsapi.PackageManagerPermissions{
				ClusterWide: options.ClusterWide,
				Namespaces: options.AdminNamespaces,
			},
		},
	}
}



func CreateSDSPackageManager(packageManager sdsapi.SDSPackageManager, namespace string, config *rest.Config) (bool, error) {
	var err error
	SetLogger()
	client := prepareRestClient(config)

	_, err = client.CmaV1alpha1().SDSPackageManagers(namespace).Create(&packageManager)
	if err != nil && !k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("PackageManager -->%s<-- Cannot be created, error was %v", packageManager.ObjectMeta.Name, err)
		return false, err
	} else if k8sutil.IsResourceAlreadyExistsError(err) {
		logger.Infof("PackageManager -->%s<-- Already exists, cannot recreate", packageManager.ObjectMeta.Name)
		return false, err
	}

	return true, err
}