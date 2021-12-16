package lke

import (
	"testing"

	"github.com/crossplane/crossplane-runtime/pkg/test"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestDecodeKubeconfig(t *testing.T) {
	cases := map[string]struct {
		in  map[string]interface{}
		out map[string][]byte
		err error
	}{
		"Successful": {
			in: map[string]interface{}{
				"kubeconfig": "aSBhbSBhIGt1YmVjb25maWc=",
			},
			out: map[string][]byte{
				"kubeconfig": []byte("i am a kubeconfig"),
			},
		},
		"Skipped": {
			in: map[string]interface{}{},
		},
		"NotString": {
			in: map[string]interface{}{
				"kubeconfig": 5,
			},
			err: errors.New(errKubeconfigNotString),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			out, err := DecodeKubeconfig(tc.in)
			if diff := cmp.Diff(tc.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nDecodeKubeconfig(...): -want, +got:\n%s", name, diff)
			}
			if diff := cmp.Diff(tc.out, out); diff != "" {
				t.Errorf("\n%s\nDecodeKubeconfig(...): -want, +got:\n%s", name, diff)
			}
		})
	}
}
