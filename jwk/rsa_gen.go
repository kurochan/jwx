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
	rsaDKey  = "d"
	rsaDPKey = "dp"
	rsaDQKey = "dq"
	rsaEKey  = "e"
	rsaNKey  = "n"
	rsaPKey  = "p"
	rsaQKey  = "q"
	rsaQIKey = "qi"
)

type RSAPrivateKey struct {
	algorithm              *string
	d                      []byte
	dp                     []byte
	dq                     []byte
	e                      []byte
	keyID                  *string
	keyType                *jwa.KeyType
	keyUsage               *string
	keyops                 KeyOperationList
	n                      []byte
	p                      []byte
	q                      []byte
	qi                     []byte
	x509CertChain          *CertificateChain
	x509CertThumbprint     *string
	x509CertThumbprintS256 *string
	x509URL                *string
	privateParams          map[string]interface{}
}

type rsaPrivateKeyMarshalProxy struct {
	Xalgorithm              *string           `json:"alg,omitempty"`
	Xd                      *string           `json:"d,omitempty"`
	Xdp                     *string           `json:"dp,omitempty"`
	Xdq                     *string           `json:"dq,omitempty"`
	Xe                      *string           `json:"e,omitempty"`
	XkeyID                  *string           `json:"kid,omitempty"`
	XkeyType                *jwa.KeyType      `json:"kty,omitempty"`
	XkeyUsage               *string           `json:"use,omitempty"`
	Xkeyops                 KeyOperationList  `json:"key_ops,omitempty"`
	Xn                      *string           `json:"n,omitempty"`
	Xp                      *string           `json:"p,omitempty"`
	Xq                      *string           `json:"q,omitempty"`
	Xqi                     *string           `json:"qi,omitempty"`
	Xx509CertChain          *CertificateChain `json:"x5c,omitempty"`
	Xx509CertThumbprint     *string           `json:"x5t,omitempty"`
	Xx509CertThumbprintS256 *string           `json:"x5t#S256,omitempty"`
	Xx509URL                *string           `json:"x5u,omitempty"`
}

func (h *RSAPrivateKey) Algorithm() string {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return ""
}

func (h *RSAPrivateKey) D() []byte {
	return h.d
}

func (h *RSAPrivateKey) DP() []byte {
	return h.dp
}

func (h *RSAPrivateKey) DQ() []byte {
	return h.dq
}

func (h *RSAPrivateKey) E() []byte {
	return h.e
}

func (h *RSAPrivateKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *RSAPrivateKey) KeyType() jwa.KeyType {
	if h.keyType != nil {
		return *(h.keyType)
	}
	return jwa.InvalidKeyType
}

func (h *RSAPrivateKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *RSAPrivateKey) KeyOps() KeyOperationList {
	return h.keyops
}

func (h *RSAPrivateKey) N() []byte {
	return h.n
}

func (h *RSAPrivateKey) P() []byte {
	return h.p
}

func (h *RSAPrivateKey) Q() []byte {
	return h.q
}

func (h *RSAPrivateKey) QI() []byte {
	return h.qi
}

func (h *RSAPrivateKey) X509CertChain() []*x509.Certificate {
	if h.x509CertChain != nil {
		return h.x509CertChain.Get()
	}
	return nil
}

func (h *RSAPrivateKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *RSAPrivateKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *RSAPrivateKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *RSAPrivateKey) iterate(ctx context.Context, ch chan *HeaderPair) {
	defer close(ch)
	var pairs []*HeaderPair
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.d != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaDKey, Value: h.d})
	}
	if h.dp != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaDPKey, Value: h.dp})
	}
	if h.dq != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaDQKey, Value: h.dq})
	}
	if h.e != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaEKey, Value: h.e})
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
	if h.n != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaNKey, Value: h.n})
	}
	if h.p != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaPKey, Value: h.p})
	}
	if h.q != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaQKey, Value: h.q})
	}
	if h.qi != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaQIKey, Value: h.qi})
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

