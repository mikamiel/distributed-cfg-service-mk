package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "distributed-cfg-service-mk/proto"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Stuct for storing config record in DB
type Config struct {
	gorm.Model
	Service    string
	Parameters []Parameter
}

// Stuct for storing key/value config pair in DB,
// tagging with 'uniqueIndex' fields that compose *multicolumn* unique index
type Parameter struct {
	gorm.Model
	Key      string `gorm:"uniqueIndex:key_plus_configid"`
	Value    string
	ConfigID uint `gorm:"uniqueIndex:key_plus_configid"`
}

// Stuct for storing config subscriber (client app) in DB,
// tagging with 'uniqueIndex' fields that compose *multicolumn* unique index
type Subscriber struct {
	gorm.Model
	ClientApp string `gorm:"uniqueIndex:clientapp_plus_service"`
	Service   string `gorm:"uniqueIndex:clientapp_plus_service"`
}

type GrpcDistributedConfigServer struct {
	db *gorm.DB
	pb.UnimplementedDistributedCfgServiceMKServer
}

// Main SQL query for retrieving filtered/merged config.
// !Warning: Postgre-specific raw SQL - probably won't work in MySQL or SQLite
const SqlGetKeyValsByTimestamp = `
			SELECT * FROM (
				SELECT DISTINCT ON ("key") configs.service,  parameters."key", parameters.value, configs.updated_at
				FROM configs
				LEFT JOIN parameters
				ON configs.id=parameters.config_id
				WHERE configs.service= ?
				AND configs.updated_at <= ?
				AND configs.deleted_at IS NULL
				ORDER BY parameters."key", configs.updated_at DESC    
				) AS sub_query
			WHERE value != ''
`

// Default environment variables
func getDefaultEnvVars() map[string]string {
	return map[string]string{
		//Postgres:
		"DB_HOST_NAME": "localhost",
		"DB_PORT":      "5432",
		"DB_USER":      "postgres",
		"DB_PASSWORD":  "postgres",
		"DB_NAME":      "postgres",
		"DB_SSLMODE":   "disable",
		"DB_TIMEZONE":  "Asia/Yekaterinburg",

		//Service listening port:
		"CFG_SERVICE_PORT": "50051",
	}
}

// Update default environment variables with actual ones
func updateDefaultEnvVarsWithActual() map[string]string {
	envVars := getDefaultEnvVars()
	for key, _ := range envVars {
		envVar := os.Getenv(key)
		if envVar == "" {
			continue
		} else {
			envVars[key] = envVar
		}
	}
	return envVars
}

// For some reason GORM db driver fails to resolve hostname of a DB inside docker container network,
// so resolve it manually as a workaround
func resolveHostNameToIp(hostname string) string {
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to resolve ip-address for host: %s   %v\n", hostname, err)
		log.Fatal(err)
	}
	return ips[0].String()
}

// Func to ensure that no duplicate key names sent by client in a single request
func TestForDuplicatesInParams(request *pb.Config) []string {
	testKeysForDuplicates := map[string]int{}
	duplicates := []string{}
	for _, param := range request.Parameters {
		if _, keyExists := testKeysForDuplicates[param.Key]; keyExists {
			testKeysForDuplicates[param.Key] = testKeysForDuplicates[param.Key] + 1
			duplicates = append(duplicates, param.Key)
		} else {
			testKeysForDuplicates[param.Key] = 1
		}
	}
	return duplicates
}

