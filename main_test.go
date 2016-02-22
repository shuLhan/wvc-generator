// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"github.com/shuLhan/wvcgen"
	"testing"
)

const (
	fInputDsv = "features_test.dsv"
)

func TestAnonim(t *testing.T) {
	main.Generate("anonim", fInputDsv)
}

func TestSizeIncrement(t *testing.T) {
	main.Generate("size_increment", fInputDsv)
}

func TestBiasFrequency(t *testing.T) {
	main.Generate("bias_frequency", fInputDsv)
}

func TestSexWordsFrequency(t *testing.T) {
	main.Generate("sex_words_frequency", fInputDsv)
}

func TestBadWordsFrequency(t *testing.T) {
	main.Generate("bad_words_frequency", fInputDsv)
}
