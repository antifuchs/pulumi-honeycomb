// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package honeycomb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Resource: Trigger
//
// Creates a trigger. For more information about triggers, check out [Alert with Triggers](https://docs.honeycomb.io/working-with-your-data/triggers/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-honeycomb/sdk/go/honeycomb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		dataset := cfg.Require("dataset")
// 		exampleGetQuerySpecification, err := honeycomb.GetQuerySpecification(ctx, &GetQuerySpecificationArgs{
// 			Calculations: []GetQuerySpecificationCalculation{
// 				GetQuerySpecificationCalculation{
// 					Op:     "AVG",
// 					Column: pulumi.StringRef("duration_ms"),
// 				},
// 			},
// 			Filters: []GetQuerySpecificationFilter{
// 				GetQuerySpecificationFilter{
// 					Column: "trace.parent_id",
// 					Op:     "does-not-exist",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleQuery, err := honeycomb.NewQuery(ctx, "exampleQuery", &honeycomb.QueryArgs{
// 			Dataset:   pulumi.String(dataset),
// 			QueryJson: pulumi.String(exampleGetQuerySpecification.Json),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = honeycomb.NewTrigger(ctx, "exampleTrigger", &honeycomb.TriggerArgs{
// 			Description: pulumi.String("Average duration of all requests for the last 10 minutes."),
// 			QueryId:     exampleQuery.ID(),
// 			Dataset:     pulumi.String(dataset),
// 			Frequency:   pulumi.Int(600),
// 			Threshold: &TriggerThresholdArgs{
// 				Op:    pulumi.String(">"),
// 				Value: pulumi.Float64(1000),
// 			},
// 			Recipients: TriggerRecipientArray{
// 				&TriggerRecipientArgs{
// 					Type:   pulumi.String("email"),
// 					Target: pulumi.String("hello@example.com"),
// 				},
// 				&TriggerRecipientArgs{
// 					Type:   pulumi.String("marker"),
// 					Target: pulumi.String("Trigger - requests are slow"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Triggers can be imported using a combination of the dataset name and their ID, e.g.
//
// ```sh
//  $ pulumi import honeycomb:index/trigger:Trigger my_trigger my-dataset/AeZzSoWws9G
// ```
//
//  You can find the ID in the URL bar when visiting the trigger from the UI.
type Trigger struct {
	pulumi.CustomResourceState

	// The dataset this trigger is associated with.
	Dataset pulumi.StringOutput `pulumi:"dataset"`
	// Description of the trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The state of the trigger. If true, the trigger will not be run. Defaults to false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
	Frequency pulumi.IntOutput `pulumi:"frequency"`
	// Name of the trigger.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the Query that the Trigger will execute.
	QueryId pulumi.StringOutput `pulumi:"queryId"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
	Recipients TriggerRecipientArrayOutput `pulumi:"recipients"`
	// A configuration block (described below) describing the threshold of the trigger.
	Threshold TriggerThresholdOutput `pulumi:"threshold"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dataset == nil {
		return nil, errors.New("invalid value for required argument 'Dataset'")
	}
	if args.QueryId == nil {
		return nil, errors.New("invalid value for required argument 'QueryId'")
	}
	if args.Threshold == nil {
		return nil, errors.New("invalid value for required argument 'Threshold'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Trigger
	err := ctx.RegisterResource("honeycomb:index/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("honeycomb:index/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	// The dataset this trigger is associated with.
	Dataset *string `pulumi:"dataset"`
	// Description of the trigger.
	Description *string `pulumi:"description"`
	// The state of the trigger. If true, the trigger will not be run. Defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
	Frequency *int `pulumi:"frequency"`
	// Name of the trigger.
	Name *string `pulumi:"name"`
	// The ID of the Query that the Trigger will execute.
	QueryId *string `pulumi:"queryId"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
	Recipients []TriggerRecipient `pulumi:"recipients"`
	// A configuration block (described below) describing the threshold of the trigger.
	Threshold *TriggerThreshold `pulumi:"threshold"`
}

type TriggerState struct {
	// The dataset this trigger is associated with.
	Dataset pulumi.StringPtrInput
	// Description of the trigger.
	Description pulumi.StringPtrInput
	// The state of the trigger. If true, the trigger will not be run. Defaults to false.
	Disabled pulumi.BoolPtrInput
	// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
	Frequency pulumi.IntPtrInput
	// Name of the trigger.
	Name pulumi.StringPtrInput
	// The ID of the Query that the Trigger will execute.
	QueryId pulumi.StringPtrInput
	// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
	Recipients TriggerRecipientArrayInput
	// A configuration block (described below) describing the threshold of the trigger.
	Threshold TriggerThresholdPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	// The dataset this trigger is associated with.
	Dataset string `pulumi:"dataset"`
	// Description of the trigger.
	Description *string `pulumi:"description"`
	// The state of the trigger. If true, the trigger will not be run. Defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
	Frequency *int `pulumi:"frequency"`
	// Name of the trigger.
	Name *string `pulumi:"name"`
	// The ID of the Query that the Trigger will execute.
	QueryId string `pulumi:"queryId"`
	// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
	Recipients []TriggerRecipient `pulumi:"recipients"`
	// A configuration block (described below) describing the threshold of the trigger.
	Threshold TriggerThreshold `pulumi:"threshold"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	// The dataset this trigger is associated with.
	Dataset pulumi.StringInput
	// Description of the trigger.
	Description pulumi.StringPtrInput
	// The state of the trigger. If true, the trigger will not be run. Defaults to false.
	Disabled pulumi.BoolPtrInput
	// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
	Frequency pulumi.IntPtrInput
	// Name of the trigger.
	Name pulumi.StringPtrInput
	// The ID of the Query that the Trigger will execute.
	QueryId pulumi.StringInput
	// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
	Recipients TriggerRecipientArrayInput
	// A configuration block (described below) describing the threshold of the trigger.
	Threshold TriggerThresholdInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}

type TriggerInput interface {
	pulumi.Input

	ToTriggerOutput() TriggerOutput
	ToTriggerOutputWithContext(ctx context.Context) TriggerOutput
}

func (*Trigger) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (i *Trigger) ToTriggerOutput() TriggerOutput {
	return i.ToTriggerOutputWithContext(context.Background())
}

func (i *Trigger) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerOutput)
}

// TriggerArrayInput is an input type that accepts TriggerArray and TriggerArrayOutput values.
// You can construct a concrete instance of `TriggerArrayInput` via:
//
//          TriggerArray{ TriggerArgs{...} }
type TriggerArrayInput interface {
	pulumi.Input

	ToTriggerArrayOutput() TriggerArrayOutput
	ToTriggerArrayOutputWithContext(context.Context) TriggerArrayOutput
}

type TriggerArray []TriggerInput

func (TriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (i TriggerArray) ToTriggerArrayOutput() TriggerArrayOutput {
	return i.ToTriggerArrayOutputWithContext(context.Background())
}

func (i TriggerArray) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerArrayOutput)
}

