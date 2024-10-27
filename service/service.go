package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	pb "github/challenge/protos/distributor/v1"

	"gopkg.in/yaml.v3"
)

type Distributor struct {
	Include []string `yaml:"include"`
	Exclude []string `yaml:"exclude"`
	Parent  string   `yaml:"parent,omitempty"`
}

type DistributorService struct {
	pb.UnimplementedCreateDistributorServiceServer
	pb.UnimplementedCheckPermissionServiceServer
	Distributors map[string]Distributor `yaml:"DISTRIBUTORS"`
	ValidRegions map[string]bool
}

func (ds *DistributorService) LoadDistributors(filename string) error {
	ds.Distributors = make(map[string]Distributor)
	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %w", err)
	}
	err = yaml.Unmarshal(file, &ds.Distributors)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML file: %w", err)
	}
	return nil
}

func (ds *DistributorService) WriteDistributors(filename string) error {
	existingDistributors := ds.Distributors
	fmt.Println(existingDistributors)
	for name, distributor := range ds.Distributors {
		existingDistributors[name] = distributor
	}
	data, err := yaml.Marshal(existingDistributors)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing YAML file: %w", err)
	}
	return nil
}

func (ds *DistributorService) LoadCSVRegions(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	ds.ValidRegions = make(map[string]bool)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		country := strings.ToUpper(record[2])
		province := fmt.Sprintf("%s-%s", strings.ToUpper(record[1]), country)
		city := fmt.Sprintf("%s-%s", strings.ToUpper(record[0]), province)

		ds.ValidRegions[country] = true
		ds.ValidRegions[province] = true
		ds.ValidRegions[city] = true
	}
	return nil
}

func (ds *DistributorService) ValidateRegionFormat(regions []string) error {
	for _, region := range regions {
		if !ds.ValidRegions[region] {
			return fmt.Errorf("invalid region format or unknown region: %s", region)
		}
	}
	return nil
}

func (ds *DistributorService) ValidateRegions(parentID string, includes, excludes []string) error {
	_, exists := ds.Distributors[parentID]
	if !exists {
		return nil
	}
	for _, region := range includes {
		req := &pb.CheckPermissionRequest{
			Name:   parentID,
			Region: region,
		}
		res, err := ds.CheckPermission(context.Background(), req)
		if err != nil || res.Message == "NO" {
			return fmt.Errorf("region %s is not allowed under parent distributor %s", region, parentID)
		}
	}
	return nil
}

func (ds *DistributorService) CreateDistributor(ctx context.Context, req *pb.CreateDistributorRequest) (*pb.CreateDistributorResponse, error) {
	loadErr := ds.LoadDistributors("distributors.yaml")
	if loadErr != nil {
		fmt.Println("error in loading distributors")
	}
	fmt.Println(ds.Distributors)
	if _, exists := ds.Distributors[req.Name]; exists {
		return &pb.CreateDistributorResponse{
			Status:  "ERROR",
			Message: "Distributor already exists",
		}, nil
	}

	if err := ds.ValidateRegionFormat(req.Include); err != nil {
		return &pb.CreateDistributorResponse{
			Status:  "ERROR",
			Message: fmt.Sprintf("Include region validation error: %v", err),
		}, nil
	}
	if err := ds.ValidateRegionFormat(req.Exclude); err != nil {
		return &pb.CreateDistributorResponse{
			Status:  "ERROR",
			Message: fmt.Sprintf("Exclude region validation error: %v", err),
		}, nil
	}

	if req.Parent != "" {
		err := ds.ValidateRegions(req.Parent, req.Include, req.Exclude)
		if err != nil {
			return &pb.CreateDistributorResponse{
				Status:  "ERROR",
				Message: fmt.Sprintf("Validation error: %v", err),
			}, nil
		}
	}

	ds.Distributors[req.Name] = Distributor{
		Parent:  req.Parent,
		Include: req.Include,
		Exclude: req.Exclude,
	}

	err := ds.WriteDistributors("distributors.yaml")
	if err != nil {
		return &pb.CreateDistributorResponse{
			Status:  "ERROR",
			Message: fmt.Sprintf("Error saving distributor to File: %v", err),
		}, err
	}
	return &pb.CreateDistributorResponse{
		Status:  "SUCCESS",
		Message: "Distributor created successfully",
	}, nil
}

func (ds *DistributorService) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	loadErr := ds.LoadDistributors("distributors.yaml")
	if loadErr != nil {
		fmt.Println("error in loading distributors")
	}

	if err := ds.ValidateRegionFormat([]string{req.Region}); err != nil {
		return &pb.CheckPermissionResponse{
			Message: fmt.Sprintf("region validation error: %v", err),
		}, err
	}

	isRegionMatch := func(region, target string) bool {
		return region == target ||
			strings.HasSuffix(region, "-"+target) ||
			strings.HasPrefix(region, target+"-")
	}

	current := req.Name
	hasInclusion := false

	for current != "" {
		dist, exists := ds.Distributors[current]
		if !exists {
			break
		}

		for _, exc := range dist.Exclude {
			if isRegionMatch(req.Region, exc) {
				return &pb.CheckPermissionResponse{Message: "NO"}, nil
			}
		}

		for _, inc := range dist.Include {
			if isRegionMatch(req.Region, inc) {
				hasInclusion = true
				break
			}
		}

		if len(dist.Include) > 0 && !hasInclusion {
			return &pb.CheckPermissionResponse{Message: "NO"}, nil
		}

		current = dist.Parent
	}

	if hasInclusion {
		return &pb.CheckPermissionResponse{Message: "YES"}, nil
	}
	return &pb.CheckPermissionResponse{Message: "NO"}, nil
}
