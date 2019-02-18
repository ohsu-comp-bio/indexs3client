package endpoints

// Service identifiers
//
// Deprecated: Use client package's EndpointID value instead of these
// ServiceIDs. These IDs are not maintained, and are out of date.
const (
	A4bServiceID                          = "a4b"                          // A4b.
	AcmServiceID                          = "acm"                          // Acm.
	AcmPcaServiceID                       = "acm-pca"                      // AcmPca.
	ApiMediatailorServiceID               = "api.mediatailor"              // ApiMediatailor.
	ApiPricingServiceID                   = "api.pricing"                  // ApiPricing.
	ApiSagemakerServiceID                 = "api.sagemaker"                // ApiSagemaker.
	ApigatewayServiceID                   = "apigateway"                   // Apigateway.
	ApplicationAutoscalingServiceID       = "application-autoscaling"      // ApplicationAutoscaling.
	Appstream2ServiceID                   = "appstream2"                   // Appstream2.
	AppsyncServiceID                      = "appsync"                      // Appsync.
	AthenaServiceID                       = "athena"                       // Athena.
	AutoscalingServiceID                  = "autoscaling"                  // Autoscaling.
	AutoscalingPlansServiceID             = "autoscaling-plans"            // AutoscalingPlans.
	BatchServiceID                        = "batch"                        // Batch.
	BudgetsServiceID                      = "budgets"                      // Budgets.
	CeServiceID                           = "ce"                           // Ce.
	ChimeServiceID                        = "chime"                        // Chime.
	Cloud9ServiceID                       = "cloud9"                       // Cloud9.
	ClouddirectoryServiceID               = "clouddirectory"               // Clouddirectory.
	CloudformationServiceID               = "cloudformation"               // Cloudformation.
	CloudfrontServiceID                   = "cloudfront"                   // Cloudfront.
	CloudhsmServiceID                     = "cloudhsm"                     // Cloudhsm.
	Cloudhsmv2ServiceID                   = "cloudhsmv2"                   // Cloudhsmv2.
	CloudsearchServiceID                  = "cloudsearch"                  // Cloudsearch.
	CloudtrailServiceID                   = "cloudtrail"                   // Cloudtrail.
	CodebuildServiceID                    = "codebuild"                    // Codebuild.
	CodecommitServiceID                   = "codecommit"                   // Codecommit.
	CodedeployServiceID                   = "codedeploy"                   // Codedeploy.
	CodepipelineServiceID                 = "codepipeline"                 // Codepipeline.
	CodestarServiceID                     = "codestar"                     // Codestar.
	CognitoIdentityServiceID              = "cognito-identity"             // CognitoIdentity.
	CognitoIdpServiceID                   = "cognito-idp"                  // CognitoIdp.
	CognitoSyncServiceID                  = "cognito-sync"                 // CognitoSync.
	ComprehendServiceID                   = "comprehend"                   // Comprehend.
	ConfigServiceID                       = "config"                       // Config.
	CurServiceID                          = "cur"                          // Cur.
	DatapipelineServiceID                 = "datapipeline"                 // Datapipeline.
	DaxServiceID                          = "dax"                          // Dax.
	DevicefarmServiceID                   = "devicefarm"                   // Devicefarm.
	DirectconnectServiceID                = "directconnect"                // Directconnect.
	DiscoveryServiceID                    = "discovery"                    // Discovery.
	DmsServiceID                          = "dms"                          // Dms.
	DsServiceID                           = "ds"                           // Ds.
	DynamodbServiceID                     = "dynamodb"                     // Dynamodb.
	Ec2ServiceID                          = "ec2"                          // Ec2.
	Ec2metadataServiceID                  = "ec2metadata"                  // Ec2metadata.
	EcrServiceID                          = "ecr"                          // Ecr.
	EcsServiceID                          = "ecs"                          // Ecs.
	ElasticacheServiceID                  = "elasticache"                  // Elasticache.
	ElasticbeanstalkServiceID             = "elasticbeanstalk"             // Elasticbeanstalk.
	ElasticfilesystemServiceID            = "elasticfilesystem"            // Elasticfilesystem.
	ElasticloadbalancingServiceID         = "elasticloadbalancing"         // Elasticloadbalancing.
	ElasticmapreduceServiceID             = "elasticmapreduce"             // Elasticmapreduce.
	ElastictranscoderServiceID            = "elastictranscoder"            // Elastictranscoder.
	EmailServiceID                        = "email"                        // Email.
	EntitlementMarketplaceServiceID       = "entitlement.marketplace"      // EntitlementMarketplace.
	EsServiceID                           = "es"                           // Es.
	EventsServiceID                       = "events"                       // Events.
	FirehoseServiceID                     = "firehose"                     // Firehose.
	FmsServiceID                          = "fms"                          // Fms.
	GameliftServiceID                     = "gamelift"                     // Gamelift.
	GlacierServiceID                      = "glacier"                      // Glacier.
	GlueServiceID                         = "glue"                         // Glue.
	GreengrassServiceID                   = "greengrass"                   // Greengrass.
	GuarddutyServiceID                    = "guardduty"                    // Guardduty.
	HealthServiceID                       = "health"                       // Health.
	IamServiceID                          = "iam"                          // Iam.
	ImportexportServiceID                 = "importexport"                 // Importexport.
	InspectorServiceID                    = "inspector"                    // Inspector.
	IotServiceID                          = "iot"                          // Iot.
	IotanalyticsServiceID                 = "iotanalytics"                 // Iotanalytics.
	KinesisServiceID                      = "kinesis"                      // Kinesis.
	KinesisanalyticsServiceID             = "kinesisanalytics"             // Kinesisanalytics.
	KinesisvideoServiceID                 = "kinesisvideo"                 // Kinesisvideo.
	KmsServiceID                          = "kms"                          // Kms.
	LambdaServiceID                       = "lambda"                       // Lambda.
	LightsailServiceID                    = "lightsail"                    // Lightsail.
	LogsServiceID                         = "logs"                         // Logs.
	MachinelearningServiceID              = "machinelearning"              // Machinelearning.
	MarketplacecommerceanalyticsServiceID = "marketplacecommerceanalytics" // Marketplacecommerceanalytics.
	MediaconvertServiceID                 = "mediaconvert"                 // Mediaconvert.
	MedialiveServiceID                    = "medialive"                    // Medialive.
	MediapackageServiceID                 = "mediapackage"                 // Mediapackage.
	MediastoreServiceID                   = "mediastore"                   // Mediastore.
	MeteringMarketplaceServiceID          = "metering.marketplace"         // MeteringMarketplace.
	MghServiceID                          = "mgh"                          // Mgh.
	MobileanalyticsServiceID              = "mobileanalytics"              // Mobileanalytics.
	ModelsLexServiceID                    = "models.lex"                   // ModelsLex.
	MonitoringServiceID                   = "monitoring"                   // Monitoring.
	MturkRequesterServiceID               = "mturk-requester"              // MturkRequester.
	NeptuneServiceID                      = "neptune"                      // Neptune.
	OpsworksServiceID                     = "opsworks"                     // Opsworks.
	OpsworksCmServiceID                   = "opsworks-cm"                  // OpsworksCm.
	OrganizationsServiceID                = "organizations"                // Organizations.
	PinpointServiceID                     = "pinpoint"                     // Pinpoint.
	PollyServiceID                        = "polly"                        // Polly.
	RdsServiceID                          = "rds"                          // Rds.
	RedshiftServiceID                     = "redshift"                     // Redshift.
	RekognitionServiceID                  = "rekognition"                  // Rekognition.
	ResourceGroupsServiceID               = "resource-groups"              // ResourceGroups.
	Route53ServiceID                      = "route53"                      // Route53.
	Route53domainsServiceID               = "route53domains"               // Route53domains.
	RuntimeLexServiceID                   = "runtime.lex"                  // RuntimeLex.
	RuntimeSagemakerServiceID             = "runtime.sagemaker"            // RuntimeSagemaker.
	S3ServiceID                           = "s3"                           // S3.
	S3ControlServiceID                    = "s3-control"                   // S3Control.
	SagemakerServiceID                    = "api.sagemaker"                // Sagemaker.
	SdbServiceID                          = "sdb"                          // Sdb.
	SecretsmanagerServiceID               = "secretsmanager"               // Secretsmanager.
	ServerlessrepoServiceID               = "serverlessrepo"               // Serverlessrepo.
	ServicecatalogServiceID               = "servicecatalog"               // Servicecatalog.
	ServicediscoveryServiceID             = "servicediscovery"             // Servicediscovery.
	ShieldServiceID                       = "shield"                       // Shield.
	SmsServiceID                          = "sms"                          // Sms.
	SnowballServiceID                     = "snowball"                     // Snowball.
	SnsServiceID                          = "sns"                          // Sns.
	SqsServiceID                          = "sqs"                          // Sqs.
	SsmServiceID                          = "ssm"                          // Ssm.
	StatesServiceID                       = "states"                       // States.
	StoragegatewayServiceID               = "storagegateway"               // Storagegateway.
	StreamsDynamodbServiceID              = "streams.dynamodb"             // StreamsDynamodb.
	StsServiceID                          = "sts"                          // Sts.
	SupportServiceID                      = "support"                      // Support.
	SwfServiceID                          = "swf"                          // Swf.
	TaggingServiceID                      = "tagging"                      // Tagging.
	TransferServiceID                     = "transfer"                     // Transfer.
	TranslateServiceID                    = "translate"                    // Translate.
	WafServiceID                          = "waf"                          // Waf.
	WafRegionalServiceID                  = "waf-regional"                 // WafRegional.
	WorkdocsServiceID                     = "workdocs"                     // Workdocs.
	WorkmailServiceID                     = "workmail"                     // Workmail.
	WorkspacesServiceID                   = "workspaces"                   // Workspaces.
	XrayServiceID                         = "xray"                         // Xray.
)