// CreateConfig() gRPC handler
func (srv *GrpcDistributedConfigServer) CreateConfig(ctx context.Context, request *pb.Config) (*pb.Timestamp, error) {
	result := srv.db.Limit(1).Where("service = ?", request.Service).Find(&Config{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists,
			"Config for service \"%s\" already exist, use UpdateConfig() to modify it", request.Service)
	}

	duplicates := TestForDuplicatesInParams(request)
	if len(duplicates) > 0 {
		log.Println("Creating config failed due to duplicate config parameter(s):", duplicates)
		return nil, status.Errorf(codes.InvalidArgument,
			"Creating config failed due to duplicate config parameter(s):%v", duplicates)
	}

	configToCreate := &Config{
		Service:    request.Service,
		Parameters: []Parameter{},
	}
	for _, param := range request.Parameters {
		configToCreate.Parameters = append(configToCreate.Parameters, Parameter{Key: param.Key, Value: param.Value})
	}

	if result := srv.db.Create(&configToCreate); result.Error != nil {
		log.Println(result.Error)
		return nil, status.Errorf(codes.Internal,
			"Unable to create config for %s", request.Service)
	}
	srv.db.First(&configToCreate, configToCreate.ID)
	timestamp := configToCreate.UpdatedAt

	return &pb.Timestamp{
			Service:   fmt.Sprintf(request.Service),
			Timestamp: timestamppb.New(timestamp),
		},
		status.Errorf(codes.OK,
			"Config for service \"%s\" created succesfully at timestamp %s", request.Service, timestamp.String())
}

// UpdateConfig() gRPC handler
func (srv *GrpcDistributedConfigServer) UpdateConfig(ctx context.Context, request *pb.Config) (*pb.Timestamp, error) {
	result := srv.db.Limit(1).Where("service = ?", request.Service).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Config for service \"%s\" not found, use CreateConfig() to create it from scratch", request.Service)
	}

	duplicates := TestForDuplicatesInParams(request)
	if len(duplicates) > 0 {
		log.Println("Updating config failed due to duplicate config parameter(s):", duplicates)
		return nil, status.Errorf(codes.InvalidArgument,
			"Updating config failed due to duplicate config parameter(s):%v", duplicates)
	}

	configToCreate := &Config{
		Service:    request.Service,
		Parameters: []Parameter{},
	}
	for _, param := range request.Parameters {
		configToCreate.Parameters = append(configToCreate.Parameters, Parameter{Key: param.Key, Value: param.Value})
	}

	if result := srv.db.Create(&configToCreate); result.Error != nil {
		log.Println(result.Error)
		return nil, status.Errorf(codes.Internal,
			"Unable to update config for ", request.Service)
	}
	srv.db.First(&configToCreate, configToCreate.ID)
	timestamp := configToCreate.UpdatedAt
	return &pb.Timestamp{
			Service:   request.Service,
			Timestamp: timestamppb.New(timestamp),
		},
		status.Errorf(codes.OK,
			"Config for service \"%s\" updated succesfully at timestamp %s", request.Service, timestamp.String())
}

// GetConfig() gRPC handler
func (srv *GrpcDistributedConfigServer) GetConfig(ctx context.Context, requestedService *pb.Service) (*pb.Config, error) {
	config := &Config{}
	configToSendBack := &pb.Config{}
	var paramsGrpc []*pb.Parameter

	// Check if config with name provided by client exists in DB:
	result := srv.db.Limit(1).Where("service = ?", requestedService.Name).Order("updated_at DESC").Find(&config)
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Config for service \"%s\" not found, use CreateConfig() to create it from scratch", requestedService.Name)
	}

	// Get requested(actual) config, based on the latest timestamp
	srv.db.Raw(SqlGetKeyValsByTimestamp, requestedService.Name, config.UpdatedAt).Scan(&paramsGrpc)

	// Send back config
	configToSendBack.Service = requestedService.Name
	configToSendBack.Parameters = paramsGrpc
	return configToSendBack, status.Errorf(codes.OK,
		"Actual config for service \"%s\" (config timestamp %s) retrieved succesfully", requestedService.Name, config.UpdatedAt.String())

}

