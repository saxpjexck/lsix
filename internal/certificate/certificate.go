package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"flag"
	"fmt"
	"jetbra-free/internal/util"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	forceDelete          bool
	binDir               = util.GetBinDir()
	certPath             = filepath.Join(binDir, "jetbra.pem")
	keyPath              = filepath.Join(binDir, "jetbra.key")
	powerPath            = filepath.Join(binDir, "power.txt")
	jaNetfilterpowerConf = filepath.Join(binDir, "static", "ja-netfilter", "config-jetbrains", "power.conf")
)

func init() {
	flag.BoolVar(&forceDelete, "f", false, "force delete existing files")
	flag.Parse()

	if forceDelete {
		DeleteFile(certPath)
		DeleteFile(keyPath)
		DeleteFile(powerPath)
	}

}

func Run() {
	println("Certificate...")

	GenCertificate()

	filePath := filepath.Join(powerPath)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := generateEqualResult()
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	startMarker := "; jetbra-free-start"
	endMarker := "; jetbra-free-end"
	fileContent, err := os.ReadFile(jaNetfilterpowerConf)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	fileStr := string(fileContent)

	if strings.Contains(fileStr, startMarker) && strings.Contains(fileStr, endMarker) {
		re := regexp.MustCompile(fmt.Sprintf(`(?s)%s.*?%s`, regexp.QuoteMeta(startMarker), regexp.QuoteMeta(endMarker)))
		fileStr = re.ReplaceAllString(fileStr, fmt.Sprintf("%s\n%s\n%s", startMarker, content, endMarker))
	} else {
		fileStr += fmt.Sprintf("\n%s\n%s\n%s", startMarker, content, endMarker)
	}

	err = os.WriteFile(jaNetfilterpowerConf, []byte(fileStr), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generate Certificate Finished")

}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DeleteFile(filename string) error {
	if FileExists(filename) {
		return os.Remove(filename)
	}
	return nil
}

func GenCertificate() error {

	if FileExists(keyPath) && FileExists(certPath) {
		return nil
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	notBefore := time.Now().AddDate(0, 0, -1)
	notAfter := time.Now().AddDate(10, 0, 0)

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		panic(err)
	}
	caTemplate := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{CommonName: "JetProfile CA"},
		Issuer:                pkix.Name{CommonName: "JetProfile CA"},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	_, err = x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	serialNumber, err = rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		panic(err)
	}

	template := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{CommonName: "jetbra-free"},
		Issuer:                pkix.Name{CommonName: "JetProfile CA"},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &caTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	certFile, err := os.Create(certPath)
	if err != nil {
		panic(err)
	}
	defer certFile.Close()
	pem.Encode(certFile, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	keyFile, err := os.Create(keyPath)
	if err != nil {
		panic(err)
	}
	defer keyFile.Close()
	pem.Encode(keyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	fmt.Printf("Private key saved to %s\n", keyPath)
	fmt.Printf("Certificate saved to %s\n", certPath)
	return nil
}

func printCertificateDetails(crt *x509.Certificate) {
	fmt.Println("Certificate Details:")
	fmt.Printf("  Subject: %s\n", crt.Subject)
	fmt.Printf("  Issuer: %s\n", crt.Issuer)
	fmt.Printf("  Serial Number: %s\n", crt.SerialNumber)
	fmt.Printf("  Not Before: %s\n", crt.NotBefore)
	fmt.Printf("  Not After: %s\n", crt.NotAfter)
	fmt.Printf("  Signature Algorithm: %s\n", crt.SignatureAlgorithm)
	fmt.Printf("  Public Key Algorithm: %s\n", crt.PublicKeyAlgorithm)

	fmt.Println("  Extensions:")
	for _, ext := range crt.Extensions {
		fmt.Printf("    OID: %s, Critical: %t, Value: %x\n", ext.Id, ext.Critical, ext.Value)
	}

	switch pub := crt.PublicKey.(type) {
	case *rsa.PublicKey:
		fmt.Printf("  RSA Public Key:\n")
		fmt.Printf("    Modulus: %x\n", pub.N)
		fmt.Printf("    Exponent: %d\n", pub.E)
	default:
		fmt.Println("  Public Key: (unsupported type)")
	}
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certPem, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	certBlock, _ := pem.Decode(certPem)
	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func extractSignature(cert *x509.Certificate) *big.Int {
	return new(big.Int).SetBytes(cert.Signature)
}

func extractPublicKeyExponent(cert *x509.Certificate) int {
	return cert.PublicKey.(*rsa.PublicKey).E
}

func extractPublicKey(cert *x509.Certificate) *rsa.PublicKey {
	return cert.PublicKey.(*rsa.PublicKey)
}

func calculateR(x, y *big.Int, jetbraPublicKey *rsa.PublicKey) *big.Int {
	r := new(big.Int)

	r.Exp(x, y, jetbraPublicKey.N)
	return r

}

func generateEqualResult() (string, error) {
	jetbraCertificate, err := loadCertificate(certPath)
	if err != nil {
		return "", err
	}
	printCertificateDetails(jetbraCertificate)

	x := extractSignature(jetbraCertificate)
	y := extractPublicKeyExponent(jetbraCertificate)

	rootCertificate, err := loadCertificate(filepath.Join(binDir, "static", "root_certificate.pem"))
	if err != nil {
		return "", err
	}

	printCertificateDetails(rootCertificate)

	z := extractPublicKey(rootCertificate).N

	r := calculateR(x, big.NewInt(int64(y)), extractPublicKey(jetbraCertificate))

	output := fmt.Sprintf("EQUAL,%d,%d,%d->%d", x, y, z, r)
	println(output)
	return output, nil
}
