// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package feature

import (
	// Import any libraries for working with dataset and computing the
	// feature values.
	"github.com/shuLhan/tabula"
)

// Template template to add new feature to this generator.
type Template Feature

func init() {
	Register(&Template{}, tabula.TInteger, "template")
}

/*
Compute describe what this feature do.
*/
func (ftr *Template) Compute(dataset tabula.DatasetInterface) {
	// Get the column from dataset. This is a reference to `InputMetadata`
	// in `features.dsv`.
	// To see the list of column that we can process, see `features.dsv`
	// for an example.
	col := dataset.GetColumnByName("editid")

	for _, rec := range col.Records {
		// This is where the computed value will be saved.
		r := &tabula.Record{}

		// Get the field value from dataset
		s := rec.String()

		// Process the field value `s`, (e.g. cleaning, etc).
		// ...

		// Set the feature value after processing
		e := r.SetValue(s, ftr.GetType())
		if e == nil {
			r.V = 0
		}

		// Save the record value
		ftr.PushBack(r)
	}
}