// GetArchivedConfig() gRPC handler
func (srv *GrpcDistributedConfigServer) GetArchivedConfig(ctx context.Context, requestedVersion *pb.Timestamp) (*pb.ConfigByTimestamp, error) {
	configToSendBack := &pb.ConfigByTimestamp{}
	var paramsGrpc []*pb.Parameter

	// Check if config with NAME provided by client exists in DB:
	result := srv.db.Limit(1).Where("service = ?", requestedVersion.Service).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"No config for service \"%s\" found, use CreateConfig() to create it from scratch", requestedVersion.Service)
	}

	// Check if config with NAME AND TIMESTAMP provided by client exists in DB:
	result = srv.db.Limit(1).Where("service = ? AND updated_at = ?", requestedVersion.Service, requestedVersion.Timestamp.AsTime()).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Config for service \"%s\" exists, but no version with provided timestamp \"%s\" found. Use ListConfigTimestamps() to list available timestamps.",
			requestedVersion.Service, requestedVersion.Timestamp.AsTime().String())
	}

	srv.db.Raw(SqlGetKeyValsByTimestamp, requestedVersion.Service, requestedVersion.Timestamp.AsTime()).Scan(&paramsGrpc)
	configToSendBack.Service = requestedVersion.Service
	configToSendBack.Timestamp = requestedVersion.Timestamp
	configToSendBack.Parameters = paramsGrpc
	return configToSendBack, status.Errorf(codes.OK,
		"Archival config for service \"%s\" (timestamp %s) retrieved succesfully", requestedVersion.Service, requestedVersion.Timestamp.AsTime().String())
}

// ListConfigTimestamps() gRPC handler
func (srv *GrpcDistributedConfigServer) ListConfigTimestamps(ctx context.Context, requestedService *pb.Service) (*pb.TimestampList, error) {
	grpcTimestampList := &pb.TimestampList{
		Service:    requestedService.Name,
		Timestamps: []*timestamp.Timestamp{},
	}
	var timestamps []time.Time

	result := srv.db.Table("configs").Select("updated_at").Where("service = ? AND deleted_at IS NULL", requestedService.Name).Order("updated_at DESC").Scan(&timestamps)
	// Check if config with NAME provided by client exists in DB:
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"No config for service \"%s\" found, use CreateConfig() to create it from scratch", requestedService.Name)
	}

	for _, tmSt := range timestamps {
		grpcTimestampList.Timestamps = append(grpcTimestampList.Timestamps, timestamppb.New(tmSt))
	}

	return grpcTimestampList, status.Errorf(codes.OK,
		"Timestamp list for config \"%s\" retrieved successfully", grpcTimestampList.Service)

}

// SubscribeClientApp() gRPC handler
func (srv *GrpcDistributedConfigServer) SubscribeClientApp(ctx context.Context, subscriptionRequest *pb.SubscriptionRequest) (*emptypb.Empty, error) {
	result := srv.db.Limit(1).Where("service = ?", subscriptionRequest.Service).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Unable to subscribe. Config for service \"%s\" not found, use CreateConfig() to create it from scratch", subscriptionRequest.Service)
	}

	result = srv.db.Limit(1).Where("service = ? AND client_app = ?", subscriptionRequest.Service, subscriptionRequest.ClientApp).Find(&Subscriber{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.AlreadyExists,
			"Client app \"%s\" is already subscribed to config for service \"%s\" ", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	}

	subscription := &Subscriber{
		Service:   subscriptionRequest.Service,
		ClientApp: subscriptionRequest.ClientApp,
	}
	if result := srv.db.Create(subscription); result.Error != nil {
		log.Println(result.Error)
		return nil, status.Errorf(codes.Internal,
			"Internal error. Unable to subscribe \"%s\" to \"%s\"", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	} else {
		return &emptypb.Empty{}, status.Errorf(codes.OK,
			"Client app \"%s\" successfully subscribed to config for service \"%s\"", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	}

}

// UnSubscribeClientApp() gRPC handler
func (srv *GrpcDistributedConfigServer) UnSubscribeClientApp(ctx context.Context, subscriptionRequest *pb.SubscriptionRequest) (*emptypb.Empty, error) {
	result := srv.db.Limit(1).Where("service = ?", subscriptionRequest.Service).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Unable to UnSubscribe from non-existent config. Config for service \"%s\" not found", subscriptionRequest.Service)
	}

	subscription := &Subscriber{}
	result = srv.db.Limit(1).Where("service = ? AND client_app = ?", subscriptionRequest.Service, subscriptionRequest.ClientApp).Find(subscription)
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Failed to UnSubscribe. Client app \"%s\" is not subscribed to config for service \"%s\" ", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	}

	if result = srv.db.Unscoped().Delete(subscription); result.Error != nil {
		log.Println(result.Error)
		return nil, status.Errorf(codes.Internal,
			"Internal error. Unable to UnSubscribe \"%s\" from \"%s\"", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	} else {
		return &emptypb.Empty{}, status.Errorf(codes.OK,
			"Client app \"%s\" successfully UnSubscribed from config for service \"%s\"", subscriptionRequest.ClientApp, subscriptionRequest.Service)
	}

}

