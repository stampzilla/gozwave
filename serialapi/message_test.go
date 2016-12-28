package serialapi

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func Test_Decode(t *testing.T) {

	raw := []string{
		"01 04 01 13 01 00",

		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 00",
		"01 11 00 04 00 02 0b 71 05 07 ff 00 ff 07 08 01 08 00 00",
		"01 11 00 04 00 02 0b 71 05 07 00 00 ff 07 00 01 08 00 00",

		"01 0c 00 04 00 03 06 31 05 03 0a 00 c2 00",
		"01 0c 00 04 00 03 06 31 05 03 0a 00 1d 00",

		"01 09 00 04 00 02 03 20 01 00 00",
		"01 09 00 04 00 02 03 20 01 ff 00",

		"01 08 00 04 04 03 02 84 07 00",
	}

	for _, val := range raw {
		data := []byte{}

		bytes := strings.Split(val, " ")
		for _, b := range bytes {
			dhex, _ := hex.DecodeString(b)
			data = append(data, dhex...)
		}

		msg := NewMessage(data)
		fmt.Printf("%+v\n\n", msg.Data)
	}

}
