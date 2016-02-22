// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
	"github.com/shuLhan/tekstus"
)

/*
UpperLowerRatio is a feature that compare uppercase and lowercase characters.
*/
type UpperLowerRatio struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&UpperLowerRatio{}, dsv.TReal, "upper_lower_ratio")
}

/*
GetValues return feature values.
*/
func (ftr *UpperLowerRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute ratio of uppercase and lowercase in new revision.
*/
func (ftr *UpperLowerRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")

	for _, rec := range adds.Records {
		ratio := tekstus.RatioUpperLowerChar(rec.String())

		ftr.PushBack(&dsv.Record{V: Round(ratio)})
	}
}