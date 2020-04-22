// This file is auto-generated. DO NOT EDIT

package jwk

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/lestrrat-go/iter/mapiter"
	"github.com/lestrrat-go/jwx/internal/base64"
	"github.com/lestrrat-go/jwx/internal/iter"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/pkg/errors"
)

const (
	symmetricOctetsKey = "k"
)

type SymmetricKey struct {
	algorithm              *string
	keyID                  *string
	keyType                *jwa.KeyType
	keyUsage               *string
	keyops                 KeyOperationList
	octets                 []byte
	x509CertChain          *CertificateChain
	x509CertThumbprint     *string
	x509CertThumbprintS256 *string
	x509URL                *string
	privateParams          map[string]interface{}
}

type symmetricSymmetricKeyMarshalProxy struct {
	Xalgorithm              *string           `json:"alg,omitempty"`
	XkeyID                  *string           `json:"kid,omitempty"`
	XkeyType                *jwa.KeyType      `json:"kty,omitempty"`
	XkeyUsage               *string           `json:"use,omitempty"`
	Xkeyops                 KeyOperationList  `json:"key_ops,omitempty"`
	Xoctets                 *string           `json:"k,omitempty"`
	Xx509CertChain          *CertificateChain `json:"x5c,omitempty"`
	Xx509CertThumbprint     *string           `json:"x5t,omitempty"`
	Xx509CertThumbprintS256 *string           `json:"x5t#S256,omitempty"`
	Xx509URL                *string           `json:"x5u,omitempty"`
}

func (h *SymmetricKey) Algorithm() string {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return ""
}

func (h *SymmetricKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *SymmetricKey) KeyType() jwa.KeyType {
	if h.keyType != nil {
		return *(h.keyType)
	}
	return jwa.InvalidKeyType
}

func (h *SymmetricKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *SymmetricKey) KeyOps() KeyOperationList {
	return h.keyops
}

func (h *SymmetricKey) Octets() []byte {
	return h.octets
}

func (h *SymmetricKey) X509CertChain() []*x509.Certificate {
	if h.x509CertChain != nil {
		return h.x509CertChain.Get()
	}
	return nil
}

func (h *SymmetricKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *SymmetricKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *SymmetricKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *SymmetricKey) iterate(ctx context.Context, ch chan *HeaderPair) {
	defer close(ch)
	var pairs []*HeaderPair
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.keyType != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyTypeKey, Value: *(h.keyType)})
	}
	if h.keyUsage != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyUsageKey, Value: *(h.keyUsage)})
	}
	if h.keyops != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyOpsKey, Value: h.keyops})
	}
	if h.octets != nil {
		pairs = append(pairs, &HeaderPair{Key: symmetricOctetsKey, Value: h.octets})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: *(h.x509CertChain)})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	for _, pair := range pairs {
		select {
		case <-ctx.Done():
			return
		case ch <- pair:
		}
	}
}

