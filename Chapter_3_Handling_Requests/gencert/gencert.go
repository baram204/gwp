package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)

	// 아주 큰 정수ㅇ로 이루어진 시리얼 번호
	serialNumber, _ := rand.Int(rand.Reader, max)
	// 구별된 이름과 인증서 제목 설정
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		// 유효기간을 1년으로 설정
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		// x509 인증서를 가리키고 서버 인증을 위해 사용됨
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		// 127.0.0.1 에서만 작동한다고 제한
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	//  crypto/rsa 이용 RSA 개인키 생성
	// 개인키는 공개키를 만들 때 사용됨
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// crypto/x509 라이브러리로 SSL 인증서를 만든다.
	// 숫자, 인증서, 공개키, 개인키를 사용해 포맷팅된 조각 생성
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// encoding/pem 라이브러리 이용해 인증서를 cert.pem 파일로 인코딩
	certOut, _ := os.Create("cert.pem")
	// pem 인코딩 블록 타입이 인증이고, 바이트 내용물은 위에서 만든거로 인코드 한 다음 파일에 기록한다.
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	// 기록 끝났으니 파일 닫는다.
	certOut.Close()

	// 개인 키 생성
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
