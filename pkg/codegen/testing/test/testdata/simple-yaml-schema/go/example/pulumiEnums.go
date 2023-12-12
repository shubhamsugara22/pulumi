// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type OutputOnlyEnumType string

const (
	OutputOnlyEnumTypeFoo = OutputOnlyEnumType("foo")
	OutputOnlyEnumTypeBar = OutputOnlyEnumType("bar")
)

type OutputOnlyEnumTypeOutput struct{ *pulumi.OutputState }

func (OutputOnlyEnumTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypeOutput() OutputOnlyEnumTypeOutput {
	return o
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypeOutputWithContext(ctx context.Context) OutputOnlyEnumTypeOutput {
	return o
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypePtrOutput() OutputOnlyEnumTypePtrOutput {
	return o.ToOutputOnlyEnumTypePtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToOutputOnlyEnumTypePtrOutputWithContext(ctx context.Context) OutputOnlyEnumTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputOnlyEnumType) *OutputOnlyEnumType {
		return &v
	}).(OutputOnlyEnumTypePtrOutput)
}

func (o OutputOnlyEnumTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputOnlyEnumType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutputOnlyEnumTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputOnlyEnumType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputOnlyEnumTypePtrOutput struct{ *pulumi.OutputState }

func (OutputOnlyEnumTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypePtrOutput) ToOutputOnlyEnumTypePtrOutput() OutputOnlyEnumTypePtrOutput {
	return o
}

func (o OutputOnlyEnumTypePtrOutput) ToOutputOnlyEnumTypePtrOutputWithContext(ctx context.Context) OutputOnlyEnumTypePtrOutput {
	return o
}

func (o OutputOnlyEnumTypePtrOutput) Elem() OutputOnlyEnumTypeOutput {
	return o.ApplyT(func(v *OutputOnlyEnumType) OutputOnlyEnumType {
		if v != nil {
			return *v
		}
		var ret OutputOnlyEnumType
		return ret
	}).(OutputOnlyEnumTypeOutput)
}

func (o OutputOnlyEnumTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputOnlyEnumTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputOnlyEnumType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputOnlyEnumTypeMapOutput struct{ *pulumi.OutputState }

func (OutputOnlyEnumTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OutputOnlyEnumType)(nil)).Elem()
}

func (o OutputOnlyEnumTypeMapOutput) ToOutputOnlyEnumTypeMapOutput() OutputOnlyEnumTypeMapOutput {
	return o
}

func (o OutputOnlyEnumTypeMapOutput) ToOutputOnlyEnumTypeMapOutputWithContext(ctx context.Context) OutputOnlyEnumTypeMapOutput {
	return o
}

func (o OutputOnlyEnumTypeMapOutput) MapIndex(k pulumi.StringInput) OutputOnlyEnumTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OutputOnlyEnumType {
		return vs[0].(map[string]OutputOnlyEnumType)[vs[1].(string)]
	}).(OutputOnlyEnumTypeOutput)
}

// types of rubber trees
type RubberTreeVariety string

const (
	// A burgundy rubber tree.
	RubberTreeVarietyBurgundy = RubberTreeVariety("Burgundy")
	// A ruby rubber tree.
	RubberTreeVarietyRuby = RubberTreeVariety("Ruby")
	// A tineke rubber tree.
	RubberTreeVarietyTineke = RubberTreeVariety("Tineke")
)

func (RubberTreeVariety) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTreeVariety)(nil)).Elem()
}

func (e RubberTreeVariety) ToRubberTreeVarietyOutput() RubberTreeVarietyOutput {
	return pulumi.ToOutput(e).(RubberTreeVarietyOutput)
}

func (e RubberTreeVariety) ToRubberTreeVarietyOutputWithContext(ctx context.Context) RubberTreeVarietyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RubberTreeVarietyOutput)
}

func (e RubberTreeVariety) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return e.ToRubberTreeVarietyPtrOutputWithContext(context.Background())
}

