package cloudfunctions

import (
	"context"
	"fmt"
	"google.golang.org/api/compute/v1"
	"log"
	"net/http"
	"os"
)

var ProjectID = ""
var Zone = ""
var Region = ""
var InstanceName = ""

// DeployInstance will use the Golang GCP API to deploy a GCE instance with given startup-script that creates a text file
// and logs the time. If the instance is there. It will shut it down, and the shutdown script will be invoked.
func DeployInstance(w http.ResponseWriter, r *http.Request) {
	ProjectID = os.Getenv("PROJECT_ID")
	Zone = os.Getenv("ZONE")
	Region = os.Getenv("REGION")
	InstanceName = os.Getenv("INSTANCE_NAME")

	cs, err := InitComputeService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	//Try retrieve the instance. On error we shall HAPHAZARDLY assume it doesnt exist and try create it.
	// There could be other reasons.
	instance, err := GetInstance(cs)
	if err != nil {
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(err.Error() + " instance may not exist yet"))
		log.Print(err)

		_, err = CreateInstance(cs)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("creating instance " + InstanceName + "in zone: " + Zone))
		}
	}else {
		switch instance.Status {
		case "RUNNING":
			stopInstance(cs, w)
		case "PROVISIONING":
			stopInstance(cs, w)
		case "STAGING":
			stopInstance(cs, w)
		case "STOPPED":
			startInstance(cs, w)
		case "TERMINATED":
			startInstance(cs, w)
		default:
			msg := "instance is in intermediate state: " + instance.Status
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(msg))
			log.Println(msg)
		}
	}
}

// InitComputeService obtains the compute service that allows us to use the compute API
func InitComputeService() (*compute.Service, error) {
	ctx := context.Background()
	return compute.NewService(ctx)
}

// GetInstance passes in the instance name supplied and retrieves it.
// An error indicates an instance that was never created.
// A non-nil error indicates an instance is present whether in the RUNNING or TERMINATED state.
func GetInstance(computeService *compute.Service) (*compute.Instance, error) {
	return computeService.Instances.Get(ProjectID, Zone, InstanceName).Do()
}

// StopInstance will delete an instance with the name specified in the InstanceName variable.
func StopInstance(computeService *compute.Service) (*compute.Operation, error) {
	return computeService.Instances.Stop(ProjectID, Zone, InstanceName).Do()
}

// StartInstance begins an instance with the given name
func StartInstance(computeService *compute.Service) (*compute.Operation, error) {
	return computeService.Instances.Start(ProjectID, Zone, InstanceName).Do()
}

// CreateInstance creates a given instance with metadata that logs its information.
func CreateInstance(computeService *compute.Service) (*compute.Operation, error) {
	startupMetadata := "#! /bin/bash\n\necho \"I am STARTING some work at $(date)\" | sudo tee -a $HOME/work.txt"
	shutdownMetadata := "#! /bin/bash\n\necho \"I am FINISHING some work at $(date)\" | sudo tee -a $(HOME)/work.txt"

	instance := &compute.Instance{
		Name:        InstanceName,
		MachineType: fmt.Sprintf("zones/%s/machineTypes/n1-standard-1", Zone),
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				Name:       "default",
				Subnetwork: fmt.Sprintf("projects/%s/regions/%s/subnetworks/default", ProjectID, Region),
				AccessConfigs: []*compute.AccessConfig{
					{
						Name:        "External NAT",
						Type:        "ONE_TO_ONE_NAT",
						NetworkTier: "PREMIUM",
					},
				},
			},
		},
		Scheduling: &compute.Scheduling{
			Preemptible: true,
		},
		Disks: []*compute.AttachedDisk{
			{
				Boot:       true,         // The first disk must be a boot disk.
				AutoDelete: false,        //Optional
				Mode:       "READ_WRITE", //Mode should be READ_WRITE or READ_ONLY
				Interface:  "SCSI",       //SCSI or NVME - NVME only for SSDs
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    "worker-instance-boot-disk",
					SourceImage: "projects/debian-cloud/global/images/family/debian-9",
				},
			},
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				{
					Key:   "startup-script",
					Value: &startupMetadata,
				},
				{
					Key:   "shutdown-script",
					Value: &shutdownMetadata,
				},
			},
		},
	}
	return computeService.Instances.Insert(ProjectID, Zone, instance).Do()
}

// startInstance is a wrapper function for the switch statement
func startInstance(cs *compute.Service, w http.ResponseWriter) {
	operation, err := StartInstance(cs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := operation.MarshalJSON()
	w.Write(data)
}

// startInstance is a wrapper function for the switch statement
func stopInstance(cs *compute.Service, w http.ResponseWriter) {
	operation, err := StopInstance(cs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	data, _ := operation.MarshalJSON()
	w.Write(data)
}
