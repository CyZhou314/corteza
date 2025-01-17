package yaml

import (
	"context"
	"io"
	"regexp"
	"strings"

	"github.com/cyzhou314/corteza/server/pkg/y7s"
	. "github.com/cyzhou314/corteza/server/pkg/y7s"

	"github.com/cyzhou314/corteza/server/pkg/envoy"
	"github.com/cyzhou314/corteza/server/pkg/envoy/resource"
	"gopkg.in/yaml.v3"
)

type (
	// decoder is a wrapper struct for yaml related methods
	decoder struct{}

	EnvoyMarshler interface {
		MarshalEnvoy() ([]resource.Interface, error)
	}
)

var (
	validReference = regexp.MustCompile(`^[A-Za-z1-9][0-9A-Za-z_\-.]*[A-Za-z0-9]$`)
)

func Decoder() *decoder {
	return &decoder{}
}

// CanDecodeFile checks if the file can be handled by this decoder
//
// @todo Add support for this; current library is unable to detect this.
func (y *decoder) CanDecodeFile(f io.Reader) bool {
	// relying on can-decode-ext for now
	return true
}

func (y *decoder) CanDecodeExt(name string) bool {
	var (
		pt  = strings.Split(name, ".")
		ext = strings.TrimSpace(pt[len(pt)-1])
	)
	return ext == "yaml" || ext == "yml"
}

func (y *decoder) Decode(ctx context.Context, r io.Reader, dctx *envoy.DecoderOpts) ([]resource.Interface, error) {
	var (
		doc = &Document{}
	)

	if err := yaml.NewDecoder(r).Decode(doc); err != nil {
		return nil, err
	}

	aaa, err := doc.Decode(ctx)

	return aaa, err

}

// Checks validity of ref node and sets the value to given arg ptr
func decodeRef(n *yaml.Node, refType string, ref *string) error {
	if n == nil {
		return nil
	}

	if !IsKind(n, yaml.ScalarNode) {
		return y7s.NodeErr(n, "%s reference must be scalar", refType)
	}

	if !IsValidRef(n.Value) {
		return y7s.NodeErr(n, "%s reference must be a valid ID or Handle", refType)
	}

	*ref = n.Value
	return nil
}

func IsValidRef(s string) bool {
	return len(s) >= 2 && validReference.MatchString(s)
}