func (e RubberTreeVariety) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return RubberTreeVariety(e).ToRubberTreeVarietyOutputWithContext(ctx).ToRubberTreeVarietyPtrOutputWithContext(ctx)
}

func (e RubberTreeVariety) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RubberTreeVariety) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RubberTreeVariety) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RubberTreeVariety) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RubberTreeVarietyOutput struct{ *pulumi.OutputState }

func (RubberTreeVarietyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyOutput() RubberTreeVarietyOutput {
	return o
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyOutputWithContext(ctx context.Context) RubberTreeVarietyOutput {
	return o
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return o.ToRubberTreeVarietyPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RubberTreeVariety) *RubberTreeVariety {
		return &v
	}).(RubberTreeVarietyPtrOutput)
}

func (o RubberTreeVarietyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RubberTreeVariety) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RubberTreeVarietyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RubberTreeVariety) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RubberTreeVarietyPtrOutput struct{ *pulumi.OutputState }

func (RubberTreeVarietyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RubberTreeVariety)(nil)).Elem()
}

func (o RubberTreeVarietyPtrOutput) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return o
}

func (o RubberTreeVarietyPtrOutput) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return o
}

func (o RubberTreeVarietyPtrOutput) Elem() RubberTreeVarietyOutput {
	return o.ApplyT(func(v *RubberTreeVariety) RubberTreeVariety {
		if v != nil {
			return *v
		}
		var ret RubberTreeVariety
		return ret
	}).(RubberTreeVarietyOutput)
}

func (o RubberTreeVarietyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RubberTreeVarietyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RubberTreeVariety) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// RubberTreeVarietyInput is an input type that accepts values of the RubberTreeVariety enum
// A concrete instance of `RubberTreeVarietyInput` can be one of the following:
//
//	RubberTreeVarietyBurgundy
//	RubberTreeVarietyRuby
//	RubberTreeVarietyTineke
type RubberTreeVarietyInput interface {
	pulumi.Input

	ToRubberTreeVarietyOutput() RubberTreeVarietyOutput
	ToRubberTreeVarietyOutputWithContext(context.Context) RubberTreeVarietyOutput
}

var rubberTreeVarietyPtrType = reflect.TypeOf((**RubberTreeVariety)(nil)).Elem()

type RubberTreeVarietyPtrInput interface {
	pulumi.Input

	ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput
	ToRubberTreeVarietyPtrOutputWithContext(context.Context) RubberTreeVarietyPtrOutput
}

type rubberTreeVarietyPtr string

func RubberTreeVarietyPtr(v string) RubberTreeVarietyPtrInput {
	return (*rubberTreeVarietyPtr)(&v)
}

func (*rubberTreeVarietyPtr) ElementType() reflect.Type {
	return rubberTreeVarietyPtrType
}

func (in *rubberTreeVarietyPtr) ToRubberTreeVarietyPtrOutput() RubberTreeVarietyPtrOutput {
	return pulumi.ToOutput(in).(RubberTreeVarietyPtrOutput)
}

func (in *rubberTreeVarietyPtr) ToRubberTreeVarietyPtrOutputWithContext(ctx context.Context) RubberTreeVarietyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RubberTreeVarietyPtrOutput)
}

func (in *rubberTreeVarietyPtr) ToOutput(ctx context.Context) pulumix.Output[*RubberTreeVariety] {
	return pulumix.Output[*RubberTreeVariety]{
		OutputState: in.ToRubberTreeVarietyPtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	pulumi.RegisterInputType(reflect.TypeOf((*RubberTreeVarietyPtrInput)(nil)).Elem(), RubberTreeVariety("Burgundy"))
	pulumi.RegisterOutputType(OutputOnlyEnumTypeOutput{})
	pulumi.RegisterOutputType(OutputOnlyEnumTypePtrOutput{})
	pulumi.RegisterOutputType(OutputOnlyEnumTypeMapOutput{})
	pulumi.RegisterOutputType(RubberTreeVarietyOutput{})
	pulumi.RegisterOutputType(RubberTreeVarietyPtrOutput{})
}
