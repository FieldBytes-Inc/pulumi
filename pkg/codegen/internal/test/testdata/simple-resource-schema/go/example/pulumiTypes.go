// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Object struct {
	Bar *string   `pulumi:"bar"`
	Foo *Resource `pulumi:"foo"`
}

// ObjectInput is an input type that accepts ObjectArgs and ObjectOutput values.
// You can construct a concrete instance of `ObjectInput` via:
//
//          ObjectArgs{...}
type ObjectInput interface {
	pulumi.Input

	ToObjectOutput() ObjectOutput
	ToObjectOutputWithContext(context.Context) ObjectOutput
}

type ObjectArgs struct {
	Bar pulumi.StringPtrInput `pulumi:"bar"`
	Foo ResourceInput         `pulumi:"foo"`
}

func (ObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (i ObjectArgs) ToObjectOutput() ObjectOutput {
	return i.ToObjectOutputWithContext(context.Background())
}

func (i ObjectArgs) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectOutput)
}

type ObjectOutput struct{ *pulumi.OutputState }

func (ObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Object)(nil)).Elem()
}

func (o ObjectOutput) ToObjectOutput() ObjectOutput {
	return o
}

func (o ObjectOutput) ToObjectOutputWithContext(ctx context.Context) ObjectOutput {
	return o
}

func (o ObjectOutput) Bar() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Object) *string { return v.Bar }).(pulumi.StringPtrOutput)
}

func (o ObjectOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v Object) *Resource { return v.Foo }).(ResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(ObjectOutput{})
}