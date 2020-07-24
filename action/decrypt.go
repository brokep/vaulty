package action

import "github.com/vaulty/vaulty/encryption"

type Decrypt struct {
	enc encryption.Encrypter
}

func (a *Decrypt) Transform(body []byte) ([]byte, error) {
	newBody, err := a.enc.Decrypt(body)
	if err != nil {
		return nil, err
	}
	return newBody, nil
}