// ListConfigSubscribers() gRPC handler
func (srv *GrpcDistributedConfigServer) ListConfigSubscribers(ctx context.Context, requestedService *pb.Service) (*pb.ConfigSubscribers, error) {
	result := srv.db.Limit(1).Where("service = ?", requestedService.Name).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Config for service \"%s\" not found. Nothing to list", requestedService.Name)
	}

	subscribers := []*Subscriber{}
	result = srv.db.Where("service = ?", requestedService.Name).Find(&subscribers)
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"No client apps subscribed to config for service \"%s\" found", requestedService.Name)
	} else {
		subscribersList := &pb.ConfigSubscribers{Service: requestedService.Name}
		for _, subscriber := range subscribers {
			subscribersList.ClientApp = append(subscribersList.ClientApp, subscriber.ClientApp)
		}
		return subscribersList, status.Errorf(codes.OK,
			"Success retrieving client apps subscribed to config for service \"%s\"", requestedService.Name)
	}

}

// DeleteConfig() gRPC handler
func (srv *GrpcDistributedConfigServer) DeleteConfig(ctx context.Context, requestedService *pb.Service) (*pb.Timestamp, error) {
	result := srv.db.Limit(1).Where("service = ?", requestedService.Name).Find(&Config{})
	if !(result.RowsAffected > 0) {
		return nil, status.Errorf(codes.NotFound,
			"Config for service \"%s\" not found. Nothing to delete", requestedService.Name)
	}

	result = srv.db.Limit(1).Where("service = ?", requestedService.Name).Find(&Subscriber{})
	if result.RowsAffected > 0 {
		return nil, status.Errorf(codes.FailedPrecondition,
			"Deleting of config \"%s\" failed because it has subscribed client app(s). Use ListConfigSubscribers() to check for subscribers", requestedService.Name)
	}

	if result = srv.db.Where("service = ?", requestedService.Name).Delete(&Config{}); result.Error != nil {
		log.Println(result.Error)
		return nil, status.Errorf(codes.Internal,
			"Internal error. Unable to delete config \"%s\"", requestedService.Name)
	} else {
		deletedConfig := &Config{}
		// get deleted_at timestamp:
		srv.db.Unscoped().Limit(1).Where("service = ? AND deleted_at IS NOT NULL", requestedService.Name).Order("deleted_at DESC").Find(deletedConfig)
		return &pb.Timestamp{Service: requestedService.Name, Timestamp: timestamppb.New(deletedConfig.DeletedAt.Time)}, status.Errorf(codes.OK,
			"Config for service \"%s\" deleted succesfully at timestamp %s", requestedService.Name, deletedConfig.DeletedAt.Time.String())
	}

}

func main() {

	envVars := updateDefaultEnvVarsWithActual()

	// resolve db host name to ip address manually
	dbIP := resolveHostNameToIp(envVars["DB_HOST_NAME"])
	//Debuging info:
	log.Printf("DB_HOST_NAME: %s resolved to ip address: %s", envVars["DB_HOST_NAME"], dbIP)
	envVars["DB_HOST_NAME"] = dbIP

	//construct db data source string from environment variables
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		envVars["DB_HOST_NAME"], envVars["DB_PORT"], envVars["DB_USER"], envVars["DB_PASSWORD"], envVars["DB_NAME"], envVars["DB_SSLMODE"], envVars["DB_TIMEZONE"])

	// Connect to DB and migrate the schema
	gormDBconn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	gormDBconn.AutoMigrate(&Config{}, &Parameter{}, Subscriber{})

	// Open listening port
	addr := fmt.Sprintf(":%s", envVars["CFG_SERVICE_PORT"])
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen on port %s", addr)
	}

	// Init and start gRPC service
	srv := grpc.NewServer()
	pb.RegisterDistributedCfgServiceMKServer(srv, &GrpcDistributedConfigServer{db: gormDBconn})
	log.Printf("Starting gRPC distributed config service at port %s ...", addr)
	if err := srv.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