func (h *SymmetricKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *SymmetricKey) Get(name string) (interface{}, bool) {
	switch name {
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case KeyTypeKey:
		if h.keyType == nil {
			return nil, false
		}
		return *(h.keyType), true
	case KeyUsageKey:
		if h.keyUsage == nil {
			return nil, false
		}
		return *(h.keyUsage), true
	case KeyOpsKey:
		if h.keyops == nil {
			return nil, false
		}
		return h.keyops, true
	case symmetricOctetsKey:
		if h.octets == nil {
			return nil, false
		}
		return h.octets, true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return *(h.x509CertChain), true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *SymmetricKey) Set(name string, value interface{}) error {
	switch name {
	case AlgorithmKey:
		switch v := value.(type) {
		case string:
			h.algorithm = &v
		case fmt.Stringer:
			tmp := v.String()
			h.algorithm = &tmp
		default:
			return errors.Errorf(`invalid type for %s key: %T`, AlgorithmKey, value)
		}
		return nil
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case KeyTypeKey:
		var acceptor jwa.KeyType
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, KeyTypeKey)
		}
		h.keyType = &acceptor
		return nil
	case KeyUsageKey:
		if v, ok := value.(string); ok {
			h.keyUsage = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, KeyUsageKey, value)
	case KeyOpsKey:
		var acceptor KeyOperationList
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, KeyOpsKey)
		}
		h.keyops = acceptor
		return nil
	case symmetricOctetsKey:
		if v, ok := value.([]byte); ok {
			h.octets = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, symmetricOctetsKey, value)
	case X509CertChainKey:
		var acceptor CertificateChain
		if err := acceptor.Accept(value); err != nil {
			return errors.Wrapf(err, `invalid value for %s key`, X509CertChainKey)
		}
		h.x509CertChain = &acceptor
		return nil
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (h *SymmetricKey) UnmarshalJSON(buf []byte) error {
	var proxy symmetricSymmetricKeyMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal SymmetricKey`)
	}
	h.algorithm = proxy.Xalgorithm
	h.keyID = proxy.XkeyID
	h.keyType = proxy.XkeyType
	h.keyUsage = proxy.XkeyUsage
	h.keyops = proxy.Xkeyops
	if proxy.Xoctets == nil {
		return errors.New(`required field k is missing`)
	}
	if h.octets = nil; proxy.Xoctets != nil {
		decoded, err := base64.DecodeString(*(proxy.Xoctets))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for octets`)
		}
		h.octets = decoded
	}
	h.x509CertChain = proxy.Xx509CertChain
	h.x509CertThumbprint = proxy.Xx509CertThumbprint
	h.x509CertThumbprintS256 = proxy.Xx509CertThumbprintS256
	h.x509URL = proxy.Xx509URL
	var m map[string]interface{}
	if err := json.Unmarshal(buf, &m); err != nil {
		return errors.Wrap(err, `failed to parse privsate parameters`)
	}
	delete(m, AlgorithmKey)
	delete(m, KeyIDKey)
	delete(m, KeyTypeKey)
	delete(m, KeyUsageKey)
	delete(m, KeyOpsKey)
	delete(m, symmetricOctetsKey)
	delete(m, X509CertChainKey)
	delete(m, X509CertThumbprintKey)
	delete(m, X509CertThumbprintS256Key)
	delete(m, X509URLKey)
	h.privateParams = m
	return nil
}

func (h SymmetricKey) MarshalJSON() ([]byte, error) {
	var proxy symmetricSymmetricKeyMarshalProxy
	proxy.Xalgorithm = h.algorithm
	proxy.XkeyID = h.keyID
	proxy.XkeyType = h.keyType
	if proxy.XkeyType == nil {
		v := jwa.OctetSeq
		proxy.XkeyType = &v
	}
	proxy.XkeyUsage = h.keyUsage
	proxy.Xkeyops = h.keyops
	if len(h.octets) > 0 {
		v := base64.EncodeToStringStd(h.octets)
		proxy.Xoctets = &v
	}
	proxy.Xx509CertChain = h.x509CertChain
	proxy.Xx509CertThumbprint = h.x509CertThumbprint
	proxy.Xx509CertThumbprintS256 = h.x509CertThumbprintS256
	proxy.Xx509URL = h.x509URL
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(proxy); err != nil {
		return nil, errors.Wrap(err, `failed to encode proxy to JSON`)
	}
	hasContent := buf.Len() > 3 // encoding/json always adds a newline, so "{}\n" is the empty hash
	if l := len(h.privateParams); l > 0 {
		buf.Truncate(buf.Len() - 2)
		keys := make([]string, 0, l)
		for k := range h.privateParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if hasContent || i > 0 {
				fmt.Fprintf(&buf, `,`)
			}
			fmt.Fprintf(&buf, `%s:`, strconv.Quote(k))
			if err := enc.Encode(h.privateParams[k]); err != nil {
				return nil, errors.Wrapf(err, `failed to encode private param %s`, k)
			}
		}
		fmt.Fprintf(&buf, `}`)
	}
	return buf.Bytes(), nil
}

func (h *SymmetricKey) Iterate(ctx context.Context) HeaderIterator {
	ch := make(chan *HeaderPair)
	go h.iterate(ctx, ch)
	return mapiter.New(ch)
}

func (h *SymmetricKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *SymmetricKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}
