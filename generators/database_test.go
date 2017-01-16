package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandClass (t *testing.T) {
	c := commandClass{
		ID: "15",
	}

	assert.Equal(t,"0x15", c.IDasHex());
}

func TestDeviceDescription (t *testing.T) {
	d := deviceDescription{
		BrandName: "a\\b\nc\rd-brand",
		ProductName: "a\\b\nc\rd-product", 
		Description: "a\\b\nc\rd-desc",
	}

	assert.Equal(t,"a\\\\b\\nc\\nd-brand", d.EscapedBrandName());
	assert.Equal(t,"a\\\\b\\nc\\nd-product", d.EscapedProductName());
	assert.Equal(t,"a\\\\b\\nc\\nd-desc", d.EscapedDescription());
}

func TestRemoveNewLines(t *testing.T) {
	s := removeNewLines("a\\b\nc\rd")

	assert.Equal(t, "a\\\\b\\nc\\nd", s);
}

func TestValue (t *testing.T) {
	v := value{
		From:"12121212",
		To:"34121212",
		size: 1,
		Desc: []lang{
			lang{Language: "ru", Body:"a\\b\nc\rd-langru"},
			lang{Language: "en", Body:"a\\b\nc\rd-langen"},
		},
	}

	assert.Equal(t, "18", v.FromDec());
	assert.Equal(t, "52", v.ToDec());

	v.size = 2

	assert.Equal(t, "4626", v.FromDec());
	assert.Equal(t, "13330", v.ToDec());

	v.size = 4

	assert.Equal(t, "303174162", v.FromDec());
	assert.Equal(t, "873599506", v.ToDec());

	assert.Equal(t, "a\\\\b\\nc\\nd-langen", v.EscapedDesc());

	v.Desc = []lang{}
	assert.Equal(t, "", v.EscapedDesc());
}

func TestParameter(t *testing.T) {
	p := parameter{
		ID: "4",
		Name: []lang{
			lang{Language: "ru", Body:"a\\b\nc\rd-nameru"},
			lang{Language: "en", Body:"a\\b\nc\rd-nameen"},
		},
		Desc: []lang{
			lang{Language: "ru", Body:"a\\b\nc\rd-descru"},
			lang{Language: "en", Body:"a\\b\nc\rd-descen"},
		},
	}

	assert.Equal(t, "a\\\\b\\nc\\nd-nameen", p.NameCombined());

	p.Name = []lang{}
	assert.Equal(t, "a\\\\b\\nc\\nd-descen", p.NameCombined());

	p.Desc = []lang{}
	assert.Equal(t, "Parameter 4", p.NameCombined());
}
