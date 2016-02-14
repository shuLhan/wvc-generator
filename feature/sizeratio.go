// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	"github.com/shuLhan/dsv"
)

/*
SizeRatio is a feature that compare the size of new with old revision.
*/
type SizeRatio struct {
	dsv.Column
}

func init() {
	// Register to list of feature
	Register(&SizeRatio{}, dsv.TReal, "sizeratio")
}

/*
GetValues return feature values.
*/
func (ftr *SizeRatio) GetValues() dsv.Column {
	return ftr.Column
}

/*
Compute ratio of length of characters in new with old revision.
*/
func (ftr *SizeRatio) Compute(dataset dsv.Dataset) {
	adds := dataset.GetColumnByName("additions")
	dels := dataset.GetColumnByName("deletions")

	delslen := dels.Len()

	for x, rec := range adds.Records {
		if x >= delslen {
			// Just in case additions is greater than deletions
			break
		}

		r := &dsv.Record{}

		newlen := len(rec.String())
		oldlen := len(dels.Records[x].String())
		ratio := float64(1+newlen) / float64(1+oldlen)

		// round it to five digits after comma.
		r.SetFloat(float64(int(ratio*100000)) / 100000)

		ftr.PushBack(r)
	}
}
