package cloudfunctions

import (
	"context"
	"google.golang.org/api/compute/v1"
	"log"
	"net/http"
	"os"
)

var ProjectID = ""
var Zone = ""
var Region = ""
var InstanceName = ""

// PollInstance will use the Golang GCP API to deploy poll a given instance.
// If the instance is stopped or terminated, it will restart the instance
func PollInstance(w http.ResponseWriter, r *http.Request) {
	ProjectID = os.Getenv("PROJECT_ID")
	Zone = os.Getenv("ZONE")
	Region = os.Getenv("REGION")
	InstanceName = os.Getenv("INSTANCE_NAME")

	cs, err := InitComputeService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}

	//Try retrieve the instance. On error we have to assume the instance was never created or the environment
	// variables supplied were faulty.
	instance, err := GetInstance(cs)
	if err != nil {
		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Write([]byte(err.Error() + " instance may not exist yet or error with supplied environment variables"))
		log.Print(err)
		return
	}

	// If the instance isa stopped, terminated or suspended state. Call the startInstance method.
	switch instance.Status {
	case "STOPPED":
		startInstance(cs, w)
	case "TERMINATED":
		startInstance(cs, w)
	case "SUSPENDED":
		startInstance(cs, w)
	default:
		msg := "lazarus is not yet dead: " + instance.Status
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(msg))
		log.Println(msg)
		return
	}

	msg := "lazarus-vm was " + instance.Status + ", it's being resurrected: " + instance.Status
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(msg))
	log.Println(msg)
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


// StartInstance begins an instance with the given name
func StartInstance(computeService *compute.Service) (*compute.Operation, error) {
	return computeService.Instances.Start(ProjectID, Zone, InstanceName).Do()
}

// startInstance is a wrapper function for the switch statement used in PollInstance method
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
