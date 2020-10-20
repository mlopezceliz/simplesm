package simplesm

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

const (
	mapper       = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"
	renderServer = "https://plantuml.arch-services.net/svg/"

	hSixtyThree = 0x3f
	hFifteen    = 0xF
	hThree      = 0x3
)

func Diagram(rootNode *Node) string {
	output := "@startuml"

	output += "\n[*] --> " + rootNode.ID
	_, output = rootNode.draw([]*Node{}, output)

	output += "\n@enduml\n"

	return output
}

func Link(rootNode *Node) string {
	return fmt.Sprint(renderServer, Encoded(Diagram(rootNode)))
}

func Encoded(diagram string) string {
	raw := []byte(diagram)
	compressed := deflate(raw)

	return base64Encode(compressed)
}

func deflate(input []byte) []byte {
	var b bytes.Buffer

	w, _ := zlib.NewWriterLevel(&b, zlib.BestCompression)
	_, _ = w.Write(input)
	_ = w.Close()

	return b.Bytes()
}

func base64Encode(input []byte) string {
	var buffer bytes.Buffer

	inputLength := len(input)

	for i := 0; i < 3-inputLength%3; i++ {
		input = append(input, byte(0))
	}

	for i := 0; i < inputLength; i += 3 {
		b1, b2, b3, b4 := input[i], input[i+1], input[i+2], byte(0)

		b4 = b3 & hSixtyThree
		b3 = ((b2 & hFifteen) << 2) | (b3 >> 6)
		b2 = ((b1 & hThree) << 4) | (b2 >> 4)
		b1 >>= 2

		for _, b := range []byte{b1, b2, b3, b4} {
			buffer.WriteByte(mapper[b])
		}
	}

	return buffer.String()
}
