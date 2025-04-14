package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"sort"
)

var result = map[string]attr.Type{
	"number_of_bins": types.Int64Type,
	"bins":           types.ListType{ElemType: types.ListType{ElemType: types.StringType}},
}

type Item struct {
	Value  string
	Weight float64
}

type Bin struct {
	Capacity  float64
	Remaining float64
	Items     []Item
}

type BinPackingFunction struct{}

func NewBinPackingFunction() function.Function {
	return &BinPackingFunction{}
}

func (f *BinPackingFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "bin_packing_string"
}

func (f *BinPackingFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "Parse an RFC3339 timestamp string into an object",
		Description: "Given an RFC3339 timestamp string, will parse and return an object representation of that date and time.",

		Parameters: []function.Parameter{
			function.ListParameter{
				ElementType: types.StringType,
				Name:        "values",
			},
			function.Int64Parameter{
				AllowNullValue:     false,
				AllowUnknownValues: false,
				CustomType:         nil,
				Name:               "bin_size",
				Validators:         nil,
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: result,
		},
	}
}

func (f *BinPackingFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var strings []string
	var binSize int64

	resp.Error = req.Arguments.Get(ctx, &strings, &binSize)
	if resp.Error != nil {
		return
	}

	items := make([]Item, len(strings))

	for i := range len(strings) {
		items[i] = Item{
			Value:  strings[i],
			Weight: float64(len(strings[i])),
		}
	}

	bins := FirstFitDecreasing(items, float64(binSize))

	var outerListElements []attr.Value
	for _, bin := range bins {
		l := make([]string, len(bin.Items))
		for i, item := range bin.Items {
			l[i] = item.Value
		}
		innerList, _ := types.ListValueFrom(ctx, types.StringType, l)
		outerListElements = append(outerListElements, innerList)
	}

	res, _ := types.ListValue(
		types.ListType{ElemType: types.StringType}, // Outer list's element type
		outerListElements,
	)

	response, diags := types.ObjectValue(
		result,
		map[string]attr.Value{
			"number_of_bins": types.Int64Value(int64(len(bins))),
			"bins":           res,
		},
	)

	resp.Error = function.FuncErrorFromDiags(ctx, diags)
	if resp.Error != nil {
		return
	}

	resp.Error = resp.Result.Set(ctx, &response)
}

func FirstFitDecreasing(items []Item, binCapacity float64) []Bin {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Weight > items[j].Weight
	})

	var bins []Bin

	for _, item := range items {
		placed := false
		// Try placing in existing bins
		for i := range bins {
			if bins[i].Remaining >= item.Weight {
				bins[i].Items = append(bins[i].Items, item)
				bins[i].Remaining -= item.Weight
				placed = true
				break
			}
		}
		// Create new bin if no placement found
		if !placed {
			bins = append(bins, Bin{
				Capacity:  binCapacity,
				Remaining: binCapacity - item.Weight,
				Items:     []Item{item},
			})
		}
	}
	return bins
}
