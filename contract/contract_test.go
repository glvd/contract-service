package contract

import (
	"dhcrypto"
	"io/ioutil"
	"testing"
	"time"

	"github.com/goextension/log"
	"go.uber.org/zap"
)

// PublicKey ...
var PublicKey = []byte(`-----BEGIN CERTIFICATE-----
MIIFBjCCAu4CCQDSOHS/Tm+8kjANBgkqhkiG9w0BAQsFADBFMQswCQYDVQQGEwJD
TjETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0
cyBQdHkgTHRkMB4XDTE5MTEwNzA2MDA0N1oXDTIwMTEwNjA2MDA0N1owRTELMAkG
A1UEBhMCQ04xEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0
IFdpZGdpdHMgUHR5IEx0ZDCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIB
AMZVZC/Hj0zZIo3cXn9qJwuHHJp7x8e3eFcptCLiFKmdkBTRVGMPyDE99VBSKnoH
luUxrPWuN3BASzTYAskxCrKLj1uWY0qOHPAnjjUXSFvRrEUJNHyuF8x1Wq5vaYzz
Mdi5+FcpyLFo9T96sukvTTSKMPc9GpAb4M6/qURdR6LLRs/0VLJ48YPdaOpBntIB
IQLss0A6Lpgkvi50Wx3jY+gLBmu+93S70GZhRPfeeUnzBMvJulBbL/f6FoioS3R7
yiDPh013p2dqMCOpP4a3JlNN0bmqdZ0AMGA6Y3UWPADf94YgJBKDt4FP7MfrNzFz
3v4iK2QIShjFz7VQ29iA9X+rt2/CuLS9N0koNcmopn8u+4tbxk6R851AUA8Fh+U8
/IV88upEQA9gBrL8L24v0MyCRobTMNSlS15vblk1H+nvL68GdYEp6s65097Qak1q
DB5Cye6C54+Ok6inofSz68AbG7MLNWw0eFiLF3acM4JKE1nOy1GMmK+G/AAlilek
AcCrQ7wE1d+m58eT+dc5AEy5pq/syhrcjJ9JXZaWR/WMCqVuCjX3Km45J2wfm0DD
aJZR1MRfyblYMNixmweLM9m8mGGQKBV6jat9ZBZ4IuZz3B4uo/A8IKENNNVvJLIp
uOcEb7qiZlClkcoyYXhQrMH0v4HuhVUyFDmBMm7nI4odAgMBAAEwDQYJKoZIhvcN
AQELBQADggIBAKvvWUlbzPKlpyGtKKmHU3lkSqLsLed/mEXCiYrCKaXqCmU3peSv
rfpjTnup/q7DNS62kdUZ283f4P3cdHwEhFawrMyi098UTQjM3Zuy7L1J+ZFM9vHo
dDDYA7PlucIiDjPEtAd1ydQKis9yTsh/lQW8i4YcTtrm0FRVXHxMJnMuZPzJFAoh
5omINPQ3bTv7HRvWme4V24rGGzPJQAznjjE9NCwo/NbapXTk85Mjtt0CkX+oqiBB
m5ZVPGrkdTy1/j3TKnXumSyvlWUl6L/hzhwQA5M4OeG61ez1XJhuedEFcwl+io/X
JDC/8/M0g5kPlt/i+e3coSsCRdZfXZNCYj//lEXnyXUmmrlHYFc9y6n9VOl79eFn
quJjX4HrffVFjz/xP/7D5fQnMSSj9m+48mJTo9dKsaqU4W4QvwGglbMEC2yL5Yrw
I/KejxpBCc7cXc4hYGvA/iIy4NimFO/nQ83CEzwHC7FgOWhwyeJqqUZrYwwMGFTe
FZ3IZSpxwLxi2LUxqD4HCfke0CwXEiyL//Q+18ntvkDzdPB9jahY1A0iV/y9ZS5J
OLhSWN0ney5s04NkpX/R2n3AfV+Uazioc0L5pfoAFuhGyf6zYdN9XPMrgXeOap6i
Zw8J8Apa6dpcllxNY3Gna6y3byc8kclUvhRykxqf4u0xQ5TnTRgOe7AL
-----END CERTIFICATE-----
`)

func init() {
	logger, e := zap.NewProduction(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if e != nil {
		panic(e)
	}
	log.Register(logger.Sugar())
	DefaultNodeAddress = "0xA2A99890eA4Bf08E4C8875bbe39685679188A8A1"
	DefaultGatway = "http://localhost:8545"
}

// TestContract_AddNodes ...
func TestContract_AddNodes(t *testing.T) {
	c := NewContract(ETHClient(DefaultGatway), HexKey("2ed78769ad77af7fa01734e5f3302f03e7e40b94c4bdb1abaa5f54615b9ea0b1"), Node(DefaultNodeAddress))
	tm := time.Now()
	idx := tm.UnixNano()
	key, e := ioutil.ReadFile("./test_key/rsa.crt")
	if e != nil {
		t.Error(e)
	}
	_ = key
	log.Info("sec", idx)

	enc := dhcrypto.NewCipherEncoder(key, int(idx), tm)
	bytes, e := enc.Encode("/ip4/13.124.213.107/tcp/4001/ipfs/QmS9knxyQkdiFrGbzEb2FNqnq2yEzvBKx2pEGNBnXTWi7h")
	if e != nil {
		t.Error(e)
	}
	t.Log("encoded:", string(bytes))
	e = c.AddNodes(true, time.Now(), string(bytes))
	if e != nil {
		t.Error(e)
	}
}

// TestContract_GetNodes ...
func TestContract_GetNodes(t *testing.T) {
	c := NewContract(ETHClient(DefaultGatway), HexKey("2ed78769ad77af7fa01734e5f3302f03e7e40b94c4bdb1abaa5f54615b9ea0b1"), Node(DefaultNodeAddress))
	key, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	_ = key

	strings, i, e := c.GetNodes(time.Time{})
	if e != nil {
		t.Error(e)
		return
	}
	t.Log("ts:", i.Int64())

	tm := time.Unix(i.Int64(), 0)

	dec := dhcrypto.NewCipherDecode(key, tm)
	for _, s := range strings {
		bytes, e := dec.Decode(s)
		if e != nil {
			t.Log(e)
			continue
		}
		t.Log("decoded:", string(bytes))
	}

}
