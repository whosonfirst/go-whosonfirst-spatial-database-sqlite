// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing maintenance window. Only specified parameters are modified.
// The value you specify for Duration determines the specific end time for the
// maintenance window based on the time it begins. No maintenance window tasks are
// permitted to start after the resulting endtime minus the number of hours you
// specify for Cutoff . For example, if the maintenance window starts at 3 PM, the
// duration is three hours, and the value you specify for Cutoff is one hour, no
// maintenance window tasks can start after 5 PM.
func (c *Client) UpdateMaintenanceWindow(ctx context.Context, params *UpdateMaintenanceWindowInput, optFns ...func(*Options)) (*UpdateMaintenanceWindowOutput, error) {
	if params == nil {
		params = &UpdateMaintenanceWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMaintenanceWindow", params, optFns, c.addOperationUpdateMaintenanceWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMaintenanceWindowInput struct {

	// The ID of the maintenance window to update.
	//
	// This member is required.
	WindowId *string

	// Whether targets must be registered with the maintenance window before tasks can
	// be defined for those targets.
	AllowUnassociatedTargets *bool

	// The number of hours before the end of the maintenance window that Amazon Web
	// Services Systems Manager stops scheduling new tasks for execution.
	Cutoff *int32

	// An optional description for the update request.
	Description *string

	// The duration of the maintenance window in hours.
	Duration *int32

	// Whether the maintenance window is enabled.
	Enabled *bool

	// The date and time, in ISO-8601 Extended format, for when you want the
	// maintenance window to become inactive. EndDate allows you to set a date and
	// time in the future when the maintenance window will no longer run.
	EndDate *string

	// The name of the maintenance window.
	Name *string

	// If True , then all fields that are required by the CreateMaintenanceWindow
	// operation are also required for this API request. Optional fields that aren't
	// specified are set to null.
	Replace *bool

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string

	// The number of days to wait after the date and time specified by a cron
	// expression before running the maintenance window. For example, the following
	// cron expression schedules a maintenance window to run the third Tuesday of every
	// month at 11:30 PM. cron(30 23 ? * TUE#3 *) If the schedule offset is 2 , the
	// maintenance window won't run until two days later.
	ScheduleOffset *int32

	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string

	// The date and time, in ISO-8601 Extended format, for when you want the
	// maintenance window to become active. StartDate allows you to delay activation
	// of the maintenance window until the specified future date.
	StartDate *string

	noSmithyDocumentSerde
}

type UpdateMaintenanceWindowOutput struct {

	// Whether targets must be registered with the maintenance window before tasks can
	// be defined for those targets.
	AllowUnassociatedTargets bool

	// The number of hours before the end of the maintenance window that Amazon Web
	// Services Systems Manager stops scheduling new tasks for execution.
	Cutoff int32

	// An optional description of the update.
	Description *string

	// The duration of the maintenance window in hours.
	Duration int32

	// Whether the maintenance window is enabled.
	Enabled bool

	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become inactive. The maintenance window won't run after this
	// specified time.
	EndDate *string

	// The name of the maintenance window.
	Name *string

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string

	// The number of days to wait to run a maintenance window after the scheduled cron
	// expression date and time.
	ScheduleOffset *int32

	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string

	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become active. The maintenance window won't run before this
	// specified time.
	StartDate *string

	// The ID of the created maintenance window.
	WindowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMaintenanceWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateMaintenanceWindow{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateMaintenanceWindowResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateMaintenanceWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMaintenanceWindow(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateMaintenanceWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateMaintenanceWindow",
	}
}

type opUpdateMaintenanceWindowResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opUpdateMaintenanceWindowResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opUpdateMaintenanceWindowResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "ssm"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "ssm"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("ssm")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addUpdateMaintenanceWindowResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opUpdateMaintenanceWindowResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}