package jsonld

import "bytes"

func UseContextCache(body []byte) (cached []byte) {
	cached = bytes.Replace(
		bytes.Replace(
			body,
			[]byte("https://www.w3.org/ns/activitystreams"),
			[]byte("jsonld/activitystreams.json"),
			1,
		),
		[]byte("https://w3id.org/security/v1"),
		[]byte("jsonld/security.json"),
		1,
	)
	return
}