// TriggerMapInput is an input type that accepts TriggerMap and TriggerMapOutput values.
// You can construct a concrete instance of `TriggerMapInput` via:
//
//          TriggerMap{ "key": TriggerArgs{...} }
type TriggerMapInput interface {
	pulumi.Input

	ToTriggerMapOutput() TriggerMapOutput
	ToTriggerMapOutputWithContext(context.Context) TriggerMapOutput
}

type TriggerMap map[string]TriggerInput

func (TriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (i TriggerMap) ToTriggerMapOutput() TriggerMapOutput {
	return i.ToTriggerMapOutputWithContext(context.Background())
}

func (i TriggerMap) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TriggerMapOutput)
}

type TriggerOutput struct{ *pulumi.OutputState }

func (TriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trigger)(nil)).Elem()
}

func (o TriggerOutput) ToTriggerOutput() TriggerOutput {
	return o
}

func (o TriggerOutput) ToTriggerOutputWithContext(ctx context.Context) TriggerOutput {
	return o
}

// The dataset this trigger is associated with.
func (o TriggerOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Dataset }).(pulumi.StringOutput)
}

// Description of the trigger.
func (o TriggerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The state of the trigger. If true, the trigger will not be run. Defaults to false.
func (o TriggerOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Trigger) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).
func (o TriggerOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *Trigger) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// Name of the trigger.
func (o TriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the Query that the Trigger will execute.
func (o TriggerOutput) QueryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Trigger) pulumi.StringOutput { return v.QueryId }).(pulumi.StringOutput)
}

// Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires.
func (o TriggerOutput) Recipients() TriggerRecipientArrayOutput {
	return o.ApplyT(func(v *Trigger) TriggerRecipientArrayOutput { return v.Recipients }).(TriggerRecipientArrayOutput)
}

// A configuration block (described below) describing the threshold of the trigger.
func (o TriggerOutput) Threshold() TriggerThresholdOutput {
	return o.ApplyT(func(v *Trigger) TriggerThresholdOutput { return v.Threshold }).(TriggerThresholdOutput)
}

type TriggerArrayOutput struct{ *pulumi.OutputState }

func (TriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trigger)(nil)).Elem()
}

func (o TriggerArrayOutput) ToTriggerArrayOutput() TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) ToTriggerArrayOutputWithContext(ctx context.Context) TriggerArrayOutput {
	return o
}

func (o TriggerArrayOutput) Index(i pulumi.IntInput) TriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].([]*Trigger)[vs[1].(int)]
	}).(TriggerOutput)
}

type TriggerMapOutput struct{ *pulumi.OutputState }

func (TriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trigger)(nil)).Elem()
}

func (o TriggerMapOutput) ToTriggerMapOutput() TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) ToTriggerMapOutputWithContext(ctx context.Context) TriggerMapOutput {
	return o
}

func (o TriggerMapOutput) MapIndex(k pulumi.StringInput) TriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trigger {
		return vs[0].(map[string]*Trigger)[vs[1].(string)]
	}).(TriggerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerInput)(nil)).Elem(), &Trigger{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerArrayInput)(nil)).Elem(), TriggerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TriggerMapInput)(nil)).Elem(), TriggerMap{})
	pulumi.RegisterOutputType(TriggerOutput{})
	pulumi.RegisterOutputType(TriggerArrayOutput{})
	pulumi.RegisterOutputType(TriggerMapOutput{})
}