func (h *RSAPrivateKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *RSAPrivateKey) Get(name string) (interface{}, bool) {
	switch name {
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case rsaDKey:
		if h.d == nil {
			return nil, false
		}
		return h.d, true
	case rsaDPKey:
		if h.dp == nil {
			return nil, false
		}
		return h.dp, true
	case rsaDQKey:
		if h.dq == nil {
			return nil, false
		}
		return h.dq, true
	case rsaEKey:
		if h.e == nil {
			return nil, false
		}
		return h.e, true
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
	case rsaNKey:
		if h.n == nil {
			return nil, false
		}
		return h.n, true
	case rsaPKey:
		if h.p == nil {
			return nil, false
		}
		return h.p, true
	case rsaQKey:
		if h.q == nil {
			return nil, false
		}
		return h.q, true
	case rsaQIKey:
		if h.qi == nil {
			return nil, false
		}
		return h.qi, true
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

func (h *RSAPrivateKey) Set(name string, value interface{}) error {
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
	case rsaDKey:
		if v, ok := value.([]byte); ok {
			h.d = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaDKey, value)
	case rsaDPKey:
		if v, ok := value.([]byte); ok {
			h.dp = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaDPKey, value)
	case rsaDQKey:
		if v, ok := value.([]byte); ok {
			h.dq = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaDQKey, value)
	case rsaEKey:
		if v, ok := value.([]byte); ok {
			h.e = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaEKey, value)
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
	case rsaNKey:
		if v, ok := value.([]byte); ok {
			h.n = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaNKey, value)
	case rsaPKey:
		if v, ok := value.([]byte); ok {
			h.p = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaPKey, value)
	case rsaQKey:
		if v, ok := value.([]byte); ok {
			h.q = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaQKey, value)
	case rsaQIKey:
		if v, ok := value.([]byte); ok {
			h.qi = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaQIKey, value)
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

func (h *RSAPrivateKey) UnmarshalJSON(buf []byte) error {
	var proxy rsaPrivateKeyMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal RSAPrivateKey`)
	}
	h.algorithm = proxy.Xalgorithm
	if proxy.Xd == nil {
		return errors.New(`required field d is missing`)
	}
	if h.d = nil; proxy.Xd != nil {
		decoded, err := base64.DecodeString(*(proxy.Xd))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for d`)
		}
		h.d = decoded
	}
	if h.dp = nil; proxy.Xdp != nil {
		decoded, err := base64.DecodeString(*(proxy.Xdp))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for dp`)
		}
		h.dp = decoded
	}
	if h.dq = nil; proxy.Xdq != nil {
		decoded, err := base64.DecodeString(*(proxy.Xdq))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for dq`)
		}
		h.dq = decoded
	}
	if proxy.Xe == nil {
		return errors.New(`required field e is missing`)
	}
	if h.e = nil; proxy.Xe != nil {
		decoded, err := base64.DecodeString(*(proxy.Xe))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for e`)
		}
		h.e = decoded
	}
	h.keyID = proxy.XkeyID
	h.keyType = proxy.XkeyType
	h.keyUsage = proxy.XkeyUsage
	h.keyops = proxy.Xkeyops
	if proxy.Xn == nil {
		return errors.New(`required field n is missing`)
	}
	if h.n = nil; proxy.Xn != nil {
		decoded, err := base64.DecodeString(*(proxy.Xn))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for n`)
		}
		h.n = decoded
	}
	if proxy.Xp == nil {
		return errors.New(`required field p is missing`)
	}
	if h.p = nil; proxy.Xp != nil {
		decoded, err := base64.DecodeString(*(proxy.Xp))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for p`)
		}
		h.p = decoded
	}
	if proxy.Xq == nil {
		return errors.New(`required field q is missing`)
	}
	if h.q = nil; proxy.Xq != nil {
		decoded, err := base64.DecodeString(*(proxy.Xq))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for q`)
		}
		h.q = decoded
	}
	if h.qi = nil; proxy.Xqi != nil {
		decoded, err := base64.DecodeString(*(proxy.Xqi))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for qi`)
		}
		h.qi = decoded
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
	delete(m, rsaDKey)
	delete(m, rsaDPKey)
	delete(m, rsaDQKey)
	delete(m, rsaEKey)
	delete(m, KeyIDKey)
	delete(m, KeyTypeKey)
	delete(m, KeyUsageKey)
	delete(m, KeyOpsKey)
	delete(m, rsaNKey)
	delete(m, rsaPKey)
	delete(m, rsaQKey)
	delete(m, rsaQIKey)
	delete(m, X509CertChainKey)
	delete(m, X509CertThumbprintKey)
	delete(m, X509CertThumbprintS256Key)
	delete(m, X509URLKey)
	h.privateParams = m
	return nil
}

func (h RSAPrivateKey) MarshalJSON() ([]byte, error) {
	var proxy rsaPrivateKeyMarshalProxy
	proxy.Xalgorithm = h.algorithm
	if len(h.d) > 0 {
		v := base64.EncodeToStringStd(h.d)
		proxy.Xd = &v
	}
	if len(h.dp) > 0 {
		v := base64.EncodeToStringStd(h.dp)
		proxy.Xdp = &v
	}
	if len(h.dq) > 0 {
		v := base64.EncodeToStringStd(h.dq)
		proxy.Xdq = &v
	}
	if len(h.e) > 0 {
		v := base64.EncodeToStringStd(h.e)
		proxy.Xe = &v
	}
	proxy.XkeyID = h.keyID
	proxy.XkeyType = h.keyType
	if proxy.XkeyType == nil {
		v := jwa.RSA
		proxy.XkeyType = &v
	}
	proxy.XkeyUsage = h.keyUsage
	proxy.Xkeyops = h.keyops
	if len(h.n) > 0 {
		v := base64.EncodeToStringStd(h.n)
		proxy.Xn = &v
	}
	if len(h.p) > 0 {
		v := base64.EncodeToStringStd(h.p)
		proxy.Xp = &v
	}
	if len(h.q) > 0 {
		v := base64.EncodeToStringStd(h.q)
		proxy.Xq = &v
	}
	if len(h.qi) > 0 {
		v := base64.EncodeToStringStd(h.qi)
		proxy.Xqi = &v
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

func (h *RSAPrivateKey) Iterate(ctx context.Context) HeaderIterator {
	ch := make(chan *HeaderPair)
	go h.iterate(ctx, ch)
	return mapiter.New(ch)
}

func (h *RSAPrivateKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *RSAPrivateKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}

type RSAPublicKey struct {
	algorithm              *string
	e                      []byte
	keyID                  *string
	keyType                *jwa.KeyType
	keyUsage               *string
	keyops                 KeyOperationList
	n                      []byte
	x509CertChain          *CertificateChain
	x509CertThumbprint     *string
	x509CertThumbprintS256 *string
	x509URL                *string
	privateParams          map[string]interface{}
}

type rsaPublicKeyMarshalProxy struct {
	Xalgorithm              *string           `json:"alg,omitempty"`
	Xe                      *string           `json:"e,omitempty"`
	XkeyID                  *string           `json:"kid,omitempty"`
	XkeyType                *jwa.KeyType      `json:"kty,omitempty"`
	XkeyUsage               *string           `json:"use,omitempty"`
	Xkeyops                 KeyOperationList  `json:"key_ops,omitempty"`
	Xn                      *string           `json:"n,omitempty"`
	Xx509CertChain          *CertificateChain `json:"x5c,omitempty"`
	Xx509CertThumbprint     *string           `json:"x5t,omitempty"`
	Xx509CertThumbprintS256 *string           `json:"x5t#S256,omitempty"`
	Xx509URL                *string           `json:"x5u,omitempty"`
}

func (h *RSAPublicKey) Algorithm() string {
	if h.algorithm != nil {
		return *(h.algorithm)
	}
	return ""
}

func (h *RSAPublicKey) E() []byte {
	return h.e
}

func (h *RSAPublicKey) KeyID() string {
	if h.keyID != nil {
		return *(h.keyID)
	}
	return ""
}

func (h *RSAPublicKey) KeyType() jwa.KeyType {
	if h.keyType != nil {
		return *(h.keyType)
	}
	return jwa.InvalidKeyType
}

func (h *RSAPublicKey) KeyUsage() string {
	if h.keyUsage != nil {
		return *(h.keyUsage)
	}
	return ""
}

func (h *RSAPublicKey) KeyOps() KeyOperationList {
	return h.keyops
}

func (h *RSAPublicKey) N() []byte {
	return h.n
}

func (h *RSAPublicKey) X509CertChain() []*x509.Certificate {
	if h.x509CertChain != nil {
		return h.x509CertChain.Get()
	}
	return nil
}

func (h *RSAPublicKey) X509CertThumbprint() string {
	if h.x509CertThumbprint != nil {
		return *(h.x509CertThumbprint)
	}
	return ""
}

func (h *RSAPublicKey) X509CertThumbprintS256() string {
	if h.x509CertThumbprintS256 != nil {
		return *(h.x509CertThumbprintS256)
	}
	return ""
}

func (h *RSAPublicKey) X509URL() string {
	if h.x509URL != nil {
		return *(h.x509URL)
	}
	return ""
}

func (h *RSAPublicKey) iterate(ctx context.Context, ch chan *HeaderPair) {
	defer close(ch)
	var pairs []*HeaderPair
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.e != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaEKey, Value: h.e})
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
	if h.n != nil {
		pairs = append(pairs, &HeaderPair{Key: rsaNKey, Value: h.n})
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

func (h *RSAPublicKey) PrivateParams() map[string]interface{} {
	return h.privateParams
}

func (h *RSAPublicKey) Get(name string) (interface{}, bool) {
	switch name {
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case rsaEKey:
		if h.e == nil {
			return nil, false
		}
		return h.e, true
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
	case rsaNKey:
		if h.n == nil {
			return nil, false
		}
		return h.n, true
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

func (h *RSAPublicKey) Set(name string, value interface{}) error {
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
	case rsaEKey:
		if v, ok := value.([]byte); ok {
			h.e = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaEKey, value)
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
	case rsaNKey:
		if v, ok := value.([]byte); ok {
			h.n = v
			return nil
		}
		return errors.Errorf(`invalid value for %s key: %T`, rsaNKey, value)
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

func (h *RSAPublicKey) UnmarshalJSON(buf []byte) error {
	var proxy rsaPublicKeyMarshalProxy
	if err := json.Unmarshal(buf, &proxy); err != nil {
		return errors.Wrap(err, `failed to unmarshal RSAPublicKey`)
	}
	h.algorithm = proxy.Xalgorithm
	if proxy.Xe == nil {
		return errors.New(`required field e is missing`)
	}
	if h.e = nil; proxy.Xe != nil {
		decoded, err := base64.DecodeString(*(proxy.Xe))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for e`)
		}
		h.e = decoded
	}
	h.keyID = proxy.XkeyID
	h.keyType = proxy.XkeyType
	h.keyUsage = proxy.XkeyUsage
	h.keyops = proxy.Xkeyops
	if proxy.Xn == nil {
		return errors.New(`required field n is missing`)
	}
	if h.n = nil; proxy.Xn != nil {
		decoded, err := base64.DecodeString(*(proxy.Xn))
		if err != nil {
			return errors.Wrap(err, `failed to decode base64 value for n`)
		}
		h.n = decoded
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
	delete(m, rsaEKey)
	delete(m, KeyIDKey)
	delete(m, KeyTypeKey)
	delete(m, KeyUsageKey)
	delete(m, KeyOpsKey)
	delete(m, rsaNKey)
	delete(m, X509CertChainKey)
	delete(m, X509CertThumbprintKey)
	delete(m, X509CertThumbprintS256Key)
	delete(m, X509URLKey)
	h.privateParams = m
	return nil
}

func (h RSAPublicKey) MarshalJSON() ([]byte, error) {
	var proxy rsaPublicKeyMarshalProxy
	proxy.Xalgorithm = h.algorithm
	if len(h.e) > 0 {
		v := base64.EncodeToStringStd(h.e)
		proxy.Xe = &v
	}
	proxy.XkeyID = h.keyID
	proxy.XkeyType = h.keyType
	if proxy.XkeyType == nil {
		v := jwa.RSA
		proxy.XkeyType = &v
	}
	proxy.XkeyUsage = h.keyUsage
	proxy.Xkeyops = h.keyops
	if len(h.n) > 0 {
		v := base64.EncodeToStringStd(h.n)
		proxy.Xn = &v
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

func (h *RSAPublicKey) Iterate(ctx context.Context) HeaderIterator {
	ch := make(chan *HeaderPair)
	go h.iterate(ctx, ch)
	return mapiter.New(ch)
}

func (h *RSAPublicKey) Walk(ctx context.Context, visitor HeaderVisitor) error {
	return iter.WalkMap(ctx, h, visitor)
}

func (h *RSAPublicKey) AsMap(ctx context.Context) (map[string]interface{}, error) {
	return iter.AsMap(ctx, h)
}
