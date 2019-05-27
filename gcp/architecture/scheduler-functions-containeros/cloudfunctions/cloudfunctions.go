package cloudfunctions

import (
	"context"
	"fmt"
	"google.golang.org/api/compute/v1"
	"log"
)

// DeployMIG checks to see if a specified Managed Instance Group(MIG) is up and running. If it is running,
// it should shut it down, if it is unavailable it will restart it.
func DeployCloudRun() {

}

const ProjectID = "martins-project-2014113"
const InstanceTemplateName = "batch-mig-template"
const Zone = "us-central1-a"
const NetworkName = "default"

func DeployContainerOS() {
	cs, err := InitComputeService()
	if err != nil {
		log.Fatal(err)
	}

	_, err := CreateContainerOSInstance(cs)
	if err != nil {
		log.Fatal(err)
	}
}

// InitComputeService obtains the compute service that allows us to use the compute API
func InitComputeService() (*compute.Service, error) {
	ctx := context.Background()
	return compute.NewService(ctx)
}

func CreateContainerOSInstance(computeService *compute.Service) (*compute.Operation, error) {
	instance := &compute.Instance{
		Name: "worker-instance",
		MachineType: fmt.Sprintf("zones/%s/machineTypes/n1-standard-1", Zone),
		NetworkInterfaces: []*compute.NetworkInterface{{Network: fmt.Sprintf("projects/%s/global/networks/default", ProjectID),}},
		Disks: []*compute.AttachedDisk{
			{
				Boot:       true,         // The first disk must be a boot disk.
				AutoDelete: false,        //Optional
				Mode:       "READ_WRITE", //Mode should be READ_WRITE or READ_ONLY
				Interface:  "SCSI",       //SCSI or NVME - NVME only for SSDs
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    "worker-instance-boot-disk",
					SourceImage: "projects/cos-cloud/global/images/family/cos-stable",
				},
			},
		},
		Metadata:

	}
	return computeService.Instances.Insert(ProjectID, Zone, instance).Do()
}

// Returns an instance template
func GetInstanceTemplate(computeService *compute.Service) (*compute.InstanceTemplate, error) {
	get := computeService.InstanceTemplates.Get(ProjectID, InstanceTemplateName)
	return get.Do(nil)
}

// StartMIG starts a managed instance group. This MIG will not have any instances.
func StartMIG(computeService *compute.Service, template *compute.InstanceTemplate) {
	instanceGroup := compute.InstanceGroup{
		Zone:Zone,
		Description: "Instance Group for Random Batch Work",
		Network: fmt.Sprintf("https://www.googleapis." +
			"com/compute/v1/projects/%s/global/networks/%s", ProjectID, NetworkName),
		NamedPorts: []*compute.NamedPort{{Name: "http", Port: 8080}},
	}

	createInstanceGroupCall := computeService.InstanceGroups.Insert(ProjectID, Zone, &instanceGroup)
	_, err := createInstanceGroupCall.Do()
	if err != nil {
		log.Fatal(err)
	}

	computeService.InstanceGroupManagers.SetInstanceTemplate(ProjectID, Zone, )

	computeService.InstanceGroupManagers.Insert(ProjectID, Zone, )
}

// ApplyInstanceTemplateToMIG will apply an instance template to a Managed Instance Group
func ApplyInstanceTemplateToMIG(instanceGroupName string) {

}