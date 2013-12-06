package goapi

import . `github.com/yak-labs/chirp-lang`
import (
	archive_tar `archive/tar`
	archive_zip `archive/zip`
	bufio `bufio`
	bytes `bytes`
	compress_bzip2 `compress/bzip2`
	compress_flate `compress/flate`
	compress_gzip `compress/gzip`
	compress_lzw `compress/lzw`
	compress_zlib `compress/zlib`
	container_heap `container/heap`
	container_list `container/list`
	container_ring `container/ring`
	crypto `crypto`
	crypto_aes `crypto/aes`
	crypto_cipher `crypto/cipher`
	crypto_des `crypto/des`
	crypto_dsa `crypto/dsa`
	crypto_ecdsa `crypto/ecdsa`
	crypto_elliptic `crypto/elliptic`
	crypto_hmac `crypto/hmac`
	crypto_md5 `crypto/md5`
	crypto_rand `crypto/rand`
	crypto_rc4 `crypto/rc4`
	crypto_rsa `crypto/rsa`
	crypto_sha1 `crypto/sha1`
	crypto_sha256 `crypto/sha256`
	crypto_sha512 `crypto/sha512`
	crypto_subtle `crypto/subtle`
	crypto_tls `crypto/tls`
	crypto_x509 `crypto/x509`
	crypto_x509_pkix `crypto/x509/pkix`
	database_sql `database/sql`
	database_sql_driver `database/sql/driver`
	debug_dwarf `debug/dwarf`
	debug_elf `debug/elf`
	debug_gosym `debug/gosym`
	debug_macho `debug/macho`
	debug_pe `debug/pe`
	encoding_ascii85 `encoding/ascii85`
	encoding_asn1 `encoding/asn1`
	encoding_base32 `encoding/base32`
	encoding_base64 `encoding/base64`
	encoding_binary `encoding/binary`
	encoding_csv `encoding/csv`
	encoding_gob `encoding/gob`
	encoding_hex `encoding/hex`
	encoding_json `encoding/json`
	encoding_pem `encoding/pem`
	encoding_xml `encoding/xml`
	errors `errors`
	expvar `expvar`
	flag `flag`
	fmt `fmt`
	go_ast `go/ast`
	go_build `go/build`
	go_doc `go/doc`
	go_parser `go/parser`
	go_printer `go/printer`
	go_scanner `go/scanner`
	go_token `go/token`
	hash_adler32 `hash/adler32`
	hash_crc32 `hash/crc32`
	hash_crc64 `hash/crc64`
	hash_fnv `hash/fnv`
	html `html`
	html_template `html/template`
	image `image`
	image_color `image/color`
	image_draw `image/draw`
	image_gif `image/gif`
	image_jpeg `image/jpeg`
	image_png `image/png`
	index_suffixarray `index/suffixarray`
	io `io`
	io_ioutil `io/ioutil`
	log `log`
	math `math`
	math_big `math/big`
	math_cmplx `math/cmplx`
	math_rand `math/rand`
	mime `mime`
	mime_multipart `mime/multipart`
	net `net`
	net_http `net/http`
	net_http_cgi `net/http/cgi`
	net_http_fcgi `net/http/fcgi`
	net_http_httptest `net/http/httptest`
	net_http_httputil `net/http/httputil`
	net_http_pprof `net/http/pprof`
	net_mail `net/mail`
	net_rpc `net/rpc`
	net_rpc_jsonrpc `net/rpc/jsonrpc`
	net_smtp `net/smtp`
	net_textproto `net/textproto`
	net_url `net/url`
	os `os`
	os_exec `os/exec`
	os_signal `os/signal`
	os_user `os/user`
	path `path`
	path_filepath `path/filepath`
	reflect `reflect`
	regexp `regexp`
	regexp_syntax `regexp/syntax`
	runtime `runtime`
	runtime_debug `runtime/debug`
	runtime_pprof `runtime/pprof`
	sort `sort`
	strconv `strconv`
	strings `strings`
	sync `sync`
	sync_atomic `sync/atomic`
	syscall `syscall`
	text_scanner `text/scanner`
	text_tabwriter `text/tabwriter`
	text_template `text/template`
	text_template_parse `text/template/parse`
	time `time`
	unicode `unicode`
	unicode_utf16 `unicode/utf16`
	unicode_utf8 `unicode/utf8`
)

func init() {
	Roots[`/archive/tar/NewReader`] = FuncRoot{Func: reflect.ValueOf(archive_tar.NewReader)}
	Roots[`/archive/tar/NewWriter`] = FuncRoot{Func: reflect.ValueOf(archive_tar.NewWriter)}
	Roots[`/archive/zip/FileInfoHeader`] = FuncRoot{Func: reflect.ValueOf(archive_zip.FileInfoHeader)}
	Roots[`/archive/zip/NewReader`] = FuncRoot{Func: reflect.ValueOf(archive_zip.NewReader)}
	Roots[`/archive/zip/NewWriter`] = FuncRoot{Func: reflect.ValueOf(archive_zip.NewWriter)}
	Roots[`/archive/zip/OpenReader`] = FuncRoot{Func: reflect.ValueOf(archive_zip.OpenReader)}
	Roots[`/bufio/NewReadWriter`] = FuncRoot{Func: reflect.ValueOf(bufio.NewReadWriter)}
	Roots[`/bufio/NewReader`] = FuncRoot{Func: reflect.ValueOf(bufio.NewReader)}
	Roots[`/bufio/NewReaderSize`] = FuncRoot{Func: reflect.ValueOf(bufio.NewReaderSize)}
	Roots[`/bufio/NewWriter`] = FuncRoot{Func: reflect.ValueOf(bufio.NewWriter)}
	Roots[`/bufio/NewWriterSize`] = FuncRoot{Func: reflect.ValueOf(bufio.NewWriterSize)}
	Roots[`/bytes/Compare`] = FuncRoot{Func: reflect.ValueOf(bytes.Compare)}
	Roots[`/bytes/Contains`] = FuncRoot{Func: reflect.ValueOf(bytes.Contains)}
	Roots[`/bytes/Count`] = FuncRoot{Func: reflect.ValueOf(bytes.Count)}
	Roots[`/bytes/Equal`] = FuncRoot{Func: reflect.ValueOf(bytes.Equal)}
	Roots[`/bytes/EqualFold`] = FuncRoot{Func: reflect.ValueOf(bytes.EqualFold)}
	Roots[`/bytes/Fields`] = FuncRoot{Func: reflect.ValueOf(bytes.Fields)}
	Roots[`/bytes/FieldsFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.FieldsFunc)}
	Roots[`/bytes/HasPrefix`] = FuncRoot{Func: reflect.ValueOf(bytes.HasPrefix)}
	Roots[`/bytes/HasSuffix`] = FuncRoot{Func: reflect.ValueOf(bytes.HasSuffix)}
	Roots[`/bytes/Index`] = FuncRoot{Func: reflect.ValueOf(bytes.Index)}
	Roots[`/bytes/IndexAny`] = FuncRoot{Func: reflect.ValueOf(bytes.IndexAny)}
	Roots[`/bytes/IndexByte`] = FuncRoot{Func: reflect.ValueOf(bytes.IndexByte)}
	Roots[`/bytes/IndexFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.IndexFunc)}
	Roots[`/bytes/IndexRune`] = FuncRoot{Func: reflect.ValueOf(bytes.IndexRune)}
	Roots[`/bytes/Join`] = FuncRoot{Func: reflect.ValueOf(bytes.Join)}
	Roots[`/bytes/LastIndex`] = FuncRoot{Func: reflect.ValueOf(bytes.LastIndex)}
	Roots[`/bytes/LastIndexAny`] = FuncRoot{Func: reflect.ValueOf(bytes.LastIndexAny)}
	Roots[`/bytes/LastIndexFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.LastIndexFunc)}
	Roots[`/bytes/Map`] = FuncRoot{Func: reflect.ValueOf(bytes.Map)}
	Roots[`/bytes/NewBuffer`] = FuncRoot{Func: reflect.ValueOf(bytes.NewBuffer)}
	Roots[`/bytes/NewBufferString`] = FuncRoot{Func: reflect.ValueOf(bytes.NewBufferString)}
	Roots[`/bytes/NewReader`] = FuncRoot{Func: reflect.ValueOf(bytes.NewReader)}
	Roots[`/bytes/Repeat`] = FuncRoot{Func: reflect.ValueOf(bytes.Repeat)}
	Roots[`/bytes/Replace`] = FuncRoot{Func: reflect.ValueOf(bytes.Replace)}
	Roots[`/bytes/Runes`] = FuncRoot{Func: reflect.ValueOf(bytes.Runes)}
	Roots[`/bytes/Split`] = FuncRoot{Func: reflect.ValueOf(bytes.Split)}
	Roots[`/bytes/SplitAfter`] = FuncRoot{Func: reflect.ValueOf(bytes.SplitAfter)}
	Roots[`/bytes/SplitAfterN`] = FuncRoot{Func: reflect.ValueOf(bytes.SplitAfterN)}
	Roots[`/bytes/SplitN`] = FuncRoot{Func: reflect.ValueOf(bytes.SplitN)}
	Roots[`/bytes/Title`] = FuncRoot{Func: reflect.ValueOf(bytes.Title)}
	Roots[`/bytes/ToLower`] = FuncRoot{Func: reflect.ValueOf(bytes.ToLower)}
	Roots[`/bytes/ToLowerSpecial`] = FuncRoot{Func: reflect.ValueOf(bytes.ToLowerSpecial)}
	Roots[`/bytes/ToTitle`] = FuncRoot{Func: reflect.ValueOf(bytes.ToTitle)}
	Roots[`/bytes/ToTitleSpecial`] = FuncRoot{Func: reflect.ValueOf(bytes.ToTitleSpecial)}
	Roots[`/bytes/ToUpper`] = FuncRoot{Func: reflect.ValueOf(bytes.ToUpper)}
	Roots[`/bytes/ToUpperSpecial`] = FuncRoot{Func: reflect.ValueOf(bytes.ToUpperSpecial)}
	Roots[`/bytes/Trim`] = FuncRoot{Func: reflect.ValueOf(bytes.Trim)}
	Roots[`/bytes/TrimFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimFunc)}
	Roots[`/bytes/TrimLeft`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimLeft)}
	Roots[`/bytes/TrimLeftFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimLeftFunc)}
	Roots[`/bytes/TrimRight`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimRight)}
	Roots[`/bytes/TrimRightFunc`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimRightFunc)}
	Roots[`/bytes/TrimSpace`] = FuncRoot{Func: reflect.ValueOf(bytes.TrimSpace)}
	Roots[`/compress/bzip2/NewReader`] = FuncRoot{Func: reflect.ValueOf(compress_bzip2.NewReader)}
	Roots[`/compress/flate/NewReader`] = FuncRoot{Func: reflect.ValueOf(compress_flate.NewReader)}
	Roots[`/compress/flate/NewReaderDict`] = FuncRoot{Func: reflect.ValueOf(compress_flate.NewReaderDict)}
	Roots[`/compress/flate/NewWriter`] = FuncRoot{Func: reflect.ValueOf(compress_flate.NewWriter)}
	Roots[`/compress/flate/NewWriterDict`] = FuncRoot{Func: reflect.ValueOf(compress_flate.NewWriterDict)}
	Roots[`/compress/gzip/NewReader`] = FuncRoot{Func: reflect.ValueOf(compress_gzip.NewReader)}
	Roots[`/compress/gzip/NewWriter`] = FuncRoot{Func: reflect.ValueOf(compress_gzip.NewWriter)}
	Roots[`/compress/gzip/NewWriterLevel`] = FuncRoot{Func: reflect.ValueOf(compress_gzip.NewWriterLevel)}
	Roots[`/compress/lzw/NewReader`] = FuncRoot{Func: reflect.ValueOf(compress_lzw.NewReader)}
	Roots[`/compress/lzw/NewWriter`] = FuncRoot{Func: reflect.ValueOf(compress_lzw.NewWriter)}
	Roots[`/compress/zlib/NewReader`] = FuncRoot{Func: reflect.ValueOf(compress_zlib.NewReader)}
	Roots[`/compress/zlib/NewReaderDict`] = FuncRoot{Func: reflect.ValueOf(compress_zlib.NewReaderDict)}
	Roots[`/compress/zlib/NewWriter`] = FuncRoot{Func: reflect.ValueOf(compress_zlib.NewWriter)}
	Roots[`/compress/zlib/NewWriterLevel`] = FuncRoot{Func: reflect.ValueOf(compress_zlib.NewWriterLevel)}
	Roots[`/compress/zlib/NewWriterLevelDict`] = FuncRoot{Func: reflect.ValueOf(compress_zlib.NewWriterLevelDict)}
	Roots[`/container/heap/Init`] = FuncRoot{Func: reflect.ValueOf(container_heap.Init)}
	Roots[`/container/heap/Pop`] = FuncRoot{Func: reflect.ValueOf(container_heap.Pop)}
	Roots[`/container/heap/Push`] = FuncRoot{Func: reflect.ValueOf(container_heap.Push)}
	Roots[`/container/heap/Remove`] = FuncRoot{Func: reflect.ValueOf(container_heap.Remove)}
	Roots[`/container/list/New`] = FuncRoot{Func: reflect.ValueOf(container_list.New)}
	Roots[`/container/ring/New`] = FuncRoot{Func: reflect.ValueOf(container_ring.New)}
	Roots[`/crypto/RegisterHash`] = FuncRoot{Func: reflect.ValueOf(crypto.RegisterHash)}
	Roots[`/crypto/aes/NewCipher`] = FuncRoot{Func: reflect.ValueOf(crypto_aes.NewCipher)}
	Roots[`/crypto/cipher/NewCBCDecrypter`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewCBCDecrypter)}
	Roots[`/crypto/cipher/NewCBCEncrypter`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewCBCEncrypter)}
	Roots[`/crypto/cipher/NewCFBDecrypter`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewCFBDecrypter)}
	Roots[`/crypto/cipher/NewCFBEncrypter`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewCFBEncrypter)}
	Roots[`/crypto/cipher/NewCTR`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewCTR)}
	Roots[`/crypto/cipher/NewOFB`] = FuncRoot{Func: reflect.ValueOf(crypto_cipher.NewOFB)}
	Roots[`/crypto/des/NewCipher`] = FuncRoot{Func: reflect.ValueOf(crypto_des.NewCipher)}
	Roots[`/crypto/des/NewTripleDESCipher`] = FuncRoot{Func: reflect.ValueOf(crypto_des.NewTripleDESCipher)}
	Roots[`/crypto/dsa/GenerateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_dsa.GenerateKey)}
	Roots[`/crypto/dsa/GenerateParameters`] = FuncRoot{Func: reflect.ValueOf(crypto_dsa.GenerateParameters)}
	Roots[`/crypto/dsa/Sign`] = FuncRoot{Func: reflect.ValueOf(crypto_dsa.Sign)}
	Roots[`/crypto/dsa/Verify`] = FuncRoot{Func: reflect.ValueOf(crypto_dsa.Verify)}
	Roots[`/crypto/ecdsa/GenerateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_ecdsa.GenerateKey)}
	Roots[`/crypto/ecdsa/Sign`] = FuncRoot{Func: reflect.ValueOf(crypto_ecdsa.Sign)}
	Roots[`/crypto/ecdsa/Verify`] = FuncRoot{Func: reflect.ValueOf(crypto_ecdsa.Verify)}
	Roots[`/crypto/elliptic/GenerateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.GenerateKey)}
	Roots[`/crypto/elliptic/Marshal`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.Marshal)}
	Roots[`/crypto/elliptic/P224`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.P224)}
	Roots[`/crypto/elliptic/P256`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.P256)}
	Roots[`/crypto/elliptic/P384`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.P384)}
	Roots[`/crypto/elliptic/P521`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.P521)}
	Roots[`/crypto/elliptic/Unmarshal`] = FuncRoot{Func: reflect.ValueOf(crypto_elliptic.Unmarshal)}
	Roots[`/crypto/hmac/New`] = FuncRoot{Func: reflect.ValueOf(crypto_hmac.New)}
	Roots[`/crypto/md5/New`] = FuncRoot{Func: reflect.ValueOf(crypto_md5.New)}
	Roots[`/crypto/rand/Int`] = FuncRoot{Func: reflect.ValueOf(crypto_rand.Int)}
	Roots[`/crypto/rand/Prime`] = FuncRoot{Func: reflect.ValueOf(crypto_rand.Prime)}
	Roots[`/crypto/rand/Read`] = FuncRoot{Func: reflect.ValueOf(crypto_rand.Read)}
	Roots[`/crypto/rc4/NewCipher`] = FuncRoot{Func: reflect.ValueOf(crypto_rc4.NewCipher)}
	Roots[`/crypto/rsa/DecryptOAEP`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.DecryptOAEP)}
	Roots[`/crypto/rsa/DecryptPKCS1v15`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.DecryptPKCS1v15)}
	Roots[`/crypto/rsa/DecryptPKCS1v15SessionKey`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.DecryptPKCS1v15SessionKey)}
	Roots[`/crypto/rsa/EncryptOAEP`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.EncryptOAEP)}
	Roots[`/crypto/rsa/EncryptPKCS1v15`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.EncryptPKCS1v15)}
	Roots[`/crypto/rsa/GenerateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.GenerateKey)}
	Roots[`/crypto/rsa/GenerateMultiPrimeKey`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.GenerateMultiPrimeKey)}
	Roots[`/crypto/rsa/SignPKCS1v15`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.SignPKCS1v15)}
	Roots[`/crypto/rsa/VerifyPKCS1v15`] = FuncRoot{Func: reflect.ValueOf(crypto_rsa.VerifyPKCS1v15)}
	Roots[`/crypto/sha1/New`] = FuncRoot{Func: reflect.ValueOf(crypto_sha1.New)}
	Roots[`/crypto/sha256/New`] = FuncRoot{Func: reflect.ValueOf(crypto_sha256.New)}
	Roots[`/crypto/sha256/New224`] = FuncRoot{Func: reflect.ValueOf(crypto_sha256.New224)}
	Roots[`/crypto/sha512/New`] = FuncRoot{Func: reflect.ValueOf(crypto_sha512.New)}
	Roots[`/crypto/sha512/New384`] = FuncRoot{Func: reflect.ValueOf(crypto_sha512.New384)}
	Roots[`/crypto/subtle/ConstantTimeByteEq`] = FuncRoot{Func: reflect.ValueOf(crypto_subtle.ConstantTimeByteEq)}
	Roots[`/crypto/subtle/ConstantTimeCompare`] = FuncRoot{Func: reflect.ValueOf(crypto_subtle.ConstantTimeCompare)}
	Roots[`/crypto/subtle/ConstantTimeCopy`] = FuncRoot{Func: reflect.ValueOf(crypto_subtle.ConstantTimeCopy)}
	Roots[`/crypto/subtle/ConstantTimeEq`] = FuncRoot{Func: reflect.ValueOf(crypto_subtle.ConstantTimeEq)}
	Roots[`/crypto/subtle/ConstantTimeSelect`] = FuncRoot{Func: reflect.ValueOf(crypto_subtle.ConstantTimeSelect)}
	Roots[`/crypto/tls/Client`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.Client)}
	Roots[`/crypto/tls/Dial`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.Dial)}
	Roots[`/crypto/tls/Listen`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.Listen)}
	Roots[`/crypto/tls/LoadX509KeyPair`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.LoadX509KeyPair)}
	Roots[`/crypto/tls/NewListener`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.NewListener)}
	Roots[`/crypto/tls/Server`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.Server)}
	Roots[`/crypto/tls/X509KeyPair`] = FuncRoot{Func: reflect.ValueOf(crypto_tls.X509KeyPair)}
	Roots[`/crypto/x509/CreateCertificate`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.CreateCertificate)}
	Roots[`/crypto/x509/MarshalPKCS1PrivateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.MarshalPKCS1PrivateKey)}
	Roots[`/crypto/x509/MarshalPKIXPublicKey`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.MarshalPKIXPublicKey)}
	Roots[`/crypto/x509/NewCertPool`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.NewCertPool)}
	Roots[`/crypto/x509/ParseCRL`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParseCRL)}
	Roots[`/crypto/x509/ParseCertificate`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParseCertificate)}
	Roots[`/crypto/x509/ParseCertificates`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParseCertificates)}
	Roots[`/crypto/x509/ParseDERCRL`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParseDERCRL)}
	Roots[`/crypto/x509/ParsePKCS1PrivateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParsePKCS1PrivateKey)}
	Roots[`/crypto/x509/ParsePKCS8PrivateKey`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParsePKCS8PrivateKey)}
	Roots[`/crypto/x509/ParsePKIXPublicKey`] = FuncRoot{Func: reflect.ValueOf(crypto_x509.ParsePKIXPublicKey)}
	Roots[`/database/sql/Open`] = FuncRoot{Func: reflect.ValueOf(database_sql.Open)}
	Roots[`/database/sql/Register`] = FuncRoot{Func: reflect.ValueOf(database_sql.Register)}
	Roots[`/database/sql/driver/IsScanValue`] = FuncRoot{Func: reflect.ValueOf(database_sql_driver.IsScanValue)}
	Roots[`/database/sql/driver/IsValue`] = FuncRoot{Func: reflect.ValueOf(database_sql_driver.IsValue)}
	Roots[`/debug/dwarf/New`] = FuncRoot{Func: reflect.ValueOf(debug_dwarf.New)}
	Roots[`/debug/elf/NewFile`] = FuncRoot{Func: reflect.ValueOf(debug_elf.NewFile)}
	Roots[`/debug/elf/Open`] = FuncRoot{Func: reflect.ValueOf(debug_elf.Open)}
	Roots[`/debug/elf/R_INFO`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_INFO)}
	Roots[`/debug/elf/R_INFO32`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_INFO32)}
	Roots[`/debug/elf/R_SYM32`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_SYM32)}
	Roots[`/debug/elf/R_SYM64`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_SYM64)}
	Roots[`/debug/elf/R_TYPE32`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_TYPE32)}
	Roots[`/debug/elf/R_TYPE64`] = FuncRoot{Func: reflect.ValueOf(debug_elf.R_TYPE64)}
	Roots[`/debug/elf/ST_BIND`] = FuncRoot{Func: reflect.ValueOf(debug_elf.ST_BIND)}
	Roots[`/debug/elf/ST_INFO`] = FuncRoot{Func: reflect.ValueOf(debug_elf.ST_INFO)}
	Roots[`/debug/elf/ST_TYPE`] = FuncRoot{Func: reflect.ValueOf(debug_elf.ST_TYPE)}
	Roots[`/debug/elf/ST_VISIBILITY`] = FuncRoot{Func: reflect.ValueOf(debug_elf.ST_VISIBILITY)}
	Roots[`/debug/gosym/NewLineTable`] = FuncRoot{Func: reflect.ValueOf(debug_gosym.NewLineTable)}
	Roots[`/debug/gosym/NewTable`] = FuncRoot{Func: reflect.ValueOf(debug_gosym.NewTable)}
	Roots[`/debug/macho/NewFile`] = FuncRoot{Func: reflect.ValueOf(debug_macho.NewFile)}
	Roots[`/debug/macho/Open`] = FuncRoot{Func: reflect.ValueOf(debug_macho.Open)}
	Roots[`/debug/pe/NewFile`] = FuncRoot{Func: reflect.ValueOf(debug_pe.NewFile)}
	Roots[`/debug/pe/Open`] = FuncRoot{Func: reflect.ValueOf(debug_pe.Open)}
	Roots[`/encoding/ascii85/Decode`] = FuncRoot{Func: reflect.ValueOf(encoding_ascii85.Decode)}
	Roots[`/encoding/ascii85/Encode`] = FuncRoot{Func: reflect.ValueOf(encoding_ascii85.Encode)}
	Roots[`/encoding/ascii85/MaxEncodedLen`] = FuncRoot{Func: reflect.ValueOf(encoding_ascii85.MaxEncodedLen)}
	Roots[`/encoding/ascii85/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_ascii85.NewDecoder)}
	Roots[`/encoding/ascii85/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_ascii85.NewEncoder)}
	Roots[`/encoding/asn1/Marshal`] = FuncRoot{Func: reflect.ValueOf(encoding_asn1.Marshal)}
	Roots[`/encoding/asn1/Unmarshal`] = FuncRoot{Func: reflect.ValueOf(encoding_asn1.Unmarshal)}
	Roots[`/encoding/asn1/UnmarshalWithParams`] = FuncRoot{Func: reflect.ValueOf(encoding_asn1.UnmarshalWithParams)}
	Roots[`/encoding/base32/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_base32.NewDecoder)}
	Roots[`/encoding/base32/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_base32.NewEncoder)}
	Roots[`/encoding/base32/NewEncoding`] = FuncRoot{Func: reflect.ValueOf(encoding_base32.NewEncoding)}
	Roots[`/encoding/base64/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_base64.NewDecoder)}
	Roots[`/encoding/base64/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_base64.NewEncoder)}
	Roots[`/encoding/base64/NewEncoding`] = FuncRoot{Func: reflect.ValueOf(encoding_base64.NewEncoding)}
	Roots[`/encoding/binary/PutUvarint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.PutUvarint)}
	Roots[`/encoding/binary/PutVarint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.PutVarint)}
	Roots[`/encoding/binary/Read`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.Read)}
	Roots[`/encoding/binary/ReadUvarint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.ReadUvarint)}
	Roots[`/encoding/binary/ReadVarint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.ReadVarint)}
	Roots[`/encoding/binary/Size`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.Size)}
	Roots[`/encoding/binary/Uvarint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.Uvarint)}
	Roots[`/encoding/binary/Varint`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.Varint)}
	Roots[`/encoding/binary/Write`] = FuncRoot{Func: reflect.ValueOf(encoding_binary.Write)}
	Roots[`/encoding/csv/NewReader`] = FuncRoot{Func: reflect.ValueOf(encoding_csv.NewReader)}
	Roots[`/encoding/csv/NewWriter`] = FuncRoot{Func: reflect.ValueOf(encoding_csv.NewWriter)}
	Roots[`/encoding/gob/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_gob.NewDecoder)}
	Roots[`/encoding/gob/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_gob.NewEncoder)}
	Roots[`/encoding/gob/Register`] = FuncRoot{Func: reflect.ValueOf(encoding_gob.Register)}
	Roots[`/encoding/gob/RegisterName`] = FuncRoot{Func: reflect.ValueOf(encoding_gob.RegisterName)}
	Roots[`/encoding/hex/Decode`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.Decode)}
	Roots[`/encoding/hex/DecodeString`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.DecodeString)}
	Roots[`/encoding/hex/DecodedLen`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.DecodedLen)}
	Roots[`/encoding/hex/Dump`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.Dump)}
	Roots[`/encoding/hex/Dumper`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.Dumper)}
	Roots[`/encoding/hex/Encode`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.Encode)}
	Roots[`/encoding/hex/EncodeToString`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.EncodeToString)}
	Roots[`/encoding/hex/EncodedLen`] = FuncRoot{Func: reflect.ValueOf(encoding_hex.EncodedLen)}
	Roots[`/encoding/json/Compact`] = FuncRoot{Func: reflect.ValueOf(encoding_json.Compact)}
	Roots[`/encoding/json/HTMLEscape`] = FuncRoot{Func: reflect.ValueOf(encoding_json.HTMLEscape)}
	Roots[`/encoding/json/Indent`] = FuncRoot{Func: reflect.ValueOf(encoding_json.Indent)}
	Roots[`/encoding/json/Marshal`] = FuncRoot{Func: reflect.ValueOf(encoding_json.Marshal)}
	Roots[`/encoding/json/MarshalIndent`] = FuncRoot{Func: reflect.ValueOf(encoding_json.MarshalIndent)}
	Roots[`/encoding/json/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_json.NewDecoder)}
	Roots[`/encoding/json/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_json.NewEncoder)}
	Roots[`/encoding/json/Unmarshal`] = FuncRoot{Func: reflect.ValueOf(encoding_json.Unmarshal)}
	Roots[`/encoding/pem/Decode`] = FuncRoot{Func: reflect.ValueOf(encoding_pem.Decode)}
	Roots[`/encoding/pem/Encode`] = FuncRoot{Func: reflect.ValueOf(encoding_pem.Encode)}
	Roots[`/encoding/pem/EncodeToMemory`] = FuncRoot{Func: reflect.ValueOf(encoding_pem.EncodeToMemory)}
	Roots[`/encoding/xml/CopyToken`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.CopyToken)}
	Roots[`/encoding/xml/Escape`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.Escape)}
	Roots[`/encoding/xml/Marshal`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.Marshal)}
	Roots[`/encoding/xml/MarshalIndent`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.MarshalIndent)}
	Roots[`/encoding/xml/NewDecoder`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.NewDecoder)}
	Roots[`/encoding/xml/NewEncoder`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.NewEncoder)}
	Roots[`/encoding/xml/Unmarshal`] = FuncRoot{Func: reflect.ValueOf(encoding_xml.Unmarshal)}
	Roots[`/errors/New`] = FuncRoot{Func: reflect.ValueOf(errors.New)}
	Roots[`/expvar/Do`] = FuncRoot{Func: reflect.ValueOf(expvar.Do)}
	Roots[`/expvar/Get`] = FuncRoot{Func: reflect.ValueOf(expvar.Get)}
	Roots[`/expvar/NewFloat`] = FuncRoot{Func: reflect.ValueOf(expvar.NewFloat)}
	Roots[`/expvar/NewInt`] = FuncRoot{Func: reflect.ValueOf(expvar.NewInt)}
	Roots[`/expvar/NewMap`] = FuncRoot{Func: reflect.ValueOf(expvar.NewMap)}
	Roots[`/expvar/NewString`] = FuncRoot{Func: reflect.ValueOf(expvar.NewString)}
	Roots[`/expvar/Publish`] = FuncRoot{Func: reflect.ValueOf(expvar.Publish)}
	Roots[`/flag/Arg`] = FuncRoot{Func: reflect.ValueOf(flag.Arg)}
	Roots[`/flag/Args`] = FuncRoot{Func: reflect.ValueOf(flag.Args)}
	Roots[`/flag/Bool`] = FuncRoot{Func: reflect.ValueOf(flag.Bool)}
	Roots[`/flag/BoolVar`] = FuncRoot{Func: reflect.ValueOf(flag.BoolVar)}
	Roots[`/flag/Duration`] = FuncRoot{Func: reflect.ValueOf(flag.Duration)}
	Roots[`/flag/DurationVar`] = FuncRoot{Func: reflect.ValueOf(flag.DurationVar)}
	Roots[`/flag/Float64`] = FuncRoot{Func: reflect.ValueOf(flag.Float64)}
	Roots[`/flag/Float64Var`] = FuncRoot{Func: reflect.ValueOf(flag.Float64Var)}
	Roots[`/flag/Int`] = FuncRoot{Func: reflect.ValueOf(flag.Int)}
	Roots[`/flag/Int64`] = FuncRoot{Func: reflect.ValueOf(flag.Int64)}
	Roots[`/flag/Int64Var`] = FuncRoot{Func: reflect.ValueOf(flag.Int64Var)}
	Roots[`/flag/IntVar`] = FuncRoot{Func: reflect.ValueOf(flag.IntVar)}
	Roots[`/flag/Lookup`] = FuncRoot{Func: reflect.ValueOf(flag.Lookup)}
	Roots[`/flag/NArg`] = FuncRoot{Func: reflect.ValueOf(flag.NArg)}
	Roots[`/flag/NFlag`] = FuncRoot{Func: reflect.ValueOf(flag.NFlag)}
	Roots[`/flag/NewFlagSet`] = FuncRoot{Func: reflect.ValueOf(flag.NewFlagSet)}
	Roots[`/flag/Parse`] = FuncRoot{Func: reflect.ValueOf(flag.Parse)}
	Roots[`/flag/Parsed`] = FuncRoot{Func: reflect.ValueOf(flag.Parsed)}
	Roots[`/flag/PrintDefaults`] = FuncRoot{Func: reflect.ValueOf(flag.PrintDefaults)}
	Roots[`/flag/Set`] = FuncRoot{Func: reflect.ValueOf(flag.Set)}
	Roots[`/flag/String`] = FuncRoot{Func: reflect.ValueOf(flag.String)}
	Roots[`/flag/StringVar`] = FuncRoot{Func: reflect.ValueOf(flag.StringVar)}
	Roots[`/flag/Uint`] = FuncRoot{Func: reflect.ValueOf(flag.Uint)}
	Roots[`/flag/Uint64`] = FuncRoot{Func: reflect.ValueOf(flag.Uint64)}
	Roots[`/flag/Uint64Var`] = FuncRoot{Func: reflect.ValueOf(flag.Uint64Var)}
	Roots[`/flag/UintVar`] = FuncRoot{Func: reflect.ValueOf(flag.UintVar)}
	Roots[`/flag/Var`] = FuncRoot{Func: reflect.ValueOf(flag.Var)}
	Roots[`/flag/Visit`] = FuncRoot{Func: reflect.ValueOf(flag.Visit)}
	Roots[`/flag/VisitAll`] = FuncRoot{Func: reflect.ValueOf(flag.VisitAll)}
	Roots[`/fmt/Errorf`] = FuncRoot{Func: reflect.ValueOf(fmt.Errorf)}
	Roots[`/fmt/Fprint`] = FuncRoot{Func: reflect.ValueOf(fmt.Fprint)}
	Roots[`/fmt/Fprintf`] = FuncRoot{Func: reflect.ValueOf(fmt.Fprintf)}
	Roots[`/fmt/Fprintln`] = FuncRoot{Func: reflect.ValueOf(fmt.Fprintln)}
	Roots[`/fmt/Fscan`] = FuncRoot{Func: reflect.ValueOf(fmt.Fscan)}
	Roots[`/fmt/Fscanf`] = FuncRoot{Func: reflect.ValueOf(fmt.Fscanf)}
	Roots[`/fmt/Fscanln`] = FuncRoot{Func: reflect.ValueOf(fmt.Fscanln)}
	Roots[`/fmt/Print`] = FuncRoot{Func: reflect.ValueOf(fmt.Print)}
	Roots[`/fmt/Printf`] = FuncRoot{Func: reflect.ValueOf(fmt.Printf)}
	Roots[`/fmt/Println`] = FuncRoot{Func: reflect.ValueOf(fmt.Println)}
	Roots[`/fmt/Scan`] = FuncRoot{Func: reflect.ValueOf(fmt.Scan)}
	Roots[`/fmt/Scanf`] = FuncRoot{Func: reflect.ValueOf(fmt.Scanf)}
	Roots[`/fmt/Scanln`] = FuncRoot{Func: reflect.ValueOf(fmt.Scanln)}
	Roots[`/fmt/Sprint`] = FuncRoot{Func: reflect.ValueOf(fmt.Sprint)}
	Roots[`/fmt/Sprintf`] = FuncRoot{Func: reflect.ValueOf(fmt.Sprintf)}
	Roots[`/fmt/Sprintln`] = FuncRoot{Func: reflect.ValueOf(fmt.Sprintln)}
	Roots[`/fmt/Sscan`] = FuncRoot{Func: reflect.ValueOf(fmt.Sscan)}
	Roots[`/fmt/Sscanf`] = FuncRoot{Func: reflect.ValueOf(fmt.Sscanf)}
	Roots[`/fmt/Sscanln`] = FuncRoot{Func: reflect.ValueOf(fmt.Sscanln)}
	Roots[`/go/ast/FileExports`] = FuncRoot{Func: reflect.ValueOf(go_ast.FileExports)}
	Roots[`/go/ast/FilterDecl`] = FuncRoot{Func: reflect.ValueOf(go_ast.FilterDecl)}
	Roots[`/go/ast/FilterFile`] = FuncRoot{Func: reflect.ValueOf(go_ast.FilterFile)}
	Roots[`/go/ast/FilterPackage`] = FuncRoot{Func: reflect.ValueOf(go_ast.FilterPackage)}
	Roots[`/go/ast/Fprint`] = FuncRoot{Func: reflect.ValueOf(go_ast.Fprint)}
	Roots[`/go/ast/Inspect`] = FuncRoot{Func: reflect.ValueOf(go_ast.Inspect)}
	Roots[`/go/ast/IsExported`] = FuncRoot{Func: reflect.ValueOf(go_ast.IsExported)}
	Roots[`/go/ast/MergePackageFiles`] = FuncRoot{Func: reflect.ValueOf(go_ast.MergePackageFiles)}
	Roots[`/go/ast/NewIdent`] = FuncRoot{Func: reflect.ValueOf(go_ast.NewIdent)}
	Roots[`/go/ast/NewObj`] = FuncRoot{Func: reflect.ValueOf(go_ast.NewObj)}
	Roots[`/go/ast/NewPackage`] = FuncRoot{Func: reflect.ValueOf(go_ast.NewPackage)}
	Roots[`/go/ast/NewScope`] = FuncRoot{Func: reflect.ValueOf(go_ast.NewScope)}
	Roots[`/go/ast/NotNilFilter`] = FuncRoot{Func: reflect.ValueOf(go_ast.NotNilFilter)}
	Roots[`/go/ast/PackageExports`] = FuncRoot{Func: reflect.ValueOf(go_ast.PackageExports)}
	Roots[`/go/ast/Print`] = FuncRoot{Func: reflect.ValueOf(go_ast.Print)}
	Roots[`/go/ast/SortImports`] = FuncRoot{Func: reflect.ValueOf(go_ast.SortImports)}
	Roots[`/go/ast/Walk`] = FuncRoot{Func: reflect.ValueOf(go_ast.Walk)}
	Roots[`/go/build/ArchChar`] = FuncRoot{Func: reflect.ValueOf(go_build.ArchChar)}
	Roots[`/go/build/Import`] = FuncRoot{Func: reflect.ValueOf(go_build.Import)}
	Roots[`/go/build/ImportDir`] = FuncRoot{Func: reflect.ValueOf(go_build.ImportDir)}
	Roots[`/go/build/IsLocalImport`] = FuncRoot{Func: reflect.ValueOf(go_build.IsLocalImport)}
	Roots[`/go/doc/Examples`] = FuncRoot{Func: reflect.ValueOf(go_doc.Examples)}
	Roots[`/go/doc/New`] = FuncRoot{Func: reflect.ValueOf(go_doc.New)}
	Roots[`/go/doc/Synopsis`] = FuncRoot{Func: reflect.ValueOf(go_doc.Synopsis)}
	Roots[`/go/doc/ToHTML`] = FuncRoot{Func: reflect.ValueOf(go_doc.ToHTML)}
	Roots[`/go/doc/ToText`] = FuncRoot{Func: reflect.ValueOf(go_doc.ToText)}
	Roots[`/go/parser/ParseDir`] = FuncRoot{Func: reflect.ValueOf(go_parser.ParseDir)}
	Roots[`/go/parser/ParseExpr`] = FuncRoot{Func: reflect.ValueOf(go_parser.ParseExpr)}
	Roots[`/go/parser/ParseFile`] = FuncRoot{Func: reflect.ValueOf(go_parser.ParseFile)}
	Roots[`/go/printer/Fprint`] = FuncRoot{Func: reflect.ValueOf(go_printer.Fprint)}
	Roots[`/go/scanner/PrintError`] = FuncRoot{Func: reflect.ValueOf(go_scanner.PrintError)}
	Roots[`/go/token/Lookup`] = FuncRoot{Func: reflect.ValueOf(go_token.Lookup)}
	Roots[`/go/token/NewFileSet`] = FuncRoot{Func: reflect.ValueOf(go_token.NewFileSet)}
	Roots[`/hash/adler32/Checksum`] = FuncRoot{Func: reflect.ValueOf(hash_adler32.Checksum)}
	Roots[`/hash/adler32/New`] = FuncRoot{Func: reflect.ValueOf(hash_adler32.New)}
	Roots[`/hash/crc32/Checksum`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.Checksum)}
	Roots[`/hash/crc32/ChecksumIEEE`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.ChecksumIEEE)}
	Roots[`/hash/crc32/MakeTable`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.MakeTable)}
	Roots[`/hash/crc32/New`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.New)}
	Roots[`/hash/crc32/NewIEEE`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.NewIEEE)}
	Roots[`/hash/crc32/Update`] = FuncRoot{Func: reflect.ValueOf(hash_crc32.Update)}
	Roots[`/hash/crc64/Checksum`] = FuncRoot{Func: reflect.ValueOf(hash_crc64.Checksum)}
	Roots[`/hash/crc64/MakeTable`] = FuncRoot{Func: reflect.ValueOf(hash_crc64.MakeTable)}
	Roots[`/hash/crc64/New`] = FuncRoot{Func: reflect.ValueOf(hash_crc64.New)}
	Roots[`/hash/crc64/Update`] = FuncRoot{Func: reflect.ValueOf(hash_crc64.Update)}
	Roots[`/hash/fnv/New32`] = FuncRoot{Func: reflect.ValueOf(hash_fnv.New32)}
	Roots[`/hash/fnv/New32a`] = FuncRoot{Func: reflect.ValueOf(hash_fnv.New32a)}
	Roots[`/hash/fnv/New64`] = FuncRoot{Func: reflect.ValueOf(hash_fnv.New64)}
	Roots[`/hash/fnv/New64a`] = FuncRoot{Func: reflect.ValueOf(hash_fnv.New64a)}
	Roots[`/html/EscapeString`] = FuncRoot{Func: reflect.ValueOf(html.EscapeString)}
	Roots[`/html/UnescapeString`] = FuncRoot{Func: reflect.ValueOf(html.UnescapeString)}
	Roots[`/html/template/HTMLEscape`] = FuncRoot{Func: reflect.ValueOf(html_template.HTMLEscape)}
	Roots[`/html/template/HTMLEscapeString`] = FuncRoot{Func: reflect.ValueOf(html_template.HTMLEscapeString)}
	Roots[`/html/template/HTMLEscaper`] = FuncRoot{Func: reflect.ValueOf(html_template.HTMLEscaper)}
	Roots[`/html/template/JSEscape`] = FuncRoot{Func: reflect.ValueOf(html_template.JSEscape)}
	Roots[`/html/template/JSEscapeString`] = FuncRoot{Func: reflect.ValueOf(html_template.JSEscapeString)}
	Roots[`/html/template/JSEscaper`] = FuncRoot{Func: reflect.ValueOf(html_template.JSEscaper)}
	Roots[`/html/template/Must`] = FuncRoot{Func: reflect.ValueOf(html_template.Must)}
	Roots[`/html/template/New`] = FuncRoot{Func: reflect.ValueOf(html_template.New)}
	Roots[`/html/template/ParseFiles`] = FuncRoot{Func: reflect.ValueOf(html_template.ParseFiles)}
	Roots[`/html/template/ParseGlob`] = FuncRoot{Func: reflect.ValueOf(html_template.ParseGlob)}
	Roots[`/html/template/URLQueryEscaper`] = FuncRoot{Func: reflect.ValueOf(html_template.URLQueryEscaper)}
	Roots[`/image/Decode`] = FuncRoot{Func: reflect.ValueOf(image.Decode)}
	Roots[`/image/DecodeConfig`] = FuncRoot{Func: reflect.ValueOf(image.DecodeConfig)}
	Roots[`/image/NewAlpha`] = FuncRoot{Func: reflect.ValueOf(image.NewAlpha)}
	Roots[`/image/NewAlpha16`] = FuncRoot{Func: reflect.ValueOf(image.NewAlpha16)}
	Roots[`/image/NewGray`] = FuncRoot{Func: reflect.ValueOf(image.NewGray)}
	Roots[`/image/NewGray16`] = FuncRoot{Func: reflect.ValueOf(image.NewGray16)}
	Roots[`/image/NewNRGBA`] = FuncRoot{Func: reflect.ValueOf(image.NewNRGBA)}
	Roots[`/image/NewNRGBA64`] = FuncRoot{Func: reflect.ValueOf(image.NewNRGBA64)}
	Roots[`/image/NewPaletted`] = FuncRoot{Func: reflect.ValueOf(image.NewPaletted)}
	Roots[`/image/NewRGBA`] = FuncRoot{Func: reflect.ValueOf(image.NewRGBA)}
	Roots[`/image/NewRGBA64`] = FuncRoot{Func: reflect.ValueOf(image.NewRGBA64)}
	Roots[`/image/NewUniform`] = FuncRoot{Func: reflect.ValueOf(image.NewUniform)}
	Roots[`/image/NewYCbCr`] = FuncRoot{Func: reflect.ValueOf(image.NewYCbCr)}
	Roots[`/image/Pt`] = FuncRoot{Func: reflect.ValueOf(image.Pt)}
	Roots[`/image/Rect`] = FuncRoot{Func: reflect.ValueOf(image.Rect)}
	Roots[`/image/RegisterFormat`] = FuncRoot{Func: reflect.ValueOf(image.RegisterFormat)}
	Roots[`/image/color/ModelFunc`] = FuncRoot{Func: reflect.ValueOf(image_color.ModelFunc)}
	Roots[`/image/color/RGBToYCbCr`] = FuncRoot{Func: reflect.ValueOf(image_color.RGBToYCbCr)}
	Roots[`/image/color/YCbCrToRGB`] = FuncRoot{Func: reflect.ValueOf(image_color.YCbCrToRGB)}
	Roots[`/image/draw/Draw`] = FuncRoot{Func: reflect.ValueOf(image_draw.Draw)}
	Roots[`/image/draw/DrawMask`] = FuncRoot{Func: reflect.ValueOf(image_draw.DrawMask)}
	Roots[`/image/gif/Decode`] = FuncRoot{Func: reflect.ValueOf(image_gif.Decode)}
	Roots[`/image/gif/DecodeAll`] = FuncRoot{Func: reflect.ValueOf(image_gif.DecodeAll)}
	Roots[`/image/gif/DecodeConfig`] = FuncRoot{Func: reflect.ValueOf(image_gif.DecodeConfig)}
	Roots[`/image/jpeg/Decode`] = FuncRoot{Func: reflect.ValueOf(image_jpeg.Decode)}
	Roots[`/image/jpeg/DecodeConfig`] = FuncRoot{Func: reflect.ValueOf(image_jpeg.DecodeConfig)}
	Roots[`/image/jpeg/Encode`] = FuncRoot{Func: reflect.ValueOf(image_jpeg.Encode)}
	Roots[`/image/png/Decode`] = FuncRoot{Func: reflect.ValueOf(image_png.Decode)}
	Roots[`/image/png/DecodeConfig`] = FuncRoot{Func: reflect.ValueOf(image_png.DecodeConfig)}
	Roots[`/image/png/Encode`] = FuncRoot{Func: reflect.ValueOf(image_png.Encode)}
	Roots[`/index/suffixarray/New`] = FuncRoot{Func: reflect.ValueOf(index_suffixarray.New)}
	Roots[`/io/Copy`] = FuncRoot{Func: reflect.ValueOf(io.Copy)}
	Roots[`/io/CopyN`] = FuncRoot{Func: reflect.ValueOf(io.CopyN)}
	Roots[`/io/LimitReader`] = FuncRoot{Func: reflect.ValueOf(io.LimitReader)}
	Roots[`/io/MultiReader`] = FuncRoot{Func: reflect.ValueOf(io.MultiReader)}
	Roots[`/io/MultiWriter`] = FuncRoot{Func: reflect.ValueOf(io.MultiWriter)}
	Roots[`/io/NewSectionReader`] = FuncRoot{Func: reflect.ValueOf(io.NewSectionReader)}
	Roots[`/io/Pipe`] = FuncRoot{Func: reflect.ValueOf(io.Pipe)}
	Roots[`/io/ReadAtLeast`] = FuncRoot{Func: reflect.ValueOf(io.ReadAtLeast)}
	Roots[`/io/ReadFull`] = FuncRoot{Func: reflect.ValueOf(io.ReadFull)}
	Roots[`/io/TeeReader`] = FuncRoot{Func: reflect.ValueOf(io.TeeReader)}
	Roots[`/io/WriteString`] = FuncRoot{Func: reflect.ValueOf(io.WriteString)}
	Roots[`/io/ioutil/NopCloser`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.NopCloser)}
	Roots[`/io/ioutil/ReadAll`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.ReadAll)}
	Roots[`/io/ioutil/ReadDir`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.ReadDir)}
	Roots[`/io/ioutil/ReadFile`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.ReadFile)}
	Roots[`/io/ioutil/TempDir`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.TempDir)}
	Roots[`/io/ioutil/TempFile`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.TempFile)}
	Roots[`/io/ioutil/WriteFile`] = FuncRoot{Func: reflect.ValueOf(io_ioutil.WriteFile)}
	Roots[`/log/Fatal`] = FuncRoot{Func: reflect.ValueOf(log.Fatal)}
	Roots[`/log/Fatalf`] = FuncRoot{Func: reflect.ValueOf(log.Fatalf)}
	Roots[`/log/Fatalln`] = FuncRoot{Func: reflect.ValueOf(log.Fatalln)}
	Roots[`/log/Flags`] = FuncRoot{Func: reflect.ValueOf(log.Flags)}
	Roots[`/log/New`] = FuncRoot{Func: reflect.ValueOf(log.New)}
	Roots[`/log/Panic`] = FuncRoot{Func: reflect.ValueOf(log.Panic)}
	Roots[`/log/Panicf`] = FuncRoot{Func: reflect.ValueOf(log.Panicf)}
	Roots[`/log/Panicln`] = FuncRoot{Func: reflect.ValueOf(log.Panicln)}
	Roots[`/log/Prefix`] = FuncRoot{Func: reflect.ValueOf(log.Prefix)}
	Roots[`/log/Print`] = FuncRoot{Func: reflect.ValueOf(log.Print)}
	Roots[`/log/Printf`] = FuncRoot{Func: reflect.ValueOf(log.Printf)}
	Roots[`/log/Println`] = FuncRoot{Func: reflect.ValueOf(log.Println)}
	Roots[`/log/SetFlags`] = FuncRoot{Func: reflect.ValueOf(log.SetFlags)}
	Roots[`/log/SetOutput`] = FuncRoot{Func: reflect.ValueOf(log.SetOutput)}
	Roots[`/log/SetPrefix`] = FuncRoot{Func: reflect.ValueOf(log.SetPrefix)}
	Roots[`/math/Abs`] = FuncRoot{Func: reflect.ValueOf(math.Abs)}
	Roots[`/math/Acos`] = FuncRoot{Func: reflect.ValueOf(math.Acos)}
	Roots[`/math/Acosh`] = FuncRoot{Func: reflect.ValueOf(math.Acosh)}
	Roots[`/math/Asin`] = FuncRoot{Func: reflect.ValueOf(math.Asin)}
	Roots[`/math/Asinh`] = FuncRoot{Func: reflect.ValueOf(math.Asinh)}
	Roots[`/math/Atan`] = FuncRoot{Func: reflect.ValueOf(math.Atan)}
	Roots[`/math/Atan2`] = FuncRoot{Func: reflect.ValueOf(math.Atan2)}
	Roots[`/math/Atanh`] = FuncRoot{Func: reflect.ValueOf(math.Atanh)}
	Roots[`/math/Cbrt`] = FuncRoot{Func: reflect.ValueOf(math.Cbrt)}
	Roots[`/math/Ceil`] = FuncRoot{Func: reflect.ValueOf(math.Ceil)}
	Roots[`/math/Copysign`] = FuncRoot{Func: reflect.ValueOf(math.Copysign)}
	Roots[`/math/Cos`] = FuncRoot{Func: reflect.ValueOf(math.Cos)}
	Roots[`/math/Cosh`] = FuncRoot{Func: reflect.ValueOf(math.Cosh)}
	Roots[`/math/Dim`] = FuncRoot{Func: reflect.ValueOf(math.Dim)}
	Roots[`/math/Erf`] = FuncRoot{Func: reflect.ValueOf(math.Erf)}
	Roots[`/math/Erfc`] = FuncRoot{Func: reflect.ValueOf(math.Erfc)}
	Roots[`/math/Exp`] = FuncRoot{Func: reflect.ValueOf(math.Exp)}
	Roots[`/math/Exp2`] = FuncRoot{Func: reflect.ValueOf(math.Exp2)}
	Roots[`/math/Expm1`] = FuncRoot{Func: reflect.ValueOf(math.Expm1)}
	Roots[`/math/Float32bits`] = FuncRoot{Func: reflect.ValueOf(math.Float32bits)}
	Roots[`/math/Float32frombits`] = FuncRoot{Func: reflect.ValueOf(math.Float32frombits)}
	Roots[`/math/Float64bits`] = FuncRoot{Func: reflect.ValueOf(math.Float64bits)}
	Roots[`/math/Float64frombits`] = FuncRoot{Func: reflect.ValueOf(math.Float64frombits)}
	Roots[`/math/Floor`] = FuncRoot{Func: reflect.ValueOf(math.Floor)}
	Roots[`/math/Frexp`] = FuncRoot{Func: reflect.ValueOf(math.Frexp)}
	Roots[`/math/Gamma`] = FuncRoot{Func: reflect.ValueOf(math.Gamma)}
	Roots[`/math/Hypot`] = FuncRoot{Func: reflect.ValueOf(math.Hypot)}
	Roots[`/math/Ilogb`] = FuncRoot{Func: reflect.ValueOf(math.Ilogb)}
	Roots[`/math/Inf`] = FuncRoot{Func: reflect.ValueOf(math.Inf)}
	Roots[`/math/IsInf`] = FuncRoot{Func: reflect.ValueOf(math.IsInf)}
	Roots[`/math/IsNaN`] = FuncRoot{Func: reflect.ValueOf(math.IsNaN)}
	Roots[`/math/J0`] = FuncRoot{Func: reflect.ValueOf(math.J0)}
	Roots[`/math/J1`] = FuncRoot{Func: reflect.ValueOf(math.J1)}
	Roots[`/math/Jn`] = FuncRoot{Func: reflect.ValueOf(math.Jn)}
	Roots[`/math/Ldexp`] = FuncRoot{Func: reflect.ValueOf(math.Ldexp)}
	Roots[`/math/Lgamma`] = FuncRoot{Func: reflect.ValueOf(math.Lgamma)}
	Roots[`/math/Log`] = FuncRoot{Func: reflect.ValueOf(math.Log)}
	Roots[`/math/Log10`] = FuncRoot{Func: reflect.ValueOf(math.Log10)}
	Roots[`/math/Log1p`] = FuncRoot{Func: reflect.ValueOf(math.Log1p)}
	Roots[`/math/Log2`] = FuncRoot{Func: reflect.ValueOf(math.Log2)}
	Roots[`/math/Logb`] = FuncRoot{Func: reflect.ValueOf(math.Logb)}
	Roots[`/math/Max`] = FuncRoot{Func: reflect.ValueOf(math.Max)}
	Roots[`/math/Min`] = FuncRoot{Func: reflect.ValueOf(math.Min)}
	Roots[`/math/Mod`] = FuncRoot{Func: reflect.ValueOf(math.Mod)}
	Roots[`/math/Modf`] = FuncRoot{Func: reflect.ValueOf(math.Modf)}
	Roots[`/math/NaN`] = FuncRoot{Func: reflect.ValueOf(math.NaN)}
	Roots[`/math/Nextafter`] = FuncRoot{Func: reflect.ValueOf(math.Nextafter)}
	Roots[`/math/Pow`] = FuncRoot{Func: reflect.ValueOf(math.Pow)}
	Roots[`/math/Pow10`] = FuncRoot{Func: reflect.ValueOf(math.Pow10)}
	Roots[`/math/Remainder`] = FuncRoot{Func: reflect.ValueOf(math.Remainder)}
	Roots[`/math/Signbit`] = FuncRoot{Func: reflect.ValueOf(math.Signbit)}
	Roots[`/math/Sin`] = FuncRoot{Func: reflect.ValueOf(math.Sin)}
	Roots[`/math/Sincos`] = FuncRoot{Func: reflect.ValueOf(math.Sincos)}
	Roots[`/math/Sinh`] = FuncRoot{Func: reflect.ValueOf(math.Sinh)}
	Roots[`/math/Sqrt`] = FuncRoot{Func: reflect.ValueOf(math.Sqrt)}
	Roots[`/math/Tan`] = FuncRoot{Func: reflect.ValueOf(math.Tan)}
	Roots[`/math/Tanh`] = FuncRoot{Func: reflect.ValueOf(math.Tanh)}
	Roots[`/math/Trunc`] = FuncRoot{Func: reflect.ValueOf(math.Trunc)}
	Roots[`/math/Y0`] = FuncRoot{Func: reflect.ValueOf(math.Y0)}
	Roots[`/math/Y1`] = FuncRoot{Func: reflect.ValueOf(math.Y1)}
	Roots[`/math/Yn`] = FuncRoot{Func: reflect.ValueOf(math.Yn)}
	Roots[`/math/big/NewInt`] = FuncRoot{Func: reflect.ValueOf(math_big.NewInt)}
	Roots[`/math/big/NewRat`] = FuncRoot{Func: reflect.ValueOf(math_big.NewRat)}
	Roots[`/math/cmplx/Abs`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Abs)}
	Roots[`/math/cmplx/Acos`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Acos)}
	Roots[`/math/cmplx/Acosh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Acosh)}
	Roots[`/math/cmplx/Asin`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Asin)}
	Roots[`/math/cmplx/Asinh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Asinh)}
	Roots[`/math/cmplx/Atan`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Atan)}
	Roots[`/math/cmplx/Atanh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Atanh)}
	Roots[`/math/cmplx/Conj`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Conj)}
	Roots[`/math/cmplx/Cos`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Cos)}
	Roots[`/math/cmplx/Cosh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Cosh)}
	Roots[`/math/cmplx/Cot`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Cot)}
	Roots[`/math/cmplx/Exp`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Exp)}
	Roots[`/math/cmplx/Inf`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Inf)}
	Roots[`/math/cmplx/IsInf`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.IsInf)}
	Roots[`/math/cmplx/IsNaN`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.IsNaN)}
	Roots[`/math/cmplx/Log`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Log)}
	Roots[`/math/cmplx/Log10`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Log10)}
	Roots[`/math/cmplx/NaN`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.NaN)}
	Roots[`/math/cmplx/Phase`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Phase)}
	Roots[`/math/cmplx/Polar`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Polar)}
	Roots[`/math/cmplx/Pow`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Pow)}
	Roots[`/math/cmplx/Rect`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Rect)}
	Roots[`/math/cmplx/Sin`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Sin)}
	Roots[`/math/cmplx/Sinh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Sinh)}
	Roots[`/math/cmplx/Sqrt`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Sqrt)}
	Roots[`/math/cmplx/Tan`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Tan)}
	Roots[`/math/cmplx/Tanh`] = FuncRoot{Func: reflect.ValueOf(math_cmplx.Tanh)}
	Roots[`/math/rand/ExpFloat64`] = FuncRoot{Func: reflect.ValueOf(math_rand.ExpFloat64)}
	Roots[`/math/rand/Float32`] = FuncRoot{Func: reflect.ValueOf(math_rand.Float32)}
	Roots[`/math/rand/Float64`] = FuncRoot{Func: reflect.ValueOf(math_rand.Float64)}
	Roots[`/math/rand/Int`] = FuncRoot{Func: reflect.ValueOf(math_rand.Int)}
	Roots[`/math/rand/Int31`] = FuncRoot{Func: reflect.ValueOf(math_rand.Int31)}
	Roots[`/math/rand/Int31n`] = FuncRoot{Func: reflect.ValueOf(math_rand.Int31n)}
	Roots[`/math/rand/Int63`] = FuncRoot{Func: reflect.ValueOf(math_rand.Int63)}
	Roots[`/math/rand/Int63n`] = FuncRoot{Func: reflect.ValueOf(math_rand.Int63n)}
	Roots[`/math/rand/Intn`] = FuncRoot{Func: reflect.ValueOf(math_rand.Intn)}
	Roots[`/math/rand/New`] = FuncRoot{Func: reflect.ValueOf(math_rand.New)}
	Roots[`/math/rand/NewSource`] = FuncRoot{Func: reflect.ValueOf(math_rand.NewSource)}
	Roots[`/math/rand/NewZipf`] = FuncRoot{Func: reflect.ValueOf(math_rand.NewZipf)}
	Roots[`/math/rand/NormFloat64`] = FuncRoot{Func: reflect.ValueOf(math_rand.NormFloat64)}
	Roots[`/math/rand/Perm`] = FuncRoot{Func: reflect.ValueOf(math_rand.Perm)}
	Roots[`/math/rand/Seed`] = FuncRoot{Func: reflect.ValueOf(math_rand.Seed)}
	Roots[`/math/rand/Uint32`] = FuncRoot{Func: reflect.ValueOf(math_rand.Uint32)}
	Roots[`/mime/AddExtensionType`] = FuncRoot{Func: reflect.ValueOf(mime.AddExtensionType)}
	Roots[`/mime/FormatMediaType`] = FuncRoot{Func: reflect.ValueOf(mime.FormatMediaType)}
	Roots[`/mime/ParseMediaType`] = FuncRoot{Func: reflect.ValueOf(mime.ParseMediaType)}
	Roots[`/mime/TypeByExtension`] = FuncRoot{Func: reflect.ValueOf(mime.TypeByExtension)}
	Roots[`/mime/multipart/NewReader`] = FuncRoot{Func: reflect.ValueOf(mime_multipart.NewReader)}
	Roots[`/mime/multipart/NewWriter`] = FuncRoot{Func: reflect.ValueOf(mime_multipart.NewWriter)}
	Roots[`/net/CIDRMask`] = FuncRoot{Func: reflect.ValueOf(net.CIDRMask)}
	Roots[`/net/Dial`] = FuncRoot{Func: reflect.ValueOf(net.Dial)}
	Roots[`/net/DialIP`] = FuncRoot{Func: reflect.ValueOf(net.DialIP)}
	Roots[`/net/DialTCP`] = FuncRoot{Func: reflect.ValueOf(net.DialTCP)}
	Roots[`/net/DialTimeout`] = FuncRoot{Func: reflect.ValueOf(net.DialTimeout)}
	Roots[`/net/DialUDP`] = FuncRoot{Func: reflect.ValueOf(net.DialUDP)}
	Roots[`/net/DialUnix`] = FuncRoot{Func: reflect.ValueOf(net.DialUnix)}
	Roots[`/net/FileConn`] = FuncRoot{Func: reflect.ValueOf(net.FileConn)}
	Roots[`/net/FileListener`] = FuncRoot{Func: reflect.ValueOf(net.FileListener)}
	Roots[`/net/FilePacketConn`] = FuncRoot{Func: reflect.ValueOf(net.FilePacketConn)}
	Roots[`/net/IPv4`] = FuncRoot{Func: reflect.ValueOf(net.IPv4)}
	Roots[`/net/IPv4Mask`] = FuncRoot{Func: reflect.ValueOf(net.IPv4Mask)}
	Roots[`/net/InterfaceAddrs`] = FuncRoot{Func: reflect.ValueOf(net.InterfaceAddrs)}
	Roots[`/net/InterfaceByIndex`] = FuncRoot{Func: reflect.ValueOf(net.InterfaceByIndex)}
	Roots[`/net/InterfaceByName`] = FuncRoot{Func: reflect.ValueOf(net.InterfaceByName)}
	Roots[`/net/Interfaces`] = FuncRoot{Func: reflect.ValueOf(net.Interfaces)}
	Roots[`/net/JoinHostPort`] = FuncRoot{Func: reflect.ValueOf(net.JoinHostPort)}
	Roots[`/net/Listen`] = FuncRoot{Func: reflect.ValueOf(net.Listen)}
	Roots[`/net/ListenIP`] = FuncRoot{Func: reflect.ValueOf(net.ListenIP)}
	Roots[`/net/ListenMulticastUDP`] = FuncRoot{Func: reflect.ValueOf(net.ListenMulticastUDP)}
	Roots[`/net/ListenPacket`] = FuncRoot{Func: reflect.ValueOf(net.ListenPacket)}
	Roots[`/net/ListenTCP`] = FuncRoot{Func: reflect.ValueOf(net.ListenTCP)}
	Roots[`/net/ListenUDP`] = FuncRoot{Func: reflect.ValueOf(net.ListenUDP)}
	Roots[`/net/ListenUnix`] = FuncRoot{Func: reflect.ValueOf(net.ListenUnix)}
	Roots[`/net/ListenUnixgram`] = FuncRoot{Func: reflect.ValueOf(net.ListenUnixgram)}
	Roots[`/net/LookupAddr`] = FuncRoot{Func: reflect.ValueOf(net.LookupAddr)}
	Roots[`/net/LookupCNAME`] = FuncRoot{Func: reflect.ValueOf(net.LookupCNAME)}
	Roots[`/net/LookupHost`] = FuncRoot{Func: reflect.ValueOf(net.LookupHost)}
	Roots[`/net/LookupIP`] = FuncRoot{Func: reflect.ValueOf(net.LookupIP)}
	Roots[`/net/LookupMX`] = FuncRoot{Func: reflect.ValueOf(net.LookupMX)}
	Roots[`/net/LookupPort`] = FuncRoot{Func: reflect.ValueOf(net.LookupPort)}
	Roots[`/net/LookupSRV`] = FuncRoot{Func: reflect.ValueOf(net.LookupSRV)}
	Roots[`/net/LookupTXT`] = FuncRoot{Func: reflect.ValueOf(net.LookupTXT)}
	Roots[`/net/ParseCIDR`] = FuncRoot{Func: reflect.ValueOf(net.ParseCIDR)}
	Roots[`/net/ParseIP`] = FuncRoot{Func: reflect.ValueOf(net.ParseIP)}
	Roots[`/net/ParseMAC`] = FuncRoot{Func: reflect.ValueOf(net.ParseMAC)}
	Roots[`/net/Pipe`] = FuncRoot{Func: reflect.ValueOf(net.Pipe)}
	Roots[`/net/ResolveIPAddr`] = FuncRoot{Func: reflect.ValueOf(net.ResolveIPAddr)}
	Roots[`/net/ResolveTCPAddr`] = FuncRoot{Func: reflect.ValueOf(net.ResolveTCPAddr)}
	Roots[`/net/ResolveUDPAddr`] = FuncRoot{Func: reflect.ValueOf(net.ResolveUDPAddr)}
	Roots[`/net/ResolveUnixAddr`] = FuncRoot{Func: reflect.ValueOf(net.ResolveUnixAddr)}
	Roots[`/net/SplitHostPort`] = FuncRoot{Func: reflect.ValueOf(net.SplitHostPort)}
	Roots[`/net/http/CanonicalHeaderKey`] = FuncRoot{Func: reflect.ValueOf(net_http.CanonicalHeaderKey)}
	Roots[`/net/http/DetectContentType`] = FuncRoot{Func: reflect.ValueOf(net_http.DetectContentType)}
	Roots[`/net/http/Error`] = FuncRoot{Func: reflect.ValueOf(net_http.Error)}
	Roots[`/net/http/FileServer`] = FuncRoot{Func: reflect.ValueOf(net_http.FileServer)}
	Roots[`/net/http/Get`] = FuncRoot{Func: reflect.ValueOf(net_http.Get)}
	Roots[`/net/http/Handle`] = FuncRoot{Func: reflect.ValueOf(net_http.Handle)}
	Roots[`/net/http/HandleFunc`] = FuncRoot{Func: reflect.ValueOf(net_http.HandleFunc)}
	Roots[`/net/http/Head`] = FuncRoot{Func: reflect.ValueOf(net_http.Head)}
	Roots[`/net/http/ListenAndServe`] = FuncRoot{Func: reflect.ValueOf(net_http.ListenAndServe)}
	Roots[`/net/http/ListenAndServeTLS`] = FuncRoot{Func: reflect.ValueOf(net_http.ListenAndServeTLS)}
	Roots[`/net/http/MaxBytesReader`] = FuncRoot{Func: reflect.ValueOf(net_http.MaxBytesReader)}
	Roots[`/net/http/NewFileTransport`] = FuncRoot{Func: reflect.ValueOf(net_http.NewFileTransport)}
	Roots[`/net/http/NewRequest`] = FuncRoot{Func: reflect.ValueOf(net_http.NewRequest)}
	Roots[`/net/http/NewServeMux`] = FuncRoot{Func: reflect.ValueOf(net_http.NewServeMux)}
	Roots[`/net/http/NotFound`] = FuncRoot{Func: reflect.ValueOf(net_http.NotFound)}
	Roots[`/net/http/NotFoundHandler`] = FuncRoot{Func: reflect.ValueOf(net_http.NotFoundHandler)}
	Roots[`/net/http/ParseHTTPVersion`] = FuncRoot{Func: reflect.ValueOf(net_http.ParseHTTPVersion)}
	Roots[`/net/http/Post`] = FuncRoot{Func: reflect.ValueOf(net_http.Post)}
	Roots[`/net/http/PostForm`] = FuncRoot{Func: reflect.ValueOf(net_http.PostForm)}
	Roots[`/net/http/ProxyFromEnvironment`] = FuncRoot{Func: reflect.ValueOf(net_http.ProxyFromEnvironment)}
	Roots[`/net/http/ProxyURL`] = FuncRoot{Func: reflect.ValueOf(net_http.ProxyURL)}
	Roots[`/net/http/ReadRequest`] = FuncRoot{Func: reflect.ValueOf(net_http.ReadRequest)}
	Roots[`/net/http/ReadResponse`] = FuncRoot{Func: reflect.ValueOf(net_http.ReadResponse)}
	Roots[`/net/http/Redirect`] = FuncRoot{Func: reflect.ValueOf(net_http.Redirect)}
	Roots[`/net/http/RedirectHandler`] = FuncRoot{Func: reflect.ValueOf(net_http.RedirectHandler)}
	Roots[`/net/http/Serve`] = FuncRoot{Func: reflect.ValueOf(net_http.Serve)}
	Roots[`/net/http/ServeContent`] = FuncRoot{Func: reflect.ValueOf(net_http.ServeContent)}
	Roots[`/net/http/ServeFile`] = FuncRoot{Func: reflect.ValueOf(net_http.ServeFile)}
	Roots[`/net/http/SetCookie`] = FuncRoot{Func: reflect.ValueOf(net_http.SetCookie)}
	Roots[`/net/http/StatusText`] = FuncRoot{Func: reflect.ValueOf(net_http.StatusText)}
	Roots[`/net/http/StripPrefix`] = FuncRoot{Func: reflect.ValueOf(net_http.StripPrefix)}
	Roots[`/net/http/TimeoutHandler`] = FuncRoot{Func: reflect.ValueOf(net_http.TimeoutHandler)}
	Roots[`/net/http/cgi/Request`] = FuncRoot{Func: reflect.ValueOf(net_http_cgi.Request)}
	Roots[`/net/http/cgi/RequestFromMap`] = FuncRoot{Func: reflect.ValueOf(net_http_cgi.RequestFromMap)}
	Roots[`/net/http/cgi/Serve`] = FuncRoot{Func: reflect.ValueOf(net_http_cgi.Serve)}
	Roots[`/net/http/fcgi/Serve`] = FuncRoot{Func: reflect.ValueOf(net_http_fcgi.Serve)}
	Roots[`/net/http/httptest/NewRecorder`] = FuncRoot{Func: reflect.ValueOf(net_http_httptest.NewRecorder)}
	Roots[`/net/http/httptest/NewServer`] = FuncRoot{Func: reflect.ValueOf(net_http_httptest.NewServer)}
	Roots[`/net/http/httptest/NewTLSServer`] = FuncRoot{Func: reflect.ValueOf(net_http_httptest.NewTLSServer)}
	Roots[`/net/http/httptest/NewUnstartedServer`] = FuncRoot{Func: reflect.ValueOf(net_http_httptest.NewUnstartedServer)}
	Roots[`/net/http/httputil/DumpRequest`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.DumpRequest)}
	Roots[`/net/http/httputil/DumpRequestOut`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.DumpRequestOut)}
	Roots[`/net/http/httputil/DumpResponse`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.DumpResponse)}
	Roots[`/net/http/httputil/NewChunkedReader`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewChunkedReader)}
	Roots[`/net/http/httputil/NewChunkedWriter`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewChunkedWriter)}
	Roots[`/net/http/httputil/NewClientConn`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewClientConn)}
	Roots[`/net/http/httputil/NewProxyClientConn`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewProxyClientConn)}
	Roots[`/net/http/httputil/NewServerConn`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewServerConn)}
	Roots[`/net/http/httputil/NewSingleHostReverseProxy`] = FuncRoot{Func: reflect.ValueOf(net_http_httputil.NewSingleHostReverseProxy)}
	Roots[`/net/http/pprof/Cmdline`] = FuncRoot{Func: reflect.ValueOf(net_http_pprof.Cmdline)}
	Roots[`/net/http/pprof/Handler`] = FuncRoot{Func: reflect.ValueOf(net_http_pprof.Handler)}
	Roots[`/net/http/pprof/Index`] = FuncRoot{Func: reflect.ValueOf(net_http_pprof.Index)}
	Roots[`/net/http/pprof/Profile`] = FuncRoot{Func: reflect.ValueOf(net_http_pprof.Profile)}
	Roots[`/net/http/pprof/Symbol`] = FuncRoot{Func: reflect.ValueOf(net_http_pprof.Symbol)}
	Roots[`/net/mail/ReadMessage`] = FuncRoot{Func: reflect.ValueOf(net_mail.ReadMessage)}
	Roots[`/net/rpc/Accept`] = FuncRoot{Func: reflect.ValueOf(net_rpc.Accept)}
	Roots[`/net/rpc/Dial`] = FuncRoot{Func: reflect.ValueOf(net_rpc.Dial)}
	Roots[`/net/rpc/DialHTTP`] = FuncRoot{Func: reflect.ValueOf(net_rpc.DialHTTP)}
	Roots[`/net/rpc/DialHTTPPath`] = FuncRoot{Func: reflect.ValueOf(net_rpc.DialHTTPPath)}
	Roots[`/net/rpc/HandleHTTP`] = FuncRoot{Func: reflect.ValueOf(net_rpc.HandleHTTP)}
	Roots[`/net/rpc/NewClient`] = FuncRoot{Func: reflect.ValueOf(net_rpc.NewClient)}
	Roots[`/net/rpc/NewClientWithCodec`] = FuncRoot{Func: reflect.ValueOf(net_rpc.NewClientWithCodec)}
	Roots[`/net/rpc/NewServer`] = FuncRoot{Func: reflect.ValueOf(net_rpc.NewServer)}
	Roots[`/net/rpc/Register`] = FuncRoot{Func: reflect.ValueOf(net_rpc.Register)}
	Roots[`/net/rpc/RegisterName`] = FuncRoot{Func: reflect.ValueOf(net_rpc.RegisterName)}
	Roots[`/net/rpc/ServeCodec`] = FuncRoot{Func: reflect.ValueOf(net_rpc.ServeCodec)}
	Roots[`/net/rpc/ServeConn`] = FuncRoot{Func: reflect.ValueOf(net_rpc.ServeConn)}
	Roots[`/net/rpc/ServeRequest`] = FuncRoot{Func: reflect.ValueOf(net_rpc.ServeRequest)}
	Roots[`/net/rpc/jsonrpc/Dial`] = FuncRoot{Func: reflect.ValueOf(net_rpc_jsonrpc.Dial)}
	Roots[`/net/rpc/jsonrpc/NewClient`] = FuncRoot{Func: reflect.ValueOf(net_rpc_jsonrpc.NewClient)}
	Roots[`/net/rpc/jsonrpc/NewClientCodec`] = FuncRoot{Func: reflect.ValueOf(net_rpc_jsonrpc.NewClientCodec)}
	Roots[`/net/rpc/jsonrpc/NewServerCodec`] = FuncRoot{Func: reflect.ValueOf(net_rpc_jsonrpc.NewServerCodec)}
	Roots[`/net/rpc/jsonrpc/ServeConn`] = FuncRoot{Func: reflect.ValueOf(net_rpc_jsonrpc.ServeConn)}
	Roots[`/net/smtp/CRAMMD5Auth`] = FuncRoot{Func: reflect.ValueOf(net_smtp.CRAMMD5Auth)}
	Roots[`/net/smtp/Dial`] = FuncRoot{Func: reflect.ValueOf(net_smtp.Dial)}
	Roots[`/net/smtp/NewClient`] = FuncRoot{Func: reflect.ValueOf(net_smtp.NewClient)}
	Roots[`/net/smtp/PlainAuth`] = FuncRoot{Func: reflect.ValueOf(net_smtp.PlainAuth)}
	Roots[`/net/smtp/SendMail`] = FuncRoot{Func: reflect.ValueOf(net_smtp.SendMail)}
	Roots[`/net/textproto/CanonicalMIMEHeaderKey`] = FuncRoot{Func: reflect.ValueOf(net_textproto.CanonicalMIMEHeaderKey)}
	Roots[`/net/textproto/Dial`] = FuncRoot{Func: reflect.ValueOf(net_textproto.Dial)}
	Roots[`/net/textproto/NewConn`] = FuncRoot{Func: reflect.ValueOf(net_textproto.NewConn)}
	Roots[`/net/textproto/NewReader`] = FuncRoot{Func: reflect.ValueOf(net_textproto.NewReader)}
	Roots[`/net/textproto/NewWriter`] = FuncRoot{Func: reflect.ValueOf(net_textproto.NewWriter)}
	Roots[`/net/url/Parse`] = FuncRoot{Func: reflect.ValueOf(net_url.Parse)}
	Roots[`/net/url/ParseQuery`] = FuncRoot{Func: reflect.ValueOf(net_url.ParseQuery)}
	Roots[`/net/url/ParseRequestURI`] = FuncRoot{Func: reflect.ValueOf(net_url.ParseRequestURI)}
	Roots[`/net/url/QueryEscape`] = FuncRoot{Func: reflect.ValueOf(net_url.QueryEscape)}
	Roots[`/net/url/QueryUnescape`] = FuncRoot{Func: reflect.ValueOf(net_url.QueryUnescape)}
	Roots[`/net/url/User`] = FuncRoot{Func: reflect.ValueOf(net_url.User)}
	Roots[`/net/url/UserPassword`] = FuncRoot{Func: reflect.ValueOf(net_url.UserPassword)}
	Roots[`/os/Chdir`] = FuncRoot{Func: reflect.ValueOf(os.Chdir)}
	Roots[`/os/Chmod`] = FuncRoot{Func: reflect.ValueOf(os.Chmod)}
	Roots[`/os/Chown`] = FuncRoot{Func: reflect.ValueOf(os.Chown)}
	Roots[`/os/Chtimes`] = FuncRoot{Func: reflect.ValueOf(os.Chtimes)}
	Roots[`/os/Clearenv`] = FuncRoot{Func: reflect.ValueOf(os.Clearenv)}
	Roots[`/os/Create`] = FuncRoot{Func: reflect.ValueOf(os.Create)}
	Roots[`/os/Environ`] = FuncRoot{Func: reflect.ValueOf(os.Environ)}
	Roots[`/os/Exit`] = FuncRoot{Func: reflect.ValueOf(os.Exit)}
	Roots[`/os/Expand`] = FuncRoot{Func: reflect.ValueOf(os.Expand)}
	Roots[`/os/ExpandEnv`] = FuncRoot{Func: reflect.ValueOf(os.ExpandEnv)}
	Roots[`/os/FindProcess`] = FuncRoot{Func: reflect.ValueOf(os.FindProcess)}
	Roots[`/os/Getegid`] = FuncRoot{Func: reflect.ValueOf(os.Getegid)}
	Roots[`/os/Getenv`] = FuncRoot{Func: reflect.ValueOf(os.Getenv)}
	Roots[`/os/Geteuid`] = FuncRoot{Func: reflect.ValueOf(os.Geteuid)}
	Roots[`/os/Getgid`] = FuncRoot{Func: reflect.ValueOf(os.Getgid)}
	Roots[`/os/Getgroups`] = FuncRoot{Func: reflect.ValueOf(os.Getgroups)}
	Roots[`/os/Getpagesize`] = FuncRoot{Func: reflect.ValueOf(os.Getpagesize)}
	Roots[`/os/Getpid`] = FuncRoot{Func: reflect.ValueOf(os.Getpid)}
	Roots[`/os/Getppid`] = FuncRoot{Func: reflect.ValueOf(os.Getppid)}
	Roots[`/os/Getuid`] = FuncRoot{Func: reflect.ValueOf(os.Getuid)}
	Roots[`/os/Getwd`] = FuncRoot{Func: reflect.ValueOf(os.Getwd)}
	Roots[`/os/Hostname`] = FuncRoot{Func: reflect.ValueOf(os.Hostname)}
	Roots[`/os/IsExist`] = FuncRoot{Func: reflect.ValueOf(os.IsExist)}
	Roots[`/os/IsNotExist`] = FuncRoot{Func: reflect.ValueOf(os.IsNotExist)}
	Roots[`/os/IsPathSeparator`] = FuncRoot{Func: reflect.ValueOf(os.IsPathSeparator)}
	Roots[`/os/IsPermission`] = FuncRoot{Func: reflect.ValueOf(os.IsPermission)}
	Roots[`/os/Lchown`] = FuncRoot{Func: reflect.ValueOf(os.Lchown)}
	Roots[`/os/Link`] = FuncRoot{Func: reflect.ValueOf(os.Link)}
	Roots[`/os/Lstat`] = FuncRoot{Func: reflect.ValueOf(os.Lstat)}
	Roots[`/os/Mkdir`] = FuncRoot{Func: reflect.ValueOf(os.Mkdir)}
	Roots[`/os/MkdirAll`] = FuncRoot{Func: reflect.ValueOf(os.MkdirAll)}
	Roots[`/os/NewFile`] = FuncRoot{Func: reflect.ValueOf(os.NewFile)}
	Roots[`/os/NewSyscallError`] = FuncRoot{Func: reflect.ValueOf(os.NewSyscallError)}
	Roots[`/os/Open`] = FuncRoot{Func: reflect.ValueOf(os.Open)}
	Roots[`/os/OpenFile`] = FuncRoot{Func: reflect.ValueOf(os.OpenFile)}
	Roots[`/os/Pipe`] = FuncRoot{Func: reflect.ValueOf(os.Pipe)}
	Roots[`/os/Readlink`] = FuncRoot{Func: reflect.ValueOf(os.Readlink)}
	Roots[`/os/Remove`] = FuncRoot{Func: reflect.ValueOf(os.Remove)}
	Roots[`/os/RemoveAll`] = FuncRoot{Func: reflect.ValueOf(os.RemoveAll)}
	Roots[`/os/Rename`] = FuncRoot{Func: reflect.ValueOf(os.Rename)}
	Roots[`/os/SameFile`] = FuncRoot{Func: reflect.ValueOf(os.SameFile)}
	Roots[`/os/Setenv`] = FuncRoot{Func: reflect.ValueOf(os.Setenv)}
	Roots[`/os/StartProcess`] = FuncRoot{Func: reflect.ValueOf(os.StartProcess)}
	Roots[`/os/Stat`] = FuncRoot{Func: reflect.ValueOf(os.Stat)}
	Roots[`/os/Symlink`] = FuncRoot{Func: reflect.ValueOf(os.Symlink)}
	Roots[`/os/TempDir`] = FuncRoot{Func: reflect.ValueOf(os.TempDir)}
	Roots[`/os/Truncate`] = FuncRoot{Func: reflect.ValueOf(os.Truncate)}
	Roots[`/os/exec/Command`] = FuncRoot{Func: reflect.ValueOf(os_exec.Command)}
	Roots[`/os/exec/LookPath`] = FuncRoot{Func: reflect.ValueOf(os_exec.LookPath)}
	Roots[`/os/signal/Notify`] = FuncRoot{Func: reflect.ValueOf(os_signal.Notify)}
	Roots[`/os/user/Current`] = FuncRoot{Func: reflect.ValueOf(os_user.Current)}
	Roots[`/os/user/Lookup`] = FuncRoot{Func: reflect.ValueOf(os_user.Lookup)}
	Roots[`/os/user/LookupId`] = FuncRoot{Func: reflect.ValueOf(os_user.LookupId)}
	Roots[`/path/Base`] = FuncRoot{Func: reflect.ValueOf(path.Base)}
	Roots[`/path/Clean`] = FuncRoot{Func: reflect.ValueOf(path.Clean)}
	Roots[`/path/Dir`] = FuncRoot{Func: reflect.ValueOf(path.Dir)}
	Roots[`/path/Ext`] = FuncRoot{Func: reflect.ValueOf(path.Ext)}
	Roots[`/path/IsAbs`] = FuncRoot{Func: reflect.ValueOf(path.IsAbs)}
	Roots[`/path/Join`] = FuncRoot{Func: reflect.ValueOf(path.Join)}
	Roots[`/path/Match`] = FuncRoot{Func: reflect.ValueOf(path.Match)}
	Roots[`/path/Split`] = FuncRoot{Func: reflect.ValueOf(path.Split)}
	Roots[`/path/filepath/Abs`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Abs)}
	Roots[`/path/filepath/Base`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Base)}
	Roots[`/path/filepath/Clean`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Clean)}
	Roots[`/path/filepath/Dir`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Dir)}
	Roots[`/path/filepath/EvalSymlinks`] = FuncRoot{Func: reflect.ValueOf(path_filepath.EvalSymlinks)}
	Roots[`/path/filepath/Ext`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Ext)}
	Roots[`/path/filepath/FromSlash`] = FuncRoot{Func: reflect.ValueOf(path_filepath.FromSlash)}
	Roots[`/path/filepath/Glob`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Glob)}
	Roots[`/path/filepath/HasPrefix`] = FuncRoot{Func: reflect.ValueOf(path_filepath.HasPrefix)}
	Roots[`/path/filepath/IsAbs`] = FuncRoot{Func: reflect.ValueOf(path_filepath.IsAbs)}
	Roots[`/path/filepath/Join`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Join)}
	Roots[`/path/filepath/Match`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Match)}
	Roots[`/path/filepath/Rel`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Rel)}
	Roots[`/path/filepath/Split`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Split)}
	Roots[`/path/filepath/SplitList`] = FuncRoot{Func: reflect.ValueOf(path_filepath.SplitList)}
	Roots[`/path/filepath/ToSlash`] = FuncRoot{Func: reflect.ValueOf(path_filepath.ToSlash)}
	Roots[`/path/filepath/VolumeName`] = FuncRoot{Func: reflect.ValueOf(path_filepath.VolumeName)}
	Roots[`/path/filepath/Walk`] = FuncRoot{Func: reflect.ValueOf(path_filepath.Walk)}
	Roots[`/reflect/Append`] = FuncRoot{Func: reflect.ValueOf(reflect.Append)}
	Roots[`/reflect/AppendSlice`] = FuncRoot{Func: reflect.ValueOf(reflect.AppendSlice)}
	Roots[`/reflect/Copy`] = FuncRoot{Func: reflect.ValueOf(reflect.Copy)}
	Roots[`/reflect/DeepEqual`] = FuncRoot{Func: reflect.ValueOf(reflect.DeepEqual)}
	Roots[`/reflect/Indirect`] = FuncRoot{Func: reflect.ValueOf(reflect.Indirect)}
	Roots[`/reflect/MakeChan`] = FuncRoot{Func: reflect.ValueOf(reflect.MakeChan)}
	Roots[`/reflect/MakeMap`] = FuncRoot{Func: reflect.ValueOf(reflect.MakeMap)}
	Roots[`/reflect/MakeSlice`] = FuncRoot{Func: reflect.ValueOf(reflect.MakeSlice)}
	Roots[`/reflect/New`] = FuncRoot{Func: reflect.ValueOf(reflect.New)}
	Roots[`/reflect/NewAt`] = FuncRoot{Func: reflect.ValueOf(reflect.NewAt)}
	Roots[`/reflect/PtrTo`] = FuncRoot{Func: reflect.ValueOf(reflect.PtrTo)}
	Roots[`/reflect/TypeOf`] = FuncRoot{Func: reflect.ValueOf(reflect.TypeOf)}
	Roots[`/reflect/ValueOf`] = FuncRoot{Func: reflect.ValueOf(reflect.ValueOf)}
	Roots[`/reflect/Zero`] = FuncRoot{Func: reflect.ValueOf(reflect.Zero)}
	Roots[`/regexp/Compile`] = FuncRoot{Func: reflect.ValueOf(regexp.Compile)}
	Roots[`/regexp/CompilePOSIX`] = FuncRoot{Func: reflect.ValueOf(regexp.CompilePOSIX)}
	Roots[`/regexp/Match`] = FuncRoot{Func: reflect.ValueOf(regexp.Match)}
	Roots[`/regexp/MatchReader`] = FuncRoot{Func: reflect.ValueOf(regexp.MatchReader)}
	Roots[`/regexp/MatchString`] = FuncRoot{Func: reflect.ValueOf(regexp.MatchString)}
	Roots[`/regexp/MustCompile`] = FuncRoot{Func: reflect.ValueOf(regexp.MustCompile)}
	Roots[`/regexp/MustCompilePOSIX`] = FuncRoot{Func: reflect.ValueOf(regexp.MustCompilePOSIX)}
	Roots[`/regexp/QuoteMeta`] = FuncRoot{Func: reflect.ValueOf(regexp.QuoteMeta)}
	Roots[`/regexp/syntax/Compile`] = FuncRoot{Func: reflect.ValueOf(regexp_syntax.Compile)}
	Roots[`/regexp/syntax/EmptyOpContext`] = FuncRoot{Func: reflect.ValueOf(regexp_syntax.EmptyOpContext)}
	Roots[`/regexp/syntax/IsWordChar`] = FuncRoot{Func: reflect.ValueOf(regexp_syntax.IsWordChar)}
	Roots[`/regexp/syntax/Parse`] = FuncRoot{Func: reflect.ValueOf(regexp_syntax.Parse)}
	Roots[`/runtime/Breakpoint`] = FuncRoot{Func: reflect.ValueOf(runtime.Breakpoint)}
	Roots[`/runtime/CPUProfile`] = FuncRoot{Func: reflect.ValueOf(runtime.CPUProfile)}
	Roots[`/runtime/Caller`] = FuncRoot{Func: reflect.ValueOf(runtime.Caller)}
	Roots[`/runtime/Callers`] = FuncRoot{Func: reflect.ValueOf(runtime.Callers)}
	Roots[`/runtime/FuncForPC`] = FuncRoot{Func: reflect.ValueOf(runtime.FuncForPC)}
	Roots[`/runtime/GC`] = FuncRoot{Func: reflect.ValueOf(runtime.GC)}
	Roots[`/runtime/GOMAXPROCS`] = FuncRoot{Func: reflect.ValueOf(runtime.GOMAXPROCS)}
	Roots[`/runtime/GOROOT`] = FuncRoot{Func: reflect.ValueOf(runtime.GOROOT)}
	Roots[`/runtime/Goexit`] = FuncRoot{Func: reflect.ValueOf(runtime.Goexit)}
	Roots[`/runtime/GoroutineProfile`] = FuncRoot{Func: reflect.ValueOf(runtime.GoroutineProfile)}
	Roots[`/runtime/Gosched`] = FuncRoot{Func: reflect.ValueOf(runtime.Gosched)}
	Roots[`/runtime/LockOSThread`] = FuncRoot{Func: reflect.ValueOf(runtime.LockOSThread)}
	Roots[`/runtime/MemProfile`] = FuncRoot{Func: reflect.ValueOf(runtime.MemProfile)}
	Roots[`/runtime/NumCPU`] = FuncRoot{Func: reflect.ValueOf(runtime.NumCPU)}
	Roots[`/runtime/NumCgoCall`] = FuncRoot{Func: reflect.ValueOf(runtime.NumCgoCall)}
	Roots[`/runtime/NumGoroutine`] = FuncRoot{Func: reflect.ValueOf(runtime.NumGoroutine)}
	Roots[`/runtime/ReadMemStats`] = FuncRoot{Func: reflect.ValueOf(runtime.ReadMemStats)}
	Roots[`/runtime/SetCPUProfileRate`] = FuncRoot{Func: reflect.ValueOf(runtime.SetCPUProfileRate)}
	Roots[`/runtime/SetFinalizer`] = FuncRoot{Func: reflect.ValueOf(runtime.SetFinalizer)}
	Roots[`/runtime/Stack`] = FuncRoot{Func: reflect.ValueOf(runtime.Stack)}
	Roots[`/runtime/ThreadCreateProfile`] = FuncRoot{Func: reflect.ValueOf(runtime.ThreadCreateProfile)}
	Roots[`/runtime/UnlockOSThread`] = FuncRoot{Func: reflect.ValueOf(runtime.UnlockOSThread)}
	Roots[`/runtime/Version`] = FuncRoot{Func: reflect.ValueOf(runtime.Version)}
	Roots[`/runtime/debug/PrintStack`] = FuncRoot{Func: reflect.ValueOf(runtime_debug.PrintStack)}
	Roots[`/runtime/debug/Stack`] = FuncRoot{Func: reflect.ValueOf(runtime_debug.Stack)}
	Roots[`/runtime/pprof/Lookup`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.Lookup)}
	Roots[`/runtime/pprof/NewProfile`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.NewProfile)}
	Roots[`/runtime/pprof/Profiles`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.Profiles)}
	Roots[`/runtime/pprof/StartCPUProfile`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.StartCPUProfile)}
	Roots[`/runtime/pprof/StopCPUProfile`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.StopCPUProfile)}
	Roots[`/runtime/pprof/WriteHeapProfile`] = FuncRoot{Func: reflect.ValueOf(runtime_pprof.WriteHeapProfile)}
	Roots[`/sort/Float64s`] = FuncRoot{Func: reflect.ValueOf(sort.Float64s)}
	Roots[`/sort/Float64sAreSorted`] = FuncRoot{Func: reflect.ValueOf(sort.Float64sAreSorted)}
	Roots[`/sort/Ints`] = FuncRoot{Func: reflect.ValueOf(sort.Ints)}
	Roots[`/sort/IntsAreSorted`] = FuncRoot{Func: reflect.ValueOf(sort.IntsAreSorted)}
	Roots[`/sort/IsSorted`] = FuncRoot{Func: reflect.ValueOf(sort.IsSorted)}
	Roots[`/sort/Search`] = FuncRoot{Func: reflect.ValueOf(sort.Search)}
	Roots[`/sort/SearchFloat64s`] = FuncRoot{Func: reflect.ValueOf(sort.SearchFloat64s)}
	Roots[`/sort/SearchInts`] = FuncRoot{Func: reflect.ValueOf(sort.SearchInts)}
	Roots[`/sort/SearchStrings`] = FuncRoot{Func: reflect.ValueOf(sort.SearchStrings)}
	Roots[`/sort/Sort`] = FuncRoot{Func: reflect.ValueOf(sort.Sort)}
	Roots[`/sort/Strings`] = FuncRoot{Func: reflect.ValueOf(sort.Strings)}
	Roots[`/sort/StringsAreSorted`] = FuncRoot{Func: reflect.ValueOf(sort.StringsAreSorted)}
	Roots[`/strconv/AppendBool`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendBool)}
	Roots[`/strconv/AppendFloat`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendFloat)}
	Roots[`/strconv/AppendInt`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendInt)}
	Roots[`/strconv/AppendQuote`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendQuote)}
	Roots[`/strconv/AppendQuoteRune`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendQuoteRune)}
	Roots[`/strconv/AppendQuoteRuneToASCII`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendQuoteRuneToASCII)}
	Roots[`/strconv/AppendQuoteToASCII`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendQuoteToASCII)}
	Roots[`/strconv/AppendUint`] = FuncRoot{Func: reflect.ValueOf(strconv.AppendUint)}
	Roots[`/strconv/Atoi`] = FuncRoot{Func: reflect.ValueOf(strconv.Atoi)}
	Roots[`/strconv/CanBackquote`] = FuncRoot{Func: reflect.ValueOf(strconv.CanBackquote)}
	Roots[`/strconv/FormatBool`] = FuncRoot{Func: reflect.ValueOf(strconv.FormatBool)}
	Roots[`/strconv/FormatFloat`] = FuncRoot{Func: reflect.ValueOf(strconv.FormatFloat)}
	Roots[`/strconv/FormatInt`] = FuncRoot{Func: reflect.ValueOf(strconv.FormatInt)}
	Roots[`/strconv/FormatUint`] = FuncRoot{Func: reflect.ValueOf(strconv.FormatUint)}
	Roots[`/strconv/IsPrint`] = FuncRoot{Func: reflect.ValueOf(strconv.IsPrint)}
	Roots[`/strconv/Itoa`] = FuncRoot{Func: reflect.ValueOf(strconv.Itoa)}
	Roots[`/strconv/ParseBool`] = FuncRoot{Func: reflect.ValueOf(strconv.ParseBool)}
	Roots[`/strconv/ParseFloat`] = FuncRoot{Func: reflect.ValueOf(strconv.ParseFloat)}
	Roots[`/strconv/ParseInt`] = FuncRoot{Func: reflect.ValueOf(strconv.ParseInt)}
	Roots[`/strconv/ParseUint`] = FuncRoot{Func: reflect.ValueOf(strconv.ParseUint)}
	Roots[`/strconv/Quote`] = FuncRoot{Func: reflect.ValueOf(strconv.Quote)}
	Roots[`/strconv/QuoteRune`] = FuncRoot{Func: reflect.ValueOf(strconv.QuoteRune)}
	Roots[`/strconv/QuoteRuneToASCII`] = FuncRoot{Func: reflect.ValueOf(strconv.QuoteRuneToASCII)}
	Roots[`/strconv/QuoteToASCII`] = FuncRoot{Func: reflect.ValueOf(strconv.QuoteToASCII)}
	Roots[`/strconv/Unquote`] = FuncRoot{Func: reflect.ValueOf(strconv.Unquote)}
	Roots[`/strconv/UnquoteChar`] = FuncRoot{Func: reflect.ValueOf(strconv.UnquoteChar)}
	Roots[`/strings/Contains`] = FuncRoot{Func: reflect.ValueOf(strings.Contains)}
	Roots[`/strings/ContainsAny`] = FuncRoot{Func: reflect.ValueOf(strings.ContainsAny)}
	Roots[`/strings/ContainsRune`] = FuncRoot{Func: reflect.ValueOf(strings.ContainsRune)}
	Roots[`/strings/Count`] = FuncRoot{Func: reflect.ValueOf(strings.Count)}
	Roots[`/strings/EqualFold`] = FuncRoot{Func: reflect.ValueOf(strings.EqualFold)}
	Roots[`/strings/Fields`] = FuncRoot{Func: reflect.ValueOf(strings.Fields)}
	Roots[`/strings/FieldsFunc`] = FuncRoot{Func: reflect.ValueOf(strings.FieldsFunc)}
	Roots[`/strings/HasPrefix`] = FuncRoot{Func: reflect.ValueOf(strings.HasPrefix)}
	Roots[`/strings/HasSuffix`] = FuncRoot{Func: reflect.ValueOf(strings.HasSuffix)}
	Roots[`/strings/Index`] = FuncRoot{Func: reflect.ValueOf(strings.Index)}
	Roots[`/strings/IndexAny`] = FuncRoot{Func: reflect.ValueOf(strings.IndexAny)}
	Roots[`/strings/IndexFunc`] = FuncRoot{Func: reflect.ValueOf(strings.IndexFunc)}
	Roots[`/strings/IndexRune`] = FuncRoot{Func: reflect.ValueOf(strings.IndexRune)}
	Roots[`/strings/Join`] = FuncRoot{Func: reflect.ValueOf(strings.Join)}
	Roots[`/strings/LastIndex`] = FuncRoot{Func: reflect.ValueOf(strings.LastIndex)}
	Roots[`/strings/LastIndexAny`] = FuncRoot{Func: reflect.ValueOf(strings.LastIndexAny)}
	Roots[`/strings/LastIndexFunc`] = FuncRoot{Func: reflect.ValueOf(strings.LastIndexFunc)}
	Roots[`/strings/Map`] = FuncRoot{Func: reflect.ValueOf(strings.Map)}
	Roots[`/strings/NewReader`] = FuncRoot{Func: reflect.ValueOf(strings.NewReader)}
	Roots[`/strings/NewReplacer`] = FuncRoot{Func: reflect.ValueOf(strings.NewReplacer)}
	Roots[`/strings/Repeat`] = FuncRoot{Func: reflect.ValueOf(strings.Repeat)}
	Roots[`/strings/Replace`] = FuncRoot{Func: reflect.ValueOf(strings.Replace)}
	Roots[`/strings/Split`] = FuncRoot{Func: reflect.ValueOf(strings.Split)}
	Roots[`/strings/SplitAfter`] = FuncRoot{Func: reflect.ValueOf(strings.SplitAfter)}
	Roots[`/strings/SplitAfterN`] = FuncRoot{Func: reflect.ValueOf(strings.SplitAfterN)}
	Roots[`/strings/SplitN`] = FuncRoot{Func: reflect.ValueOf(strings.SplitN)}
	Roots[`/strings/Title`] = FuncRoot{Func: reflect.ValueOf(strings.Title)}
	Roots[`/strings/ToLower`] = FuncRoot{Func: reflect.ValueOf(strings.ToLower)}
	Roots[`/strings/ToLowerSpecial`] = FuncRoot{Func: reflect.ValueOf(strings.ToLowerSpecial)}
	Roots[`/strings/ToTitle`] = FuncRoot{Func: reflect.ValueOf(strings.ToTitle)}
	Roots[`/strings/ToTitleSpecial`] = FuncRoot{Func: reflect.ValueOf(strings.ToTitleSpecial)}
	Roots[`/strings/ToUpper`] = FuncRoot{Func: reflect.ValueOf(strings.ToUpper)}
	Roots[`/strings/ToUpperSpecial`] = FuncRoot{Func: reflect.ValueOf(strings.ToUpperSpecial)}
	Roots[`/strings/Trim`] = FuncRoot{Func: reflect.ValueOf(strings.Trim)}
	Roots[`/strings/TrimFunc`] = FuncRoot{Func: reflect.ValueOf(strings.TrimFunc)}
	Roots[`/strings/TrimLeft`] = FuncRoot{Func: reflect.ValueOf(strings.TrimLeft)}
	Roots[`/strings/TrimLeftFunc`] = FuncRoot{Func: reflect.ValueOf(strings.TrimLeftFunc)}
	Roots[`/strings/TrimRight`] = FuncRoot{Func: reflect.ValueOf(strings.TrimRight)}
	Roots[`/strings/TrimRightFunc`] = FuncRoot{Func: reflect.ValueOf(strings.TrimRightFunc)}
	Roots[`/strings/TrimSpace`] = FuncRoot{Func: reflect.ValueOf(strings.TrimSpace)}
	Roots[`/sync/NewCond`] = FuncRoot{Func: reflect.ValueOf(sync.NewCond)}
	Roots[`/sync/atomic/AddInt32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.AddInt32)}
	Roots[`/sync/atomic/AddInt64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.AddInt64)}
	Roots[`/sync/atomic/AddUint32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.AddUint32)}
	Roots[`/sync/atomic/AddUint64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.AddUint64)}
	Roots[`/sync/atomic/AddUintptr`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.AddUintptr)}
	Roots[`/sync/atomic/CompareAndSwapInt32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapInt32)}
	Roots[`/sync/atomic/CompareAndSwapInt64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapInt64)}
	Roots[`/sync/atomic/CompareAndSwapPointer`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapPointer)}
	Roots[`/sync/atomic/CompareAndSwapUint32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapUint32)}
	Roots[`/sync/atomic/CompareAndSwapUint64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapUint64)}
	Roots[`/sync/atomic/CompareAndSwapUintptr`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.CompareAndSwapUintptr)}
	Roots[`/sync/atomic/LoadInt32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadInt32)}
	Roots[`/sync/atomic/LoadInt64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadInt64)}
	Roots[`/sync/atomic/LoadPointer`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadPointer)}
	Roots[`/sync/atomic/LoadUint32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadUint32)}
	Roots[`/sync/atomic/LoadUint64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadUint64)}
	Roots[`/sync/atomic/LoadUintptr`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.LoadUintptr)}
	Roots[`/sync/atomic/StoreInt32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StoreInt32)}
	Roots[`/sync/atomic/StoreInt64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StoreInt64)}
	Roots[`/sync/atomic/StorePointer`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StorePointer)}
	Roots[`/sync/atomic/StoreUint32`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StoreUint32)}
	Roots[`/sync/atomic/StoreUint64`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StoreUint64)}
	Roots[`/sync/atomic/StoreUintptr`] = FuncRoot{Func: reflect.ValueOf(sync_atomic.StoreUintptr)}
	Roots[`/syscall/Chdir`] = FuncRoot{Func: reflect.ValueOf(syscall.Chdir)}
	Roots[`/syscall/Chmod`] = FuncRoot{Func: reflect.ValueOf(syscall.Chmod)}
	Roots[`/syscall/Chown`] = FuncRoot{Func: reflect.ValueOf(syscall.Chown)}
	Roots[`/syscall/Clearenv`] = FuncRoot{Func: reflect.ValueOf(syscall.Clearenv)}
	Roots[`/syscall/Environ`] = FuncRoot{Func: reflect.ValueOf(syscall.Environ)}
	Roots[`/syscall/Exec`] = FuncRoot{Func: reflect.ValueOf(syscall.Exec)}
	Roots[`/syscall/Exit`] = FuncRoot{Func: reflect.ValueOf(syscall.Exit)}
	Roots[`/syscall/Getegid`] = FuncRoot{Func: reflect.ValueOf(syscall.Getegid)}
	Roots[`/syscall/Getenv`] = FuncRoot{Func: reflect.ValueOf(syscall.Getenv)}
	Roots[`/syscall/Geteuid`] = FuncRoot{Func: reflect.ValueOf(syscall.Geteuid)}
	Roots[`/syscall/Getgid`] = FuncRoot{Func: reflect.ValueOf(syscall.Getgid)}
	Roots[`/syscall/Getgroups`] = FuncRoot{Func: reflect.ValueOf(syscall.Getgroups)}
	Roots[`/syscall/Getpagesize`] = FuncRoot{Func: reflect.ValueOf(syscall.Getpagesize)}
	Roots[`/syscall/Getpid`] = FuncRoot{Func: reflect.ValueOf(syscall.Getpid)}
	Roots[`/syscall/Getppid`] = FuncRoot{Func: reflect.ValueOf(syscall.Getppid)}
	Roots[`/syscall/Gettimeofday`] = FuncRoot{Func: reflect.ValueOf(syscall.Gettimeofday)}
	Roots[`/syscall/Getuid`] = FuncRoot{Func: reflect.ValueOf(syscall.Getuid)}
	Roots[`/syscall/Getwd`] = FuncRoot{Func: reflect.ValueOf(syscall.Getwd)}
	Roots[`/syscall/Lchown`] = FuncRoot{Func: reflect.ValueOf(syscall.Lchown)}
	Roots[`/syscall/Link`] = FuncRoot{Func: reflect.ValueOf(syscall.Link)}
	Roots[`/syscall/Mkdir`] = FuncRoot{Func: reflect.ValueOf(syscall.Mkdir)}
	Roots[`/syscall/NsecToTimeval`] = FuncRoot{Func: reflect.ValueOf(syscall.NsecToTimeval)}
	Roots[`/syscall/Readlink`] = FuncRoot{Func: reflect.ValueOf(syscall.Readlink)}
	Roots[`/syscall/Rename`] = FuncRoot{Func: reflect.ValueOf(syscall.Rename)}
	Roots[`/syscall/Rmdir`] = FuncRoot{Func: reflect.ValueOf(syscall.Rmdir)}
	Roots[`/syscall/Setenv`] = FuncRoot{Func: reflect.ValueOf(syscall.Setenv)}
	Roots[`/syscall/StartProcess`] = FuncRoot{Func: reflect.ValueOf(syscall.StartProcess)}
	Roots[`/syscall/StringBytePtr`] = FuncRoot{Func: reflect.ValueOf(syscall.StringBytePtr)}
	Roots[`/syscall/StringByteSlice`] = FuncRoot{Func: reflect.ValueOf(syscall.StringByteSlice)}
	Roots[`/syscall/Symlink`] = FuncRoot{Func: reflect.ValueOf(syscall.Symlink)}
	Roots[`/syscall/Unlink`] = FuncRoot{Func: reflect.ValueOf(syscall.Unlink)}
	Roots[`/syscall/Utimes`] = FuncRoot{Func: reflect.ValueOf(syscall.Utimes)}
	Roots[`/text/scanner/TokenString`] = FuncRoot{Func: reflect.ValueOf(text_scanner.TokenString)}
	Roots[`/text/tabwriter/NewWriter`] = FuncRoot{Func: reflect.ValueOf(text_tabwriter.NewWriter)}
	Roots[`/text/template/HTMLEscape`] = FuncRoot{Func: reflect.ValueOf(text_template.HTMLEscape)}
	Roots[`/text/template/HTMLEscapeString`] = FuncRoot{Func: reflect.ValueOf(text_template.HTMLEscapeString)}
	Roots[`/text/template/HTMLEscaper`] = FuncRoot{Func: reflect.ValueOf(text_template.HTMLEscaper)}
	Roots[`/text/template/JSEscape`] = FuncRoot{Func: reflect.ValueOf(text_template.JSEscape)}
	Roots[`/text/template/JSEscapeString`] = FuncRoot{Func: reflect.ValueOf(text_template.JSEscapeString)}
	Roots[`/text/template/JSEscaper`] = FuncRoot{Func: reflect.ValueOf(text_template.JSEscaper)}
	Roots[`/text/template/Must`] = FuncRoot{Func: reflect.ValueOf(text_template.Must)}
	Roots[`/text/template/New`] = FuncRoot{Func: reflect.ValueOf(text_template.New)}
	Roots[`/text/template/ParseFiles`] = FuncRoot{Func: reflect.ValueOf(text_template.ParseFiles)}
	Roots[`/text/template/ParseGlob`] = FuncRoot{Func: reflect.ValueOf(text_template.ParseGlob)}
	Roots[`/text/template/URLQueryEscaper`] = FuncRoot{Func: reflect.ValueOf(text_template.URLQueryEscaper)}
	Roots[`/text/template/parse/IsEmptyTree`] = FuncRoot{Func: reflect.ValueOf(text_template_parse.IsEmptyTree)}
	Roots[`/text/template/parse/New`] = FuncRoot{Func: reflect.ValueOf(text_template_parse.New)}
	Roots[`/text/template/parse/NewIdentifier`] = FuncRoot{Func: reflect.ValueOf(text_template_parse.NewIdentifier)}
	Roots[`/text/template/parse/Parse`] = FuncRoot{Func: reflect.ValueOf(text_template_parse.Parse)}
	Roots[`/time/After`] = FuncRoot{Func: reflect.ValueOf(time.After)}
	Roots[`/time/AfterFunc`] = FuncRoot{Func: reflect.ValueOf(time.AfterFunc)}
	Roots[`/time/Date`] = FuncRoot{Func: reflect.ValueOf(time.Date)}
	Roots[`/time/FixedZone`] = FuncRoot{Func: reflect.ValueOf(time.FixedZone)}
	Roots[`/time/LoadLocation`] = FuncRoot{Func: reflect.ValueOf(time.LoadLocation)}
	Roots[`/time/NewTicker`] = FuncRoot{Func: reflect.ValueOf(time.NewTicker)}
	Roots[`/time/NewTimer`] = FuncRoot{Func: reflect.ValueOf(time.NewTimer)}
	Roots[`/time/Now`] = FuncRoot{Func: reflect.ValueOf(time.Now)}
	Roots[`/time/Parse`] = FuncRoot{Func: reflect.ValueOf(time.Parse)}
	Roots[`/time/ParseDuration`] = FuncRoot{Func: reflect.ValueOf(time.ParseDuration)}
	Roots[`/time/Since`] = FuncRoot{Func: reflect.ValueOf(time.Since)}
	Roots[`/time/Sleep`] = FuncRoot{Func: reflect.ValueOf(time.Sleep)}
	Roots[`/time/Tick`] = FuncRoot{Func: reflect.ValueOf(time.Tick)}
	Roots[`/time/Unix`] = FuncRoot{Func: reflect.ValueOf(time.Unix)}
	Roots[`/unicode/Is`] = FuncRoot{Func: reflect.ValueOf(unicode.Is)}
	Roots[`/unicode/IsControl`] = FuncRoot{Func: reflect.ValueOf(unicode.IsControl)}
	Roots[`/unicode/IsDigit`] = FuncRoot{Func: reflect.ValueOf(unicode.IsDigit)}
	Roots[`/unicode/IsGraphic`] = FuncRoot{Func: reflect.ValueOf(unicode.IsGraphic)}
	Roots[`/unicode/IsLetter`] = FuncRoot{Func: reflect.ValueOf(unicode.IsLetter)}
	Roots[`/unicode/IsLower`] = FuncRoot{Func: reflect.ValueOf(unicode.IsLower)}
	Roots[`/unicode/IsMark`] = FuncRoot{Func: reflect.ValueOf(unicode.IsMark)}
	Roots[`/unicode/IsNumber`] = FuncRoot{Func: reflect.ValueOf(unicode.IsNumber)}
	Roots[`/unicode/IsOneOf`] = FuncRoot{Func: reflect.ValueOf(unicode.IsOneOf)}
	Roots[`/unicode/IsPrint`] = FuncRoot{Func: reflect.ValueOf(unicode.IsPrint)}
	Roots[`/unicode/IsPunct`] = FuncRoot{Func: reflect.ValueOf(unicode.IsPunct)}
	Roots[`/unicode/IsSpace`] = FuncRoot{Func: reflect.ValueOf(unicode.IsSpace)}
	Roots[`/unicode/IsSymbol`] = FuncRoot{Func: reflect.ValueOf(unicode.IsSymbol)}
	Roots[`/unicode/IsTitle`] = FuncRoot{Func: reflect.ValueOf(unicode.IsTitle)}
	Roots[`/unicode/IsUpper`] = FuncRoot{Func: reflect.ValueOf(unicode.IsUpper)}
	Roots[`/unicode/SimpleFold`] = FuncRoot{Func: reflect.ValueOf(unicode.SimpleFold)}
	Roots[`/unicode/To`] = FuncRoot{Func: reflect.ValueOf(unicode.To)}
	Roots[`/unicode/ToLower`] = FuncRoot{Func: reflect.ValueOf(unicode.ToLower)}
	Roots[`/unicode/ToTitle`] = FuncRoot{Func: reflect.ValueOf(unicode.ToTitle)}
	Roots[`/unicode/ToUpper`] = FuncRoot{Func: reflect.ValueOf(unicode.ToUpper)}
	Roots[`/unicode/utf16/Decode`] = FuncRoot{Func: reflect.ValueOf(unicode_utf16.Decode)}
	Roots[`/unicode/utf16/DecodeRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf16.DecodeRune)}
	Roots[`/unicode/utf16/Encode`] = FuncRoot{Func: reflect.ValueOf(unicode_utf16.Encode)}
	Roots[`/unicode/utf16/EncodeRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf16.EncodeRune)}
	Roots[`/unicode/utf16/IsSurrogate`] = FuncRoot{Func: reflect.ValueOf(unicode_utf16.IsSurrogate)}
	Roots[`/unicode/utf8/DecodeLastRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.DecodeLastRune)}
	Roots[`/unicode/utf8/DecodeLastRuneInString`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.DecodeLastRuneInString)}
	Roots[`/unicode/utf8/DecodeRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.DecodeRune)}
	Roots[`/unicode/utf8/DecodeRuneInString`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.DecodeRuneInString)}
	Roots[`/unicode/utf8/EncodeRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.EncodeRune)}
	Roots[`/unicode/utf8/FullRune`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.FullRune)}
	Roots[`/unicode/utf8/FullRuneInString`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.FullRuneInString)}
	Roots[`/unicode/utf8/RuneCount`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.RuneCount)}
	Roots[`/unicode/utf8/RuneCountInString`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.RuneCountInString)}
	Roots[`/unicode/utf8/RuneLen`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.RuneLen)}
	Roots[`/unicode/utf8/RuneStart`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.RuneStart)}
	Roots[`/unicode/utf8/Valid`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.Valid)}
	Roots[`/unicode/utf8/ValidString`] = FuncRoot{Func: reflect.ValueOf(unicode_utf8.ValidString)}
	Roots[`/archive/tar/ErrFieldTooLong`] = VarRoot{Var: reflect.ValueOf(&archive_tar.ErrFieldTooLong)}
	Roots[`/archive/tar/ErrHeader`] = VarRoot{Var: reflect.ValueOf(&archive_tar.ErrHeader)}
	Roots[`/archive/tar/ErrWriteAfterClose`] = VarRoot{Var: reflect.ValueOf(&archive_tar.ErrWriteAfterClose)}
	Roots[`/archive/tar/ErrWriteTooLong`] = VarRoot{Var: reflect.ValueOf(&archive_tar.ErrWriteTooLong)}
	Roots[`/archive/zip/ErrAlgorithm`] = VarRoot{Var: reflect.ValueOf(&archive_zip.ErrAlgorithm)}
	Roots[`/archive/zip/ErrChecksum`] = VarRoot{Var: reflect.ValueOf(&archive_zip.ErrChecksum)}
	Roots[`/archive/zip/ErrFormat`] = VarRoot{Var: reflect.ValueOf(&archive_zip.ErrFormat)}
	Roots[`/bufio/ErrBufferFull`] = VarRoot{Var: reflect.ValueOf(&bufio.ErrBufferFull)}
	Roots[`/bufio/ErrInvalidUnreadByte`] = VarRoot{Var: reflect.ValueOf(&bufio.ErrInvalidUnreadByte)}
	Roots[`/bufio/ErrInvalidUnreadRune`] = VarRoot{Var: reflect.ValueOf(&bufio.ErrInvalidUnreadRune)}
	Roots[`/bufio/ErrNegativeCount`] = VarRoot{Var: reflect.ValueOf(&bufio.ErrNegativeCount)}
	Roots[`/bytes/ErrTooLarge`] = VarRoot{Var: reflect.ValueOf(&bytes.ErrTooLarge)}
	Roots[`/compress/gzip/ErrChecksum`] = VarRoot{Var: reflect.ValueOf(&compress_gzip.ErrChecksum)}
	Roots[`/compress/gzip/ErrHeader`] = VarRoot{Var: reflect.ValueOf(&compress_gzip.ErrHeader)}
	Roots[`/compress/zlib/ErrChecksum`] = VarRoot{Var: reflect.ValueOf(&compress_zlib.ErrChecksum)}
	Roots[`/compress/zlib/ErrDictionary`] = VarRoot{Var: reflect.ValueOf(&compress_zlib.ErrDictionary)}
	Roots[`/compress/zlib/ErrHeader`] = VarRoot{Var: reflect.ValueOf(&compress_zlib.ErrHeader)}
	Roots[`/crypto/dsa/ErrInvalidPublicKey`] = VarRoot{Var: reflect.ValueOf(&crypto_dsa.ErrInvalidPublicKey)}
	Roots[`/crypto/rand/Reader`] = VarRoot{Var: reflect.ValueOf(&crypto_rand.Reader)}
	Roots[`/crypto/rsa/ErrDecryption`] = VarRoot{Var: reflect.ValueOf(&crypto_rsa.ErrDecryption)}
	Roots[`/crypto/rsa/ErrMessageTooLong`] = VarRoot{Var: reflect.ValueOf(&crypto_rsa.ErrMessageTooLong)}
	Roots[`/crypto/rsa/ErrVerification`] = VarRoot{Var: reflect.ValueOf(&crypto_rsa.ErrVerification)}
	Roots[`/crypto/x509/ErrUnsupportedAlgorithm`] = VarRoot{Var: reflect.ValueOf(&crypto_x509.ErrUnsupportedAlgorithm)}
	Roots[`/database/sql/ErrNoRows`] = VarRoot{Var: reflect.ValueOf(&database_sql.ErrNoRows)}
	Roots[`/database/sql/ErrTxDone`] = VarRoot{Var: reflect.ValueOf(&database_sql.ErrTxDone)}
	Roots[`/database/sql/driver/Bool`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.Bool)}
	Roots[`/database/sql/driver/DefaultParameterConverter`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.DefaultParameterConverter)}
	Roots[`/database/sql/driver/ErrBadConn`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.ErrBadConn)}
	Roots[`/database/sql/driver/ErrSkip`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.ErrSkip)}
	Roots[`/database/sql/driver/Int32`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.Int32)}
	Roots[`/database/sql/driver/ResultNoRows`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.ResultNoRows)}
	Roots[`/database/sql/driver/String`] = VarRoot{Var: reflect.ValueOf(&database_sql_driver.String)}
	Roots[`/encoding/base32/HexEncoding`] = VarRoot{Var: reflect.ValueOf(&encoding_base32.HexEncoding)}
	Roots[`/encoding/base32/StdEncoding`] = VarRoot{Var: reflect.ValueOf(&encoding_base32.StdEncoding)}
	Roots[`/encoding/base64/StdEncoding`] = VarRoot{Var: reflect.ValueOf(&encoding_base64.StdEncoding)}
	Roots[`/encoding/base64/URLEncoding`] = VarRoot{Var: reflect.ValueOf(&encoding_base64.URLEncoding)}
	Roots[`/encoding/binary/BigEndian`] = VarRoot{Var: reflect.ValueOf(&encoding_binary.BigEndian)}
	Roots[`/encoding/binary/LittleEndian`] = VarRoot{Var: reflect.ValueOf(&encoding_binary.LittleEndian)}
	Roots[`/encoding/csv/ErrBareQuote`] = VarRoot{Var: reflect.ValueOf(&encoding_csv.ErrBareQuote)}
	Roots[`/encoding/csv/ErrFieldCount`] = VarRoot{Var: reflect.ValueOf(&encoding_csv.ErrFieldCount)}
	Roots[`/encoding/csv/ErrQuote`] = VarRoot{Var: reflect.ValueOf(&encoding_csv.ErrQuote)}
	Roots[`/encoding/csv/ErrTrailingComma`] = VarRoot{Var: reflect.ValueOf(&encoding_csv.ErrTrailingComma)}
	Roots[`/encoding/hex/ErrLength`] = VarRoot{Var: reflect.ValueOf(&encoding_hex.ErrLength)}
	Roots[`/encoding/xml/HTMLAutoClose`] = VarRoot{Var: reflect.ValueOf(&encoding_xml.HTMLAutoClose)}
	Roots[`/encoding/xml/HTMLEntity`] = VarRoot{Var: reflect.ValueOf(&encoding_xml.HTMLEntity)}
	Roots[`/flag/ErrHelp`] = VarRoot{Var: reflect.ValueOf(&flag.ErrHelp)}
	Roots[`/flag/Usage`] = VarRoot{Var: reflect.ValueOf(&flag.Usage)}
	Roots[`/go/build/Default`] = VarRoot{Var: reflect.ValueOf(&go_build.Default)}
	Roots[`/go/build/ToolDir`] = VarRoot{Var: reflect.ValueOf(&go_build.ToolDir)}
	Roots[`/hash/crc32/IEEETable`] = VarRoot{Var: reflect.ValueOf(&hash_crc32.IEEETable)}
	Roots[`/image/Black`] = VarRoot{Var: reflect.ValueOf(&image.Black)}
	Roots[`/image/ErrFormat`] = VarRoot{Var: reflect.ValueOf(&image.ErrFormat)}
	Roots[`/image/Opaque`] = VarRoot{Var: reflect.ValueOf(&image.Opaque)}
	Roots[`/image/Transparent`] = VarRoot{Var: reflect.ValueOf(&image.Transparent)}
	Roots[`/image/White`] = VarRoot{Var: reflect.ValueOf(&image.White)}
	Roots[`/image/ZP`] = VarRoot{Var: reflect.ValueOf(&image.ZP)}
	Roots[`/image/ZR`] = VarRoot{Var: reflect.ValueOf(&image.ZR)}
	Roots[`/image/color/Alpha16Model`] = VarRoot{Var: reflect.ValueOf(&image_color.Alpha16Model)}
	Roots[`/image/color/AlphaModel`] = VarRoot{Var: reflect.ValueOf(&image_color.AlphaModel)}
	Roots[`/image/color/Black`] = VarRoot{Var: reflect.ValueOf(&image_color.Black)}
	Roots[`/image/color/Gray16Model`] = VarRoot{Var: reflect.ValueOf(&image_color.Gray16Model)}
	Roots[`/image/color/GrayModel`] = VarRoot{Var: reflect.ValueOf(&image_color.GrayModel)}
	Roots[`/image/color/NRGBA64Model`] = VarRoot{Var: reflect.ValueOf(&image_color.NRGBA64Model)}
	Roots[`/image/color/NRGBAModel`] = VarRoot{Var: reflect.ValueOf(&image_color.NRGBAModel)}
	Roots[`/image/color/Opaque`] = VarRoot{Var: reflect.ValueOf(&image_color.Opaque)}
	Roots[`/image/color/RGBA64Model`] = VarRoot{Var: reflect.ValueOf(&image_color.RGBA64Model)}
	Roots[`/image/color/RGBAModel`] = VarRoot{Var: reflect.ValueOf(&image_color.RGBAModel)}
	Roots[`/image/color/Transparent`] = VarRoot{Var: reflect.ValueOf(&image_color.Transparent)}
	Roots[`/image/color/White`] = VarRoot{Var: reflect.ValueOf(&image_color.White)}
	Roots[`/image/color/YCbCrModel`] = VarRoot{Var: reflect.ValueOf(&image_color.YCbCrModel)}
	Roots[`/io/EOF`] = VarRoot{Var: reflect.ValueOf(&io.EOF)}
	Roots[`/io/ErrClosedPipe`] = VarRoot{Var: reflect.ValueOf(&io.ErrClosedPipe)}
	Roots[`/io/ErrShortBuffer`] = VarRoot{Var: reflect.ValueOf(&io.ErrShortBuffer)}
	Roots[`/io/ErrShortWrite`] = VarRoot{Var: reflect.ValueOf(&io.ErrShortWrite)}
	Roots[`/io/ErrUnexpectedEOF`] = VarRoot{Var: reflect.ValueOf(&io.ErrUnexpectedEOF)}
	Roots[`/io/ioutil/Discard`] = VarRoot{Var: reflect.ValueOf(&io_ioutil.Discard)}
	Roots[`/net/ErrWriteToConnected`] = VarRoot{Var: reflect.ValueOf(&net.ErrWriteToConnected)}
	Roots[`/net/IPv4allrouter`] = VarRoot{Var: reflect.ValueOf(&net.IPv4allrouter)}
	Roots[`/net/IPv4allsys`] = VarRoot{Var: reflect.ValueOf(&net.IPv4allsys)}
	Roots[`/net/IPv4bcast`] = VarRoot{Var: reflect.ValueOf(&net.IPv4bcast)}
	Roots[`/net/IPv4zero`] = VarRoot{Var: reflect.ValueOf(&net.IPv4zero)}
	Roots[`/net/IPv6interfacelocalallnodes`] = VarRoot{Var: reflect.ValueOf(&net.IPv6interfacelocalallnodes)}
	Roots[`/net/IPv6linklocalallnodes`] = VarRoot{Var: reflect.ValueOf(&net.IPv6linklocalallnodes)}
	Roots[`/net/IPv6linklocalallrouters`] = VarRoot{Var: reflect.ValueOf(&net.IPv6linklocalallrouters)}
	Roots[`/net/IPv6loopback`] = VarRoot{Var: reflect.ValueOf(&net.IPv6loopback)}
	Roots[`/net/IPv6unspecified`] = VarRoot{Var: reflect.ValueOf(&net.IPv6unspecified)}
	Roots[`/net/IPv6zero`] = VarRoot{Var: reflect.ValueOf(&net.IPv6zero)}
	Roots[`/net/http/DefaultClient`] = VarRoot{Var: reflect.ValueOf(&net_http.DefaultClient)}
	Roots[`/net/http/DefaultServeMux`] = VarRoot{Var: reflect.ValueOf(&net_http.DefaultServeMux)}
	Roots[`/net/http/DefaultTransport`] = VarRoot{Var: reflect.ValueOf(&net_http.DefaultTransport)}
	Roots[`/net/http/ErrBodyNotAllowed`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrBodyNotAllowed)}
	Roots[`/net/http/ErrBodyReadAfterClose`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrBodyReadAfterClose)}
	Roots[`/net/http/ErrContentLength`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrContentLength)}
	Roots[`/net/http/ErrHandlerTimeout`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrHandlerTimeout)}
	Roots[`/net/http/ErrHeaderTooLong`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrHeaderTooLong)}
	Roots[`/net/http/ErrHijacked`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrHijacked)}
	Roots[`/net/http/ErrLineTooLong`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrLineTooLong)}
	Roots[`/net/http/ErrMissingBoundary`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrMissingBoundary)}
	Roots[`/net/http/ErrMissingContentLength`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrMissingContentLength)}
	Roots[`/net/http/ErrMissingFile`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrMissingFile)}
	Roots[`/net/http/ErrNoCookie`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrNoCookie)}
	Roots[`/net/http/ErrNoLocation`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrNoLocation)}
	Roots[`/net/http/ErrNotMultipart`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrNotMultipart)}
	Roots[`/net/http/ErrNotSupported`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrNotSupported)}
	Roots[`/net/http/ErrShortBody`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrShortBody)}
	Roots[`/net/http/ErrUnexpectedTrailer`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrUnexpectedTrailer)}
	Roots[`/net/http/ErrWriteAfterFlush`] = VarRoot{Var: reflect.ValueOf(&net_http.ErrWriteAfterFlush)}
	Roots[`/net/http/httputil/ErrClosed`] = VarRoot{Var: reflect.ValueOf(&net_http_httputil.ErrClosed)}
	Roots[`/net/http/httputil/ErrLineTooLong`] = VarRoot{Var: reflect.ValueOf(&net_http_httputil.ErrLineTooLong)}
	Roots[`/net/http/httputil/ErrPersistEOF`] = VarRoot{Var: reflect.ValueOf(&net_http_httputil.ErrPersistEOF)}
	Roots[`/net/http/httputil/ErrPipeline`] = VarRoot{Var: reflect.ValueOf(&net_http_httputil.ErrPipeline)}
	Roots[`/net/mail/ErrHeaderNotPresent`] = VarRoot{Var: reflect.ValueOf(&net_mail.ErrHeaderNotPresent)}
	Roots[`/net/rpc/DefaultServer`] = VarRoot{Var: reflect.ValueOf(&net_rpc.DefaultServer)}
	Roots[`/net/rpc/ErrShutdown`] = VarRoot{Var: reflect.ValueOf(&net_rpc.ErrShutdown)}
	Roots[`/os/Args`] = VarRoot{Var: reflect.ValueOf(&os.Args)}
	Roots[`/os/ErrExist`] = VarRoot{Var: reflect.ValueOf(&os.ErrExist)}
	Roots[`/os/ErrInvalid`] = VarRoot{Var: reflect.ValueOf(&os.ErrInvalid)}
	Roots[`/os/ErrNotExist`] = VarRoot{Var: reflect.ValueOf(&os.ErrNotExist)}
	Roots[`/os/ErrPermission`] = VarRoot{Var: reflect.ValueOf(&os.ErrPermission)}
	Roots[`/os/Interrupt`] = VarRoot{Var: reflect.ValueOf(&os.Interrupt)}
	Roots[`/os/Kill`] = VarRoot{Var: reflect.ValueOf(&os.Kill)}
	Roots[`/os/Stderr`] = VarRoot{Var: reflect.ValueOf(&os.Stderr)}
	Roots[`/os/Stdin`] = VarRoot{Var: reflect.ValueOf(&os.Stdin)}
	Roots[`/os/Stdout`] = VarRoot{Var: reflect.ValueOf(&os.Stdout)}
	Roots[`/os/exec/ErrNotFound`] = VarRoot{Var: reflect.ValueOf(&os_exec.ErrNotFound)}
	Roots[`/path/ErrBadPattern`] = VarRoot{Var: reflect.ValueOf(&path.ErrBadPattern)}
	Roots[`/path/filepath/ErrBadPattern`] = VarRoot{Var: reflect.ValueOf(&path_filepath.ErrBadPattern)}
	Roots[`/path/filepath/SkipDir`] = VarRoot{Var: reflect.ValueOf(&path_filepath.SkipDir)}
	Roots[`/runtime/MemProfileRate`] = VarRoot{Var: reflect.ValueOf(&runtime.MemProfileRate)}
	Roots[`/strconv/ErrRange`] = VarRoot{Var: reflect.ValueOf(&strconv.ErrRange)}
	Roots[`/strconv/ErrSyntax`] = VarRoot{Var: reflect.ValueOf(&strconv.ErrSyntax)}
	Roots[`/syscall/ForkLock`] = VarRoot{Var: reflect.ValueOf(&syscall.ForkLock)}
	Roots[`/syscall/SocketDisableIPv6`] = VarRoot{Var: reflect.ValueOf(&syscall.SocketDisableIPv6)}
	Roots[`/time/Local`] = VarRoot{Var: reflect.ValueOf(&time.Local)}
	Roots[`/time/UTC`] = VarRoot{Var: reflect.ValueOf(&time.UTC)}
	Roots[`/unicode/ASCII_Hex_Digit`] = VarRoot{Var: reflect.ValueOf(&unicode.ASCII_Hex_Digit)}
	Roots[`/unicode/Arabic`] = VarRoot{Var: reflect.ValueOf(&unicode.Arabic)}
	Roots[`/unicode/Armenian`] = VarRoot{Var: reflect.ValueOf(&unicode.Armenian)}
	Roots[`/unicode/Avestan`] = VarRoot{Var: reflect.ValueOf(&unicode.Avestan)}
	Roots[`/unicode/AzeriCase`] = VarRoot{Var: reflect.ValueOf(&unicode.AzeriCase)}
	Roots[`/unicode/Balinese`] = VarRoot{Var: reflect.ValueOf(&unicode.Balinese)}
	Roots[`/unicode/Bamum`] = VarRoot{Var: reflect.ValueOf(&unicode.Bamum)}
	Roots[`/unicode/Batak`] = VarRoot{Var: reflect.ValueOf(&unicode.Batak)}
	Roots[`/unicode/Bengali`] = VarRoot{Var: reflect.ValueOf(&unicode.Bengali)}
	Roots[`/unicode/Bidi_Control`] = VarRoot{Var: reflect.ValueOf(&unicode.Bidi_Control)}
	Roots[`/unicode/Bopomofo`] = VarRoot{Var: reflect.ValueOf(&unicode.Bopomofo)}
	Roots[`/unicode/Brahmi`] = VarRoot{Var: reflect.ValueOf(&unicode.Brahmi)}
	Roots[`/unicode/Braille`] = VarRoot{Var: reflect.ValueOf(&unicode.Braille)}
	Roots[`/unicode/Buginese`] = VarRoot{Var: reflect.ValueOf(&unicode.Buginese)}
	Roots[`/unicode/Buhid`] = VarRoot{Var: reflect.ValueOf(&unicode.Buhid)}
	Roots[`/unicode/C`] = VarRoot{Var: reflect.ValueOf(&unicode.C)}
	Roots[`/unicode/Canadian_Aboriginal`] = VarRoot{Var: reflect.ValueOf(&unicode.Canadian_Aboriginal)}
	Roots[`/unicode/Carian`] = VarRoot{Var: reflect.ValueOf(&unicode.Carian)}
	Roots[`/unicode/CaseRanges`] = VarRoot{Var: reflect.ValueOf(&unicode.CaseRanges)}
	Roots[`/unicode/Categories`] = VarRoot{Var: reflect.ValueOf(&unicode.Categories)}
	Roots[`/unicode/Cc`] = VarRoot{Var: reflect.ValueOf(&unicode.Cc)}
	Roots[`/unicode/Cf`] = VarRoot{Var: reflect.ValueOf(&unicode.Cf)}
	Roots[`/unicode/Cham`] = VarRoot{Var: reflect.ValueOf(&unicode.Cham)}
	Roots[`/unicode/Cherokee`] = VarRoot{Var: reflect.ValueOf(&unicode.Cherokee)}
	Roots[`/unicode/Co`] = VarRoot{Var: reflect.ValueOf(&unicode.Co)}
	Roots[`/unicode/Common`] = VarRoot{Var: reflect.ValueOf(&unicode.Common)}
	Roots[`/unicode/Coptic`] = VarRoot{Var: reflect.ValueOf(&unicode.Coptic)}
	Roots[`/unicode/Cs`] = VarRoot{Var: reflect.ValueOf(&unicode.Cs)}
	Roots[`/unicode/Cuneiform`] = VarRoot{Var: reflect.ValueOf(&unicode.Cuneiform)}
	Roots[`/unicode/Cypriot`] = VarRoot{Var: reflect.ValueOf(&unicode.Cypriot)}
	Roots[`/unicode/Cyrillic`] = VarRoot{Var: reflect.ValueOf(&unicode.Cyrillic)}
	Roots[`/unicode/Dash`] = VarRoot{Var: reflect.ValueOf(&unicode.Dash)}
	Roots[`/unicode/Deprecated`] = VarRoot{Var: reflect.ValueOf(&unicode.Deprecated)}
	Roots[`/unicode/Deseret`] = VarRoot{Var: reflect.ValueOf(&unicode.Deseret)}
	Roots[`/unicode/Devanagari`] = VarRoot{Var: reflect.ValueOf(&unicode.Devanagari)}
	Roots[`/unicode/Diacritic`] = VarRoot{Var: reflect.ValueOf(&unicode.Diacritic)}
	Roots[`/unicode/Digit`] = VarRoot{Var: reflect.ValueOf(&unicode.Digit)}
	Roots[`/unicode/Egyptian_Hieroglyphs`] = VarRoot{Var: reflect.ValueOf(&unicode.Egyptian_Hieroglyphs)}
	Roots[`/unicode/Ethiopic`] = VarRoot{Var: reflect.ValueOf(&unicode.Ethiopic)}
	Roots[`/unicode/Extender`] = VarRoot{Var: reflect.ValueOf(&unicode.Extender)}
	Roots[`/unicode/FoldCategory`] = VarRoot{Var: reflect.ValueOf(&unicode.FoldCategory)}
	Roots[`/unicode/FoldScript`] = VarRoot{Var: reflect.ValueOf(&unicode.FoldScript)}
	Roots[`/unicode/Georgian`] = VarRoot{Var: reflect.ValueOf(&unicode.Georgian)}
	Roots[`/unicode/Glagolitic`] = VarRoot{Var: reflect.ValueOf(&unicode.Glagolitic)}
	Roots[`/unicode/Gothic`] = VarRoot{Var: reflect.ValueOf(&unicode.Gothic)}
	Roots[`/unicode/GraphicRanges`] = VarRoot{Var: reflect.ValueOf(&unicode.GraphicRanges)}
	Roots[`/unicode/Greek`] = VarRoot{Var: reflect.ValueOf(&unicode.Greek)}
	Roots[`/unicode/Gujarati`] = VarRoot{Var: reflect.ValueOf(&unicode.Gujarati)}
	Roots[`/unicode/Gurmukhi`] = VarRoot{Var: reflect.ValueOf(&unicode.Gurmukhi)}
	Roots[`/unicode/Han`] = VarRoot{Var: reflect.ValueOf(&unicode.Han)}
	Roots[`/unicode/Hangul`] = VarRoot{Var: reflect.ValueOf(&unicode.Hangul)}
	Roots[`/unicode/Hanunoo`] = VarRoot{Var: reflect.ValueOf(&unicode.Hanunoo)}
	Roots[`/unicode/Hebrew`] = VarRoot{Var: reflect.ValueOf(&unicode.Hebrew)}
	Roots[`/unicode/Hex_Digit`] = VarRoot{Var: reflect.ValueOf(&unicode.Hex_Digit)}
	Roots[`/unicode/Hiragana`] = VarRoot{Var: reflect.ValueOf(&unicode.Hiragana)}
	Roots[`/unicode/Hyphen`] = VarRoot{Var: reflect.ValueOf(&unicode.Hyphen)}
	Roots[`/unicode/IDS_Binary_Operator`] = VarRoot{Var: reflect.ValueOf(&unicode.IDS_Binary_Operator)}
	Roots[`/unicode/IDS_Trinary_Operator`] = VarRoot{Var: reflect.ValueOf(&unicode.IDS_Trinary_Operator)}
	Roots[`/unicode/Ideographic`] = VarRoot{Var: reflect.ValueOf(&unicode.Ideographic)}
	Roots[`/unicode/Imperial_Aramaic`] = VarRoot{Var: reflect.ValueOf(&unicode.Imperial_Aramaic)}
	Roots[`/unicode/Inherited`] = VarRoot{Var: reflect.ValueOf(&unicode.Inherited)}
	Roots[`/unicode/Inscriptional_Pahlavi`] = VarRoot{Var: reflect.ValueOf(&unicode.Inscriptional_Pahlavi)}
	Roots[`/unicode/Inscriptional_Parthian`] = VarRoot{Var: reflect.ValueOf(&unicode.Inscriptional_Parthian)}
	Roots[`/unicode/Javanese`] = VarRoot{Var: reflect.ValueOf(&unicode.Javanese)}
	Roots[`/unicode/Join_Control`] = VarRoot{Var: reflect.ValueOf(&unicode.Join_Control)}
	Roots[`/unicode/Kaithi`] = VarRoot{Var: reflect.ValueOf(&unicode.Kaithi)}
	Roots[`/unicode/Kannada`] = VarRoot{Var: reflect.ValueOf(&unicode.Kannada)}
	Roots[`/unicode/Katakana`] = VarRoot{Var: reflect.ValueOf(&unicode.Katakana)}
	Roots[`/unicode/Kayah_Li`] = VarRoot{Var: reflect.ValueOf(&unicode.Kayah_Li)}
	Roots[`/unicode/Kharoshthi`] = VarRoot{Var: reflect.ValueOf(&unicode.Kharoshthi)}
	Roots[`/unicode/Khmer`] = VarRoot{Var: reflect.ValueOf(&unicode.Khmer)}
	Roots[`/unicode/L`] = VarRoot{Var: reflect.ValueOf(&unicode.L)}
	Roots[`/unicode/Lao`] = VarRoot{Var: reflect.ValueOf(&unicode.Lao)}
	Roots[`/unicode/Latin`] = VarRoot{Var: reflect.ValueOf(&unicode.Latin)}
	Roots[`/unicode/Lepcha`] = VarRoot{Var: reflect.ValueOf(&unicode.Lepcha)}
	Roots[`/unicode/Letter`] = VarRoot{Var: reflect.ValueOf(&unicode.Letter)}
	Roots[`/unicode/Limbu`] = VarRoot{Var: reflect.ValueOf(&unicode.Limbu)}
	Roots[`/unicode/Linear_B`] = VarRoot{Var: reflect.ValueOf(&unicode.Linear_B)}
	Roots[`/unicode/Lisu`] = VarRoot{Var: reflect.ValueOf(&unicode.Lisu)}
	Roots[`/unicode/Ll`] = VarRoot{Var: reflect.ValueOf(&unicode.Ll)}
	Roots[`/unicode/Lm`] = VarRoot{Var: reflect.ValueOf(&unicode.Lm)}
	Roots[`/unicode/Lo`] = VarRoot{Var: reflect.ValueOf(&unicode.Lo)}
	Roots[`/unicode/Logical_Order_Exception`] = VarRoot{Var: reflect.ValueOf(&unicode.Logical_Order_Exception)}
	Roots[`/unicode/Lower`] = VarRoot{Var: reflect.ValueOf(&unicode.Lower)}
	Roots[`/unicode/Lt`] = VarRoot{Var: reflect.ValueOf(&unicode.Lt)}
	Roots[`/unicode/Lu`] = VarRoot{Var: reflect.ValueOf(&unicode.Lu)}
	Roots[`/unicode/Lycian`] = VarRoot{Var: reflect.ValueOf(&unicode.Lycian)}
	Roots[`/unicode/Lydian`] = VarRoot{Var: reflect.ValueOf(&unicode.Lydian)}
	Roots[`/unicode/M`] = VarRoot{Var: reflect.ValueOf(&unicode.M)}
	Roots[`/unicode/Malayalam`] = VarRoot{Var: reflect.ValueOf(&unicode.Malayalam)}
	Roots[`/unicode/Mandaic`] = VarRoot{Var: reflect.ValueOf(&unicode.Mandaic)}
	Roots[`/unicode/Mark`] = VarRoot{Var: reflect.ValueOf(&unicode.Mark)}
	Roots[`/unicode/Mc`] = VarRoot{Var: reflect.ValueOf(&unicode.Mc)}
	Roots[`/unicode/Me`] = VarRoot{Var: reflect.ValueOf(&unicode.Me)}
	Roots[`/unicode/Meetei_Mayek`] = VarRoot{Var: reflect.ValueOf(&unicode.Meetei_Mayek)}
	Roots[`/unicode/Mn`] = VarRoot{Var: reflect.ValueOf(&unicode.Mn)}
	Roots[`/unicode/Mongolian`] = VarRoot{Var: reflect.ValueOf(&unicode.Mongolian)}
	Roots[`/unicode/Myanmar`] = VarRoot{Var: reflect.ValueOf(&unicode.Myanmar)}
	Roots[`/unicode/N`] = VarRoot{Var: reflect.ValueOf(&unicode.N)}
	Roots[`/unicode/Nd`] = VarRoot{Var: reflect.ValueOf(&unicode.Nd)}
	Roots[`/unicode/New_Tai_Lue`] = VarRoot{Var: reflect.ValueOf(&unicode.New_Tai_Lue)}
	Roots[`/unicode/Nko`] = VarRoot{Var: reflect.ValueOf(&unicode.Nko)}
	Roots[`/unicode/Nl`] = VarRoot{Var: reflect.ValueOf(&unicode.Nl)}
	Roots[`/unicode/No`] = VarRoot{Var: reflect.ValueOf(&unicode.No)}
	Roots[`/unicode/Noncharacter_Code_Point`] = VarRoot{Var: reflect.ValueOf(&unicode.Noncharacter_Code_Point)}
	Roots[`/unicode/Number`] = VarRoot{Var: reflect.ValueOf(&unicode.Number)}
	Roots[`/unicode/Ogham`] = VarRoot{Var: reflect.ValueOf(&unicode.Ogham)}
	Roots[`/unicode/Ol_Chiki`] = VarRoot{Var: reflect.ValueOf(&unicode.Ol_Chiki)}
	Roots[`/unicode/Old_Italic`] = VarRoot{Var: reflect.ValueOf(&unicode.Old_Italic)}
	Roots[`/unicode/Old_Persian`] = VarRoot{Var: reflect.ValueOf(&unicode.Old_Persian)}
	Roots[`/unicode/Old_South_Arabian`] = VarRoot{Var: reflect.ValueOf(&unicode.Old_South_Arabian)}
	Roots[`/unicode/Old_Turkic`] = VarRoot{Var: reflect.ValueOf(&unicode.Old_Turkic)}
	Roots[`/unicode/Oriya`] = VarRoot{Var: reflect.ValueOf(&unicode.Oriya)}
	Roots[`/unicode/Osmanya`] = VarRoot{Var: reflect.ValueOf(&unicode.Osmanya)}
	Roots[`/unicode/Other`] = VarRoot{Var: reflect.ValueOf(&unicode.Other)}
	Roots[`/unicode/Other_Alphabetic`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Alphabetic)}
	Roots[`/unicode/Other_Default_Ignorable_Code_Point`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Default_Ignorable_Code_Point)}
	Roots[`/unicode/Other_Grapheme_Extend`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Grapheme_Extend)}
	Roots[`/unicode/Other_ID_Continue`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_ID_Continue)}
	Roots[`/unicode/Other_ID_Start`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_ID_Start)}
	Roots[`/unicode/Other_Lowercase`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Lowercase)}
	Roots[`/unicode/Other_Math`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Math)}
	Roots[`/unicode/Other_Uppercase`] = VarRoot{Var: reflect.ValueOf(&unicode.Other_Uppercase)}
	Roots[`/unicode/P`] = VarRoot{Var: reflect.ValueOf(&unicode.P)}
	Roots[`/unicode/Pattern_Syntax`] = VarRoot{Var: reflect.ValueOf(&unicode.Pattern_Syntax)}
	Roots[`/unicode/Pattern_White_Space`] = VarRoot{Var: reflect.ValueOf(&unicode.Pattern_White_Space)}
	Roots[`/unicode/Pc`] = VarRoot{Var: reflect.ValueOf(&unicode.Pc)}
	Roots[`/unicode/Pd`] = VarRoot{Var: reflect.ValueOf(&unicode.Pd)}
	Roots[`/unicode/Pe`] = VarRoot{Var: reflect.ValueOf(&unicode.Pe)}
	Roots[`/unicode/Pf`] = VarRoot{Var: reflect.ValueOf(&unicode.Pf)}
	Roots[`/unicode/Phags_Pa`] = VarRoot{Var: reflect.ValueOf(&unicode.Phags_Pa)}
	Roots[`/unicode/Phoenician`] = VarRoot{Var: reflect.ValueOf(&unicode.Phoenician)}
	Roots[`/unicode/Pi`] = VarRoot{Var: reflect.ValueOf(&unicode.Pi)}
	Roots[`/unicode/Po`] = VarRoot{Var: reflect.ValueOf(&unicode.Po)}
	Roots[`/unicode/PrintRanges`] = VarRoot{Var: reflect.ValueOf(&unicode.PrintRanges)}
	Roots[`/unicode/Properties`] = VarRoot{Var: reflect.ValueOf(&unicode.Properties)}
	Roots[`/unicode/Ps`] = VarRoot{Var: reflect.ValueOf(&unicode.Ps)}
	Roots[`/unicode/Punct`] = VarRoot{Var: reflect.ValueOf(&unicode.Punct)}
	Roots[`/unicode/Quotation_Mark`] = VarRoot{Var: reflect.ValueOf(&unicode.Quotation_Mark)}
	Roots[`/unicode/Radical`] = VarRoot{Var: reflect.ValueOf(&unicode.Radical)}
	Roots[`/unicode/Rejang`] = VarRoot{Var: reflect.ValueOf(&unicode.Rejang)}
	Roots[`/unicode/Runic`] = VarRoot{Var: reflect.ValueOf(&unicode.Runic)}
	Roots[`/unicode/S`] = VarRoot{Var: reflect.ValueOf(&unicode.S)}
	Roots[`/unicode/STerm`] = VarRoot{Var: reflect.ValueOf(&unicode.STerm)}
	Roots[`/unicode/Samaritan`] = VarRoot{Var: reflect.ValueOf(&unicode.Samaritan)}
	Roots[`/unicode/Saurashtra`] = VarRoot{Var: reflect.ValueOf(&unicode.Saurashtra)}
	Roots[`/unicode/Sc`] = VarRoot{Var: reflect.ValueOf(&unicode.Sc)}
	Roots[`/unicode/Scripts`] = VarRoot{Var: reflect.ValueOf(&unicode.Scripts)}
	Roots[`/unicode/Shavian`] = VarRoot{Var: reflect.ValueOf(&unicode.Shavian)}
	Roots[`/unicode/Sinhala`] = VarRoot{Var: reflect.ValueOf(&unicode.Sinhala)}
	Roots[`/unicode/Sk`] = VarRoot{Var: reflect.ValueOf(&unicode.Sk)}
	Roots[`/unicode/Sm`] = VarRoot{Var: reflect.ValueOf(&unicode.Sm)}
	Roots[`/unicode/So`] = VarRoot{Var: reflect.ValueOf(&unicode.So)}
	Roots[`/unicode/Soft_Dotted`] = VarRoot{Var: reflect.ValueOf(&unicode.Soft_Dotted)}
	Roots[`/unicode/Space`] = VarRoot{Var: reflect.ValueOf(&unicode.Space)}
	Roots[`/unicode/Sundanese`] = VarRoot{Var: reflect.ValueOf(&unicode.Sundanese)}
	Roots[`/unicode/Syloti_Nagri`] = VarRoot{Var: reflect.ValueOf(&unicode.Syloti_Nagri)}
	Roots[`/unicode/Symbol`] = VarRoot{Var: reflect.ValueOf(&unicode.Symbol)}
	Roots[`/unicode/Syriac`] = VarRoot{Var: reflect.ValueOf(&unicode.Syriac)}
	Roots[`/unicode/Tagalog`] = VarRoot{Var: reflect.ValueOf(&unicode.Tagalog)}
	Roots[`/unicode/Tagbanwa`] = VarRoot{Var: reflect.ValueOf(&unicode.Tagbanwa)}
	Roots[`/unicode/Tai_Le`] = VarRoot{Var: reflect.ValueOf(&unicode.Tai_Le)}
	Roots[`/unicode/Tai_Tham`] = VarRoot{Var: reflect.ValueOf(&unicode.Tai_Tham)}
	Roots[`/unicode/Tai_Viet`] = VarRoot{Var: reflect.ValueOf(&unicode.Tai_Viet)}
	Roots[`/unicode/Tamil`] = VarRoot{Var: reflect.ValueOf(&unicode.Tamil)}
	Roots[`/unicode/Telugu`] = VarRoot{Var: reflect.ValueOf(&unicode.Telugu)}
	Roots[`/unicode/Terminal_Punctuation`] = VarRoot{Var: reflect.ValueOf(&unicode.Terminal_Punctuation)}
	Roots[`/unicode/Thaana`] = VarRoot{Var: reflect.ValueOf(&unicode.Thaana)}
	Roots[`/unicode/Thai`] = VarRoot{Var: reflect.ValueOf(&unicode.Thai)}
	Roots[`/unicode/Tibetan`] = VarRoot{Var: reflect.ValueOf(&unicode.Tibetan)}
	Roots[`/unicode/Tifinagh`] = VarRoot{Var: reflect.ValueOf(&unicode.Tifinagh)}
	Roots[`/unicode/Title`] = VarRoot{Var: reflect.ValueOf(&unicode.Title)}
	Roots[`/unicode/TurkishCase`] = VarRoot{Var: reflect.ValueOf(&unicode.TurkishCase)}
	Roots[`/unicode/Ugaritic`] = VarRoot{Var: reflect.ValueOf(&unicode.Ugaritic)}
	Roots[`/unicode/Unified_Ideograph`] = VarRoot{Var: reflect.ValueOf(&unicode.Unified_Ideograph)}
	Roots[`/unicode/Upper`] = VarRoot{Var: reflect.ValueOf(&unicode.Upper)}
	Roots[`/unicode/Vai`] = VarRoot{Var: reflect.ValueOf(&unicode.Vai)}
	Roots[`/unicode/Variation_Selector`] = VarRoot{Var: reflect.ValueOf(&unicode.Variation_Selector)}
	Roots[`/unicode/White_Space`] = VarRoot{Var: reflect.ValueOf(&unicode.White_Space)}
	Roots[`/unicode/Yi`] = VarRoot{Var: reflect.ValueOf(&unicode.Yi)}
	Roots[`/unicode/Z`] = VarRoot{Var: reflect.ValueOf(&unicode.Z)}
	Roots[`/unicode/Zl`] = VarRoot{Var: reflect.ValueOf(&unicode.Zl)}
	Roots[`/unicode/Zp`] = VarRoot{Var: reflect.ValueOf(&unicode.Zp)}
	Roots[`/unicode/Zs`] = VarRoot{Var: reflect.ValueOf(&unicode.Zs)}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Header
		Roots[`/archive/tar/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Reader
		Roots[`/archive/tar/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_tar.Writer
		Roots[`/archive/tar/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.File
		Roots[`/archive/zip/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.File
		Roots[`/archive/zip/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.FileHeader
		Roots[`/archive/zip/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.ReadCloser
		Roots[`/archive/zip/ReadCloser`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.ReadCloser
		Roots[`/archive/zip/ReadCloser`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.Reader
		Roots[`/archive/zip/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.Reader
		Roots[`/archive/zip/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.Reader
		Roots[`/archive/zip/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *archive_zip.Writer
		Roots[`/archive/zip/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bufio.ReadWriter
		Roots[`/bufio/ReadWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bufio.ReadWriter
		Roots[`/bufio/ReadWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bufio.ReadWriter
		Roots[`/bufio/ReadWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bufio.Reader
		Roots[`/bufio/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bufio.Writer
		Roots[`/bufio/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bytes.Buffer
		Roots[`/bytes/Buffer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *bytes.Reader
		Roots[`/bytes/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.ReadError
		Roots[`/compress/flate/ReadError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.ReadError
		Roots[`/compress/flate/ReadError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.ReadError
		Roots[`/compress/flate/ReadError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.WriteError
		Roots[`/compress/flate/WriteError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.WriteError
		Roots[`/compress/flate/WriteError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.WriteError
		Roots[`/compress/flate/WriteError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_flate.Writer
		Roots[`/compress/flate/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Header
		Roots[`/compress/gzip/Header`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Reader
		Roots[`/compress/gzip/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Reader
		Roots[`/compress/gzip/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Writer
		Roots[`/compress/gzip/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_gzip.Writer
		Roots[`/compress/gzip/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *compress_zlib.Writer
		Roots[`/compress/zlib/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *container_list.Element
		Roots[`/container/list/Element`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *container_list.Element
		Roots[`/container/list/Element`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *container_list.List
		Roots[`/container/list/List`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *container_ring.Ring
		Roots[`/container/ring/Ring`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *container_ring.Ring
		Roots[`/container/ring/Ring`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamReader
		Roots[`/crypto/cipher/StreamReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamReader
		Roots[`/crypto/cipher/StreamReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamReader
		Roots[`/crypto/cipher/StreamReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamWriter
		Roots[`/crypto/cipher/StreamWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamWriter
		Roots[`/crypto/cipher/StreamWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamWriter
		Roots[`/crypto/cipher/StreamWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_cipher.StreamWriter
		Roots[`/crypto/cipher/StreamWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.Parameters
		Roots[`/crypto/dsa/Parameters`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.Parameters
		Roots[`/crypto/dsa/Parameters`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.Parameters
		Roots[`/crypto/dsa/Parameters`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.Parameters
		Roots[`/crypto/dsa/Parameters`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PrivateKey
		Roots[`/crypto/dsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PrivateKey
		Roots[`/crypto/dsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PrivateKey
		Roots[`/crypto/dsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PublicKey
		Roots[`/crypto/dsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PublicKey
		Roots[`/crypto/dsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_dsa.PublicKey
		Roots[`/crypto/dsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PrivateKey
		Roots[`/crypto/ecdsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PrivateKey
		Roots[`/crypto/ecdsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PrivateKey
		Roots[`/crypto/ecdsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PublicKey
		Roots[`/crypto/ecdsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PublicKey
		Roots[`/crypto/ecdsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PublicKey
		Roots[`/crypto/ecdsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_ecdsa.PublicKey
		Roots[`/crypto/ecdsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_elliptic.CurveParams
		Roots[`/crypto/elliptic/CurveParams`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rc4.Cipher
		Roots[`/crypto/rc4/Cipher`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.CRTValue
		Roots[`/crypto/rsa/CRTValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.CRTValue
		Roots[`/crypto/rsa/CRTValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.CRTValue
		Roots[`/crypto/rsa/CRTValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.CRTValue
		Roots[`/crypto/rsa/CRTValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrecomputedValues
		Roots[`/crypto/rsa/PrecomputedValues`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrecomputedValues
		Roots[`/crypto/rsa/PrecomputedValues`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrecomputedValues
		Roots[`/crypto/rsa/PrecomputedValues`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrecomputedValues
		Roots[`/crypto/rsa/PrecomputedValues`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrecomputedValues
		Roots[`/crypto/rsa/PrecomputedValues`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrivateKey
		Roots[`/crypto/rsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrivateKey
		Roots[`/crypto/rsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrivateKey
		Roots[`/crypto/rsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrivateKey
		Roots[`/crypto/rsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PrivateKey
		Roots[`/crypto/rsa/PrivateKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PublicKey
		Roots[`/crypto/rsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PublicKey
		Roots[`/crypto/rsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_rsa.PublicKey
		Roots[`/crypto/rsa/PublicKey`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Certificate
		Roots[`/crypto/tls/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Certificate
		Roots[`/crypto/tls/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Certificate
		Roots[`/crypto/tls/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Certificate
		Roots[`/crypto/tls/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Certificate
		Roots[`/crypto/tls/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Config
		Roots[`/crypto/tls/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.Conn
		Roots[`/crypto/tls/Conn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_tls.ConnectionState
		Roots[`/crypto/tls/ConnectionState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.CertPool
		Roots[`/crypto/x509/CertPool`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.Certificate
		Roots[`/crypto/x509/Certificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.CertificateInvalidError
		Roots[`/crypto/x509/CertificateInvalidError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.CertificateInvalidError
		Roots[`/crypto/x509/CertificateInvalidError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.CertificateInvalidError
		Roots[`/crypto/x509/CertificateInvalidError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.ConstraintViolationError
		Roots[`/crypto/x509/ConstraintViolationError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.HostnameError
		Roots[`/crypto/x509/HostnameError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.HostnameError
		Roots[`/crypto/x509/HostnameError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.HostnameError
		Roots[`/crypto/x509/HostnameError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.UnhandledCriticalExtension
		Roots[`/crypto/x509/UnhandledCriticalExtension`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.UnknownAuthorityError
		Roots[`/crypto/x509/UnknownAuthorityError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.VerifyOptions
		Roots[`/crypto/x509/VerifyOptions`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.VerifyOptions
		Roots[`/crypto/x509/VerifyOptions`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.VerifyOptions
		Roots[`/crypto/x509/VerifyOptions`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.VerifyOptions
		Roots[`/crypto/x509/VerifyOptions`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509.VerifyOptions
		Roots[`/crypto/x509/VerifyOptions`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AlgorithmIdentifier
		Roots[`/crypto/x509/pkix/AlgorithmIdentifier`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AlgorithmIdentifier
		Roots[`/crypto/x509/pkix/AlgorithmIdentifier`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AlgorithmIdentifier
		Roots[`/crypto/x509/pkix/AlgorithmIdentifier`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AttributeTypeAndValue
		Roots[`/crypto/x509/pkix/AttributeTypeAndValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AttributeTypeAndValue
		Roots[`/crypto/x509/pkix/AttributeTypeAndValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.AttributeTypeAndValue
		Roots[`/crypto/x509/pkix/AttributeTypeAndValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.CertificateList
		Roots[`/crypto/x509/pkix/CertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.CertificateList
		Roots[`/crypto/x509/pkix/CertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.CertificateList
		Roots[`/crypto/x509/pkix/CertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.CertificateList
		Roots[`/crypto/x509/pkix/CertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Extension
		Roots[`/crypto/x509/pkix/Extension`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Extension
		Roots[`/crypto/x509/pkix/Extension`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Extension
		Roots[`/crypto/x509/pkix/Extension`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Extension
		Roots[`/crypto/x509/pkix/Extension`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.Name
		Roots[`/crypto/x509/pkix/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.RevokedCertificate
		Roots[`/crypto/x509/pkix/RevokedCertificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.RevokedCertificate
		Roots[`/crypto/x509/pkix/RevokedCertificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.RevokedCertificate
		Roots[`/crypto/x509/pkix/RevokedCertificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.RevokedCertificate
		Roots[`/crypto/x509/pkix/RevokedCertificate`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *crypto_x509_pkix.TBSCertificateList
		Roots[`/crypto/x509/pkix/TBSCertificateList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.DB
		Roots[`/database/sql/DB`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullBool
		Roots[`/database/sql/NullBool`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullBool
		Roots[`/database/sql/NullBool`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullBool
		Roots[`/database/sql/NullBool`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullFloat64
		Roots[`/database/sql/NullFloat64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullFloat64
		Roots[`/database/sql/NullFloat64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullFloat64
		Roots[`/database/sql/NullFloat64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullInt64
		Roots[`/database/sql/NullInt64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullInt64
		Roots[`/database/sql/NullInt64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullInt64
		Roots[`/database/sql/NullInt64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullString
		Roots[`/database/sql/NullString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullString
		Roots[`/database/sql/NullString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.NullString
		Roots[`/database/sql/NullString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.Row
		Roots[`/database/sql/Row`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.Rows
		Roots[`/database/sql/Rows`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.Stmt
		Roots[`/database/sql/Stmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql.Tx
		Roots[`/database/sql/Tx`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql_driver.NotNull
		Roots[`/database/sql/driver/NotNull`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql_driver.NotNull
		Roots[`/database/sql/driver/NotNull`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql_driver.Null
		Roots[`/database/sql/driver/Null`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *database_sql_driver.Null
		Roots[`/database/sql/driver/Null`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.AddrType
		Roots[`/debug/dwarf/AddrType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.AddrType
		Roots[`/debug/dwarf/AddrType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ArrayType
		Roots[`/debug/dwarf/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ArrayType
		Roots[`/debug/dwarf/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ArrayType
		Roots[`/debug/dwarf/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ArrayType
		Roots[`/debug/dwarf/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ArrayType
		Roots[`/debug/dwarf/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BasicType
		Roots[`/debug/dwarf/BasicType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BasicType
		Roots[`/debug/dwarf/BasicType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BasicType
		Roots[`/debug/dwarf/BasicType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BasicType
		Roots[`/debug/dwarf/BasicType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BoolType
		Roots[`/debug/dwarf/BoolType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.BoolType
		Roots[`/debug/dwarf/BoolType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.CharType
		Roots[`/debug/dwarf/CharType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.CharType
		Roots[`/debug/dwarf/CharType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.CommonType
		Roots[`/debug/dwarf/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.CommonType
		Roots[`/debug/dwarf/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.CommonType
		Roots[`/debug/dwarf/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ComplexType
		Roots[`/debug/dwarf/ComplexType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.ComplexType
		Roots[`/debug/dwarf/ComplexType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Data
		Roots[`/debug/dwarf/Data`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DecodeError
		Roots[`/debug/dwarf/DecodeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DecodeError
		Roots[`/debug/dwarf/DecodeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DecodeError
		Roots[`/debug/dwarf/DecodeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DecodeError
		Roots[`/debug/dwarf/DecodeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DotDotDotType
		Roots[`/debug/dwarf/DotDotDotType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.DotDotDotType
		Roots[`/debug/dwarf/DotDotDotType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Entry
		Roots[`/debug/dwarf/Entry`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Entry
		Roots[`/debug/dwarf/Entry`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Entry
		Roots[`/debug/dwarf/Entry`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Entry
		Roots[`/debug/dwarf/Entry`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Entry
		Roots[`/debug/dwarf/Entry`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumType
		Roots[`/debug/dwarf/EnumType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumType
		Roots[`/debug/dwarf/EnumType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumType
		Roots[`/debug/dwarf/EnumType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumType
		Roots[`/debug/dwarf/EnumType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumValue
		Roots[`/debug/dwarf/EnumValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumValue
		Roots[`/debug/dwarf/EnumValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.EnumValue
		Roots[`/debug/dwarf/EnumValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Field
		Roots[`/debug/dwarf/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Field
		Roots[`/debug/dwarf/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Field
		Roots[`/debug/dwarf/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FloatType
		Roots[`/debug/dwarf/FloatType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FloatType
		Roots[`/debug/dwarf/FloatType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FuncType
		Roots[`/debug/dwarf/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FuncType
		Roots[`/debug/dwarf/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FuncType
		Roots[`/debug/dwarf/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.FuncType
		Roots[`/debug/dwarf/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.IntType
		Roots[`/debug/dwarf/IntType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.IntType
		Roots[`/debug/dwarf/IntType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.PtrType
		Roots[`/debug/dwarf/PtrType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.PtrType
		Roots[`/debug/dwarf/PtrType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.PtrType
		Roots[`/debug/dwarf/PtrType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.QualType
		Roots[`/debug/dwarf/QualType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.QualType
		Roots[`/debug/dwarf/QualType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.QualType
		Roots[`/debug/dwarf/QualType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.QualType
		Roots[`/debug/dwarf/QualType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.Reader
		Roots[`/debug/dwarf/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructField
		Roots[`/debug/dwarf/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.StructType
		Roots[`/debug/dwarf/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.TypedefType
		Roots[`/debug/dwarf/TypedefType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.TypedefType
		Roots[`/debug/dwarf/TypedefType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.TypedefType
		Roots[`/debug/dwarf/TypedefType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.UcharType
		Roots[`/debug/dwarf/UcharType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.UcharType
		Roots[`/debug/dwarf/UcharType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.UintType
		Roots[`/debug/dwarf/UintType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.UintType
		Roots[`/debug/dwarf/UintType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.VoidType
		Roots[`/debug/dwarf/VoidType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_dwarf.VoidType
		Roots[`/debug/dwarf/VoidType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn32
		Roots[`/debug/elf/Dyn32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn32
		Roots[`/debug/elf/Dyn32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn32
		Roots[`/debug/elf/Dyn32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn64
		Roots[`/debug/elf/Dyn64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn64
		Roots[`/debug/elf/Dyn64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Dyn64
		Roots[`/debug/elf/Dyn64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.File
		Roots[`/debug/elf/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.File
		Roots[`/debug/elf/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.File
		Roots[`/debug/elf/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.File
		Roots[`/debug/elf/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FileHeader
		Roots[`/debug/elf/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.FormatError
		Roots[`/debug/elf/FormatError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header32
		Roots[`/debug/elf/Header32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Header64
		Roots[`/debug/elf/Header64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ImportedSymbol
		Roots[`/debug/elf/ImportedSymbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ImportedSymbol
		Roots[`/debug/elf/ImportedSymbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ImportedSymbol
		Roots[`/debug/elf/ImportedSymbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ImportedSymbol
		Roots[`/debug/elf/ImportedSymbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog
		Roots[`/debug/elf/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog
		Roots[`/debug/elf/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog
		Roots[`/debug/elf/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog32
		Roots[`/debug/elf/Prog32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Prog64
		Roots[`/debug/elf/Prog64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.ProgHeader
		Roots[`/debug/elf/ProgHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel32
		Roots[`/debug/elf/Rel32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel32
		Roots[`/debug/elf/Rel32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel32
		Roots[`/debug/elf/Rel32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel64
		Roots[`/debug/elf/Rel64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel64
		Roots[`/debug/elf/Rel64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rel64
		Roots[`/debug/elf/Rel64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela32
		Roots[`/debug/elf/Rela32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela32
		Roots[`/debug/elf/Rela32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela32
		Roots[`/debug/elf/Rela32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela32
		Roots[`/debug/elf/Rela32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela64
		Roots[`/debug/elf/Rela64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela64
		Roots[`/debug/elf/Rela64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela64
		Roots[`/debug/elf/Rela64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Rela64
		Roots[`/debug/elf/Rela64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section
		Roots[`/debug/elf/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section
		Roots[`/debug/elf/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section
		Roots[`/debug/elf/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section32
		Roots[`/debug/elf/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Section64
		Roots[`/debug/elf/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.SectionHeader
		Roots[`/debug/elf/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym32
		Roots[`/debug/elf/Sym32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Sym64
		Roots[`/debug/elf/Sym64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_elf.Symbol
		Roots[`/debug/elf/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.DecodingError
		Roots[`/debug/gosym/DecodingError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Func
		Roots[`/debug/gosym/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.LineTable
		Roots[`/debug/gosym/LineTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.LineTable
		Roots[`/debug/gosym/LineTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.LineTable
		Roots[`/debug/gosym/LineTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.LineTable
		Roots[`/debug/gosym/LineTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Obj
		Roots[`/debug/gosym/Obj`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Obj
		Roots[`/debug/gosym/Obj`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Obj
		Roots[`/debug/gosym/Obj`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Sym
		Roots[`/debug/gosym/Sym`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Table
		Roots[`/debug/gosym/Table`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Table
		Roots[`/debug/gosym/Table`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Table
		Roots[`/debug/gosym/Table`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Table
		Roots[`/debug/gosym/Table`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.Table
		Roots[`/debug/gosym/Table`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.UnknownLineError
		Roots[`/debug/gosym/UnknownLineError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.UnknownLineError
		Roots[`/debug/gosym/UnknownLineError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_gosym.UnknownLineError
		Roots[`/debug/gosym/UnknownLineError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dylib
		Roots[`/debug/macho/Dylib`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DylibCmd
		Roots[`/debug/macho/DylibCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dysymtab
		Roots[`/debug/macho/Dysymtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dysymtab
		Roots[`/debug/macho/Dysymtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dysymtab
		Roots[`/debug/macho/Dysymtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Dysymtab
		Roots[`/debug/macho/Dysymtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.DysymtabCmd
		Roots[`/debug/macho/DysymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.File
		Roots[`/debug/macho/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FileHeader
		Roots[`/debug/macho/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.FormatError
		Roots[`/debug/macho/FormatError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist32
		Roots[`/debug/macho/Nlist32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Nlist64
		Roots[`/debug/macho/Nlist64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Regs386
		Roots[`/debug/macho/Regs386`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.RegsAMD64
		Roots[`/debug/macho/RegsAMD64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section
		Roots[`/debug/macho/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section
		Roots[`/debug/macho/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section
		Roots[`/debug/macho/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section32
		Roots[`/debug/macho/Section32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Section64
		Roots[`/debug/macho/Section64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SectionHeader
		Roots[`/debug/macho/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment
		Roots[`/debug/macho/Segment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment
		Roots[`/debug/macho/Segment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment
		Roots[`/debug/macho/Segment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment
		Roots[`/debug/macho/Segment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment32
		Roots[`/debug/macho/Segment32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Segment64
		Roots[`/debug/macho/Segment64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SegmentHeader
		Roots[`/debug/macho/SegmentHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symbol
		Roots[`/debug/macho/Symbol`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symtab
		Roots[`/debug/macho/Symtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symtab
		Roots[`/debug/macho/Symtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symtab
		Roots[`/debug/macho/Symtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Symtab
		Roots[`/debug/macho/Symtab`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.SymtabCmd
		Roots[`/debug/macho/SymtabCmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Thread
		Roots[`/debug/macho/Thread`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Thread
		Roots[`/debug/macho/Thread`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Thread
		Roots[`/debug/macho/Thread`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Thread
		Roots[`/debug/macho/Thread`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_macho.Thread
		Roots[`/debug/macho/Thread`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.File
		Roots[`/debug/pe/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.File
		Roots[`/debug/pe/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.File
		Roots[`/debug/pe/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FileHeader
		Roots[`/debug/pe/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.FormatError
		Roots[`/debug/pe/FormatError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.ImportDirectory
		Roots[`/debug/pe/ImportDirectory`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.Section
		Roots[`/debug/pe/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.Section
		Roots[`/debug/pe/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.Section
		Roots[`/debug/pe/Section`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader
		Roots[`/debug/pe/SectionHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *debug_pe.SectionHeader32
		Roots[`/debug/pe/SectionHeader32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.BitString
		Roots[`/encoding/asn1/BitString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.BitString
		Roots[`/encoding/asn1/BitString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.BitString
		Roots[`/encoding/asn1/BitString`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.RawValue
		Roots[`/encoding/asn1/RawValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.StructuralError
		Roots[`/encoding/asn1/StructuralError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.StructuralError
		Roots[`/encoding/asn1/StructuralError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.SyntaxError
		Roots[`/encoding/asn1/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_asn1.SyntaxError
		Roots[`/encoding/asn1/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_base32.Encoding
		Roots[`/encoding/base32/Encoding`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_base64.Encoding
		Roots[`/encoding/base64/Encoding`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.ParseError
		Roots[`/encoding/csv/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.ParseError
		Roots[`/encoding/csv/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.ParseError
		Roots[`/encoding/csv/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.ParseError
		Roots[`/encoding/csv/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Reader
		Roots[`/encoding/csv/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Writer
		Roots[`/encoding/csv/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Writer
		Roots[`/encoding/csv/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_csv.Writer
		Roots[`/encoding/csv/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_gob.CommonType
		Roots[`/encoding/gob/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_gob.CommonType
		Roots[`/encoding/gob/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_gob.CommonType
		Roots[`/encoding/gob/CommonType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_gob.Decoder
		Roots[`/encoding/gob/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_gob.Encoder
		Roots[`/encoding/gob/Encoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.Decoder
		Roots[`/encoding/json/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.Encoder
		Roots[`/encoding/json/Encoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.InvalidUTF8Error
		Roots[`/encoding/json/InvalidUTF8Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.InvalidUTF8Error
		Roots[`/encoding/json/InvalidUTF8Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.InvalidUnmarshalError
		Roots[`/encoding/json/InvalidUnmarshalError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.InvalidUnmarshalError
		Roots[`/encoding/json/InvalidUnmarshalError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.MarshalerError
		Roots[`/encoding/json/MarshalerError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.MarshalerError
		Roots[`/encoding/json/MarshalerError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.MarshalerError
		Roots[`/encoding/json/MarshalerError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.SyntaxError
		Roots[`/encoding/json/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.SyntaxError
		Roots[`/encoding/json/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalFieldError
		Roots[`/encoding/json/UnmarshalFieldError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalFieldError
		Roots[`/encoding/json/UnmarshalFieldError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalFieldError
		Roots[`/encoding/json/UnmarshalFieldError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalFieldError
		Roots[`/encoding/json/UnmarshalFieldError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalTypeError
		Roots[`/encoding/json/UnmarshalTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalTypeError
		Roots[`/encoding/json/UnmarshalTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnmarshalTypeError
		Roots[`/encoding/json/UnmarshalTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnsupportedTypeError
		Roots[`/encoding/json/UnsupportedTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnsupportedTypeError
		Roots[`/encoding/json/UnsupportedTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnsupportedValueError
		Roots[`/encoding/json/UnsupportedValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnsupportedValueError
		Roots[`/encoding/json/UnsupportedValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_json.UnsupportedValueError
		Roots[`/encoding/json/UnsupportedValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_pem.Block
		Roots[`/encoding/pem/Block`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_pem.Block
		Roots[`/encoding/pem/Block`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_pem.Block
		Roots[`/encoding/pem/Block`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_pem.Block
		Roots[`/encoding/pem/Block`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Attr
		Roots[`/encoding/xml/Attr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Attr
		Roots[`/encoding/xml/Attr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Attr
		Roots[`/encoding/xml/Attr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Decoder
		Roots[`/encoding/xml/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Decoder
		Roots[`/encoding/xml/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Decoder
		Roots[`/encoding/xml/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Decoder
		Roots[`/encoding/xml/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Decoder
		Roots[`/encoding/xml/Decoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Encoder
		Roots[`/encoding/xml/Encoder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.EndElement
		Roots[`/encoding/xml/EndElement`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.EndElement
		Roots[`/encoding/xml/EndElement`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Name
		Roots[`/encoding/xml/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Name
		Roots[`/encoding/xml/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.Name
		Roots[`/encoding/xml/Name`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.ProcInst
		Roots[`/encoding/xml/ProcInst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.ProcInst
		Roots[`/encoding/xml/ProcInst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.ProcInst
		Roots[`/encoding/xml/ProcInst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.StartElement
		Roots[`/encoding/xml/StartElement`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.StartElement
		Roots[`/encoding/xml/StartElement`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.StartElement
		Roots[`/encoding/xml/StartElement`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.SyntaxError
		Roots[`/encoding/xml/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.SyntaxError
		Roots[`/encoding/xml/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.SyntaxError
		Roots[`/encoding/xml/SyntaxError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.TagPathError
		Roots[`/encoding/xml/TagPathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.UnsupportedTypeError
		Roots[`/encoding/xml/UnsupportedTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *encoding_xml.UnsupportedTypeError
		Roots[`/encoding/xml/UnsupportedTypeError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.Float
		Roots[`/expvar/Float`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.Int
		Roots[`/expvar/Int`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.KeyValue
		Roots[`/expvar/KeyValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.KeyValue
		Roots[`/expvar/KeyValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.KeyValue
		Roots[`/expvar/KeyValue`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.Map
		Roots[`/expvar/Map`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *expvar.String
		Roots[`/expvar/String`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.Flag
		Roots[`/flag/Flag`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.Flag
		Roots[`/flag/Flag`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.Flag
		Roots[`/flag/Flag`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.Flag
		Roots[`/flag/Flag`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.Flag
		Roots[`/flag/Flag`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.FlagSet
		Roots[`/flag/FlagSet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *flag.FlagSet
		Roots[`/flag/FlagSet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ArrayType
		Roots[`/go/ast/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ArrayType
		Roots[`/go/ast/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ArrayType
		Roots[`/go/ast/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ArrayType
		Roots[`/go/ast/ArrayType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.AssignStmt
		Roots[`/go/ast/AssignStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.AssignStmt
		Roots[`/go/ast/AssignStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.AssignStmt
		Roots[`/go/ast/AssignStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.AssignStmt
		Roots[`/go/ast/AssignStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.AssignStmt
		Roots[`/go/ast/AssignStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadDecl
		Roots[`/go/ast/BadDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadDecl
		Roots[`/go/ast/BadDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadDecl
		Roots[`/go/ast/BadDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadExpr
		Roots[`/go/ast/BadExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadExpr
		Roots[`/go/ast/BadExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadExpr
		Roots[`/go/ast/BadExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadStmt
		Roots[`/go/ast/BadStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadStmt
		Roots[`/go/ast/BadStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BadStmt
		Roots[`/go/ast/BadStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BasicLit
		Roots[`/go/ast/BasicLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BasicLit
		Roots[`/go/ast/BasicLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BasicLit
		Roots[`/go/ast/BasicLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BasicLit
		Roots[`/go/ast/BasicLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BinaryExpr
		Roots[`/go/ast/BinaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BinaryExpr
		Roots[`/go/ast/BinaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BinaryExpr
		Roots[`/go/ast/BinaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BinaryExpr
		Roots[`/go/ast/BinaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BinaryExpr
		Roots[`/go/ast/BinaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BlockStmt
		Roots[`/go/ast/BlockStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BlockStmt
		Roots[`/go/ast/BlockStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BlockStmt
		Roots[`/go/ast/BlockStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BlockStmt
		Roots[`/go/ast/BlockStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BranchStmt
		Roots[`/go/ast/BranchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BranchStmt
		Roots[`/go/ast/BranchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BranchStmt
		Roots[`/go/ast/BranchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.BranchStmt
		Roots[`/go/ast/BranchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CallExpr
		Roots[`/go/ast/CallExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CaseClause
		Roots[`/go/ast/CaseClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CaseClause
		Roots[`/go/ast/CaseClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CaseClause
		Roots[`/go/ast/CaseClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CaseClause
		Roots[`/go/ast/CaseClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CaseClause
		Roots[`/go/ast/CaseClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ChanType
		Roots[`/go/ast/ChanType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ChanType
		Roots[`/go/ast/ChanType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ChanType
		Roots[`/go/ast/ChanType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ChanType
		Roots[`/go/ast/ChanType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommClause
		Roots[`/go/ast/CommClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommClause
		Roots[`/go/ast/CommClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommClause
		Roots[`/go/ast/CommClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommClause
		Roots[`/go/ast/CommClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommClause
		Roots[`/go/ast/CommClause`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Comment
		Roots[`/go/ast/Comment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Comment
		Roots[`/go/ast/Comment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Comment
		Roots[`/go/ast/Comment`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommentGroup
		Roots[`/go/ast/CommentGroup`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CommentGroup
		Roots[`/go/ast/CommentGroup`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CompositeLit
		Roots[`/go/ast/CompositeLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CompositeLit
		Roots[`/go/ast/CompositeLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CompositeLit
		Roots[`/go/ast/CompositeLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CompositeLit
		Roots[`/go/ast/CompositeLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.CompositeLit
		Roots[`/go/ast/CompositeLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.DeclStmt
		Roots[`/go/ast/DeclStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.DeclStmt
		Roots[`/go/ast/DeclStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.DeferStmt
		Roots[`/go/ast/DeferStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.DeferStmt
		Roots[`/go/ast/DeferStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.DeferStmt
		Roots[`/go/ast/DeferStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ellipsis
		Roots[`/go/ast/Ellipsis`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ellipsis
		Roots[`/go/ast/Ellipsis`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ellipsis
		Roots[`/go/ast/Ellipsis`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.EmptyStmt
		Roots[`/go/ast/EmptyStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.EmptyStmt
		Roots[`/go/ast/EmptyStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ExprStmt
		Roots[`/go/ast/ExprStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ExprStmt
		Roots[`/go/ast/ExprStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Field
		Roots[`/go/ast/Field`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FieldList
		Roots[`/go/ast/FieldList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FieldList
		Roots[`/go/ast/FieldList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FieldList
		Roots[`/go/ast/FieldList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FieldList
		Roots[`/go/ast/FieldList`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.File
		Roots[`/go/ast/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ForStmt
		Roots[`/go/ast/ForStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncDecl
		Roots[`/go/ast/FuncDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncLit
		Roots[`/go/ast/FuncLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncLit
		Roots[`/go/ast/FuncLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncLit
		Roots[`/go/ast/FuncLit`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncType
		Roots[`/go/ast/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncType
		Roots[`/go/ast/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncType
		Roots[`/go/ast/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.FuncType
		Roots[`/go/ast/FuncType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GenDecl
		Roots[`/go/ast/GenDecl`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GoStmt
		Roots[`/go/ast/GoStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GoStmt
		Roots[`/go/ast/GoStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.GoStmt
		Roots[`/go/ast/GoStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ident
		Roots[`/go/ast/Ident`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ident
		Roots[`/go/ast/Ident`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ident
		Roots[`/go/ast/Ident`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Ident
		Roots[`/go/ast/Ident`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IfStmt
		Roots[`/go/ast/IfStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ImportSpec
		Roots[`/go/ast/ImportSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IncDecStmt
		Roots[`/go/ast/IncDecStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IncDecStmt
		Roots[`/go/ast/IncDecStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IncDecStmt
		Roots[`/go/ast/IncDecStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IncDecStmt
		Roots[`/go/ast/IncDecStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IndexExpr
		Roots[`/go/ast/IndexExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IndexExpr
		Roots[`/go/ast/IndexExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IndexExpr
		Roots[`/go/ast/IndexExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IndexExpr
		Roots[`/go/ast/IndexExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.IndexExpr
		Roots[`/go/ast/IndexExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.InterfaceType
		Roots[`/go/ast/InterfaceType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.InterfaceType
		Roots[`/go/ast/InterfaceType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.InterfaceType
		Roots[`/go/ast/InterfaceType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.InterfaceType
		Roots[`/go/ast/InterfaceType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.KeyValueExpr
		Roots[`/go/ast/KeyValueExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.KeyValueExpr
		Roots[`/go/ast/KeyValueExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.KeyValueExpr
		Roots[`/go/ast/KeyValueExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.KeyValueExpr
		Roots[`/go/ast/KeyValueExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.LabeledStmt
		Roots[`/go/ast/LabeledStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.LabeledStmt
		Roots[`/go/ast/LabeledStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.LabeledStmt
		Roots[`/go/ast/LabeledStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.LabeledStmt
		Roots[`/go/ast/LabeledStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.MapType
		Roots[`/go/ast/MapType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.MapType
		Roots[`/go/ast/MapType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.MapType
		Roots[`/go/ast/MapType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.MapType
		Roots[`/go/ast/MapType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Object
		Roots[`/go/ast/Object`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Package
		Roots[`/go/ast/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Package
		Roots[`/go/ast/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Package
		Roots[`/go/ast/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Package
		Roots[`/go/ast/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Package
		Roots[`/go/ast/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ParenExpr
		Roots[`/go/ast/ParenExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ParenExpr
		Roots[`/go/ast/ParenExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ParenExpr
		Roots[`/go/ast/ParenExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ParenExpr
		Roots[`/go/ast/ParenExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.RangeStmt
		Roots[`/go/ast/RangeStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ReturnStmt
		Roots[`/go/ast/ReturnStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ReturnStmt
		Roots[`/go/ast/ReturnStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ReturnStmt
		Roots[`/go/ast/ReturnStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Scope
		Roots[`/go/ast/Scope`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Scope
		Roots[`/go/ast/Scope`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.Scope
		Roots[`/go/ast/Scope`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectStmt
		Roots[`/go/ast/SelectStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectStmt
		Roots[`/go/ast/SelectStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectStmt
		Roots[`/go/ast/SelectStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectorExpr
		Roots[`/go/ast/SelectorExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectorExpr
		Roots[`/go/ast/SelectorExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SelectorExpr
		Roots[`/go/ast/SelectorExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SendStmt
		Roots[`/go/ast/SendStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SendStmt
		Roots[`/go/ast/SendStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SendStmt
		Roots[`/go/ast/SendStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SendStmt
		Roots[`/go/ast/SendStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SliceExpr
		Roots[`/go/ast/SliceExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StarExpr
		Roots[`/go/ast/StarExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StarExpr
		Roots[`/go/ast/StarExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StarExpr
		Roots[`/go/ast/StarExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StructType
		Roots[`/go/ast/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StructType
		Roots[`/go/ast/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StructType
		Roots[`/go/ast/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.StructType
		Roots[`/go/ast/StructType`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SwitchStmt
		Roots[`/go/ast/SwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SwitchStmt
		Roots[`/go/ast/SwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SwitchStmt
		Roots[`/go/ast/SwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SwitchStmt
		Roots[`/go/ast/SwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.SwitchStmt
		Roots[`/go/ast/SwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeAssertExpr
		Roots[`/go/ast/TypeAssertExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeAssertExpr
		Roots[`/go/ast/TypeAssertExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeAssertExpr
		Roots[`/go/ast/TypeAssertExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSpec
		Roots[`/go/ast/TypeSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSpec
		Roots[`/go/ast/TypeSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSpec
		Roots[`/go/ast/TypeSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSpec
		Roots[`/go/ast/TypeSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSpec
		Roots[`/go/ast/TypeSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSwitchStmt
		Roots[`/go/ast/TypeSwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSwitchStmt
		Roots[`/go/ast/TypeSwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSwitchStmt
		Roots[`/go/ast/TypeSwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSwitchStmt
		Roots[`/go/ast/TypeSwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.TypeSwitchStmt
		Roots[`/go/ast/TypeSwitchStmt`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.UnaryExpr
		Roots[`/go/ast/UnaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.UnaryExpr
		Roots[`/go/ast/UnaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.UnaryExpr
		Roots[`/go/ast/UnaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.UnaryExpr
		Roots[`/go/ast/UnaryExpr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_ast.ValueSpec
		Roots[`/go/ast/ValueSpec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Context
		Roots[`/go/build/Context`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.NoGoError
		Roots[`/go/build/NoGoError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.NoGoError
		Roots[`/go/build/NoGoError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_build.Package
		Roots[`/go/build/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Example
		Roots[`/go/doc/Example`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Func
		Roots[`/go/doc/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Package
		Roots[`/go/doc/Package`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Type
		Roots[`/go/doc/Type`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Value
		Roots[`/go/doc/Value`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Value
		Roots[`/go/doc/Value`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Value
		Roots[`/go/doc/Value`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_doc.Value
		Roots[`/go/doc/Value`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.CommentedNode
		Roots[`/go/printer/CommentedNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.CommentedNode
		Roots[`/go/printer/CommentedNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.CommentedNode
		Roots[`/go/printer/CommentedNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.Config
		Roots[`/go/printer/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.Config
		Roots[`/go/printer/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_printer.Config
		Roots[`/go/printer/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_scanner.Error
		Roots[`/go/scanner/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_scanner.Error
		Roots[`/go/scanner/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_scanner.Error
		Roots[`/go/scanner/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_scanner.Scanner
		Roots[`/go/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_scanner.Scanner
		Roots[`/go/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.File
		Roots[`/go/token/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.FileSet
		Roots[`/go/token/FileSet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.Position
		Roots[`/go/token/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.Position
		Roots[`/go/token/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.Position
		Roots[`/go/token/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.Position
		Roots[`/go/token/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *go_token.Position
		Roots[`/go/token/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Error
		Roots[`/html/template/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Error
		Roots[`/html/template/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Error
		Roots[`/html/template/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Error
		Roots[`/html/template/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Error
		Roots[`/html/template/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *html_template.Template
		Roots[`/html/template/Template`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha
		Roots[`/image/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha
		Roots[`/image/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha
		Roots[`/image/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha
		Roots[`/image/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha16
		Roots[`/image/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha16
		Roots[`/image/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha16
		Roots[`/image/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Alpha16
		Roots[`/image/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Config
		Roots[`/image/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Config
		Roots[`/image/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Config
		Roots[`/image/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Config
		Roots[`/image/Config`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray
		Roots[`/image/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray
		Roots[`/image/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray
		Roots[`/image/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray
		Roots[`/image/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray16
		Roots[`/image/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray16
		Roots[`/image/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray16
		Roots[`/image/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Gray16
		Roots[`/image/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA
		Roots[`/image/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA
		Roots[`/image/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA
		Roots[`/image/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA
		Roots[`/image/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA64
		Roots[`/image/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA64
		Roots[`/image/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA64
		Roots[`/image/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.NRGBA64
		Roots[`/image/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Paletted
		Roots[`/image/Paletted`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Paletted
		Roots[`/image/Paletted`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Paletted
		Roots[`/image/Paletted`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Paletted
		Roots[`/image/Paletted`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Paletted
		Roots[`/image/Paletted`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Point
		Roots[`/image/Point`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Point
		Roots[`/image/Point`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Point
		Roots[`/image/Point`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA
		Roots[`/image/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA
		Roots[`/image/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA
		Roots[`/image/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA
		Roots[`/image/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA64
		Roots[`/image/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA64
		Roots[`/image/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA64
		Roots[`/image/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.RGBA64
		Roots[`/image/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Rectangle
		Roots[`/image/Rectangle`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Rectangle
		Roots[`/image/Rectangle`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Rectangle
		Roots[`/image/Rectangle`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Uniform
		Roots[`/image/Uniform`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.Uniform
		Roots[`/image/Uniform`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image.YCbCr
		Roots[`/image/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Alpha
		Roots[`/image/color/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Alpha
		Roots[`/image/color/Alpha`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Alpha16
		Roots[`/image/color/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Alpha16
		Roots[`/image/color/Alpha16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Gray
		Roots[`/image/color/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Gray
		Roots[`/image/color/Gray`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Gray16
		Roots[`/image/color/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.Gray16
		Roots[`/image/color/Gray16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA
		Roots[`/image/color/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA
		Roots[`/image/color/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA
		Roots[`/image/color/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA
		Roots[`/image/color/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA
		Roots[`/image/color/NRGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA64
		Roots[`/image/color/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA64
		Roots[`/image/color/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA64
		Roots[`/image/color/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA64
		Roots[`/image/color/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.NRGBA64
		Roots[`/image/color/NRGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA
		Roots[`/image/color/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA
		Roots[`/image/color/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA
		Roots[`/image/color/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA
		Roots[`/image/color/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA
		Roots[`/image/color/RGBA`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA64
		Roots[`/image/color/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA64
		Roots[`/image/color/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA64
		Roots[`/image/color/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA64
		Roots[`/image/color/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.RGBA64
		Roots[`/image/color/RGBA64`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.YCbCr
		Roots[`/image/color/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.YCbCr
		Roots[`/image/color/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.YCbCr
		Roots[`/image/color/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_color.YCbCr
		Roots[`/image/color/YCbCr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_gif.GIF
		Roots[`/image/gif/GIF`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_gif.GIF
		Roots[`/image/gif/GIF`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_gif.GIF
		Roots[`/image/gif/GIF`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_gif.GIF
		Roots[`/image/gif/GIF`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_jpeg.Options
		Roots[`/image/jpeg/Options`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *image_jpeg.Options
		Roots[`/image/jpeg/Options`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *index_suffixarray.Index
		Roots[`/index/suffixarray/Index`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.LimitedReader
		Roots[`/io/LimitedReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.LimitedReader
		Roots[`/io/LimitedReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.LimitedReader
		Roots[`/io/LimitedReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.PipeReader
		Roots[`/io/PipeReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.PipeWriter
		Roots[`/io/PipeWriter`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *io.SectionReader
		Roots[`/io/SectionReader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *log.Logger
		Roots[`/log/Logger`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *math_big.Int
		Roots[`/math/big/Int`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *math_big.Rat
		Roots[`/math/big/Rat`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *math_rand.Rand
		Roots[`/math/rand/Rand`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *math_rand.Zipf
		Roots[`/math/rand/Zipf`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.FileHeader
		Roots[`/mime/multipart/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.FileHeader
		Roots[`/mime/multipart/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.FileHeader
		Roots[`/mime/multipart/FileHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Form
		Roots[`/mime/multipart/Form`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Form
		Roots[`/mime/multipart/Form`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Form
		Roots[`/mime/multipart/Form`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Part
		Roots[`/mime/multipart/Part`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Part
		Roots[`/mime/multipart/Part`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Reader
		Roots[`/mime/multipart/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *mime_multipart.Writer
		Roots[`/mime/multipart/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.AddrError
		Roots[`/net/AddrError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.AddrError
		Roots[`/net/AddrError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.AddrError
		Roots[`/net/AddrError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSConfigError
		Roots[`/net/DNSConfigError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSConfigError
		Roots[`/net/DNSConfigError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSError
		Roots[`/net/DNSError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSError
		Roots[`/net/DNSError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSError
		Roots[`/net/DNSError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSError
		Roots[`/net/DNSError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.DNSError
		Roots[`/net/DNSError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPAddr
		Roots[`/net/IPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPAddr
		Roots[`/net/IPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPConn
		Roots[`/net/IPConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPNet
		Roots[`/net/IPNet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPNet
		Roots[`/net/IPNet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.IPNet
		Roots[`/net/IPNet`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.Interface
		Roots[`/net/Interface`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.MX
		Roots[`/net/MX`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.MX
		Roots[`/net/MX`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.MX
		Roots[`/net/MX`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.OpError
		Roots[`/net/OpError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.OpError
		Roots[`/net/OpError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.OpError
		Roots[`/net/OpError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.OpError
		Roots[`/net/OpError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.OpError
		Roots[`/net/OpError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.ParseError
		Roots[`/net/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.ParseError
		Roots[`/net/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.ParseError
		Roots[`/net/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.SRV
		Roots[`/net/SRV`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.SRV
		Roots[`/net/SRV`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.SRV
		Roots[`/net/SRV`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.SRV
		Roots[`/net/SRV`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.SRV
		Roots[`/net/SRV`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.TCPAddr
		Roots[`/net/TCPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.TCPAddr
		Roots[`/net/TCPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.TCPAddr
		Roots[`/net/TCPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.TCPConn
		Roots[`/net/TCPConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.TCPListener
		Roots[`/net/TCPListener`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UDPAddr
		Roots[`/net/UDPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UDPAddr
		Roots[`/net/UDPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UDPAddr
		Roots[`/net/UDPAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UDPConn
		Roots[`/net/UDPConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UnixAddr
		Roots[`/net/UnixAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UnixAddr
		Roots[`/net/UnixAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UnixAddr
		Roots[`/net/UnixAddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UnixConn
		Roots[`/net/UnixConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net.UnixListener
		Roots[`/net/UnixListener`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Client
		Roots[`/net/http/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Client
		Roots[`/net/http/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Client
		Roots[`/net/http/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Client
		Roots[`/net/http/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Cookie
		Roots[`/net/http/Cookie`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.ProtocolError
		Roots[`/net/http/ProtocolError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.ProtocolError
		Roots[`/net/http/ProtocolError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Request
		Roots[`/net/http/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Response
		Roots[`/net/http/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.ServeMux
		Roots[`/net/http/ServeMux`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Server
		Roots[`/net/http/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http.Transport
		Roots[`/net/http/Transport`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_cgi.Handler
		Roots[`/net/http/cgi/Handler`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.ResponseRecorder
		Roots[`/net/http/httptest/ResponseRecorder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.ResponseRecorder
		Roots[`/net/http/httptest/ResponseRecorder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.ResponseRecorder
		Roots[`/net/http/httptest/ResponseRecorder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.ResponseRecorder
		Roots[`/net/http/httptest/ResponseRecorder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.ResponseRecorder
		Roots[`/net/http/httptest/ResponseRecorder`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.Server
		Roots[`/net/http/httptest/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.Server
		Roots[`/net/http/httptest/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.Server
		Roots[`/net/http/httptest/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.Server
		Roots[`/net/http/httptest/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httptest.Server
		Roots[`/net/http/httptest/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ClientConn
		Roots[`/net/http/httputil/ClientConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ReverseProxy
		Roots[`/net/http/httputil/ReverseProxy`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ReverseProxy
		Roots[`/net/http/httputil/ReverseProxy`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ReverseProxy
		Roots[`/net/http/httputil/ReverseProxy`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ReverseProxy
		Roots[`/net/http/httputil/ReverseProxy`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_http_httputil.ServerConn
		Roots[`/net/http/httputil/ServerConn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Address
		Roots[`/net/mail/Address`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Address
		Roots[`/net/mail/Address`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Address
		Roots[`/net/mail/Address`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Message
		Roots[`/net/mail/Message`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Message
		Roots[`/net/mail/Message`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_mail.Message
		Roots[`/net/mail/Message`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Call
		Roots[`/net/rpc/Call`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Client
		Roots[`/net/rpc/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Request
		Roots[`/net/rpc/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Request
		Roots[`/net/rpc/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Request
		Roots[`/net/rpc/Request`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Response
		Roots[`/net/rpc/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Response
		Roots[`/net/rpc/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Response
		Roots[`/net/rpc/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Response
		Roots[`/net/rpc/Response`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_rpc.Server
		Roots[`/net/rpc/Server`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.Client
		Roots[`/net/smtp/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.Client
		Roots[`/net/smtp/Client`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.ServerInfo
		Roots[`/net/smtp/ServerInfo`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.ServerInfo
		Roots[`/net/smtp/ServerInfo`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.ServerInfo
		Roots[`/net/smtp/ServerInfo`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_smtp.ServerInfo
		Roots[`/net/smtp/ServerInfo`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Conn
		Roots[`/net/textproto/Conn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Conn
		Roots[`/net/textproto/Conn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Conn
		Roots[`/net/textproto/Conn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Conn
		Roots[`/net/textproto/Conn`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Error
		Roots[`/net/textproto/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Error
		Roots[`/net/textproto/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Error
		Roots[`/net/textproto/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Pipeline
		Roots[`/net/textproto/Pipeline`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Reader
		Roots[`/net/textproto/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Reader
		Roots[`/net/textproto/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Writer
		Roots[`/net/textproto/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_textproto.Writer
		Roots[`/net/textproto/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.Error
		Roots[`/net/url/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.Error
		Roots[`/net/url/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.Error
		Roots[`/net/url/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.Error
		Roots[`/net/url/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.URL
		Roots[`/net/url/URL`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *net_url.Userinfo
		Roots[`/net/url/Userinfo`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.File
		Roots[`/os/File`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.LinkError
		Roots[`/os/LinkError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.LinkError
		Roots[`/os/LinkError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.LinkError
		Roots[`/os/LinkError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.LinkError
		Roots[`/os/LinkError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.LinkError
		Roots[`/os/LinkError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.PathError
		Roots[`/os/PathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.PathError
		Roots[`/os/PathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.PathError
		Roots[`/os/PathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.PathError
		Roots[`/os/PathError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcAttr
		Roots[`/os/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcAttr
		Roots[`/os/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcAttr
		Roots[`/os/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcAttr
		Roots[`/os/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcAttr
		Roots[`/os/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.Process
		Roots[`/os/Process`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.Process
		Roots[`/os/Process`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.ProcessState
		Roots[`/os/ProcessState`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.SyscallError
		Roots[`/os/SyscallError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.SyscallError
		Roots[`/os/SyscallError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os.SyscallError
		Roots[`/os/SyscallError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Cmd
		Roots[`/os/exec/Cmd`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Error
		Roots[`/os/exec/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Error
		Roots[`/os/exec/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.Error
		Roots[`/os/exec/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.ExitError
		Roots[`/os/exec/ExitError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_exec.ExitError
		Roots[`/os/exec/ExitError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *os_user.User
		Roots[`/os/user/User`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Method
		Roots[`/reflect/Method`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.SliceHeader
		Roots[`/reflect/SliceHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.SliceHeader
		Roots[`/reflect/SliceHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.SliceHeader
		Roots[`/reflect/SliceHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.SliceHeader
		Roots[`/reflect/SliceHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StringHeader
		Roots[`/reflect/StringHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StringHeader
		Roots[`/reflect/StringHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StringHeader
		Roots[`/reflect/StringHeader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.StructField
		Roots[`/reflect/StructField`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.Value
		Roots[`/reflect/Value`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.ValueError
		Roots[`/reflect/ValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.ValueError
		Roots[`/reflect/ValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *reflect.ValueError
		Roots[`/reflect/ValueError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp.Regexp
		Roots[`/regexp/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Error
		Roots[`/regexp/syntax/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Error
		Roots[`/regexp/syntax/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Error
		Roots[`/regexp/syntax/Error`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Inst
		Roots[`/regexp/syntax/Inst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Inst
		Roots[`/regexp/syntax/Inst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Inst
		Roots[`/regexp/syntax/Inst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Inst
		Roots[`/regexp/syntax/Inst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Inst
		Roots[`/regexp/syntax/Inst`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Prog
		Roots[`/regexp/syntax/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Prog
		Roots[`/regexp/syntax/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Prog
		Roots[`/regexp/syntax/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Prog
		Roots[`/regexp/syntax/Prog`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *regexp_syntax.Regexp
		Roots[`/regexp/syntax/Regexp`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.Func
		Roots[`/runtime/Func`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemProfileRecord
		Roots[`/runtime/MemProfileRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.MemStats
		Roots[`/runtime/MemStats`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.StackRecord
		Roots[`/runtime/StackRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.StackRecord
		Roots[`/runtime/StackRecord`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime.TypeAssertionError
		Roots[`/runtime/TypeAssertionError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *runtime_pprof.Profile
		Roots[`/runtime/pprof/Profile`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strconv.NumError
		Roots[`/strconv/NumError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strconv.NumError
		Roots[`/strconv/NumError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strconv.NumError
		Roots[`/strconv/NumError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strconv.NumError
		Roots[`/strconv/NumError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strings.Reader
		Roots[`/strings/Reader`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *strings.Replacer
		Roots[`/strings/Replacer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.Cond
		Roots[`/sync/Cond`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.Cond
		Roots[`/sync/Cond`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.Mutex
		Roots[`/sync/Mutex`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.Once
		Roots[`/sync/Once`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.RWMutex
		Roots[`/sync/RWMutex`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *sync.WaitGroup
		Roots[`/sync/WaitGroup`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPMreq
		Roots[`/syscall/IPMreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPMreq
		Roots[`/syscall/IPMreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPMreq
		Roots[`/syscall/IPMreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPv6Mreq
		Roots[`/syscall/IPv6Mreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPv6Mreq
		Roots[`/syscall/IPv6Mreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.IPv6Mreq
		Roots[`/syscall/IPv6Mreq`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Linger
		Roots[`/syscall/Linger`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Linger
		Roots[`/syscall/Linger`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Linger
		Roots[`/syscall/Linger`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.ProcAttr
		Roots[`/syscall/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.ProcAttr
		Roots[`/syscall/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.ProcAttr
		Roots[`/syscall/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.ProcAttr
		Roots[`/syscall/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.ProcAttr
		Roots[`/syscall/ProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddr
		Roots[`/syscall/RawSockaddr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddrAny
		Roots[`/syscall/RawSockaddrAny`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddrAny
		Roots[`/syscall/RawSockaddrAny`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddrInet4
		Roots[`/syscall/RawSockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddrInet4
		Roots[`/syscall/RawSockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.RawSockaddrInet4
		Roots[`/syscall/RawSockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Rusage
		Roots[`/syscall/Rusage`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet4
		Roots[`/syscall/SockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet4
		Roots[`/syscall/SockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet4
		Roots[`/syscall/SockaddrInet4`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet6
		Roots[`/syscall/SockaddrInet6`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet6
		Roots[`/syscall/SockaddrInet6`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet6
		Roots[`/syscall/SockaddrInet6`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrInet6
		Roots[`/syscall/SockaddrInet6`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrUnix
		Roots[`/syscall/SockaddrUnix`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SockaddrUnix
		Roots[`/syscall/SockaddrUnix`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.SysProcAttr
		Roots[`/syscall/SysProcAttr`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Timespec
		Roots[`/syscall/Timespec`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *syscall.Timeval
		Roots[`/syscall/Timeval`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Position
		Roots[`/text/scanner/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Position
		Roots[`/text/scanner/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Position
		Roots[`/text/scanner/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Position
		Roots[`/text/scanner/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Position
		Roots[`/text/scanner/Position`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_scanner.Scanner
		Roots[`/text/scanner/Scanner`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_tabwriter.Writer
		Roots[`/text/tabwriter/Writer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template.Template
		Roots[`/text/template/Template`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template.Template
		Roots[`/text/template/Template`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ActionNode
		Roots[`/text/template/parse/ActionNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ActionNode
		Roots[`/text/template/parse/ActionNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ActionNode
		Roots[`/text/template/parse/ActionNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ActionNode
		Roots[`/text/template/parse/ActionNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BoolNode
		Roots[`/text/template/parse/BoolNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BoolNode
		Roots[`/text/template/parse/BoolNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BoolNode
		Roots[`/text/template/parse/BoolNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.BranchNode
		Roots[`/text/template/parse/BranchNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.CommandNode
		Roots[`/text/template/parse/CommandNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.CommandNode
		Roots[`/text/template/parse/CommandNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.CommandNode
		Roots[`/text/template/parse/CommandNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.FieldNode
		Roots[`/text/template/parse/FieldNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.FieldNode
		Roots[`/text/template/parse/FieldNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.FieldNode
		Roots[`/text/template/parse/FieldNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.IdentifierNode
		Roots[`/text/template/parse/IdentifierNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.IdentifierNode
		Roots[`/text/template/parse/IdentifierNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.IdentifierNode
		Roots[`/text/template/parse/IdentifierNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.IfNode
		Roots[`/text/template/parse/IfNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.IfNode
		Roots[`/text/template/parse/IfNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ListNode
		Roots[`/text/template/parse/ListNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ListNode
		Roots[`/text/template/parse/ListNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.ListNode
		Roots[`/text/template/parse/ListNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.NumberNode
		Roots[`/text/template/parse/NumberNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.PipeNode
		Roots[`/text/template/parse/PipeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.PipeNode
		Roots[`/text/template/parse/PipeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.PipeNode
		Roots[`/text/template/parse/PipeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.PipeNode
		Roots[`/text/template/parse/PipeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.PipeNode
		Roots[`/text/template/parse/PipeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.RangeNode
		Roots[`/text/template/parse/RangeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.RangeNode
		Roots[`/text/template/parse/RangeNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.StringNode
		Roots[`/text/template/parse/StringNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.StringNode
		Roots[`/text/template/parse/StringNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.StringNode
		Roots[`/text/template/parse/StringNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.StringNode
		Roots[`/text/template/parse/StringNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TemplateNode
		Roots[`/text/template/parse/TemplateNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TemplateNode
		Roots[`/text/template/parse/TemplateNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TemplateNode
		Roots[`/text/template/parse/TemplateNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TemplateNode
		Roots[`/text/template/parse/TemplateNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TemplateNode
		Roots[`/text/template/parse/TemplateNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TextNode
		Roots[`/text/template/parse/TextNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TextNode
		Roots[`/text/template/parse/TextNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.TextNode
		Roots[`/text/template/parse/TextNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.Tree
		Roots[`/text/template/parse/Tree`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.Tree
		Roots[`/text/template/parse/Tree`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.Tree
		Roots[`/text/template/parse/Tree`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.VariableNode
		Roots[`/text/template/parse/VariableNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.VariableNode
		Roots[`/text/template/parse/VariableNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.VariableNode
		Roots[`/text/template/parse/VariableNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.WithNode
		Roots[`/text/template/parse/WithNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *text_template_parse.WithNode
		Roots[`/text/template/parse/WithNode`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Location
		Roots[`/time/Location`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.ParseError
		Roots[`/time/ParseError`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Ticker
		Roots[`/time/Ticker`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Ticker
		Roots[`/time/Ticker`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Time
		Roots[`/time/Time`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Timer
		Roots[`/time/Timer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *time.Timer
		Roots[`/time/Timer`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.CaseRange
		Roots[`/unicode/CaseRange`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.CaseRange
		Roots[`/unicode/CaseRange`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.CaseRange
		Roots[`/unicode/CaseRange`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.CaseRange
		Roots[`/unicode/CaseRange`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range16
		Roots[`/unicode/Range16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range16
		Roots[`/unicode/Range16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range16
		Roots[`/unicode/Range16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range16
		Roots[`/unicode/Range16`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range32
		Roots[`/unicode/Range32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range32
		Roots[`/unicode/Range32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range32
		Roots[`/unicode/Range32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.Range32
		Roots[`/unicode/Range32`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.RangeTable
		Roots[`/unicode/RangeTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.RangeTable
		Roots[`/unicode/RangeTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	{
		var tmp *unicode.RangeTable
		Roots[`/unicode/RangeTable`] = TypeRoot{Type: reflect.ValueOf(tmp).Type().Elem()}
	}
	Roots[`/archive/tar/TypeBlock`] = ConstRoot{Const: archive_tar.TypeBlock}
	Roots[`/archive/tar/TypeChar`] = ConstRoot{Const: archive_tar.TypeChar}
	Roots[`/archive/tar/TypeCont`] = ConstRoot{Const: archive_tar.TypeCont}
	Roots[`/archive/tar/TypeDir`] = ConstRoot{Const: archive_tar.TypeDir}
	Roots[`/archive/tar/TypeFifo`] = ConstRoot{Const: archive_tar.TypeFifo}
	Roots[`/archive/tar/TypeLink`] = ConstRoot{Const: archive_tar.TypeLink}
	Roots[`/archive/tar/TypeReg`] = ConstRoot{Const: archive_tar.TypeReg}
	Roots[`/archive/tar/TypeRegA`] = ConstRoot{Const: archive_tar.TypeRegA}
	Roots[`/archive/tar/TypeSymlink`] = ConstRoot{Const: archive_tar.TypeSymlink}
	Roots[`/archive/tar/TypeXGlobalHeader`] = ConstRoot{Const: archive_tar.TypeXGlobalHeader}
	Roots[`/archive/tar/TypeXHeader`] = ConstRoot{Const: archive_tar.TypeXHeader}
	Roots[`/archive/zip/Deflate`] = ConstRoot{Const: archive_zip.Deflate}
	Roots[`/archive/zip/Store`] = ConstRoot{Const: archive_zip.Store}
	Roots[`/bytes/MinRead`] = ConstRoot{Const: int64(bytes.MinRead)}
	Roots[`/compress/flate/BestCompression`] = ConstRoot{Const: int64(compress_flate.BestCompression)}
	Roots[`/compress/flate/BestSpeed`] = ConstRoot{Const: int64(compress_flate.BestSpeed)}
	Roots[`/compress/flate/DefaultCompression`] = ConstRoot{Const: int64(compress_flate.DefaultCompression)}
	Roots[`/compress/flate/NoCompression`] = ConstRoot{Const: int64(compress_flate.NoCompression)}
	Roots[`/compress/gzip/BestCompression`] = ConstRoot{Const: int64(compress_gzip.BestCompression)}
	Roots[`/compress/gzip/BestSpeed`] = ConstRoot{Const: int64(compress_gzip.BestSpeed)}
	Roots[`/compress/gzip/DefaultCompression`] = ConstRoot{Const: int64(compress_gzip.DefaultCompression)}
	Roots[`/compress/gzip/NoCompression`] = ConstRoot{Const: int64(compress_gzip.NoCompression)}
	Roots[`/compress/lzw/LSB`] = ConstRoot{Const: compress_lzw.LSB}
	Roots[`/compress/lzw/MSB`] = ConstRoot{Const: compress_lzw.MSB}
	Roots[`/compress/zlib/BestCompression`] = ConstRoot{Const: int64(compress_zlib.BestCompression)}
	Roots[`/compress/zlib/BestSpeed`] = ConstRoot{Const: int64(compress_zlib.BestSpeed)}
	Roots[`/compress/zlib/DefaultCompression`] = ConstRoot{Const: int64(compress_zlib.DefaultCompression)}
	Roots[`/compress/zlib/NoCompression`] = ConstRoot{Const: int64(compress_zlib.NoCompression)}
	Roots[`/crypto/MD4`] = ConstRoot{Const: crypto.MD4}
	Roots[`/crypto/MD5`] = ConstRoot{Const: crypto.MD5}
	Roots[`/crypto/MD5SHA1`] = ConstRoot{Const: crypto.MD5SHA1}
	Roots[`/crypto/RIPEMD160`] = ConstRoot{Const: crypto.RIPEMD160}
	Roots[`/crypto/SHA1`] = ConstRoot{Const: crypto.SHA1}
	Roots[`/crypto/SHA224`] = ConstRoot{Const: crypto.SHA224}
	Roots[`/crypto/SHA256`] = ConstRoot{Const: crypto.SHA256}
	Roots[`/crypto/SHA384`] = ConstRoot{Const: crypto.SHA384}
	Roots[`/crypto/SHA512`] = ConstRoot{Const: crypto.SHA512}
	Roots[`/crypto/aes/BlockSize`] = ConstRoot{Const: int64(crypto_aes.BlockSize)}
	Roots[`/crypto/des/BlockSize`] = ConstRoot{Const: int64(crypto_des.BlockSize)}
	Roots[`/crypto/dsa/L1024N160`] = ConstRoot{Const: crypto_dsa.L1024N160}
	Roots[`/crypto/dsa/L2048N224`] = ConstRoot{Const: crypto_dsa.L2048N224}
	Roots[`/crypto/dsa/L2048N256`] = ConstRoot{Const: crypto_dsa.L2048N256}
	Roots[`/crypto/dsa/L3072N256`] = ConstRoot{Const: crypto_dsa.L3072N256}
	Roots[`/crypto/md5/BlockSize`] = ConstRoot{Const: int64(crypto_md5.BlockSize)}
	Roots[`/crypto/md5/Size`] = ConstRoot{Const: int64(crypto_md5.Size)}
	Roots[`/crypto/sha1/BlockSize`] = ConstRoot{Const: int64(crypto_sha1.BlockSize)}
	Roots[`/crypto/sha1/Size`] = ConstRoot{Const: int64(crypto_sha1.Size)}
	Roots[`/crypto/sha256/BlockSize`] = ConstRoot{Const: int64(crypto_sha256.BlockSize)}
	Roots[`/crypto/sha256/Size`] = ConstRoot{Const: int64(crypto_sha256.Size)}
	Roots[`/crypto/sha256/Size224`] = ConstRoot{Const: int64(crypto_sha256.Size224)}
	Roots[`/crypto/sha512/BlockSize`] = ConstRoot{Const: int64(crypto_sha512.BlockSize)}
	Roots[`/crypto/sha512/Size`] = ConstRoot{Const: int64(crypto_sha512.Size)}
	Roots[`/crypto/sha512/Size384`] = ConstRoot{Const: int64(crypto_sha512.Size384)}
	Roots[`/crypto/tls/NoClientCert`] = ConstRoot{Const: crypto_tls.NoClientCert}
	Roots[`/crypto/tls/RequestClientCert`] = ConstRoot{Const: crypto_tls.RequestClientCert}
	Roots[`/crypto/tls/RequireAndVerifyClientCert`] = ConstRoot{Const: crypto_tls.RequireAndVerifyClientCert}
	Roots[`/crypto/tls/RequireAnyClientCert`] = ConstRoot{Const: crypto_tls.RequireAnyClientCert}
	Roots[`/crypto/tls/TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA`] = ConstRoot{Const: crypto_tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA}
	Roots[`/crypto/tls/TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA`] = ConstRoot{Const: crypto_tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA}
	Roots[`/crypto/tls/TLS_ECDHE_RSA_WITH_RC4_128_SHA`] = ConstRoot{Const: crypto_tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA}
	Roots[`/crypto/tls/TLS_RSA_WITH_3DES_EDE_CBC_SHA`] = ConstRoot{Const: crypto_tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA}
	Roots[`/crypto/tls/TLS_RSA_WITH_AES_128_CBC_SHA`] = ConstRoot{Const: crypto_tls.TLS_RSA_WITH_AES_128_CBC_SHA}
	Roots[`/crypto/tls/TLS_RSA_WITH_RC4_128_SHA`] = ConstRoot{Const: crypto_tls.TLS_RSA_WITH_RC4_128_SHA}
	Roots[`/crypto/tls/VerifyClientCertIfGiven`] = ConstRoot{Const: crypto_tls.VerifyClientCertIfGiven}
	Roots[`/crypto/x509/CANotAuthorizedForThisName`] = ConstRoot{Const: crypto_x509.CANotAuthorizedForThisName}
	Roots[`/crypto/x509/DSA`] = ConstRoot{Const: crypto_x509.DSA}
	Roots[`/crypto/x509/DSAWithSHA1`] = ConstRoot{Const: crypto_x509.DSAWithSHA1}
	Roots[`/crypto/x509/DSAWithSHA256`] = ConstRoot{Const: crypto_x509.DSAWithSHA256}
	Roots[`/crypto/x509/Expired`] = ConstRoot{Const: crypto_x509.Expired}
	Roots[`/crypto/x509/ExtKeyUsageAny`] = ConstRoot{Const: crypto_x509.ExtKeyUsageAny}
	Roots[`/crypto/x509/ExtKeyUsageClientAuth`] = ConstRoot{Const: crypto_x509.ExtKeyUsageClientAuth}
	Roots[`/crypto/x509/ExtKeyUsageCodeSigning`] = ConstRoot{Const: crypto_x509.ExtKeyUsageCodeSigning}
	Roots[`/crypto/x509/ExtKeyUsageEmailProtection`] = ConstRoot{Const: crypto_x509.ExtKeyUsageEmailProtection}
	Roots[`/crypto/x509/ExtKeyUsageOCSPSigning`] = ConstRoot{Const: crypto_x509.ExtKeyUsageOCSPSigning}
	Roots[`/crypto/x509/ExtKeyUsageServerAuth`] = ConstRoot{Const: crypto_x509.ExtKeyUsageServerAuth}
	Roots[`/crypto/x509/ExtKeyUsageTimeStamping`] = ConstRoot{Const: crypto_x509.ExtKeyUsageTimeStamping}
	Roots[`/crypto/x509/KeyUsageCRLSign`] = ConstRoot{Const: crypto_x509.KeyUsageCRLSign}
	Roots[`/crypto/x509/KeyUsageCertSign`] = ConstRoot{Const: crypto_x509.KeyUsageCertSign}
	Roots[`/crypto/x509/KeyUsageContentCommitment`] = ConstRoot{Const: crypto_x509.KeyUsageContentCommitment}
	Roots[`/crypto/x509/KeyUsageDataEncipherment`] = ConstRoot{Const: crypto_x509.KeyUsageDataEncipherment}
	Roots[`/crypto/x509/KeyUsageDecipherOnly`] = ConstRoot{Const: crypto_x509.KeyUsageDecipherOnly}
	Roots[`/crypto/x509/KeyUsageDigitalSignature`] = ConstRoot{Const: crypto_x509.KeyUsageDigitalSignature}
	Roots[`/crypto/x509/KeyUsageEncipherOnly`] = ConstRoot{Const: crypto_x509.KeyUsageEncipherOnly}
	Roots[`/crypto/x509/KeyUsageKeyAgreement`] = ConstRoot{Const: crypto_x509.KeyUsageKeyAgreement}
	Roots[`/crypto/x509/KeyUsageKeyEncipherment`] = ConstRoot{Const: crypto_x509.KeyUsageKeyEncipherment}
	Roots[`/crypto/x509/MD2WithRSA`] = ConstRoot{Const: crypto_x509.MD2WithRSA}
	Roots[`/crypto/x509/MD5WithRSA`] = ConstRoot{Const: crypto_x509.MD5WithRSA}
	Roots[`/crypto/x509/NotAuthorizedToSign`] = ConstRoot{Const: crypto_x509.NotAuthorizedToSign}
	Roots[`/crypto/x509/RSA`] = ConstRoot{Const: crypto_x509.RSA}
	Roots[`/crypto/x509/SHA1WithRSA`] = ConstRoot{Const: crypto_x509.SHA1WithRSA}
	Roots[`/crypto/x509/SHA256WithRSA`] = ConstRoot{Const: crypto_x509.SHA256WithRSA}
	Roots[`/crypto/x509/SHA384WithRSA`] = ConstRoot{Const: crypto_x509.SHA384WithRSA}
	Roots[`/crypto/x509/SHA512WithRSA`] = ConstRoot{Const: crypto_x509.SHA512WithRSA}
	Roots[`/crypto/x509/TooManyIntermediates`] = ConstRoot{Const: crypto_x509.TooManyIntermediates}
	Roots[`/crypto/x509/UnknownPublicKeyAlgorithm`] = ConstRoot{Const: crypto_x509.UnknownPublicKeyAlgorithm}
	Roots[`/crypto/x509/UnknownSignatureAlgorithm`] = ConstRoot{Const: crypto_x509.UnknownSignatureAlgorithm}
	Roots[`/debug/dwarf/AttrAbstractOrigin`] = ConstRoot{Const: debug_dwarf.AttrAbstractOrigin}
	Roots[`/debug/dwarf/AttrAccessibility`] = ConstRoot{Const: debug_dwarf.AttrAccessibility}
	Roots[`/debug/dwarf/AttrAddrClass`] = ConstRoot{Const: debug_dwarf.AttrAddrClass}
	Roots[`/debug/dwarf/AttrAllocated`] = ConstRoot{Const: debug_dwarf.AttrAllocated}
	Roots[`/debug/dwarf/AttrArtificial`] = ConstRoot{Const: debug_dwarf.AttrArtificial}
	Roots[`/debug/dwarf/AttrAssociated`] = ConstRoot{Const: debug_dwarf.AttrAssociated}
	Roots[`/debug/dwarf/AttrBaseTypes`] = ConstRoot{Const: debug_dwarf.AttrBaseTypes}
	Roots[`/debug/dwarf/AttrBitOffset`] = ConstRoot{Const: debug_dwarf.AttrBitOffset}
	Roots[`/debug/dwarf/AttrBitSize`] = ConstRoot{Const: debug_dwarf.AttrBitSize}
	Roots[`/debug/dwarf/AttrByteSize`] = ConstRoot{Const: debug_dwarf.AttrByteSize}
	Roots[`/debug/dwarf/AttrCallColumn`] = ConstRoot{Const: debug_dwarf.AttrCallColumn}
	Roots[`/debug/dwarf/AttrCallFile`] = ConstRoot{Const: debug_dwarf.AttrCallFile}
	Roots[`/debug/dwarf/AttrCallLine`] = ConstRoot{Const: debug_dwarf.AttrCallLine}
	Roots[`/debug/dwarf/AttrCalling`] = ConstRoot{Const: debug_dwarf.AttrCalling}
	Roots[`/debug/dwarf/AttrCommonRef`] = ConstRoot{Const: debug_dwarf.AttrCommonRef}
	Roots[`/debug/dwarf/AttrCompDir`] = ConstRoot{Const: debug_dwarf.AttrCompDir}
	Roots[`/debug/dwarf/AttrConstValue`] = ConstRoot{Const: debug_dwarf.AttrConstValue}
	Roots[`/debug/dwarf/AttrContainingType`] = ConstRoot{Const: debug_dwarf.AttrContainingType}
	Roots[`/debug/dwarf/AttrCount`] = ConstRoot{Const: debug_dwarf.AttrCount}
	Roots[`/debug/dwarf/AttrDataLocation`] = ConstRoot{Const: debug_dwarf.AttrDataLocation}
	Roots[`/debug/dwarf/AttrDataMemberLoc`] = ConstRoot{Const: debug_dwarf.AttrDataMemberLoc}
	Roots[`/debug/dwarf/AttrDeclColumn`] = ConstRoot{Const: debug_dwarf.AttrDeclColumn}
	Roots[`/debug/dwarf/AttrDeclFile`] = ConstRoot{Const: debug_dwarf.AttrDeclFile}
	Roots[`/debug/dwarf/AttrDeclLine`] = ConstRoot{Const: debug_dwarf.AttrDeclLine}
	Roots[`/debug/dwarf/AttrDeclaration`] = ConstRoot{Const: debug_dwarf.AttrDeclaration}
	Roots[`/debug/dwarf/AttrDefaultValue`] = ConstRoot{Const: debug_dwarf.AttrDefaultValue}
	Roots[`/debug/dwarf/AttrDescription`] = ConstRoot{Const: debug_dwarf.AttrDescription}
	Roots[`/debug/dwarf/AttrDiscr`] = ConstRoot{Const: debug_dwarf.AttrDiscr}
	Roots[`/debug/dwarf/AttrDiscrList`] = ConstRoot{Const: debug_dwarf.AttrDiscrList}
	Roots[`/debug/dwarf/AttrDiscrValue`] = ConstRoot{Const: debug_dwarf.AttrDiscrValue}
	Roots[`/debug/dwarf/AttrEncoding`] = ConstRoot{Const: debug_dwarf.AttrEncoding}
	Roots[`/debug/dwarf/AttrEntrypc`] = ConstRoot{Const: debug_dwarf.AttrEntrypc}
	Roots[`/debug/dwarf/AttrExtension`] = ConstRoot{Const: debug_dwarf.AttrExtension}
	Roots[`/debug/dwarf/AttrExternal`] = ConstRoot{Const: debug_dwarf.AttrExternal}
	Roots[`/debug/dwarf/AttrFrameBase`] = ConstRoot{Const: debug_dwarf.AttrFrameBase}
	Roots[`/debug/dwarf/AttrFriend`] = ConstRoot{Const: debug_dwarf.AttrFriend}
	Roots[`/debug/dwarf/AttrHighpc`] = ConstRoot{Const: debug_dwarf.AttrHighpc}
	Roots[`/debug/dwarf/AttrIdentifierCase`] = ConstRoot{Const: debug_dwarf.AttrIdentifierCase}
	Roots[`/debug/dwarf/AttrImport`] = ConstRoot{Const: debug_dwarf.AttrImport}
	Roots[`/debug/dwarf/AttrInline`] = ConstRoot{Const: debug_dwarf.AttrInline}
	Roots[`/debug/dwarf/AttrIsOptional`] = ConstRoot{Const: debug_dwarf.AttrIsOptional}
	Roots[`/debug/dwarf/AttrLanguage`] = ConstRoot{Const: debug_dwarf.AttrLanguage}
	Roots[`/debug/dwarf/AttrLocation`] = ConstRoot{Const: debug_dwarf.AttrLocation}
	Roots[`/debug/dwarf/AttrLowerBound`] = ConstRoot{Const: debug_dwarf.AttrLowerBound}
	Roots[`/debug/dwarf/AttrLowpc`] = ConstRoot{Const: debug_dwarf.AttrLowpc}
	Roots[`/debug/dwarf/AttrMacroInfo`] = ConstRoot{Const: debug_dwarf.AttrMacroInfo}
	Roots[`/debug/dwarf/AttrName`] = ConstRoot{Const: debug_dwarf.AttrName}
	Roots[`/debug/dwarf/AttrNamelistItem`] = ConstRoot{Const: debug_dwarf.AttrNamelistItem}
	Roots[`/debug/dwarf/AttrOrdering`] = ConstRoot{Const: debug_dwarf.AttrOrdering}
	Roots[`/debug/dwarf/AttrPriority`] = ConstRoot{Const: debug_dwarf.AttrPriority}
	Roots[`/debug/dwarf/AttrProducer`] = ConstRoot{Const: debug_dwarf.AttrProducer}
	Roots[`/debug/dwarf/AttrPrototyped`] = ConstRoot{Const: debug_dwarf.AttrPrototyped}
	Roots[`/debug/dwarf/AttrRanges`] = ConstRoot{Const: debug_dwarf.AttrRanges}
	Roots[`/debug/dwarf/AttrReturnAddr`] = ConstRoot{Const: debug_dwarf.AttrReturnAddr}
	Roots[`/debug/dwarf/AttrSegment`] = ConstRoot{Const: debug_dwarf.AttrSegment}
	Roots[`/debug/dwarf/AttrSibling`] = ConstRoot{Const: debug_dwarf.AttrSibling}
	Roots[`/debug/dwarf/AttrSpecification`] = ConstRoot{Const: debug_dwarf.AttrSpecification}
	Roots[`/debug/dwarf/AttrStartScope`] = ConstRoot{Const: debug_dwarf.AttrStartScope}
	Roots[`/debug/dwarf/AttrStaticLink`] = ConstRoot{Const: debug_dwarf.AttrStaticLink}
	Roots[`/debug/dwarf/AttrStmtList`] = ConstRoot{Const: debug_dwarf.AttrStmtList}
	Roots[`/debug/dwarf/AttrStride`] = ConstRoot{Const: debug_dwarf.AttrStride}
	Roots[`/debug/dwarf/AttrStrideSize`] = ConstRoot{Const: debug_dwarf.AttrStrideSize}
	Roots[`/debug/dwarf/AttrStringLength`] = ConstRoot{Const: debug_dwarf.AttrStringLength}
	Roots[`/debug/dwarf/AttrTrampoline`] = ConstRoot{Const: debug_dwarf.AttrTrampoline}
	Roots[`/debug/dwarf/AttrType`] = ConstRoot{Const: debug_dwarf.AttrType}
	Roots[`/debug/dwarf/AttrUpperBound`] = ConstRoot{Const: debug_dwarf.AttrUpperBound}
	Roots[`/debug/dwarf/AttrUseLocation`] = ConstRoot{Const: debug_dwarf.AttrUseLocation}
	Roots[`/debug/dwarf/AttrUseUTF8`] = ConstRoot{Const: debug_dwarf.AttrUseUTF8}
	Roots[`/debug/dwarf/AttrVarParam`] = ConstRoot{Const: debug_dwarf.AttrVarParam}
	Roots[`/debug/dwarf/AttrVirtuality`] = ConstRoot{Const: debug_dwarf.AttrVirtuality}
	Roots[`/debug/dwarf/AttrVisibility`] = ConstRoot{Const: debug_dwarf.AttrVisibility}
	Roots[`/debug/dwarf/AttrVtableElemLoc`] = ConstRoot{Const: debug_dwarf.AttrVtableElemLoc}
	Roots[`/debug/dwarf/TagAccessDeclaration`] = ConstRoot{Const: debug_dwarf.TagAccessDeclaration}
	Roots[`/debug/dwarf/TagArrayType`] = ConstRoot{Const: debug_dwarf.TagArrayType}
	Roots[`/debug/dwarf/TagBaseType`] = ConstRoot{Const: debug_dwarf.TagBaseType}
	Roots[`/debug/dwarf/TagCatchDwarfBlock`] = ConstRoot{Const: debug_dwarf.TagCatchDwarfBlock}
	Roots[`/debug/dwarf/TagClassType`] = ConstRoot{Const: debug_dwarf.TagClassType}
	Roots[`/debug/dwarf/TagCommonDwarfBlock`] = ConstRoot{Const: debug_dwarf.TagCommonDwarfBlock}
	Roots[`/debug/dwarf/TagCommonInclusion`] = ConstRoot{Const: debug_dwarf.TagCommonInclusion}
	Roots[`/debug/dwarf/TagCompileUnit`] = ConstRoot{Const: debug_dwarf.TagCompileUnit}
	Roots[`/debug/dwarf/TagConstType`] = ConstRoot{Const: debug_dwarf.TagConstType}
	Roots[`/debug/dwarf/TagConstant`] = ConstRoot{Const: debug_dwarf.TagConstant}
	Roots[`/debug/dwarf/TagDwarfProcedure`] = ConstRoot{Const: debug_dwarf.TagDwarfProcedure}
	Roots[`/debug/dwarf/TagEntryPoint`] = ConstRoot{Const: debug_dwarf.TagEntryPoint}
	Roots[`/debug/dwarf/TagEnumerationType`] = ConstRoot{Const: debug_dwarf.TagEnumerationType}
	Roots[`/debug/dwarf/TagEnumerator`] = ConstRoot{Const: debug_dwarf.TagEnumerator}
	Roots[`/debug/dwarf/TagFileType`] = ConstRoot{Const: debug_dwarf.TagFileType}
	Roots[`/debug/dwarf/TagFormalParameter`] = ConstRoot{Const: debug_dwarf.TagFormalParameter}
	Roots[`/debug/dwarf/TagFriend`] = ConstRoot{Const: debug_dwarf.TagFriend}
	Roots[`/debug/dwarf/TagImportedDeclaration`] = ConstRoot{Const: debug_dwarf.TagImportedDeclaration}
	Roots[`/debug/dwarf/TagImportedModule`] = ConstRoot{Const: debug_dwarf.TagImportedModule}
	Roots[`/debug/dwarf/TagImportedUnit`] = ConstRoot{Const: debug_dwarf.TagImportedUnit}
	Roots[`/debug/dwarf/TagInheritance`] = ConstRoot{Const: debug_dwarf.TagInheritance}
	Roots[`/debug/dwarf/TagInlinedSubroutine`] = ConstRoot{Const: debug_dwarf.TagInlinedSubroutine}
	Roots[`/debug/dwarf/TagInterfaceType`] = ConstRoot{Const: debug_dwarf.TagInterfaceType}
	Roots[`/debug/dwarf/TagLabel`] = ConstRoot{Const: debug_dwarf.TagLabel}
	Roots[`/debug/dwarf/TagLexDwarfBlock`] = ConstRoot{Const: debug_dwarf.TagLexDwarfBlock}
	Roots[`/debug/dwarf/TagMember`] = ConstRoot{Const: debug_dwarf.TagMember}
	Roots[`/debug/dwarf/TagModule`] = ConstRoot{Const: debug_dwarf.TagModule}
	Roots[`/debug/dwarf/TagMutableType`] = ConstRoot{Const: debug_dwarf.TagMutableType}
	Roots[`/debug/dwarf/TagNamelist`] = ConstRoot{Const: debug_dwarf.TagNamelist}
	Roots[`/debug/dwarf/TagNamelistItem`] = ConstRoot{Const: debug_dwarf.TagNamelistItem}
	Roots[`/debug/dwarf/TagNamespace`] = ConstRoot{Const: debug_dwarf.TagNamespace}
	Roots[`/debug/dwarf/TagPackedType`] = ConstRoot{Const: debug_dwarf.TagPackedType}
	Roots[`/debug/dwarf/TagPartialUnit`] = ConstRoot{Const: debug_dwarf.TagPartialUnit}
	Roots[`/debug/dwarf/TagPointerType`] = ConstRoot{Const: debug_dwarf.TagPointerType}
	Roots[`/debug/dwarf/TagPtrToMemberType`] = ConstRoot{Const: debug_dwarf.TagPtrToMemberType}
	Roots[`/debug/dwarf/TagReferenceType`] = ConstRoot{Const: debug_dwarf.TagReferenceType}
	Roots[`/debug/dwarf/TagRestrictType`] = ConstRoot{Const: debug_dwarf.TagRestrictType}
	Roots[`/debug/dwarf/TagSetType`] = ConstRoot{Const: debug_dwarf.TagSetType}
	Roots[`/debug/dwarf/TagStringType`] = ConstRoot{Const: debug_dwarf.TagStringType}
	Roots[`/debug/dwarf/TagStructType`] = ConstRoot{Const: debug_dwarf.TagStructType}
	Roots[`/debug/dwarf/TagSubprogram`] = ConstRoot{Const: debug_dwarf.TagSubprogram}
	Roots[`/debug/dwarf/TagSubrangeType`] = ConstRoot{Const: debug_dwarf.TagSubrangeType}
	Roots[`/debug/dwarf/TagSubroutineType`] = ConstRoot{Const: debug_dwarf.TagSubroutineType}
	Roots[`/debug/dwarf/TagTemplateTypeParameter`] = ConstRoot{Const: debug_dwarf.TagTemplateTypeParameter}
	Roots[`/debug/dwarf/TagTemplateValueParameter`] = ConstRoot{Const: debug_dwarf.TagTemplateValueParameter}
	Roots[`/debug/dwarf/TagThrownType`] = ConstRoot{Const: debug_dwarf.TagThrownType}
	Roots[`/debug/dwarf/TagTryDwarfBlock`] = ConstRoot{Const: debug_dwarf.TagTryDwarfBlock}
	Roots[`/debug/dwarf/TagTypedef`] = ConstRoot{Const: debug_dwarf.TagTypedef}
	Roots[`/debug/dwarf/TagUnionType`] = ConstRoot{Const: debug_dwarf.TagUnionType}
	Roots[`/debug/dwarf/TagUnspecifiedParameters`] = ConstRoot{Const: debug_dwarf.TagUnspecifiedParameters}
	Roots[`/debug/dwarf/TagUnspecifiedType`] = ConstRoot{Const: debug_dwarf.TagUnspecifiedType}
	Roots[`/debug/dwarf/TagVariable`] = ConstRoot{Const: debug_dwarf.TagVariable}
	Roots[`/debug/dwarf/TagVariant`] = ConstRoot{Const: debug_dwarf.TagVariant}
	Roots[`/debug/dwarf/TagVariantPart`] = ConstRoot{Const: debug_dwarf.TagVariantPart}
	Roots[`/debug/dwarf/TagVolatileType`] = ConstRoot{Const: debug_dwarf.TagVolatileType}
	Roots[`/debug/dwarf/TagWithStmt`] = ConstRoot{Const: debug_dwarf.TagWithStmt}
	Roots[`/debug/elf/ARM_MAGIC_TRAMP_NUMBER`] = ConstRoot{Const: int64(debug_elf.ARM_MAGIC_TRAMP_NUMBER)}
	Roots[`/debug/elf/DF_BIND_NOW`] = ConstRoot{Const: debug_elf.DF_BIND_NOW}
	Roots[`/debug/elf/DF_ORIGIN`] = ConstRoot{Const: debug_elf.DF_ORIGIN}
	Roots[`/debug/elf/DF_STATIC_TLS`] = ConstRoot{Const: debug_elf.DF_STATIC_TLS}
	Roots[`/debug/elf/DF_SYMBOLIC`] = ConstRoot{Const: debug_elf.DF_SYMBOLIC}
	Roots[`/debug/elf/DF_TEXTREL`] = ConstRoot{Const: debug_elf.DF_TEXTREL}
	Roots[`/debug/elf/DT_BIND_NOW`] = ConstRoot{Const: debug_elf.DT_BIND_NOW}
	Roots[`/debug/elf/DT_DEBUG`] = ConstRoot{Const: debug_elf.DT_DEBUG}
	Roots[`/debug/elf/DT_ENCODING`] = ConstRoot{Const: debug_elf.DT_ENCODING}
	Roots[`/debug/elf/DT_FINI`] = ConstRoot{Const: debug_elf.DT_FINI}
	Roots[`/debug/elf/DT_FINI_ARRAY`] = ConstRoot{Const: debug_elf.DT_FINI_ARRAY}
	Roots[`/debug/elf/DT_FINI_ARRAYSZ`] = ConstRoot{Const: debug_elf.DT_FINI_ARRAYSZ}
	Roots[`/debug/elf/DT_FLAGS`] = ConstRoot{Const: debug_elf.DT_FLAGS}
	Roots[`/debug/elf/DT_HASH`] = ConstRoot{Const: debug_elf.DT_HASH}
	Roots[`/debug/elf/DT_HIOS`] = ConstRoot{Const: debug_elf.DT_HIOS}
	Roots[`/debug/elf/DT_HIPROC`] = ConstRoot{Const: debug_elf.DT_HIPROC}
	Roots[`/debug/elf/DT_INIT`] = ConstRoot{Const: debug_elf.DT_INIT}
	Roots[`/debug/elf/DT_INIT_ARRAY`] = ConstRoot{Const: debug_elf.DT_INIT_ARRAY}
	Roots[`/debug/elf/DT_INIT_ARRAYSZ`] = ConstRoot{Const: debug_elf.DT_INIT_ARRAYSZ}
	Roots[`/debug/elf/DT_JMPREL`] = ConstRoot{Const: debug_elf.DT_JMPREL}
	Roots[`/debug/elf/DT_LOOS`] = ConstRoot{Const: debug_elf.DT_LOOS}
	Roots[`/debug/elf/DT_LOPROC`] = ConstRoot{Const: debug_elf.DT_LOPROC}
	Roots[`/debug/elf/DT_NEEDED`] = ConstRoot{Const: debug_elf.DT_NEEDED}
	Roots[`/debug/elf/DT_NULL`] = ConstRoot{Const: debug_elf.DT_NULL}
	Roots[`/debug/elf/DT_PLTGOT`] = ConstRoot{Const: debug_elf.DT_PLTGOT}
	Roots[`/debug/elf/DT_PLTREL`] = ConstRoot{Const: debug_elf.DT_PLTREL}
	Roots[`/debug/elf/DT_PLTRELSZ`] = ConstRoot{Const: debug_elf.DT_PLTRELSZ}
	Roots[`/debug/elf/DT_PREINIT_ARRAY`] = ConstRoot{Const: debug_elf.DT_PREINIT_ARRAY}
	Roots[`/debug/elf/DT_PREINIT_ARRAYSZ`] = ConstRoot{Const: debug_elf.DT_PREINIT_ARRAYSZ}
	Roots[`/debug/elf/DT_REL`] = ConstRoot{Const: debug_elf.DT_REL}
	Roots[`/debug/elf/DT_RELA`] = ConstRoot{Const: debug_elf.DT_RELA}
	Roots[`/debug/elf/DT_RELAENT`] = ConstRoot{Const: debug_elf.DT_RELAENT}
	Roots[`/debug/elf/DT_RELASZ`] = ConstRoot{Const: debug_elf.DT_RELASZ}
	Roots[`/debug/elf/DT_RELENT`] = ConstRoot{Const: debug_elf.DT_RELENT}
	Roots[`/debug/elf/DT_RELSZ`] = ConstRoot{Const: debug_elf.DT_RELSZ}
	Roots[`/debug/elf/DT_RPATH`] = ConstRoot{Const: debug_elf.DT_RPATH}
	Roots[`/debug/elf/DT_RUNPATH`] = ConstRoot{Const: debug_elf.DT_RUNPATH}
	Roots[`/debug/elf/DT_SONAME`] = ConstRoot{Const: debug_elf.DT_SONAME}
	Roots[`/debug/elf/DT_STRSZ`] = ConstRoot{Const: debug_elf.DT_STRSZ}
	Roots[`/debug/elf/DT_STRTAB`] = ConstRoot{Const: debug_elf.DT_STRTAB}
	Roots[`/debug/elf/DT_SYMBOLIC`] = ConstRoot{Const: debug_elf.DT_SYMBOLIC}
	Roots[`/debug/elf/DT_SYMENT`] = ConstRoot{Const: debug_elf.DT_SYMENT}
	Roots[`/debug/elf/DT_SYMTAB`] = ConstRoot{Const: debug_elf.DT_SYMTAB}
	Roots[`/debug/elf/DT_TEXTREL`] = ConstRoot{Const: debug_elf.DT_TEXTREL}
	Roots[`/debug/elf/DT_VERNEED`] = ConstRoot{Const: debug_elf.DT_VERNEED}
	Roots[`/debug/elf/DT_VERNEEDNUM`] = ConstRoot{Const: debug_elf.DT_VERNEEDNUM}
	Roots[`/debug/elf/DT_VERSYM`] = ConstRoot{Const: debug_elf.DT_VERSYM}
	Roots[`/debug/elf/EI_ABIVERSION`] = ConstRoot{Const: int64(debug_elf.EI_ABIVERSION)}
	Roots[`/debug/elf/EI_CLASS`] = ConstRoot{Const: int64(debug_elf.EI_CLASS)}
	Roots[`/debug/elf/EI_DATA`] = ConstRoot{Const: int64(debug_elf.EI_DATA)}
	Roots[`/debug/elf/EI_NIDENT`] = ConstRoot{Const: int64(debug_elf.EI_NIDENT)}
	Roots[`/debug/elf/EI_OSABI`] = ConstRoot{Const: int64(debug_elf.EI_OSABI)}
	Roots[`/debug/elf/EI_PAD`] = ConstRoot{Const: int64(debug_elf.EI_PAD)}
	Roots[`/debug/elf/EI_VERSION`] = ConstRoot{Const: int64(debug_elf.EI_VERSION)}
	Roots[`/debug/elf/ELFCLASS32`] = ConstRoot{Const: debug_elf.ELFCLASS32}
	Roots[`/debug/elf/ELFCLASS64`] = ConstRoot{Const: debug_elf.ELFCLASS64}
	Roots[`/debug/elf/ELFCLASSNONE`] = ConstRoot{Const: debug_elf.ELFCLASSNONE}
	Roots[`/debug/elf/ELFDATA2LSB`] = ConstRoot{Const: debug_elf.ELFDATA2LSB}
	Roots[`/debug/elf/ELFDATA2MSB`] = ConstRoot{Const: debug_elf.ELFDATA2MSB}
	Roots[`/debug/elf/ELFDATANONE`] = ConstRoot{Const: debug_elf.ELFDATANONE}
	Roots[`/debug/elf/ELFMAG`] = ConstRoot{Const: debug_elf.ELFMAG}
	Roots[`/debug/elf/ELFOSABI_86OPEN`] = ConstRoot{Const: debug_elf.ELFOSABI_86OPEN}
	Roots[`/debug/elf/ELFOSABI_AIX`] = ConstRoot{Const: debug_elf.ELFOSABI_AIX}
	Roots[`/debug/elf/ELFOSABI_ARM`] = ConstRoot{Const: debug_elf.ELFOSABI_ARM}
	Roots[`/debug/elf/ELFOSABI_FREEBSD`] = ConstRoot{Const: debug_elf.ELFOSABI_FREEBSD}
	Roots[`/debug/elf/ELFOSABI_HPUX`] = ConstRoot{Const: debug_elf.ELFOSABI_HPUX}
	Roots[`/debug/elf/ELFOSABI_HURD`] = ConstRoot{Const: debug_elf.ELFOSABI_HURD}
	Roots[`/debug/elf/ELFOSABI_IRIX`] = ConstRoot{Const: debug_elf.ELFOSABI_IRIX}
	Roots[`/debug/elf/ELFOSABI_LINUX`] = ConstRoot{Const: debug_elf.ELFOSABI_LINUX}
	Roots[`/debug/elf/ELFOSABI_MODESTO`] = ConstRoot{Const: debug_elf.ELFOSABI_MODESTO}
	Roots[`/debug/elf/ELFOSABI_NETBSD`] = ConstRoot{Const: debug_elf.ELFOSABI_NETBSD}
	Roots[`/debug/elf/ELFOSABI_NONE`] = ConstRoot{Const: debug_elf.ELFOSABI_NONE}
	Roots[`/debug/elf/ELFOSABI_NSK`] = ConstRoot{Const: debug_elf.ELFOSABI_NSK}
	Roots[`/debug/elf/ELFOSABI_OPENBSD`] = ConstRoot{Const: debug_elf.ELFOSABI_OPENBSD}
	Roots[`/debug/elf/ELFOSABI_OPENVMS`] = ConstRoot{Const: debug_elf.ELFOSABI_OPENVMS}
	Roots[`/debug/elf/ELFOSABI_SOLARIS`] = ConstRoot{Const: debug_elf.ELFOSABI_SOLARIS}
	Roots[`/debug/elf/ELFOSABI_STANDALONE`] = ConstRoot{Const: debug_elf.ELFOSABI_STANDALONE}
	Roots[`/debug/elf/ELFOSABI_TRU64`] = ConstRoot{Const: debug_elf.ELFOSABI_TRU64}
	Roots[`/debug/elf/EM_386`] = ConstRoot{Const: debug_elf.EM_386}
	Roots[`/debug/elf/EM_486`] = ConstRoot{Const: debug_elf.EM_486}
	Roots[`/debug/elf/EM_68HC12`] = ConstRoot{Const: debug_elf.EM_68HC12}
	Roots[`/debug/elf/EM_68K`] = ConstRoot{Const: debug_elf.EM_68K}
	Roots[`/debug/elf/EM_860`] = ConstRoot{Const: debug_elf.EM_860}
	Roots[`/debug/elf/EM_88K`] = ConstRoot{Const: debug_elf.EM_88K}
	Roots[`/debug/elf/EM_960`] = ConstRoot{Const: debug_elf.EM_960}
	Roots[`/debug/elf/EM_ALPHA`] = ConstRoot{Const: debug_elf.EM_ALPHA}
	Roots[`/debug/elf/EM_ALPHA_STD`] = ConstRoot{Const: debug_elf.EM_ALPHA_STD}
	Roots[`/debug/elf/EM_ARC`] = ConstRoot{Const: debug_elf.EM_ARC}
	Roots[`/debug/elf/EM_ARM`] = ConstRoot{Const: debug_elf.EM_ARM}
	Roots[`/debug/elf/EM_COLDFIRE`] = ConstRoot{Const: debug_elf.EM_COLDFIRE}
	Roots[`/debug/elf/EM_FR20`] = ConstRoot{Const: debug_elf.EM_FR20}
	Roots[`/debug/elf/EM_H8S`] = ConstRoot{Const: debug_elf.EM_H8S}
	Roots[`/debug/elf/EM_H8_300`] = ConstRoot{Const: debug_elf.EM_H8_300}
	Roots[`/debug/elf/EM_H8_300H`] = ConstRoot{Const: debug_elf.EM_H8_300H}
	Roots[`/debug/elf/EM_H8_500`] = ConstRoot{Const: debug_elf.EM_H8_500}
	Roots[`/debug/elf/EM_IA_64`] = ConstRoot{Const: debug_elf.EM_IA_64}
	Roots[`/debug/elf/EM_M32`] = ConstRoot{Const: debug_elf.EM_M32}
	Roots[`/debug/elf/EM_ME16`] = ConstRoot{Const: debug_elf.EM_ME16}
	Roots[`/debug/elf/EM_MIPS`] = ConstRoot{Const: debug_elf.EM_MIPS}
	Roots[`/debug/elf/EM_MIPS_RS3_LE`] = ConstRoot{Const: debug_elf.EM_MIPS_RS3_LE}
	Roots[`/debug/elf/EM_MIPS_RS4_BE`] = ConstRoot{Const: debug_elf.EM_MIPS_RS4_BE}
	Roots[`/debug/elf/EM_MIPS_X`] = ConstRoot{Const: debug_elf.EM_MIPS_X}
	Roots[`/debug/elf/EM_MMA`] = ConstRoot{Const: debug_elf.EM_MMA}
	Roots[`/debug/elf/EM_NCPU`] = ConstRoot{Const: debug_elf.EM_NCPU}
	Roots[`/debug/elf/EM_NDR1`] = ConstRoot{Const: debug_elf.EM_NDR1}
	Roots[`/debug/elf/EM_NONE`] = ConstRoot{Const: debug_elf.EM_NONE}
	Roots[`/debug/elf/EM_PARISC`] = ConstRoot{Const: debug_elf.EM_PARISC}
	Roots[`/debug/elf/EM_PCP`] = ConstRoot{Const: debug_elf.EM_PCP}
	Roots[`/debug/elf/EM_PPC`] = ConstRoot{Const: debug_elf.EM_PPC}
	Roots[`/debug/elf/EM_PPC64`] = ConstRoot{Const: debug_elf.EM_PPC64}
	Roots[`/debug/elf/EM_RCE`] = ConstRoot{Const: debug_elf.EM_RCE}
	Roots[`/debug/elf/EM_RH32`] = ConstRoot{Const: debug_elf.EM_RH32}
	Roots[`/debug/elf/EM_S370`] = ConstRoot{Const: debug_elf.EM_S370}
	Roots[`/debug/elf/EM_S390`] = ConstRoot{Const: debug_elf.EM_S390}
	Roots[`/debug/elf/EM_SH`] = ConstRoot{Const: debug_elf.EM_SH}
	Roots[`/debug/elf/EM_SPARC`] = ConstRoot{Const: debug_elf.EM_SPARC}
	Roots[`/debug/elf/EM_SPARC32PLUS`] = ConstRoot{Const: debug_elf.EM_SPARC32PLUS}
	Roots[`/debug/elf/EM_SPARCV9`] = ConstRoot{Const: debug_elf.EM_SPARCV9}
	Roots[`/debug/elf/EM_ST100`] = ConstRoot{Const: debug_elf.EM_ST100}
	Roots[`/debug/elf/EM_STARCORE`] = ConstRoot{Const: debug_elf.EM_STARCORE}
	Roots[`/debug/elf/EM_TINYJ`] = ConstRoot{Const: debug_elf.EM_TINYJ}
	Roots[`/debug/elf/EM_TRICORE`] = ConstRoot{Const: debug_elf.EM_TRICORE}
	Roots[`/debug/elf/EM_V800`] = ConstRoot{Const: debug_elf.EM_V800}
	Roots[`/debug/elf/EM_VPP500`] = ConstRoot{Const: debug_elf.EM_VPP500}
	Roots[`/debug/elf/EM_X86_64`] = ConstRoot{Const: debug_elf.EM_X86_64}
	Roots[`/debug/elf/ET_CORE`] = ConstRoot{Const: debug_elf.ET_CORE}
	Roots[`/debug/elf/ET_DYN`] = ConstRoot{Const: debug_elf.ET_DYN}
	Roots[`/debug/elf/ET_EXEC`] = ConstRoot{Const: debug_elf.ET_EXEC}
	Roots[`/debug/elf/ET_HIOS`] = ConstRoot{Const: debug_elf.ET_HIOS}
	Roots[`/debug/elf/ET_HIPROC`] = ConstRoot{Const: debug_elf.ET_HIPROC}
	Roots[`/debug/elf/ET_LOOS`] = ConstRoot{Const: debug_elf.ET_LOOS}
	Roots[`/debug/elf/ET_LOPROC`] = ConstRoot{Const: debug_elf.ET_LOPROC}
	Roots[`/debug/elf/ET_NONE`] = ConstRoot{Const: debug_elf.ET_NONE}
	Roots[`/debug/elf/ET_REL`] = ConstRoot{Const: debug_elf.ET_REL}
	Roots[`/debug/elf/EV_CURRENT`] = ConstRoot{Const: debug_elf.EV_CURRENT}
	Roots[`/debug/elf/EV_NONE`] = ConstRoot{Const: debug_elf.EV_NONE}
	Roots[`/debug/elf/NT_FPREGSET`] = ConstRoot{Const: debug_elf.NT_FPREGSET}
	Roots[`/debug/elf/NT_PRPSINFO`] = ConstRoot{Const: debug_elf.NT_PRPSINFO}
	Roots[`/debug/elf/NT_PRSTATUS`] = ConstRoot{Const: debug_elf.NT_PRSTATUS}
	Roots[`/debug/elf/PF_MASKOS`] = ConstRoot{Const: debug_elf.PF_MASKOS}
	Roots[`/debug/elf/PF_MASKPROC`] = ConstRoot{Const: debug_elf.PF_MASKPROC}
	Roots[`/debug/elf/PF_R`] = ConstRoot{Const: debug_elf.PF_R}
	Roots[`/debug/elf/PF_W`] = ConstRoot{Const: debug_elf.PF_W}
	Roots[`/debug/elf/PF_X`] = ConstRoot{Const: debug_elf.PF_X}
	Roots[`/debug/elf/PT_DYNAMIC`] = ConstRoot{Const: debug_elf.PT_DYNAMIC}
	Roots[`/debug/elf/PT_HIOS`] = ConstRoot{Const: debug_elf.PT_HIOS}
	Roots[`/debug/elf/PT_HIPROC`] = ConstRoot{Const: debug_elf.PT_HIPROC}
	Roots[`/debug/elf/PT_INTERP`] = ConstRoot{Const: debug_elf.PT_INTERP}
	Roots[`/debug/elf/PT_LOAD`] = ConstRoot{Const: debug_elf.PT_LOAD}
	Roots[`/debug/elf/PT_LOOS`] = ConstRoot{Const: debug_elf.PT_LOOS}
	Roots[`/debug/elf/PT_LOPROC`] = ConstRoot{Const: debug_elf.PT_LOPROC}
	Roots[`/debug/elf/PT_NOTE`] = ConstRoot{Const: debug_elf.PT_NOTE}
	Roots[`/debug/elf/PT_NULL`] = ConstRoot{Const: debug_elf.PT_NULL}
	Roots[`/debug/elf/PT_PHDR`] = ConstRoot{Const: debug_elf.PT_PHDR}
	Roots[`/debug/elf/PT_SHLIB`] = ConstRoot{Const: debug_elf.PT_SHLIB}
	Roots[`/debug/elf/PT_TLS`] = ConstRoot{Const: debug_elf.PT_TLS}
	Roots[`/debug/elf/R_386_32`] = ConstRoot{Const: debug_elf.R_386_32}
	Roots[`/debug/elf/R_386_COPY`] = ConstRoot{Const: debug_elf.R_386_COPY}
	Roots[`/debug/elf/R_386_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_386_GLOB_DAT}
	Roots[`/debug/elf/R_386_GOT32`] = ConstRoot{Const: debug_elf.R_386_GOT32}
	Roots[`/debug/elf/R_386_GOTOFF`] = ConstRoot{Const: debug_elf.R_386_GOTOFF}
	Roots[`/debug/elf/R_386_GOTPC`] = ConstRoot{Const: debug_elf.R_386_GOTPC}
	Roots[`/debug/elf/R_386_JMP_SLOT`] = ConstRoot{Const: debug_elf.R_386_JMP_SLOT}
	Roots[`/debug/elf/R_386_NONE`] = ConstRoot{Const: debug_elf.R_386_NONE}
	Roots[`/debug/elf/R_386_PC32`] = ConstRoot{Const: debug_elf.R_386_PC32}
	Roots[`/debug/elf/R_386_PLT32`] = ConstRoot{Const: debug_elf.R_386_PLT32}
	Roots[`/debug/elf/R_386_RELATIVE`] = ConstRoot{Const: debug_elf.R_386_RELATIVE}
	Roots[`/debug/elf/R_386_TLS_DTPMOD32`] = ConstRoot{Const: debug_elf.R_386_TLS_DTPMOD32}
	Roots[`/debug/elf/R_386_TLS_DTPOFF32`] = ConstRoot{Const: debug_elf.R_386_TLS_DTPOFF32}
	Roots[`/debug/elf/R_386_TLS_GD`] = ConstRoot{Const: debug_elf.R_386_TLS_GD}
	Roots[`/debug/elf/R_386_TLS_GD_32`] = ConstRoot{Const: debug_elf.R_386_TLS_GD_32}
	Roots[`/debug/elf/R_386_TLS_GD_CALL`] = ConstRoot{Const: debug_elf.R_386_TLS_GD_CALL}
	Roots[`/debug/elf/R_386_TLS_GD_POP`] = ConstRoot{Const: debug_elf.R_386_TLS_GD_POP}
	Roots[`/debug/elf/R_386_TLS_GD_PUSH`] = ConstRoot{Const: debug_elf.R_386_TLS_GD_PUSH}
	Roots[`/debug/elf/R_386_TLS_GOTIE`] = ConstRoot{Const: debug_elf.R_386_TLS_GOTIE}
	Roots[`/debug/elf/R_386_TLS_IE`] = ConstRoot{Const: debug_elf.R_386_TLS_IE}
	Roots[`/debug/elf/R_386_TLS_IE_32`] = ConstRoot{Const: debug_elf.R_386_TLS_IE_32}
	Roots[`/debug/elf/R_386_TLS_LDM`] = ConstRoot{Const: debug_elf.R_386_TLS_LDM}
	Roots[`/debug/elf/R_386_TLS_LDM_32`] = ConstRoot{Const: debug_elf.R_386_TLS_LDM_32}
	Roots[`/debug/elf/R_386_TLS_LDM_CALL`] = ConstRoot{Const: debug_elf.R_386_TLS_LDM_CALL}
	Roots[`/debug/elf/R_386_TLS_LDM_POP`] = ConstRoot{Const: debug_elf.R_386_TLS_LDM_POP}
	Roots[`/debug/elf/R_386_TLS_LDM_PUSH`] = ConstRoot{Const: debug_elf.R_386_TLS_LDM_PUSH}
	Roots[`/debug/elf/R_386_TLS_LDO_32`] = ConstRoot{Const: debug_elf.R_386_TLS_LDO_32}
	Roots[`/debug/elf/R_386_TLS_LE`] = ConstRoot{Const: debug_elf.R_386_TLS_LE}
	Roots[`/debug/elf/R_386_TLS_LE_32`] = ConstRoot{Const: debug_elf.R_386_TLS_LE_32}
	Roots[`/debug/elf/R_386_TLS_TPOFF`] = ConstRoot{Const: debug_elf.R_386_TLS_TPOFF}
	Roots[`/debug/elf/R_386_TLS_TPOFF32`] = ConstRoot{Const: debug_elf.R_386_TLS_TPOFF32}
	Roots[`/debug/elf/R_ALPHA_BRADDR`] = ConstRoot{Const: debug_elf.R_ALPHA_BRADDR}
	Roots[`/debug/elf/R_ALPHA_COPY`] = ConstRoot{Const: debug_elf.R_ALPHA_COPY}
	Roots[`/debug/elf/R_ALPHA_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_ALPHA_GLOB_DAT}
	Roots[`/debug/elf/R_ALPHA_GPDISP`] = ConstRoot{Const: debug_elf.R_ALPHA_GPDISP}
	Roots[`/debug/elf/R_ALPHA_GPREL32`] = ConstRoot{Const: debug_elf.R_ALPHA_GPREL32}
	Roots[`/debug/elf/R_ALPHA_GPRELHIGH`] = ConstRoot{Const: debug_elf.R_ALPHA_GPRELHIGH}
	Roots[`/debug/elf/R_ALPHA_GPRELLOW`] = ConstRoot{Const: debug_elf.R_ALPHA_GPRELLOW}
	Roots[`/debug/elf/R_ALPHA_GPVALUE`] = ConstRoot{Const: debug_elf.R_ALPHA_GPVALUE}
	Roots[`/debug/elf/R_ALPHA_HINT`] = ConstRoot{Const: debug_elf.R_ALPHA_HINT}
	Roots[`/debug/elf/R_ALPHA_IMMED_BR_HI32`] = ConstRoot{Const: debug_elf.R_ALPHA_IMMED_BR_HI32}
	Roots[`/debug/elf/R_ALPHA_IMMED_GP_16`] = ConstRoot{Const: debug_elf.R_ALPHA_IMMED_GP_16}
	Roots[`/debug/elf/R_ALPHA_IMMED_GP_HI32`] = ConstRoot{Const: debug_elf.R_ALPHA_IMMED_GP_HI32}
	Roots[`/debug/elf/R_ALPHA_IMMED_LO32`] = ConstRoot{Const: debug_elf.R_ALPHA_IMMED_LO32}
	Roots[`/debug/elf/R_ALPHA_IMMED_SCN_HI32`] = ConstRoot{Const: debug_elf.R_ALPHA_IMMED_SCN_HI32}
	Roots[`/debug/elf/R_ALPHA_JMP_SLOT`] = ConstRoot{Const: debug_elf.R_ALPHA_JMP_SLOT}
	Roots[`/debug/elf/R_ALPHA_LITERAL`] = ConstRoot{Const: debug_elf.R_ALPHA_LITERAL}
	Roots[`/debug/elf/R_ALPHA_LITUSE`] = ConstRoot{Const: debug_elf.R_ALPHA_LITUSE}
	Roots[`/debug/elf/R_ALPHA_NONE`] = ConstRoot{Const: debug_elf.R_ALPHA_NONE}
	Roots[`/debug/elf/R_ALPHA_OP_PRSHIFT`] = ConstRoot{Const: debug_elf.R_ALPHA_OP_PRSHIFT}
	Roots[`/debug/elf/R_ALPHA_OP_PSUB`] = ConstRoot{Const: debug_elf.R_ALPHA_OP_PSUB}
	Roots[`/debug/elf/R_ALPHA_OP_PUSH`] = ConstRoot{Const: debug_elf.R_ALPHA_OP_PUSH}
	Roots[`/debug/elf/R_ALPHA_OP_STORE`] = ConstRoot{Const: debug_elf.R_ALPHA_OP_STORE}
	Roots[`/debug/elf/R_ALPHA_REFLONG`] = ConstRoot{Const: debug_elf.R_ALPHA_REFLONG}
	Roots[`/debug/elf/R_ALPHA_REFQUAD`] = ConstRoot{Const: debug_elf.R_ALPHA_REFQUAD}
	Roots[`/debug/elf/R_ALPHA_RELATIVE`] = ConstRoot{Const: debug_elf.R_ALPHA_RELATIVE}
	Roots[`/debug/elf/R_ALPHA_SREL16`] = ConstRoot{Const: debug_elf.R_ALPHA_SREL16}
	Roots[`/debug/elf/R_ALPHA_SREL32`] = ConstRoot{Const: debug_elf.R_ALPHA_SREL32}
	Roots[`/debug/elf/R_ALPHA_SREL64`] = ConstRoot{Const: debug_elf.R_ALPHA_SREL64}
	Roots[`/debug/elf/R_ARM_ABS12`] = ConstRoot{Const: debug_elf.R_ARM_ABS12}
	Roots[`/debug/elf/R_ARM_ABS16`] = ConstRoot{Const: debug_elf.R_ARM_ABS16}
	Roots[`/debug/elf/R_ARM_ABS32`] = ConstRoot{Const: debug_elf.R_ARM_ABS32}
	Roots[`/debug/elf/R_ARM_ABS8`] = ConstRoot{Const: debug_elf.R_ARM_ABS8}
	Roots[`/debug/elf/R_ARM_AMP_VCALL9`] = ConstRoot{Const: debug_elf.R_ARM_AMP_VCALL9}
	Roots[`/debug/elf/R_ARM_COPY`] = ConstRoot{Const: debug_elf.R_ARM_COPY}
	Roots[`/debug/elf/R_ARM_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_ARM_GLOB_DAT}
	Roots[`/debug/elf/R_ARM_GNU_VTENTRY`] = ConstRoot{Const: debug_elf.R_ARM_GNU_VTENTRY}
	Roots[`/debug/elf/R_ARM_GNU_VTINHERIT`] = ConstRoot{Const: debug_elf.R_ARM_GNU_VTINHERIT}
	Roots[`/debug/elf/R_ARM_GOT32`] = ConstRoot{Const: debug_elf.R_ARM_GOT32}
	Roots[`/debug/elf/R_ARM_GOTOFF`] = ConstRoot{Const: debug_elf.R_ARM_GOTOFF}
	Roots[`/debug/elf/R_ARM_GOTPC`] = ConstRoot{Const: debug_elf.R_ARM_GOTPC}
	Roots[`/debug/elf/R_ARM_JUMP_SLOT`] = ConstRoot{Const: debug_elf.R_ARM_JUMP_SLOT}
	Roots[`/debug/elf/R_ARM_NONE`] = ConstRoot{Const: debug_elf.R_ARM_NONE}
	Roots[`/debug/elf/R_ARM_PC13`] = ConstRoot{Const: debug_elf.R_ARM_PC13}
	Roots[`/debug/elf/R_ARM_PC24`] = ConstRoot{Const: debug_elf.R_ARM_PC24}
	Roots[`/debug/elf/R_ARM_PLT32`] = ConstRoot{Const: debug_elf.R_ARM_PLT32}
	Roots[`/debug/elf/R_ARM_RABS32`] = ConstRoot{Const: debug_elf.R_ARM_RABS32}
	Roots[`/debug/elf/R_ARM_RBASE`] = ConstRoot{Const: debug_elf.R_ARM_RBASE}
	Roots[`/debug/elf/R_ARM_REL32`] = ConstRoot{Const: debug_elf.R_ARM_REL32}
	Roots[`/debug/elf/R_ARM_RELATIVE`] = ConstRoot{Const: debug_elf.R_ARM_RELATIVE}
	Roots[`/debug/elf/R_ARM_RPC24`] = ConstRoot{Const: debug_elf.R_ARM_RPC24}
	Roots[`/debug/elf/R_ARM_RREL32`] = ConstRoot{Const: debug_elf.R_ARM_RREL32}
	Roots[`/debug/elf/R_ARM_RSBREL32`] = ConstRoot{Const: debug_elf.R_ARM_RSBREL32}
	Roots[`/debug/elf/R_ARM_SBREL32`] = ConstRoot{Const: debug_elf.R_ARM_SBREL32}
	Roots[`/debug/elf/R_ARM_SWI24`] = ConstRoot{Const: debug_elf.R_ARM_SWI24}
	Roots[`/debug/elf/R_ARM_THM_ABS5`] = ConstRoot{Const: debug_elf.R_ARM_THM_ABS5}
	Roots[`/debug/elf/R_ARM_THM_PC22`] = ConstRoot{Const: debug_elf.R_ARM_THM_PC22}
	Roots[`/debug/elf/R_ARM_THM_PC8`] = ConstRoot{Const: debug_elf.R_ARM_THM_PC8}
	Roots[`/debug/elf/R_ARM_THM_RPC22`] = ConstRoot{Const: debug_elf.R_ARM_THM_RPC22}
	Roots[`/debug/elf/R_ARM_THM_SWI8`] = ConstRoot{Const: debug_elf.R_ARM_THM_SWI8}
	Roots[`/debug/elf/R_ARM_THM_XPC22`] = ConstRoot{Const: debug_elf.R_ARM_THM_XPC22}
	Roots[`/debug/elf/R_ARM_XPC25`] = ConstRoot{Const: debug_elf.R_ARM_XPC25}
	Roots[`/debug/elf/R_PPC_ADDR14`] = ConstRoot{Const: debug_elf.R_PPC_ADDR14}
	Roots[`/debug/elf/R_PPC_ADDR14_BRNTAKEN`] = ConstRoot{Const: debug_elf.R_PPC_ADDR14_BRNTAKEN}
	Roots[`/debug/elf/R_PPC_ADDR14_BRTAKEN`] = ConstRoot{Const: debug_elf.R_PPC_ADDR14_BRTAKEN}
	Roots[`/debug/elf/R_PPC_ADDR16`] = ConstRoot{Const: debug_elf.R_PPC_ADDR16}
	Roots[`/debug/elf/R_PPC_ADDR16_HA`] = ConstRoot{Const: debug_elf.R_PPC_ADDR16_HA}
	Roots[`/debug/elf/R_PPC_ADDR16_HI`] = ConstRoot{Const: debug_elf.R_PPC_ADDR16_HI}
	Roots[`/debug/elf/R_PPC_ADDR16_LO`] = ConstRoot{Const: debug_elf.R_PPC_ADDR16_LO}
	Roots[`/debug/elf/R_PPC_ADDR24`] = ConstRoot{Const: debug_elf.R_PPC_ADDR24}
	Roots[`/debug/elf/R_PPC_ADDR32`] = ConstRoot{Const: debug_elf.R_PPC_ADDR32}
	Roots[`/debug/elf/R_PPC_COPY`] = ConstRoot{Const: debug_elf.R_PPC_COPY}
	Roots[`/debug/elf/R_PPC_DTPMOD32`] = ConstRoot{Const: debug_elf.R_PPC_DTPMOD32}
	Roots[`/debug/elf/R_PPC_DTPREL16`] = ConstRoot{Const: debug_elf.R_PPC_DTPREL16}
	Roots[`/debug/elf/R_PPC_DTPREL16_HA`] = ConstRoot{Const: debug_elf.R_PPC_DTPREL16_HA}
	Roots[`/debug/elf/R_PPC_DTPREL16_HI`] = ConstRoot{Const: debug_elf.R_PPC_DTPREL16_HI}
	Roots[`/debug/elf/R_PPC_DTPREL16_LO`] = ConstRoot{Const: debug_elf.R_PPC_DTPREL16_LO}
	Roots[`/debug/elf/R_PPC_DTPREL32`] = ConstRoot{Const: debug_elf.R_PPC_DTPREL32}
	Roots[`/debug/elf/R_PPC_EMB_BIT_FLD`] = ConstRoot{Const: debug_elf.R_PPC_EMB_BIT_FLD}
	Roots[`/debug/elf/R_PPC_EMB_MRKREF`] = ConstRoot{Const: debug_elf.R_PPC_EMB_MRKREF}
	Roots[`/debug/elf/R_PPC_EMB_NADDR16`] = ConstRoot{Const: debug_elf.R_PPC_EMB_NADDR16}
	Roots[`/debug/elf/R_PPC_EMB_NADDR16_HA`] = ConstRoot{Const: debug_elf.R_PPC_EMB_NADDR16_HA}
	Roots[`/debug/elf/R_PPC_EMB_NADDR16_HI`] = ConstRoot{Const: debug_elf.R_PPC_EMB_NADDR16_HI}
	Roots[`/debug/elf/R_PPC_EMB_NADDR16_LO`] = ConstRoot{Const: debug_elf.R_PPC_EMB_NADDR16_LO}
	Roots[`/debug/elf/R_PPC_EMB_NADDR32`] = ConstRoot{Const: debug_elf.R_PPC_EMB_NADDR32}
	Roots[`/debug/elf/R_PPC_EMB_RELSDA`] = ConstRoot{Const: debug_elf.R_PPC_EMB_RELSDA}
	Roots[`/debug/elf/R_PPC_EMB_RELSEC16`] = ConstRoot{Const: debug_elf.R_PPC_EMB_RELSEC16}
	Roots[`/debug/elf/R_PPC_EMB_RELST_HA`] = ConstRoot{Const: debug_elf.R_PPC_EMB_RELST_HA}
	Roots[`/debug/elf/R_PPC_EMB_RELST_HI`] = ConstRoot{Const: debug_elf.R_PPC_EMB_RELST_HI}
	Roots[`/debug/elf/R_PPC_EMB_RELST_LO`] = ConstRoot{Const: debug_elf.R_PPC_EMB_RELST_LO}
	Roots[`/debug/elf/R_PPC_EMB_SDA21`] = ConstRoot{Const: debug_elf.R_PPC_EMB_SDA21}
	Roots[`/debug/elf/R_PPC_EMB_SDA2I16`] = ConstRoot{Const: debug_elf.R_PPC_EMB_SDA2I16}
	Roots[`/debug/elf/R_PPC_EMB_SDA2REL`] = ConstRoot{Const: debug_elf.R_PPC_EMB_SDA2REL}
	Roots[`/debug/elf/R_PPC_EMB_SDAI16`] = ConstRoot{Const: debug_elf.R_PPC_EMB_SDAI16}
	Roots[`/debug/elf/R_PPC_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_PPC_GLOB_DAT}
	Roots[`/debug/elf/R_PPC_GOT16`] = ConstRoot{Const: debug_elf.R_PPC_GOT16}
	Roots[`/debug/elf/R_PPC_GOT16_HA`] = ConstRoot{Const: debug_elf.R_PPC_GOT16_HA}
	Roots[`/debug/elf/R_PPC_GOT16_HI`] = ConstRoot{Const: debug_elf.R_PPC_GOT16_HI}
	Roots[`/debug/elf/R_PPC_GOT16_LO`] = ConstRoot{Const: debug_elf.R_PPC_GOT16_LO}
	Roots[`/debug/elf/R_PPC_GOT_TLSGD16`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSGD16}
	Roots[`/debug/elf/R_PPC_GOT_TLSGD16_HA`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSGD16_HA}
	Roots[`/debug/elf/R_PPC_GOT_TLSGD16_HI`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSGD16_HI}
	Roots[`/debug/elf/R_PPC_GOT_TLSGD16_LO`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSGD16_LO}
	Roots[`/debug/elf/R_PPC_GOT_TLSLD16`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSLD16}
	Roots[`/debug/elf/R_PPC_GOT_TLSLD16_HA`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSLD16_HA}
	Roots[`/debug/elf/R_PPC_GOT_TLSLD16_HI`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSLD16_HI}
	Roots[`/debug/elf/R_PPC_GOT_TLSLD16_LO`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TLSLD16_LO}
	Roots[`/debug/elf/R_PPC_GOT_TPREL16`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TPREL16}
	Roots[`/debug/elf/R_PPC_GOT_TPREL16_HA`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TPREL16_HA}
	Roots[`/debug/elf/R_PPC_GOT_TPREL16_HI`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TPREL16_HI}
	Roots[`/debug/elf/R_PPC_GOT_TPREL16_LO`] = ConstRoot{Const: debug_elf.R_PPC_GOT_TPREL16_LO}
	Roots[`/debug/elf/R_PPC_JMP_SLOT`] = ConstRoot{Const: debug_elf.R_PPC_JMP_SLOT}
	Roots[`/debug/elf/R_PPC_LOCAL24PC`] = ConstRoot{Const: debug_elf.R_PPC_LOCAL24PC}
	Roots[`/debug/elf/R_PPC_NONE`] = ConstRoot{Const: debug_elf.R_PPC_NONE}
	Roots[`/debug/elf/R_PPC_PLT16_HA`] = ConstRoot{Const: debug_elf.R_PPC_PLT16_HA}
	Roots[`/debug/elf/R_PPC_PLT16_HI`] = ConstRoot{Const: debug_elf.R_PPC_PLT16_HI}
	Roots[`/debug/elf/R_PPC_PLT16_LO`] = ConstRoot{Const: debug_elf.R_PPC_PLT16_LO}
	Roots[`/debug/elf/R_PPC_PLT32`] = ConstRoot{Const: debug_elf.R_PPC_PLT32}
	Roots[`/debug/elf/R_PPC_PLTREL24`] = ConstRoot{Const: debug_elf.R_PPC_PLTREL24}
	Roots[`/debug/elf/R_PPC_PLTREL32`] = ConstRoot{Const: debug_elf.R_PPC_PLTREL32}
	Roots[`/debug/elf/R_PPC_REL14`] = ConstRoot{Const: debug_elf.R_PPC_REL14}
	Roots[`/debug/elf/R_PPC_REL14_BRNTAKEN`] = ConstRoot{Const: debug_elf.R_PPC_REL14_BRNTAKEN}
	Roots[`/debug/elf/R_PPC_REL14_BRTAKEN`] = ConstRoot{Const: debug_elf.R_PPC_REL14_BRTAKEN}
	Roots[`/debug/elf/R_PPC_REL24`] = ConstRoot{Const: debug_elf.R_PPC_REL24}
	Roots[`/debug/elf/R_PPC_REL32`] = ConstRoot{Const: debug_elf.R_PPC_REL32}
	Roots[`/debug/elf/R_PPC_RELATIVE`] = ConstRoot{Const: debug_elf.R_PPC_RELATIVE}
	Roots[`/debug/elf/R_PPC_SDAREL16`] = ConstRoot{Const: debug_elf.R_PPC_SDAREL16}
	Roots[`/debug/elf/R_PPC_SECTOFF`] = ConstRoot{Const: debug_elf.R_PPC_SECTOFF}
	Roots[`/debug/elf/R_PPC_SECTOFF_HA`] = ConstRoot{Const: debug_elf.R_PPC_SECTOFF_HA}
	Roots[`/debug/elf/R_PPC_SECTOFF_HI`] = ConstRoot{Const: debug_elf.R_PPC_SECTOFF_HI}
	Roots[`/debug/elf/R_PPC_SECTOFF_LO`] = ConstRoot{Const: debug_elf.R_PPC_SECTOFF_LO}
	Roots[`/debug/elf/R_PPC_TLS`] = ConstRoot{Const: debug_elf.R_PPC_TLS}
	Roots[`/debug/elf/R_PPC_TPREL16`] = ConstRoot{Const: debug_elf.R_PPC_TPREL16}
	Roots[`/debug/elf/R_PPC_TPREL16_HA`] = ConstRoot{Const: debug_elf.R_PPC_TPREL16_HA}
	Roots[`/debug/elf/R_PPC_TPREL16_HI`] = ConstRoot{Const: debug_elf.R_PPC_TPREL16_HI}
	Roots[`/debug/elf/R_PPC_TPREL16_LO`] = ConstRoot{Const: debug_elf.R_PPC_TPREL16_LO}
	Roots[`/debug/elf/R_PPC_TPREL32`] = ConstRoot{Const: debug_elf.R_PPC_TPREL32}
	Roots[`/debug/elf/R_PPC_UADDR16`] = ConstRoot{Const: debug_elf.R_PPC_UADDR16}
	Roots[`/debug/elf/R_PPC_UADDR32`] = ConstRoot{Const: debug_elf.R_PPC_UADDR32}
	Roots[`/debug/elf/R_SPARC_10`] = ConstRoot{Const: debug_elf.R_SPARC_10}
	Roots[`/debug/elf/R_SPARC_11`] = ConstRoot{Const: debug_elf.R_SPARC_11}
	Roots[`/debug/elf/R_SPARC_13`] = ConstRoot{Const: debug_elf.R_SPARC_13}
	Roots[`/debug/elf/R_SPARC_16`] = ConstRoot{Const: debug_elf.R_SPARC_16}
	Roots[`/debug/elf/R_SPARC_22`] = ConstRoot{Const: debug_elf.R_SPARC_22}
	Roots[`/debug/elf/R_SPARC_32`] = ConstRoot{Const: debug_elf.R_SPARC_32}
	Roots[`/debug/elf/R_SPARC_5`] = ConstRoot{Const: debug_elf.R_SPARC_5}
	Roots[`/debug/elf/R_SPARC_6`] = ConstRoot{Const: debug_elf.R_SPARC_6}
	Roots[`/debug/elf/R_SPARC_64`] = ConstRoot{Const: debug_elf.R_SPARC_64}
	Roots[`/debug/elf/R_SPARC_7`] = ConstRoot{Const: debug_elf.R_SPARC_7}
	Roots[`/debug/elf/R_SPARC_8`] = ConstRoot{Const: debug_elf.R_SPARC_8}
	Roots[`/debug/elf/R_SPARC_COPY`] = ConstRoot{Const: debug_elf.R_SPARC_COPY}
	Roots[`/debug/elf/R_SPARC_DISP16`] = ConstRoot{Const: debug_elf.R_SPARC_DISP16}
	Roots[`/debug/elf/R_SPARC_DISP32`] = ConstRoot{Const: debug_elf.R_SPARC_DISP32}
	Roots[`/debug/elf/R_SPARC_DISP64`] = ConstRoot{Const: debug_elf.R_SPARC_DISP64}
	Roots[`/debug/elf/R_SPARC_DISP8`] = ConstRoot{Const: debug_elf.R_SPARC_DISP8}
	Roots[`/debug/elf/R_SPARC_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_SPARC_GLOB_DAT}
	Roots[`/debug/elf/R_SPARC_GLOB_JMP`] = ConstRoot{Const: debug_elf.R_SPARC_GLOB_JMP}
	Roots[`/debug/elf/R_SPARC_GOT10`] = ConstRoot{Const: debug_elf.R_SPARC_GOT10}
	Roots[`/debug/elf/R_SPARC_GOT13`] = ConstRoot{Const: debug_elf.R_SPARC_GOT13}
	Roots[`/debug/elf/R_SPARC_GOT22`] = ConstRoot{Const: debug_elf.R_SPARC_GOT22}
	Roots[`/debug/elf/R_SPARC_H44`] = ConstRoot{Const: debug_elf.R_SPARC_H44}
	Roots[`/debug/elf/R_SPARC_HH22`] = ConstRoot{Const: debug_elf.R_SPARC_HH22}
	Roots[`/debug/elf/R_SPARC_HI22`] = ConstRoot{Const: debug_elf.R_SPARC_HI22}
	Roots[`/debug/elf/R_SPARC_HIPLT22`] = ConstRoot{Const: debug_elf.R_SPARC_HIPLT22}
	Roots[`/debug/elf/R_SPARC_HIX22`] = ConstRoot{Const: debug_elf.R_SPARC_HIX22}
	Roots[`/debug/elf/R_SPARC_HM10`] = ConstRoot{Const: debug_elf.R_SPARC_HM10}
	Roots[`/debug/elf/R_SPARC_JMP_SLOT`] = ConstRoot{Const: debug_elf.R_SPARC_JMP_SLOT}
	Roots[`/debug/elf/R_SPARC_L44`] = ConstRoot{Const: debug_elf.R_SPARC_L44}
	Roots[`/debug/elf/R_SPARC_LM22`] = ConstRoot{Const: debug_elf.R_SPARC_LM22}
	Roots[`/debug/elf/R_SPARC_LO10`] = ConstRoot{Const: debug_elf.R_SPARC_LO10}
	Roots[`/debug/elf/R_SPARC_LOPLT10`] = ConstRoot{Const: debug_elf.R_SPARC_LOPLT10}
	Roots[`/debug/elf/R_SPARC_LOX10`] = ConstRoot{Const: debug_elf.R_SPARC_LOX10}
	Roots[`/debug/elf/R_SPARC_M44`] = ConstRoot{Const: debug_elf.R_SPARC_M44}
	Roots[`/debug/elf/R_SPARC_NONE`] = ConstRoot{Const: debug_elf.R_SPARC_NONE}
	Roots[`/debug/elf/R_SPARC_OLO10`] = ConstRoot{Const: debug_elf.R_SPARC_OLO10}
	Roots[`/debug/elf/R_SPARC_PC10`] = ConstRoot{Const: debug_elf.R_SPARC_PC10}
	Roots[`/debug/elf/R_SPARC_PC22`] = ConstRoot{Const: debug_elf.R_SPARC_PC22}
	Roots[`/debug/elf/R_SPARC_PCPLT10`] = ConstRoot{Const: debug_elf.R_SPARC_PCPLT10}
	Roots[`/debug/elf/R_SPARC_PCPLT22`] = ConstRoot{Const: debug_elf.R_SPARC_PCPLT22}
	Roots[`/debug/elf/R_SPARC_PCPLT32`] = ConstRoot{Const: debug_elf.R_SPARC_PCPLT32}
	Roots[`/debug/elf/R_SPARC_PC_HH22`] = ConstRoot{Const: debug_elf.R_SPARC_PC_HH22}
	Roots[`/debug/elf/R_SPARC_PC_HM10`] = ConstRoot{Const: debug_elf.R_SPARC_PC_HM10}
	Roots[`/debug/elf/R_SPARC_PC_LM22`] = ConstRoot{Const: debug_elf.R_SPARC_PC_LM22}
	Roots[`/debug/elf/R_SPARC_PLT32`] = ConstRoot{Const: debug_elf.R_SPARC_PLT32}
	Roots[`/debug/elf/R_SPARC_PLT64`] = ConstRoot{Const: debug_elf.R_SPARC_PLT64}
	Roots[`/debug/elf/R_SPARC_REGISTER`] = ConstRoot{Const: debug_elf.R_SPARC_REGISTER}
	Roots[`/debug/elf/R_SPARC_RELATIVE`] = ConstRoot{Const: debug_elf.R_SPARC_RELATIVE}
	Roots[`/debug/elf/R_SPARC_UA16`] = ConstRoot{Const: debug_elf.R_SPARC_UA16}
	Roots[`/debug/elf/R_SPARC_UA32`] = ConstRoot{Const: debug_elf.R_SPARC_UA32}
	Roots[`/debug/elf/R_SPARC_UA64`] = ConstRoot{Const: debug_elf.R_SPARC_UA64}
	Roots[`/debug/elf/R_SPARC_WDISP16`] = ConstRoot{Const: debug_elf.R_SPARC_WDISP16}
	Roots[`/debug/elf/R_SPARC_WDISP19`] = ConstRoot{Const: debug_elf.R_SPARC_WDISP19}
	Roots[`/debug/elf/R_SPARC_WDISP22`] = ConstRoot{Const: debug_elf.R_SPARC_WDISP22}
	Roots[`/debug/elf/R_SPARC_WDISP30`] = ConstRoot{Const: debug_elf.R_SPARC_WDISP30}
	Roots[`/debug/elf/R_SPARC_WPLT30`] = ConstRoot{Const: debug_elf.R_SPARC_WPLT30}
	Roots[`/debug/elf/R_X86_64_16`] = ConstRoot{Const: debug_elf.R_X86_64_16}
	Roots[`/debug/elf/R_X86_64_32`] = ConstRoot{Const: debug_elf.R_X86_64_32}
	Roots[`/debug/elf/R_X86_64_32S`] = ConstRoot{Const: debug_elf.R_X86_64_32S}
	Roots[`/debug/elf/R_X86_64_64`] = ConstRoot{Const: debug_elf.R_X86_64_64}
	Roots[`/debug/elf/R_X86_64_8`] = ConstRoot{Const: debug_elf.R_X86_64_8}
	Roots[`/debug/elf/R_X86_64_COPY`] = ConstRoot{Const: debug_elf.R_X86_64_COPY}
	Roots[`/debug/elf/R_X86_64_DTPMOD64`] = ConstRoot{Const: debug_elf.R_X86_64_DTPMOD64}
	Roots[`/debug/elf/R_X86_64_DTPOFF32`] = ConstRoot{Const: debug_elf.R_X86_64_DTPOFF32}
	Roots[`/debug/elf/R_X86_64_DTPOFF64`] = ConstRoot{Const: debug_elf.R_X86_64_DTPOFF64}
	Roots[`/debug/elf/R_X86_64_GLOB_DAT`] = ConstRoot{Const: debug_elf.R_X86_64_GLOB_DAT}
	Roots[`/debug/elf/R_X86_64_GOT32`] = ConstRoot{Const: debug_elf.R_X86_64_GOT32}
	Roots[`/debug/elf/R_X86_64_GOTPCREL`] = ConstRoot{Const: debug_elf.R_X86_64_GOTPCREL}
	Roots[`/debug/elf/R_X86_64_GOTTPOFF`] = ConstRoot{Const: debug_elf.R_X86_64_GOTTPOFF}
	Roots[`/debug/elf/R_X86_64_JMP_SLOT`] = ConstRoot{Const: debug_elf.R_X86_64_JMP_SLOT}
	Roots[`/debug/elf/R_X86_64_NONE`] = ConstRoot{Const: debug_elf.R_X86_64_NONE}
	Roots[`/debug/elf/R_X86_64_PC16`] = ConstRoot{Const: debug_elf.R_X86_64_PC16}
	Roots[`/debug/elf/R_X86_64_PC32`] = ConstRoot{Const: debug_elf.R_X86_64_PC32}
	Roots[`/debug/elf/R_X86_64_PC8`] = ConstRoot{Const: debug_elf.R_X86_64_PC8}
	Roots[`/debug/elf/R_X86_64_PLT32`] = ConstRoot{Const: debug_elf.R_X86_64_PLT32}
	Roots[`/debug/elf/R_X86_64_RELATIVE`] = ConstRoot{Const: debug_elf.R_X86_64_RELATIVE}
	Roots[`/debug/elf/R_X86_64_TLSGD`] = ConstRoot{Const: debug_elf.R_X86_64_TLSGD}
	Roots[`/debug/elf/R_X86_64_TLSLD`] = ConstRoot{Const: debug_elf.R_X86_64_TLSLD}
	Roots[`/debug/elf/R_X86_64_TPOFF32`] = ConstRoot{Const: debug_elf.R_X86_64_TPOFF32}
	Roots[`/debug/elf/R_X86_64_TPOFF64`] = ConstRoot{Const: debug_elf.R_X86_64_TPOFF64}
	Roots[`/debug/elf/SHF_ALLOC`] = ConstRoot{Const: debug_elf.SHF_ALLOC}
	Roots[`/debug/elf/SHF_EXECINSTR`] = ConstRoot{Const: debug_elf.SHF_EXECINSTR}
	Roots[`/debug/elf/SHF_GROUP`] = ConstRoot{Const: debug_elf.SHF_GROUP}
	Roots[`/debug/elf/SHF_INFO_LINK`] = ConstRoot{Const: debug_elf.SHF_INFO_LINK}
	Roots[`/debug/elf/SHF_LINK_ORDER`] = ConstRoot{Const: debug_elf.SHF_LINK_ORDER}
	Roots[`/debug/elf/SHF_MASKOS`] = ConstRoot{Const: debug_elf.SHF_MASKOS}
	Roots[`/debug/elf/SHF_MASKPROC`] = ConstRoot{Const: debug_elf.SHF_MASKPROC}
	Roots[`/debug/elf/SHF_MERGE`] = ConstRoot{Const: debug_elf.SHF_MERGE}
	Roots[`/debug/elf/SHF_OS_NONCONFORMING`] = ConstRoot{Const: debug_elf.SHF_OS_NONCONFORMING}
	Roots[`/debug/elf/SHF_STRINGS`] = ConstRoot{Const: debug_elf.SHF_STRINGS}
	Roots[`/debug/elf/SHF_TLS`] = ConstRoot{Const: debug_elf.SHF_TLS}
	Roots[`/debug/elf/SHF_WRITE`] = ConstRoot{Const: debug_elf.SHF_WRITE}
	Roots[`/debug/elf/SHN_ABS`] = ConstRoot{Const: debug_elf.SHN_ABS}
	Roots[`/debug/elf/SHN_COMMON`] = ConstRoot{Const: debug_elf.SHN_COMMON}
	Roots[`/debug/elf/SHN_HIOS`] = ConstRoot{Const: debug_elf.SHN_HIOS}
	Roots[`/debug/elf/SHN_HIPROC`] = ConstRoot{Const: debug_elf.SHN_HIPROC}
	Roots[`/debug/elf/SHN_HIRESERVE`] = ConstRoot{Const: debug_elf.SHN_HIRESERVE}
	Roots[`/debug/elf/SHN_LOOS`] = ConstRoot{Const: debug_elf.SHN_LOOS}
	Roots[`/debug/elf/SHN_LOPROC`] = ConstRoot{Const: debug_elf.SHN_LOPROC}
	Roots[`/debug/elf/SHN_LORESERVE`] = ConstRoot{Const: debug_elf.SHN_LORESERVE}
	Roots[`/debug/elf/SHN_UNDEF`] = ConstRoot{Const: debug_elf.SHN_UNDEF}
	Roots[`/debug/elf/SHN_XINDEX`] = ConstRoot{Const: debug_elf.SHN_XINDEX}
	Roots[`/debug/elf/SHT_DYNAMIC`] = ConstRoot{Const: debug_elf.SHT_DYNAMIC}
	Roots[`/debug/elf/SHT_DYNSYM`] = ConstRoot{Const: debug_elf.SHT_DYNSYM}
	Roots[`/debug/elf/SHT_FINI_ARRAY`] = ConstRoot{Const: debug_elf.SHT_FINI_ARRAY}
	Roots[`/debug/elf/SHT_GNU_ATTRIBUTES`] = ConstRoot{Const: debug_elf.SHT_GNU_ATTRIBUTES}
	Roots[`/debug/elf/SHT_GNU_HASH`] = ConstRoot{Const: debug_elf.SHT_GNU_HASH}
	Roots[`/debug/elf/SHT_GNU_LIBLIST`] = ConstRoot{Const: debug_elf.SHT_GNU_LIBLIST}
	Roots[`/debug/elf/SHT_GNU_VERDEF`] = ConstRoot{Const: debug_elf.SHT_GNU_VERDEF}
	Roots[`/debug/elf/SHT_GNU_VERNEED`] = ConstRoot{Const: debug_elf.SHT_GNU_VERNEED}
	Roots[`/debug/elf/SHT_GNU_VERSYM`] = ConstRoot{Const: debug_elf.SHT_GNU_VERSYM}
	Roots[`/debug/elf/SHT_GROUP`] = ConstRoot{Const: debug_elf.SHT_GROUP}
	Roots[`/debug/elf/SHT_HASH`] = ConstRoot{Const: debug_elf.SHT_HASH}
	Roots[`/debug/elf/SHT_HIOS`] = ConstRoot{Const: debug_elf.SHT_HIOS}
	Roots[`/debug/elf/SHT_HIPROC`] = ConstRoot{Const: debug_elf.SHT_HIPROC}
	Roots[`/debug/elf/SHT_HIUSER`] = ConstRoot{Const: debug_elf.SHT_HIUSER}
	Roots[`/debug/elf/SHT_INIT_ARRAY`] = ConstRoot{Const: debug_elf.SHT_INIT_ARRAY}
	Roots[`/debug/elf/SHT_LOOS`] = ConstRoot{Const: debug_elf.SHT_LOOS}
	Roots[`/debug/elf/SHT_LOPROC`] = ConstRoot{Const: debug_elf.SHT_LOPROC}
	Roots[`/debug/elf/SHT_LOUSER`] = ConstRoot{Const: debug_elf.SHT_LOUSER}
	Roots[`/debug/elf/SHT_NOBITS`] = ConstRoot{Const: debug_elf.SHT_NOBITS}
	Roots[`/debug/elf/SHT_NOTE`] = ConstRoot{Const: debug_elf.SHT_NOTE}
	Roots[`/debug/elf/SHT_NULL`] = ConstRoot{Const: debug_elf.SHT_NULL}
	Roots[`/debug/elf/SHT_PREINIT_ARRAY`] = ConstRoot{Const: debug_elf.SHT_PREINIT_ARRAY}
	Roots[`/debug/elf/SHT_PROGBITS`] = ConstRoot{Const: debug_elf.SHT_PROGBITS}
	Roots[`/debug/elf/SHT_REL`] = ConstRoot{Const: debug_elf.SHT_REL}
	Roots[`/debug/elf/SHT_RELA`] = ConstRoot{Const: debug_elf.SHT_RELA}
	Roots[`/debug/elf/SHT_SHLIB`] = ConstRoot{Const: debug_elf.SHT_SHLIB}
	Roots[`/debug/elf/SHT_STRTAB`] = ConstRoot{Const: debug_elf.SHT_STRTAB}
	Roots[`/debug/elf/SHT_SYMTAB`] = ConstRoot{Const: debug_elf.SHT_SYMTAB}
	Roots[`/debug/elf/SHT_SYMTAB_SHNDX`] = ConstRoot{Const: debug_elf.SHT_SYMTAB_SHNDX}
	Roots[`/debug/elf/STB_GLOBAL`] = ConstRoot{Const: debug_elf.STB_GLOBAL}
	Roots[`/debug/elf/STB_HIOS`] = ConstRoot{Const: debug_elf.STB_HIOS}
	Roots[`/debug/elf/STB_HIPROC`] = ConstRoot{Const: debug_elf.STB_HIPROC}
	Roots[`/debug/elf/STB_LOCAL`] = ConstRoot{Const: debug_elf.STB_LOCAL}
	Roots[`/debug/elf/STB_LOOS`] = ConstRoot{Const: debug_elf.STB_LOOS}
	Roots[`/debug/elf/STB_LOPROC`] = ConstRoot{Const: debug_elf.STB_LOPROC}
	Roots[`/debug/elf/STB_WEAK`] = ConstRoot{Const: debug_elf.STB_WEAK}
	Roots[`/debug/elf/STT_COMMON`] = ConstRoot{Const: debug_elf.STT_COMMON}
	Roots[`/debug/elf/STT_FILE`] = ConstRoot{Const: debug_elf.STT_FILE}
	Roots[`/debug/elf/STT_FUNC`] = ConstRoot{Const: debug_elf.STT_FUNC}
	Roots[`/debug/elf/STT_HIOS`] = ConstRoot{Const: debug_elf.STT_HIOS}
	Roots[`/debug/elf/STT_HIPROC`] = ConstRoot{Const: debug_elf.STT_HIPROC}
	Roots[`/debug/elf/STT_LOOS`] = ConstRoot{Const: debug_elf.STT_LOOS}
	Roots[`/debug/elf/STT_LOPROC`] = ConstRoot{Const: debug_elf.STT_LOPROC}
	Roots[`/debug/elf/STT_NOTYPE`] = ConstRoot{Const: debug_elf.STT_NOTYPE}
	Roots[`/debug/elf/STT_OBJECT`] = ConstRoot{Const: debug_elf.STT_OBJECT}
	Roots[`/debug/elf/STT_SECTION`] = ConstRoot{Const: debug_elf.STT_SECTION}
	Roots[`/debug/elf/STT_TLS`] = ConstRoot{Const: debug_elf.STT_TLS}
	Roots[`/debug/elf/STV_DEFAULT`] = ConstRoot{Const: debug_elf.STV_DEFAULT}
	Roots[`/debug/elf/STV_HIDDEN`] = ConstRoot{Const: debug_elf.STV_HIDDEN}
	Roots[`/debug/elf/STV_INTERNAL`] = ConstRoot{Const: debug_elf.STV_INTERNAL}
	Roots[`/debug/elf/STV_PROTECTED`] = ConstRoot{Const: debug_elf.STV_PROTECTED}
	Roots[`/debug/elf/Sym32Size`] = ConstRoot{Const: int64(debug_elf.Sym32Size)}
	Roots[`/debug/elf/Sym64Size`] = ConstRoot{Const: int64(debug_elf.Sym64Size)}
	Roots[`/debug/macho/Cpu386`] = ConstRoot{Const: debug_macho.Cpu386}
	Roots[`/debug/macho/CpuAmd64`] = ConstRoot{Const: debug_macho.CpuAmd64}
	Roots[`/debug/macho/LoadCmdDylib`] = ConstRoot{Const: debug_macho.LoadCmdDylib}
	Roots[`/debug/macho/LoadCmdDylinker`] = ConstRoot{Const: debug_macho.LoadCmdDylinker}
	Roots[`/debug/macho/LoadCmdDysymtab`] = ConstRoot{Const: debug_macho.LoadCmdDysymtab}
	Roots[`/debug/macho/LoadCmdSegment`] = ConstRoot{Const: debug_macho.LoadCmdSegment}
	Roots[`/debug/macho/LoadCmdSegment64`] = ConstRoot{Const: debug_macho.LoadCmdSegment64}
	Roots[`/debug/macho/LoadCmdSymtab`] = ConstRoot{Const: debug_macho.LoadCmdSymtab}
	Roots[`/debug/macho/LoadCmdThread`] = ConstRoot{Const: debug_macho.LoadCmdThread}
	Roots[`/debug/macho/LoadCmdUnixThread`] = ConstRoot{Const: debug_macho.LoadCmdUnixThread}
	Roots[`/debug/macho/Magic32`] = ConstRoot{Const: debug_macho.Magic32}
	Roots[`/debug/macho/Magic64`] = ConstRoot{Const: debug_macho.Magic64}
	Roots[`/debug/macho/TypeExec`] = ConstRoot{Const: debug_macho.TypeExec}
	Roots[`/debug/macho/TypeObj`] = ConstRoot{Const: debug_macho.TypeObj}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_AM33`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_AM33)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_AMD64`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_AMD64)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_ARM`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_ARM)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_EBC`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_EBC)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_I386`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_I386)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_IA64`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_IA64)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_M32R`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_M32R)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_MIPS16`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_MIPS16)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_MIPSFPU`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_MIPSFPU)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_MIPSFPU16`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_MIPSFPU16)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_POWERPC`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_POWERPC)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_POWERPCFP`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_POWERPCFP)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_R4000`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_R4000)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_SH3`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_SH3)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_SH3DSP`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_SH3DSP)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_SH4`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_SH4)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_SH5`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_SH5)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_THUMB`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_THUMB)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_UNKNOWN`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_UNKNOWN)}
	Roots[`/debug/pe/IMAGE_FILE_MACHINE_WCEMIPSV2`] = ConstRoot{Const: int64(debug_pe.IMAGE_FILE_MACHINE_WCEMIPSV2)}
	Roots[`/encoding/binary/MaxVarintLen16`] = ConstRoot{Const: int64(encoding_binary.MaxVarintLen16)}
	Roots[`/encoding/binary/MaxVarintLen32`] = ConstRoot{Const: int64(encoding_binary.MaxVarintLen32)}
	Roots[`/encoding/binary/MaxVarintLen64`] = ConstRoot{Const: int64(encoding_binary.MaxVarintLen64)}
	Roots[`/encoding/xml/Header`] = ConstRoot{Const: encoding_xml.Header}
	Roots[`/flag/ContinueOnError`] = ConstRoot{Const: flag.ContinueOnError}
	Roots[`/flag/ExitOnError`] = ConstRoot{Const: flag.ExitOnError}
	Roots[`/flag/PanicOnError`] = ConstRoot{Const: flag.PanicOnError}
	Roots[`/go/ast/Bad`] = ConstRoot{Const: go_ast.Bad}
	Roots[`/go/ast/Con`] = ConstRoot{Const: go_ast.Con}
	Roots[`/go/ast/FilterFuncDuplicates`] = ConstRoot{Const: go_ast.FilterFuncDuplicates}
	Roots[`/go/ast/FilterImportDuplicates`] = ConstRoot{Const: go_ast.FilterImportDuplicates}
	Roots[`/go/ast/FilterUnassociatedComments`] = ConstRoot{Const: go_ast.FilterUnassociatedComments}
	Roots[`/go/ast/Fun`] = ConstRoot{Const: go_ast.Fun}
	Roots[`/go/ast/Lbl`] = ConstRoot{Const: go_ast.Lbl}
	Roots[`/go/ast/Pkg`] = ConstRoot{Const: go_ast.Pkg}
	Roots[`/go/ast/RECV`] = ConstRoot{Const: go_ast.RECV}
	Roots[`/go/ast/SEND`] = ConstRoot{Const: go_ast.SEND}
	Roots[`/go/ast/Typ`] = ConstRoot{Const: go_ast.Typ}
	Roots[`/go/ast/Var`] = ConstRoot{Const: go_ast.Var}
	Roots[`/go/build/AllowBinary`] = ConstRoot{Const: go_build.AllowBinary}
	Roots[`/go/build/FindOnly`] = ConstRoot{Const: go_build.FindOnly}
	Roots[`/go/doc/AllDecls`] = ConstRoot{Const: go_doc.AllDecls}
	Roots[`/go/doc/AllMethods`] = ConstRoot{Const: go_doc.AllMethods}
	Roots[`/go/parser/DeclarationErrors`] = ConstRoot{Const: go_parser.DeclarationErrors}
	Roots[`/go/parser/ImportsOnly`] = ConstRoot{Const: go_parser.ImportsOnly}
	Roots[`/go/parser/PackageClauseOnly`] = ConstRoot{Const: go_parser.PackageClauseOnly}
	Roots[`/go/parser/ParseComments`] = ConstRoot{Const: go_parser.ParseComments}
	Roots[`/go/parser/SpuriousErrors`] = ConstRoot{Const: go_parser.SpuriousErrors}
	Roots[`/go/parser/Trace`] = ConstRoot{Const: go_parser.Trace}
	Roots[`/go/printer/RawFormat`] = ConstRoot{Const: go_printer.RawFormat}
	Roots[`/go/printer/SourcePos`] = ConstRoot{Const: go_printer.SourcePos}
	Roots[`/go/printer/TabIndent`] = ConstRoot{Const: go_printer.TabIndent}
	Roots[`/go/printer/UseSpaces`] = ConstRoot{Const: go_printer.UseSpaces}
	Roots[`/go/scanner/ScanComments`] = ConstRoot{Const: go_scanner.ScanComments}
	Roots[`/go/token/ADD`] = ConstRoot{Const: go_token.ADD}
	Roots[`/go/token/ADD_ASSIGN`] = ConstRoot{Const: go_token.ADD_ASSIGN}
	Roots[`/go/token/AND`] = ConstRoot{Const: go_token.AND}
	Roots[`/go/token/AND_ASSIGN`] = ConstRoot{Const: go_token.AND_ASSIGN}
	Roots[`/go/token/AND_NOT`] = ConstRoot{Const: go_token.AND_NOT}
	Roots[`/go/token/AND_NOT_ASSIGN`] = ConstRoot{Const: go_token.AND_NOT_ASSIGN}
	Roots[`/go/token/ARROW`] = ConstRoot{Const: go_token.ARROW}
	Roots[`/go/token/ASSIGN`] = ConstRoot{Const: go_token.ASSIGN}
	Roots[`/go/token/BREAK`] = ConstRoot{Const: go_token.BREAK}
	Roots[`/go/token/CASE`] = ConstRoot{Const: go_token.CASE}
	Roots[`/go/token/CHAN`] = ConstRoot{Const: go_token.CHAN}
	Roots[`/go/token/CHAR`] = ConstRoot{Const: go_token.CHAR}
	Roots[`/go/token/COLON`] = ConstRoot{Const: go_token.COLON}
	Roots[`/go/token/COMMA`] = ConstRoot{Const: go_token.COMMA}
	Roots[`/go/token/COMMENT`] = ConstRoot{Const: go_token.COMMENT}
	Roots[`/go/token/CONST`] = ConstRoot{Const: go_token.CONST}
	Roots[`/go/token/CONTINUE`] = ConstRoot{Const: go_token.CONTINUE}
	Roots[`/go/token/DEC`] = ConstRoot{Const: go_token.DEC}
	Roots[`/go/token/DEFAULT`] = ConstRoot{Const: go_token.DEFAULT}
	Roots[`/go/token/DEFER`] = ConstRoot{Const: go_token.DEFER}
	Roots[`/go/token/DEFINE`] = ConstRoot{Const: go_token.DEFINE}
	Roots[`/go/token/ELLIPSIS`] = ConstRoot{Const: go_token.ELLIPSIS}
	Roots[`/go/token/ELSE`] = ConstRoot{Const: go_token.ELSE}
	Roots[`/go/token/EOF`] = ConstRoot{Const: go_token.EOF}
	Roots[`/go/token/EQL`] = ConstRoot{Const: go_token.EQL}
	Roots[`/go/token/FALLTHROUGH`] = ConstRoot{Const: go_token.FALLTHROUGH}
	Roots[`/go/token/FLOAT`] = ConstRoot{Const: go_token.FLOAT}
	Roots[`/go/token/FOR`] = ConstRoot{Const: go_token.FOR}
	Roots[`/go/token/FUNC`] = ConstRoot{Const: go_token.FUNC}
	Roots[`/go/token/GEQ`] = ConstRoot{Const: go_token.GEQ}
	Roots[`/go/token/GO`] = ConstRoot{Const: go_token.GO}
	Roots[`/go/token/GOTO`] = ConstRoot{Const: go_token.GOTO}
	Roots[`/go/token/GTR`] = ConstRoot{Const: go_token.GTR}
	Roots[`/go/token/HighestPrec`] = ConstRoot{Const: int64(go_token.HighestPrec)}
	Roots[`/go/token/IDENT`] = ConstRoot{Const: go_token.IDENT}
	Roots[`/go/token/IF`] = ConstRoot{Const: go_token.IF}
	Roots[`/go/token/ILLEGAL`] = ConstRoot{Const: go_token.ILLEGAL}
	Roots[`/go/token/IMAG`] = ConstRoot{Const: go_token.IMAG}
	Roots[`/go/token/IMPORT`] = ConstRoot{Const: go_token.IMPORT}
	Roots[`/go/token/INC`] = ConstRoot{Const: go_token.INC}
	Roots[`/go/token/INT`] = ConstRoot{Const: go_token.INT}
	Roots[`/go/token/INTERFACE`] = ConstRoot{Const: go_token.INTERFACE}
	Roots[`/go/token/LAND`] = ConstRoot{Const: go_token.LAND}
	Roots[`/go/token/LBRACE`] = ConstRoot{Const: go_token.LBRACE}
	Roots[`/go/token/LBRACK`] = ConstRoot{Const: go_token.LBRACK}
	Roots[`/go/token/LEQ`] = ConstRoot{Const: go_token.LEQ}
	Roots[`/go/token/LOR`] = ConstRoot{Const: go_token.LOR}
	Roots[`/go/token/LPAREN`] = ConstRoot{Const: go_token.LPAREN}
	Roots[`/go/token/LSS`] = ConstRoot{Const: go_token.LSS}
	Roots[`/go/token/LowestPrec`] = ConstRoot{Const: int64(go_token.LowestPrec)}
	Roots[`/go/token/MAP`] = ConstRoot{Const: go_token.MAP}
	Roots[`/go/token/MUL`] = ConstRoot{Const: go_token.MUL}
	Roots[`/go/token/MUL_ASSIGN`] = ConstRoot{Const: go_token.MUL_ASSIGN}
	Roots[`/go/token/NEQ`] = ConstRoot{Const: go_token.NEQ}
	Roots[`/go/token/NOT`] = ConstRoot{Const: go_token.NOT}
	Roots[`/go/token/NoPos`] = ConstRoot{Const: go_token.NoPos}
	Roots[`/go/token/OR`] = ConstRoot{Const: go_token.OR}
	Roots[`/go/token/OR_ASSIGN`] = ConstRoot{Const: go_token.OR_ASSIGN}
	Roots[`/go/token/PACKAGE`] = ConstRoot{Const: go_token.PACKAGE}
	Roots[`/go/token/PERIOD`] = ConstRoot{Const: go_token.PERIOD}
	Roots[`/go/token/QUO`] = ConstRoot{Const: go_token.QUO}
	Roots[`/go/token/QUO_ASSIGN`] = ConstRoot{Const: go_token.QUO_ASSIGN}
	Roots[`/go/token/RANGE`] = ConstRoot{Const: go_token.RANGE}
	Roots[`/go/token/RBRACE`] = ConstRoot{Const: go_token.RBRACE}
	Roots[`/go/token/RBRACK`] = ConstRoot{Const: go_token.RBRACK}
	Roots[`/go/token/REM`] = ConstRoot{Const: go_token.REM}
	Roots[`/go/token/REM_ASSIGN`] = ConstRoot{Const: go_token.REM_ASSIGN}
	Roots[`/go/token/RETURN`] = ConstRoot{Const: go_token.RETURN}
	Roots[`/go/token/RPAREN`] = ConstRoot{Const: go_token.RPAREN}
	Roots[`/go/token/SELECT`] = ConstRoot{Const: go_token.SELECT}
	Roots[`/go/token/SEMICOLON`] = ConstRoot{Const: go_token.SEMICOLON}
	Roots[`/go/token/SHL`] = ConstRoot{Const: go_token.SHL}
	Roots[`/go/token/SHL_ASSIGN`] = ConstRoot{Const: go_token.SHL_ASSIGN}
	Roots[`/go/token/SHR`] = ConstRoot{Const: go_token.SHR}
	Roots[`/go/token/SHR_ASSIGN`] = ConstRoot{Const: go_token.SHR_ASSIGN}
	Roots[`/go/token/STRING`] = ConstRoot{Const: go_token.STRING}
	Roots[`/go/token/STRUCT`] = ConstRoot{Const: go_token.STRUCT}
	Roots[`/go/token/SUB`] = ConstRoot{Const: go_token.SUB}
	Roots[`/go/token/SUB_ASSIGN`] = ConstRoot{Const: go_token.SUB_ASSIGN}
	Roots[`/go/token/SWITCH`] = ConstRoot{Const: go_token.SWITCH}
	Roots[`/go/token/TYPE`] = ConstRoot{Const: go_token.TYPE}
	Roots[`/go/token/UnaryPrec`] = ConstRoot{Const: int64(go_token.UnaryPrec)}
	Roots[`/go/token/VAR`] = ConstRoot{Const: go_token.VAR}
	Roots[`/go/token/XOR`] = ConstRoot{Const: go_token.XOR}
	Roots[`/go/token/XOR_ASSIGN`] = ConstRoot{Const: go_token.XOR_ASSIGN}
	Roots[`/hash/adler32/Size`] = ConstRoot{Const: int64(hash_adler32.Size)}
	Roots[`/hash/crc32/Castagnoli`] = ConstRoot{Const: int64(hash_crc32.Castagnoli)}
	Roots[`/hash/crc32/IEEE`] = ConstRoot{Const: int64(hash_crc32.IEEE)}
	Roots[`/hash/crc32/Koopman`] = ConstRoot{Const: int64(hash_crc32.Koopman)}
	Roots[`/hash/crc32/Size`] = ConstRoot{Const: int64(hash_crc32.Size)}
	Roots[`/hash/crc64/ECMA`] = ConstRoot{Const: uint64(hash_crc64.ECMA)}
	Roots[`/hash/crc64/ISO`] = ConstRoot{Const: uint64(hash_crc64.ISO)}
	Roots[`/hash/crc64/Size`] = ConstRoot{Const: uint64(hash_crc64.Size)}
	Roots[`/html/template/ErrAmbigContext`] = ConstRoot{Const: html_template.ErrAmbigContext}
	Roots[`/html/template/ErrBadHTML`] = ConstRoot{Const: html_template.ErrBadHTML}
	Roots[`/html/template/ErrBranchEnd`] = ConstRoot{Const: html_template.ErrBranchEnd}
	Roots[`/html/template/ErrEndContext`] = ConstRoot{Const: html_template.ErrEndContext}
	Roots[`/html/template/ErrNoSuchTemplate`] = ConstRoot{Const: html_template.ErrNoSuchTemplate}
	Roots[`/html/template/ErrOutputContext`] = ConstRoot{Const: html_template.ErrOutputContext}
	Roots[`/html/template/ErrPartialCharset`] = ConstRoot{Const: html_template.ErrPartialCharset}
	Roots[`/html/template/ErrPartialEscape`] = ConstRoot{Const: html_template.ErrPartialEscape}
	Roots[`/html/template/ErrRangeLoopReentry`] = ConstRoot{Const: html_template.ErrRangeLoopReentry}
	Roots[`/html/template/ErrSlashAmbig`] = ConstRoot{Const: html_template.ErrSlashAmbig}
	Roots[`/html/template/OK`] = ConstRoot{Const: html_template.OK}
	Roots[`/image/YCbCrSubsampleRatio420`] = ConstRoot{Const: image.YCbCrSubsampleRatio420}
	Roots[`/image/YCbCrSubsampleRatio422`] = ConstRoot{Const: image.YCbCrSubsampleRatio422}
	Roots[`/image/YCbCrSubsampleRatio444`] = ConstRoot{Const: image.YCbCrSubsampleRatio444}
	Roots[`/image/draw/Over`] = ConstRoot{Const: image_draw.Over}
	Roots[`/image/draw/Src`] = ConstRoot{Const: image_draw.Src}
	Roots[`/image/jpeg/DefaultQuality`] = ConstRoot{Const: int64(image_jpeg.DefaultQuality)}
	Roots[`/log/Ldate`] = ConstRoot{Const: int64(log.Ldate)}
	Roots[`/log/Llongfile`] = ConstRoot{Const: int64(log.Llongfile)}
	Roots[`/log/Lmicroseconds`] = ConstRoot{Const: int64(log.Lmicroseconds)}
	Roots[`/log/Lshortfile`] = ConstRoot{Const: int64(log.Lshortfile)}
	Roots[`/log/LstdFlags`] = ConstRoot{Const: int64(log.LstdFlags)}
	Roots[`/log/Ltime`] = ConstRoot{Const: int64(log.Ltime)}
	Roots[`/math/E`] = ConstRoot{Const: math.E}
	Roots[`/math/Ln10`] = ConstRoot{Const: math.Ln10}
	Roots[`/math/Ln2`] = ConstRoot{Const: math.Ln2}
	Roots[`/math/Log10E`] = ConstRoot{Const: math.Log10E}
	Roots[`/math/Log2E`] = ConstRoot{Const: math.Log2E}
	Roots[`/math/MaxFloat32`] = ConstRoot{Const: math.MaxFloat32}
	Roots[`/math/MaxFloat64`] = ConstRoot{Const: math.MaxFloat64}
	Roots[`/math/MaxInt16`] = ConstRoot{Const: int64(math.MaxInt16)}
	Roots[`/math/MaxInt32`] = ConstRoot{Const: int64(math.MaxInt32)}
	Roots[`/math/MaxInt64`] = ConstRoot{Const: int64(math.MaxInt64)}
	Roots[`/math/MaxInt8`] = ConstRoot{Const: int64(math.MaxInt8)}
	Roots[`/math/MaxUint16`] = ConstRoot{Const: int64(math.MaxUint16)}
	Roots[`/math/MaxUint32`] = ConstRoot{Const: int64(math.MaxUint32)}
	Roots[`/math/MaxUint64`] = ConstRoot{Const: uint64(math.MaxUint64)}
	Roots[`/math/MaxUint8`] = ConstRoot{Const: int64(math.MaxUint8)}
	Roots[`/math/MinInt16`] = ConstRoot{Const: int64(math.MinInt16)}
	Roots[`/math/MinInt32`] = ConstRoot{Const: int64(math.MinInt32)}
	Roots[`/math/MinInt64`] = ConstRoot{Const: int64(math.MinInt64)}
	Roots[`/math/MinInt8`] = ConstRoot{Const: int64(math.MinInt8)}
	Roots[`/math/Phi`] = ConstRoot{Const: math.Phi}
	Roots[`/math/Pi`] = ConstRoot{Const: math.Pi}
	Roots[`/math/SmallestNonzeroFloat32`] = ConstRoot{Const: math.SmallestNonzeroFloat32}
	Roots[`/math/SmallestNonzeroFloat64`] = ConstRoot{Const: math.SmallestNonzeroFloat64}
	Roots[`/math/Sqrt2`] = ConstRoot{Const: math.Sqrt2}
	Roots[`/math/SqrtE`] = ConstRoot{Const: math.SqrtE}
	Roots[`/math/SqrtPhi`] = ConstRoot{Const: math.SqrtPhi}
	Roots[`/math/SqrtPi`] = ConstRoot{Const: math.SqrtPi}
	Roots[`/math/big/MaxBase`] = ConstRoot{Const: int64(math_big.MaxBase)}
	Roots[`/net/FlagBroadcast`] = ConstRoot{Const: net.FlagBroadcast}
	Roots[`/net/FlagLoopback`] = ConstRoot{Const: net.FlagLoopback}
	Roots[`/net/FlagMulticast`] = ConstRoot{Const: net.FlagMulticast}
	Roots[`/net/FlagPointToPoint`] = ConstRoot{Const: net.FlagPointToPoint}
	Roots[`/net/FlagUp`] = ConstRoot{Const: net.FlagUp}
	Roots[`/net/IPv4len`] = ConstRoot{Const: int64(net.IPv4len)}
	Roots[`/net/IPv6len`] = ConstRoot{Const: int64(net.IPv6len)}
	Roots[`/net/http/DefaultMaxHeaderBytes`] = ConstRoot{Const: int64(net_http.DefaultMaxHeaderBytes)}
	Roots[`/net/http/DefaultMaxIdleConnsPerHost`] = ConstRoot{Const: int64(net_http.DefaultMaxIdleConnsPerHost)}
	Roots[`/net/http/StatusAccepted`] = ConstRoot{Const: int64(net_http.StatusAccepted)}
	Roots[`/net/http/StatusBadGateway`] = ConstRoot{Const: int64(net_http.StatusBadGateway)}
	Roots[`/net/http/StatusBadRequest`] = ConstRoot{Const: int64(net_http.StatusBadRequest)}
	Roots[`/net/http/StatusConflict`] = ConstRoot{Const: int64(net_http.StatusConflict)}
	Roots[`/net/http/StatusContinue`] = ConstRoot{Const: int64(net_http.StatusContinue)}
	Roots[`/net/http/StatusCreated`] = ConstRoot{Const: int64(net_http.StatusCreated)}
	Roots[`/net/http/StatusExpectationFailed`] = ConstRoot{Const: int64(net_http.StatusExpectationFailed)}
	Roots[`/net/http/StatusForbidden`] = ConstRoot{Const: int64(net_http.StatusForbidden)}
	Roots[`/net/http/StatusFound`] = ConstRoot{Const: int64(net_http.StatusFound)}
	Roots[`/net/http/StatusGatewayTimeout`] = ConstRoot{Const: int64(net_http.StatusGatewayTimeout)}
	Roots[`/net/http/StatusGone`] = ConstRoot{Const: int64(net_http.StatusGone)}
	Roots[`/net/http/StatusHTTPVersionNotSupported`] = ConstRoot{Const: int64(net_http.StatusHTTPVersionNotSupported)}
	Roots[`/net/http/StatusInternalServerError`] = ConstRoot{Const: int64(net_http.StatusInternalServerError)}
	Roots[`/net/http/StatusLengthRequired`] = ConstRoot{Const: int64(net_http.StatusLengthRequired)}
	Roots[`/net/http/StatusMethodNotAllowed`] = ConstRoot{Const: int64(net_http.StatusMethodNotAllowed)}
	Roots[`/net/http/StatusMovedPermanently`] = ConstRoot{Const: int64(net_http.StatusMovedPermanently)}
	Roots[`/net/http/StatusMultipleChoices`] = ConstRoot{Const: int64(net_http.StatusMultipleChoices)}
	Roots[`/net/http/StatusNoContent`] = ConstRoot{Const: int64(net_http.StatusNoContent)}
	Roots[`/net/http/StatusNonAuthoritativeInfo`] = ConstRoot{Const: int64(net_http.StatusNonAuthoritativeInfo)}
	Roots[`/net/http/StatusNotAcceptable`] = ConstRoot{Const: int64(net_http.StatusNotAcceptable)}
	Roots[`/net/http/StatusNotFound`] = ConstRoot{Const: int64(net_http.StatusNotFound)}
	Roots[`/net/http/StatusNotImplemented`] = ConstRoot{Const: int64(net_http.StatusNotImplemented)}
	Roots[`/net/http/StatusNotModified`] = ConstRoot{Const: int64(net_http.StatusNotModified)}
	Roots[`/net/http/StatusOK`] = ConstRoot{Const: int64(net_http.StatusOK)}
	Roots[`/net/http/StatusPartialContent`] = ConstRoot{Const: int64(net_http.StatusPartialContent)}
	Roots[`/net/http/StatusPaymentRequired`] = ConstRoot{Const: int64(net_http.StatusPaymentRequired)}
	Roots[`/net/http/StatusPreconditionFailed`] = ConstRoot{Const: int64(net_http.StatusPreconditionFailed)}
	Roots[`/net/http/StatusProxyAuthRequired`] = ConstRoot{Const: int64(net_http.StatusProxyAuthRequired)}
	Roots[`/net/http/StatusRequestEntityTooLarge`] = ConstRoot{Const: int64(net_http.StatusRequestEntityTooLarge)}
	Roots[`/net/http/StatusRequestTimeout`] = ConstRoot{Const: int64(net_http.StatusRequestTimeout)}
	Roots[`/net/http/StatusRequestURITooLong`] = ConstRoot{Const: int64(net_http.StatusRequestURITooLong)}
	Roots[`/net/http/StatusRequestedRangeNotSatisfiable`] = ConstRoot{Const: int64(net_http.StatusRequestedRangeNotSatisfiable)}
	Roots[`/net/http/StatusResetContent`] = ConstRoot{Const: int64(net_http.StatusResetContent)}
	Roots[`/net/http/StatusSeeOther`] = ConstRoot{Const: int64(net_http.StatusSeeOther)}
	Roots[`/net/http/StatusServiceUnavailable`] = ConstRoot{Const: int64(net_http.StatusServiceUnavailable)}
	Roots[`/net/http/StatusSwitchingProtocols`] = ConstRoot{Const: int64(net_http.StatusSwitchingProtocols)}
	Roots[`/net/http/StatusTeapot`] = ConstRoot{Const: int64(net_http.StatusTeapot)}
	Roots[`/net/http/StatusTemporaryRedirect`] = ConstRoot{Const: int64(net_http.StatusTemporaryRedirect)}
	Roots[`/net/http/StatusUnauthorized`] = ConstRoot{Const: int64(net_http.StatusUnauthorized)}
	Roots[`/net/http/StatusUnsupportedMediaType`] = ConstRoot{Const: int64(net_http.StatusUnsupportedMediaType)}
	Roots[`/net/http/StatusUseProxy`] = ConstRoot{Const: int64(net_http.StatusUseProxy)}
	Roots[`/net/http/TimeFormat`] = ConstRoot{Const: net_http.TimeFormat}
	Roots[`/net/http/httptest/DefaultRemoteAddr`] = ConstRoot{Const: net_http_httptest.DefaultRemoteAddr}
	Roots[`/net/rpc/DefaultDebugPath`] = ConstRoot{Const: net_rpc.DefaultDebugPath}
	Roots[`/net/rpc/DefaultRPCPath`] = ConstRoot{Const: net_rpc.DefaultRPCPath}
	Roots[`/os/DevNull`] = ConstRoot{Const: os.DevNull}
	Roots[`/os/ModeAppend`] = ConstRoot{Const: os.ModeAppend}
	Roots[`/os/ModeCharDevice`] = ConstRoot{Const: os.ModeCharDevice}
	Roots[`/os/ModeDevice`] = ConstRoot{Const: os.ModeDevice}
	Roots[`/os/ModeDir`] = ConstRoot{Const: os.ModeDir}
	Roots[`/os/ModeExclusive`] = ConstRoot{Const: os.ModeExclusive}
	Roots[`/os/ModeNamedPipe`] = ConstRoot{Const: os.ModeNamedPipe}
	Roots[`/os/ModePerm`] = ConstRoot{Const: os.ModePerm}
	Roots[`/os/ModeSetgid`] = ConstRoot{Const: os.ModeSetgid}
	Roots[`/os/ModeSetuid`] = ConstRoot{Const: os.ModeSetuid}
	Roots[`/os/ModeSocket`] = ConstRoot{Const: os.ModeSocket}
	Roots[`/os/ModeSticky`] = ConstRoot{Const: os.ModeSticky}
	Roots[`/os/ModeSymlink`] = ConstRoot{Const: os.ModeSymlink}
	Roots[`/os/ModeTemporary`] = ConstRoot{Const: os.ModeTemporary}
	Roots[`/os/ModeType`] = ConstRoot{Const: os.ModeType}
	Roots[`/os/O_APPEND`] = ConstRoot{Const: os.O_APPEND}
	Roots[`/os/O_CREATE`] = ConstRoot{Const: os.O_CREATE}
	Roots[`/os/O_EXCL`] = ConstRoot{Const: os.O_EXCL}
	Roots[`/os/O_RDONLY`] = ConstRoot{Const: os.O_RDONLY}
	Roots[`/os/O_RDWR`] = ConstRoot{Const: os.O_RDWR}
	Roots[`/os/O_SYNC`] = ConstRoot{Const: os.O_SYNC}
	Roots[`/os/O_TRUNC`] = ConstRoot{Const: os.O_TRUNC}
	Roots[`/os/O_WRONLY`] = ConstRoot{Const: os.O_WRONLY}
	Roots[`/os/PathListSeparator`] = ConstRoot{Const: os.PathListSeparator}
	Roots[`/os/PathSeparator`] = ConstRoot{Const: os.PathSeparator}
	Roots[`/os/SEEK_CUR`] = ConstRoot{Const: os.SEEK_CUR}
	Roots[`/os/SEEK_END`] = ConstRoot{Const: os.SEEK_END}
	Roots[`/os/SEEK_SET`] = ConstRoot{Const: os.SEEK_SET}
	Roots[`/path/filepath/ListSeparator`] = ConstRoot{Const: path_filepath.ListSeparator}
	Roots[`/path/filepath/Separator`] = ConstRoot{Const: path_filepath.Separator}
	Roots[`/reflect/Array`] = ConstRoot{Const: reflect.Array}
	Roots[`/reflect/Bool`] = ConstRoot{Const: reflect.Bool}
	Roots[`/reflect/BothDir`] = ConstRoot{Const: reflect.BothDir}
	Roots[`/reflect/Chan`] = ConstRoot{Const: reflect.Chan}
	Roots[`/reflect/Complex128`] = ConstRoot{Const: reflect.Complex128}
	Roots[`/reflect/Complex64`] = ConstRoot{Const: reflect.Complex64}
	Roots[`/reflect/Float32`] = ConstRoot{Const: reflect.Float32}
	Roots[`/reflect/Float64`] = ConstRoot{Const: reflect.Float64}
	Roots[`/reflect/Func`] = ConstRoot{Const: reflect.Func}
	Roots[`/reflect/Int`] = ConstRoot{Const: reflect.Int}
	Roots[`/reflect/Int16`] = ConstRoot{Const: reflect.Int16}
	Roots[`/reflect/Int32`] = ConstRoot{Const: reflect.Int32}
	Roots[`/reflect/Int64`] = ConstRoot{Const: reflect.Int64}
	Roots[`/reflect/Int8`] = ConstRoot{Const: reflect.Int8}
	Roots[`/reflect/Interface`] = ConstRoot{Const: reflect.Interface}
	Roots[`/reflect/Invalid`] = ConstRoot{Const: reflect.Invalid}
	Roots[`/reflect/Map`] = ConstRoot{Const: reflect.Map}
	Roots[`/reflect/Ptr`] = ConstRoot{Const: reflect.Ptr}
	Roots[`/reflect/RecvDir`] = ConstRoot{Const: reflect.RecvDir}
	Roots[`/reflect/SendDir`] = ConstRoot{Const: reflect.SendDir}
	Roots[`/reflect/Slice`] = ConstRoot{Const: reflect.Slice}
	Roots[`/reflect/String`] = ConstRoot{Const: reflect.String}
	Roots[`/reflect/Struct`] = ConstRoot{Const: reflect.Struct}
	Roots[`/reflect/Uint`] = ConstRoot{Const: reflect.Uint}
	Roots[`/reflect/Uint16`] = ConstRoot{Const: reflect.Uint16}
	Roots[`/reflect/Uint32`] = ConstRoot{Const: reflect.Uint32}
	Roots[`/reflect/Uint64`] = ConstRoot{Const: reflect.Uint64}
	Roots[`/reflect/Uint8`] = ConstRoot{Const: reflect.Uint8}
	Roots[`/reflect/Uintptr`] = ConstRoot{Const: reflect.Uintptr}
	Roots[`/reflect/UnsafePointer`] = ConstRoot{Const: reflect.UnsafePointer}
	Roots[`/regexp/syntax/ClassNL`] = ConstRoot{Const: regexp_syntax.ClassNL}
	Roots[`/regexp/syntax/DotNL`] = ConstRoot{Const: regexp_syntax.DotNL}
	Roots[`/regexp/syntax/EmptyBeginLine`] = ConstRoot{Const: regexp_syntax.EmptyBeginLine}
	Roots[`/regexp/syntax/EmptyBeginText`] = ConstRoot{Const: regexp_syntax.EmptyBeginText}
	Roots[`/regexp/syntax/EmptyEndLine`] = ConstRoot{Const: regexp_syntax.EmptyEndLine}
	Roots[`/regexp/syntax/EmptyEndText`] = ConstRoot{Const: regexp_syntax.EmptyEndText}
	Roots[`/regexp/syntax/EmptyNoWordBoundary`] = ConstRoot{Const: regexp_syntax.EmptyNoWordBoundary}
	Roots[`/regexp/syntax/EmptyWordBoundary`] = ConstRoot{Const: regexp_syntax.EmptyWordBoundary}
	Roots[`/regexp/syntax/ErrInternalError`] = ConstRoot{Const: regexp_syntax.ErrInternalError}
	Roots[`/regexp/syntax/ErrInvalidCharClass`] = ConstRoot{Const: regexp_syntax.ErrInvalidCharClass}
	Roots[`/regexp/syntax/ErrInvalidCharRange`] = ConstRoot{Const: regexp_syntax.ErrInvalidCharRange}
	Roots[`/regexp/syntax/ErrInvalidEscape`] = ConstRoot{Const: regexp_syntax.ErrInvalidEscape}
	Roots[`/regexp/syntax/ErrInvalidNamedCapture`] = ConstRoot{Const: regexp_syntax.ErrInvalidNamedCapture}
	Roots[`/regexp/syntax/ErrInvalidPerlOp`] = ConstRoot{Const: regexp_syntax.ErrInvalidPerlOp}
	Roots[`/regexp/syntax/ErrInvalidRepeatOp`] = ConstRoot{Const: regexp_syntax.ErrInvalidRepeatOp}
	Roots[`/regexp/syntax/ErrInvalidRepeatSize`] = ConstRoot{Const: regexp_syntax.ErrInvalidRepeatSize}
	Roots[`/regexp/syntax/ErrInvalidUTF8`] = ConstRoot{Const: regexp_syntax.ErrInvalidUTF8}
	Roots[`/regexp/syntax/ErrMissingBracket`] = ConstRoot{Const: regexp_syntax.ErrMissingBracket}
	Roots[`/regexp/syntax/ErrMissingParen`] = ConstRoot{Const: regexp_syntax.ErrMissingParen}
	Roots[`/regexp/syntax/ErrMissingRepeatArgument`] = ConstRoot{Const: regexp_syntax.ErrMissingRepeatArgument}
	Roots[`/regexp/syntax/ErrTrailingBackslash`] = ConstRoot{Const: regexp_syntax.ErrTrailingBackslash}
	Roots[`/regexp/syntax/FoldCase`] = ConstRoot{Const: regexp_syntax.FoldCase}
	Roots[`/regexp/syntax/InstAlt`] = ConstRoot{Const: regexp_syntax.InstAlt}
	Roots[`/regexp/syntax/InstAltMatch`] = ConstRoot{Const: regexp_syntax.InstAltMatch}
	Roots[`/regexp/syntax/InstCapture`] = ConstRoot{Const: regexp_syntax.InstCapture}
	Roots[`/regexp/syntax/InstEmptyWidth`] = ConstRoot{Const: regexp_syntax.InstEmptyWidth}
	Roots[`/regexp/syntax/InstFail`] = ConstRoot{Const: regexp_syntax.InstFail}
	Roots[`/regexp/syntax/InstMatch`] = ConstRoot{Const: regexp_syntax.InstMatch}
	Roots[`/regexp/syntax/InstNop`] = ConstRoot{Const: regexp_syntax.InstNop}
	Roots[`/regexp/syntax/InstRune`] = ConstRoot{Const: regexp_syntax.InstRune}
	Roots[`/regexp/syntax/InstRune1`] = ConstRoot{Const: regexp_syntax.InstRune1}
	Roots[`/regexp/syntax/InstRuneAny`] = ConstRoot{Const: regexp_syntax.InstRuneAny}
	Roots[`/regexp/syntax/InstRuneAnyNotNL`] = ConstRoot{Const: regexp_syntax.InstRuneAnyNotNL}
	Roots[`/regexp/syntax/Literal`] = ConstRoot{Const: regexp_syntax.Literal}
	Roots[`/regexp/syntax/MatchNL`] = ConstRoot{Const: regexp_syntax.MatchNL}
	Roots[`/regexp/syntax/NonGreedy`] = ConstRoot{Const: regexp_syntax.NonGreedy}
	Roots[`/regexp/syntax/OneLine`] = ConstRoot{Const: regexp_syntax.OneLine}
	Roots[`/regexp/syntax/OpAlternate`] = ConstRoot{Const: regexp_syntax.OpAlternate}
	Roots[`/regexp/syntax/OpAnyChar`] = ConstRoot{Const: regexp_syntax.OpAnyChar}
	Roots[`/regexp/syntax/OpAnyCharNotNL`] = ConstRoot{Const: regexp_syntax.OpAnyCharNotNL}
	Roots[`/regexp/syntax/OpBeginLine`] = ConstRoot{Const: regexp_syntax.OpBeginLine}
	Roots[`/regexp/syntax/OpBeginText`] = ConstRoot{Const: regexp_syntax.OpBeginText}
	Roots[`/regexp/syntax/OpCapture`] = ConstRoot{Const: regexp_syntax.OpCapture}
	Roots[`/regexp/syntax/OpCharClass`] = ConstRoot{Const: regexp_syntax.OpCharClass}
	Roots[`/regexp/syntax/OpConcat`] = ConstRoot{Const: regexp_syntax.OpConcat}
	Roots[`/regexp/syntax/OpEmptyMatch`] = ConstRoot{Const: regexp_syntax.OpEmptyMatch}
	Roots[`/regexp/syntax/OpEndLine`] = ConstRoot{Const: regexp_syntax.OpEndLine}
	Roots[`/regexp/syntax/OpEndText`] = ConstRoot{Const: regexp_syntax.OpEndText}
	Roots[`/regexp/syntax/OpLiteral`] = ConstRoot{Const: regexp_syntax.OpLiteral}
	Roots[`/regexp/syntax/OpNoMatch`] = ConstRoot{Const: regexp_syntax.OpNoMatch}
	Roots[`/regexp/syntax/OpNoWordBoundary`] = ConstRoot{Const: regexp_syntax.OpNoWordBoundary}
	Roots[`/regexp/syntax/OpPlus`] = ConstRoot{Const: regexp_syntax.OpPlus}
	Roots[`/regexp/syntax/OpQuest`] = ConstRoot{Const: regexp_syntax.OpQuest}
	Roots[`/regexp/syntax/OpRepeat`] = ConstRoot{Const: regexp_syntax.OpRepeat}
	Roots[`/regexp/syntax/OpStar`] = ConstRoot{Const: regexp_syntax.OpStar}
	Roots[`/regexp/syntax/OpWordBoundary`] = ConstRoot{Const: regexp_syntax.OpWordBoundary}
	Roots[`/regexp/syntax/POSIX`] = ConstRoot{Const: regexp_syntax.POSIX}
	Roots[`/regexp/syntax/Perl`] = ConstRoot{Const: regexp_syntax.Perl}
	Roots[`/regexp/syntax/PerlX`] = ConstRoot{Const: regexp_syntax.PerlX}
	Roots[`/regexp/syntax/Simple`] = ConstRoot{Const: regexp_syntax.Simple}
	Roots[`/regexp/syntax/UnicodeGroups`] = ConstRoot{Const: regexp_syntax.UnicodeGroups}
	Roots[`/regexp/syntax/WasDollar`] = ConstRoot{Const: regexp_syntax.WasDollar}
	Roots[`/runtime/Compiler`] = ConstRoot{Const: runtime.Compiler}
	Roots[`/runtime/GOARCH`] = ConstRoot{Const: runtime.GOARCH}
	Roots[`/runtime/GOOS`] = ConstRoot{Const: runtime.GOOS}
	Roots[`/strconv/IntSize`] = ConstRoot{Const: int64(strconv.IntSize)}
	Roots[`/syscall/AF_INET`] = ConstRoot{Const: int64(syscall.AF_INET)}
	Roots[`/syscall/AF_INET6`] = ConstRoot{Const: int64(syscall.AF_INET6)}
	Roots[`/syscall/AF_UNIX`] = ConstRoot{Const: int64(syscall.AF_UNIX)}
	Roots[`/syscall/AF_UNSPEC`] = ConstRoot{Const: int64(syscall.AF_UNSPEC)}
	Roots[`/syscall/E2BIG`] = ConstRoot{Const: syscall.E2BIG}
	Roots[`/syscall/EACCES`] = ConstRoot{Const: syscall.EACCES}
	Roots[`/syscall/EADDRINUSE`] = ConstRoot{Const: syscall.EADDRINUSE}
	Roots[`/syscall/EADDRNOTAVAIL`] = ConstRoot{Const: syscall.EADDRNOTAVAIL}
	Roots[`/syscall/EAFNOSUPPORT`] = ConstRoot{Const: syscall.EAFNOSUPPORT}
	Roots[`/syscall/EAGAIN`] = ConstRoot{Const: syscall.EAGAIN}
	Roots[`/syscall/EALREADY`] = ConstRoot{Const: syscall.EALREADY}
	Roots[`/syscall/EBADF`] = ConstRoot{Const: syscall.EBADF}
	Roots[`/syscall/EBADMSG`] = ConstRoot{Const: syscall.EBADMSG}
	Roots[`/syscall/EBUSY`] = ConstRoot{Const: syscall.EBUSY}
	Roots[`/syscall/ECANCELED`] = ConstRoot{Const: syscall.ECANCELED}
	Roots[`/syscall/ECHILD`] = ConstRoot{Const: syscall.ECHILD}
	Roots[`/syscall/ECONNABORTED`] = ConstRoot{Const: syscall.ECONNABORTED}
	Roots[`/syscall/ECONNREFUSED`] = ConstRoot{Const: syscall.ECONNREFUSED}
	Roots[`/syscall/ECONNRESET`] = ConstRoot{Const: syscall.ECONNRESET}
	Roots[`/syscall/EDEADLK`] = ConstRoot{Const: syscall.EDEADLK}
	Roots[`/syscall/EDESTADDRREQ`] = ConstRoot{Const: syscall.EDESTADDRREQ}
	Roots[`/syscall/EDOM`] = ConstRoot{Const: syscall.EDOM}
	Roots[`/syscall/EDQUOT`] = ConstRoot{Const: syscall.EDQUOT}
	Roots[`/syscall/EEXIST`] = ConstRoot{Const: syscall.EEXIST}
	Roots[`/syscall/EFAULT`] = ConstRoot{Const: syscall.EFAULT}
	Roots[`/syscall/EFBIG`] = ConstRoot{Const: syscall.EFBIG}
	Roots[`/syscall/EHOSTDOWN`] = ConstRoot{Const: syscall.EHOSTDOWN}
	Roots[`/syscall/EHOSTUNREACH`] = ConstRoot{Const: syscall.EHOSTUNREACH}
	Roots[`/syscall/EIDRM`] = ConstRoot{Const: syscall.EIDRM}
	Roots[`/syscall/EILSEQ`] = ConstRoot{Const: syscall.EILSEQ}
	Roots[`/syscall/EINPROGRESS`] = ConstRoot{Const: syscall.EINPROGRESS}
	Roots[`/syscall/EINTR`] = ConstRoot{Const: syscall.EINTR}
	Roots[`/syscall/EINVAL`] = ConstRoot{Const: syscall.EINVAL}
	Roots[`/syscall/EIO`] = ConstRoot{Const: syscall.EIO}
	Roots[`/syscall/EISCONN`] = ConstRoot{Const: syscall.EISCONN}
	Roots[`/syscall/EISDIR`] = ConstRoot{Const: syscall.EISDIR}
	Roots[`/syscall/ELOOP`] = ConstRoot{Const: syscall.ELOOP}
	Roots[`/syscall/EMFILE`] = ConstRoot{Const: syscall.EMFILE}
	Roots[`/syscall/EMLINK`] = ConstRoot{Const: syscall.EMLINK}
	Roots[`/syscall/EMSGSIZE`] = ConstRoot{Const: syscall.EMSGSIZE}
	Roots[`/syscall/EMULTIHOP`] = ConstRoot{Const: syscall.EMULTIHOP}
	Roots[`/syscall/ENAMETOOLONG`] = ConstRoot{Const: syscall.ENAMETOOLONG}
	Roots[`/syscall/ENETDOWN`] = ConstRoot{Const: syscall.ENETDOWN}
	Roots[`/syscall/ENETRESET`] = ConstRoot{Const: syscall.ENETRESET}
	Roots[`/syscall/ENETUNREACH`] = ConstRoot{Const: syscall.ENETUNREACH}
	Roots[`/syscall/ENFILE`] = ConstRoot{Const: syscall.ENFILE}
	Roots[`/syscall/ENOBUFS`] = ConstRoot{Const: syscall.ENOBUFS}
	Roots[`/syscall/ENODEV`] = ConstRoot{Const: syscall.ENODEV}
	Roots[`/syscall/ENOENT`] = ConstRoot{Const: syscall.ENOENT}
	Roots[`/syscall/ENOEXEC`] = ConstRoot{Const: syscall.ENOEXEC}
	Roots[`/syscall/ENOLCK`] = ConstRoot{Const: syscall.ENOLCK}
	Roots[`/syscall/ENOLINK`] = ConstRoot{Const: syscall.ENOLINK}
	Roots[`/syscall/ENOMEM`] = ConstRoot{Const: syscall.ENOMEM}
	Roots[`/syscall/ENOMSG`] = ConstRoot{Const: syscall.ENOMSG}
	Roots[`/syscall/ENOPROTOOPT`] = ConstRoot{Const: syscall.ENOPROTOOPT}
	Roots[`/syscall/ENOSPC`] = ConstRoot{Const: syscall.ENOSPC}
	Roots[`/syscall/ENOSYS`] = ConstRoot{Const: syscall.ENOSYS}
	Roots[`/syscall/ENOTBLK`] = ConstRoot{Const: syscall.ENOTBLK}
	Roots[`/syscall/ENOTCONN`] = ConstRoot{Const: syscall.ENOTCONN}
	Roots[`/syscall/ENOTDIR`] = ConstRoot{Const: syscall.ENOTDIR}
	Roots[`/syscall/ENOTEMPTY`] = ConstRoot{Const: syscall.ENOTEMPTY}
	Roots[`/syscall/ENOTSOCK`] = ConstRoot{Const: syscall.ENOTSOCK}
	Roots[`/syscall/ENOTSUP`] = ConstRoot{Const: syscall.ENOTSUP}
	Roots[`/syscall/ENOTTY`] = ConstRoot{Const: syscall.ENOTTY}
	Roots[`/syscall/ENXIO`] = ConstRoot{Const: syscall.ENXIO}
	Roots[`/syscall/EOPNOTSUPP`] = ConstRoot{Const: syscall.EOPNOTSUPP}
	Roots[`/syscall/EOVERFLOW`] = ConstRoot{Const: syscall.EOVERFLOW}
	Roots[`/syscall/EPERM`] = ConstRoot{Const: syscall.EPERM}
	Roots[`/syscall/EPFNOSUPPORT`] = ConstRoot{Const: syscall.EPFNOSUPPORT}
	Roots[`/syscall/EPIPE`] = ConstRoot{Const: syscall.EPIPE}
	Roots[`/syscall/EPROTO`] = ConstRoot{Const: syscall.EPROTO}
	Roots[`/syscall/EPROTONOSUPPORT`] = ConstRoot{Const: syscall.EPROTONOSUPPORT}
	Roots[`/syscall/EPROTOTYPE`] = ConstRoot{Const: syscall.EPROTOTYPE}
	Roots[`/syscall/ERANGE`] = ConstRoot{Const: syscall.ERANGE}
	Roots[`/syscall/EREMOTE`] = ConstRoot{Const: syscall.EREMOTE}
	Roots[`/syscall/EROFS`] = ConstRoot{Const: syscall.EROFS}
	Roots[`/syscall/ESHUTDOWN`] = ConstRoot{Const: syscall.ESHUTDOWN}
	Roots[`/syscall/ESOCKTNOSUPPORT`] = ConstRoot{Const: syscall.ESOCKTNOSUPPORT}
	Roots[`/syscall/ESPIPE`] = ConstRoot{Const: syscall.ESPIPE}
	Roots[`/syscall/ESRCH`] = ConstRoot{Const: syscall.ESRCH}
	Roots[`/syscall/ESTALE`] = ConstRoot{Const: syscall.ESTALE}
	Roots[`/syscall/ETIMEDOUT`] = ConstRoot{Const: syscall.ETIMEDOUT}
	Roots[`/syscall/ETOOMANYREFS`] = ConstRoot{Const: syscall.ETOOMANYREFS}
	Roots[`/syscall/ETXTBSY`] = ConstRoot{Const: syscall.ETXTBSY}
	Roots[`/syscall/EUSERS`] = ConstRoot{Const: syscall.EUSERS}
	Roots[`/syscall/EWOULDBLOCK`] = ConstRoot{Const: syscall.EWOULDBLOCK}
	Roots[`/syscall/EXDEV`] = ConstRoot{Const: syscall.EXDEV}
	Roots[`/syscall/IFF_BROADCAST`] = ConstRoot{Const: int64(syscall.IFF_BROADCAST)}
	Roots[`/syscall/IFF_LOOPBACK`] = ConstRoot{Const: int64(syscall.IFF_LOOPBACK)}
	Roots[`/syscall/IFF_MULTICAST`] = ConstRoot{Const: int64(syscall.IFF_MULTICAST)}
	Roots[`/syscall/IFF_UP`] = ConstRoot{Const: int64(syscall.IFF_UP)}
	Roots[`/syscall/IPPROTO_IP`] = ConstRoot{Const: int64(syscall.IPPROTO_IP)}
	Roots[`/syscall/IPPROTO_IPV6`] = ConstRoot{Const: int64(syscall.IPPROTO_IPV6)}
	Roots[`/syscall/IPPROTO_TCP`] = ConstRoot{Const: int64(syscall.IPPROTO_TCP)}
	Roots[`/syscall/IPPROTO_UDP`] = ConstRoot{Const: int64(syscall.IPPROTO_UDP)}
	Roots[`/syscall/IPV6_JOIN_GROUP`] = ConstRoot{Const: int64(syscall.IPV6_JOIN_GROUP)}
	Roots[`/syscall/IPV6_LEAVE_GROUP`] = ConstRoot{Const: int64(syscall.IPV6_LEAVE_GROUP)}
	Roots[`/syscall/IPV6_MULTICAST_HOPS`] = ConstRoot{Const: int64(syscall.IPV6_MULTICAST_HOPS)}
	Roots[`/syscall/IPV6_MULTICAST_IF`] = ConstRoot{Const: int64(syscall.IPV6_MULTICAST_IF)}
	Roots[`/syscall/IPV6_MULTICAST_LOOP`] = ConstRoot{Const: int64(syscall.IPV6_MULTICAST_LOOP)}
	Roots[`/syscall/IPV6_UNICAST_HOPS`] = ConstRoot{Const: int64(syscall.IPV6_UNICAST_HOPS)}
	Roots[`/syscall/IPV6_V6ONLY`] = ConstRoot{Const: int64(syscall.IPV6_V6ONLY)}
	Roots[`/syscall/IP_ADD_MEMBERSHIP`] = ConstRoot{Const: int64(syscall.IP_ADD_MEMBERSHIP)}
	Roots[`/syscall/IP_DROP_MEMBERSHIP`] = ConstRoot{Const: int64(syscall.IP_DROP_MEMBERSHIP)}
	Roots[`/syscall/IP_MULTICAST_IF`] = ConstRoot{Const: int64(syscall.IP_MULTICAST_IF)}
	Roots[`/syscall/IP_MULTICAST_LOOP`] = ConstRoot{Const: int64(syscall.IP_MULTICAST_LOOP)}
	Roots[`/syscall/IP_MULTICAST_TTL`] = ConstRoot{Const: int64(syscall.IP_MULTICAST_TTL)}
	Roots[`/syscall/IP_TOS`] = ConstRoot{Const: int64(syscall.IP_TOS)}
	Roots[`/syscall/IP_TTL`] = ConstRoot{Const: int64(syscall.IP_TTL)}
	Roots[`/syscall/ImplementsGetwd`] = ConstRoot{Const: syscall.ImplementsGetwd}
	Roots[`/syscall/O_APPEND`] = ConstRoot{Const: int64(syscall.O_APPEND)}
	Roots[`/syscall/O_ASYNC`] = ConstRoot{Const: int64(syscall.O_ASYNC)}
	Roots[`/syscall/O_CLOEXEC`] = ConstRoot{Const: int64(syscall.O_CLOEXEC)}
	Roots[`/syscall/O_CREAT`] = ConstRoot{Const: int64(syscall.O_CREAT)}
	Roots[`/syscall/O_EXCL`] = ConstRoot{Const: int64(syscall.O_EXCL)}
	Roots[`/syscall/O_NOCTTY`] = ConstRoot{Const: int64(syscall.O_NOCTTY)}
	Roots[`/syscall/O_NONBLOCK`] = ConstRoot{Const: int64(syscall.O_NONBLOCK)}
	Roots[`/syscall/O_RDONLY`] = ConstRoot{Const: int64(syscall.O_RDONLY)}
	Roots[`/syscall/O_RDWR`] = ConstRoot{Const: int64(syscall.O_RDWR)}
	Roots[`/syscall/O_SYNC`] = ConstRoot{Const: int64(syscall.O_SYNC)}
	Roots[`/syscall/O_TRUNC`] = ConstRoot{Const: int64(syscall.O_TRUNC)}
	Roots[`/syscall/O_WRONLY`] = ConstRoot{Const: int64(syscall.O_WRONLY)}
	Roots[`/syscall/SHUT_RD`] = ConstRoot{Const: int64(syscall.SHUT_RD)}
	Roots[`/syscall/SHUT_RDWR`] = ConstRoot{Const: int64(syscall.SHUT_RDWR)}
	Roots[`/syscall/SHUT_WR`] = ConstRoot{Const: int64(syscall.SHUT_WR)}
	Roots[`/syscall/SIGABRT`] = ConstRoot{Const: syscall.SIGABRT}
	Roots[`/syscall/SIGALRM`] = ConstRoot{Const: syscall.SIGALRM}
	Roots[`/syscall/SIGBUS`] = ConstRoot{Const: syscall.SIGBUS}
	Roots[`/syscall/SIGFPE`] = ConstRoot{Const: syscall.SIGFPE}
	Roots[`/syscall/SIGHUP`] = ConstRoot{Const: syscall.SIGHUP}
	Roots[`/syscall/SIGILL`] = ConstRoot{Const: syscall.SIGILL}
	Roots[`/syscall/SIGINT`] = ConstRoot{Const: syscall.SIGINT}
	Roots[`/syscall/SIGKILL`] = ConstRoot{Const: syscall.SIGKILL}
	Roots[`/syscall/SIGPIPE`] = ConstRoot{Const: syscall.SIGPIPE}
	Roots[`/syscall/SIGQUIT`] = ConstRoot{Const: syscall.SIGQUIT}
	Roots[`/syscall/SIGSEGV`] = ConstRoot{Const: syscall.SIGSEGV}
	Roots[`/syscall/SIGTERM`] = ConstRoot{Const: syscall.SIGTERM}
	Roots[`/syscall/SIGTRAP`] = ConstRoot{Const: syscall.SIGTRAP}
	Roots[`/syscall/SOCK_DGRAM`] = ConstRoot{Const: int64(syscall.SOCK_DGRAM)}
	Roots[`/syscall/SOCK_RAW`] = ConstRoot{Const: int64(syscall.SOCK_RAW)}
	Roots[`/syscall/SOCK_SEQPACKET`] = ConstRoot{Const: int64(syscall.SOCK_SEQPACKET)}
	Roots[`/syscall/SOCK_STREAM`] = ConstRoot{Const: int64(syscall.SOCK_STREAM)}
	Roots[`/syscall/SOL_SOCKET`] = ConstRoot{Const: int64(syscall.SOL_SOCKET)}
	Roots[`/syscall/SOMAXCONN`] = ConstRoot{Const: int64(syscall.SOMAXCONN)}
	Roots[`/syscall/SO_BROADCAST`] = ConstRoot{Const: int64(syscall.SO_BROADCAST)}
	Roots[`/syscall/SO_DONTROUTE`] = ConstRoot{Const: int64(syscall.SO_DONTROUTE)}
	Roots[`/syscall/SO_KEEPALIVE`] = ConstRoot{Const: int64(syscall.SO_KEEPALIVE)}
	Roots[`/syscall/SO_LINGER`] = ConstRoot{Const: int64(syscall.SO_LINGER)}
	Roots[`/syscall/SO_RCVBUF`] = ConstRoot{Const: int64(syscall.SO_RCVBUF)}
	Roots[`/syscall/SO_REUSEADDR`] = ConstRoot{Const: int64(syscall.SO_REUSEADDR)}
	Roots[`/syscall/SO_SNDBUF`] = ConstRoot{Const: int64(syscall.SO_SNDBUF)}
	Roots[`/syscall/S_IFBLK`] = ConstRoot{Const: int64(syscall.S_IFBLK)}
	Roots[`/syscall/S_IFCHR`] = ConstRoot{Const: int64(syscall.S_IFCHR)}
	Roots[`/syscall/S_IFDIR`] = ConstRoot{Const: int64(syscall.S_IFDIR)}
	Roots[`/syscall/S_IFIFO`] = ConstRoot{Const: int64(syscall.S_IFIFO)}
	Roots[`/syscall/S_IFLNK`] = ConstRoot{Const: int64(syscall.S_IFLNK)}
	Roots[`/syscall/S_IFMT`] = ConstRoot{Const: int64(syscall.S_IFMT)}
	Roots[`/syscall/S_IFREG`] = ConstRoot{Const: int64(syscall.S_IFREG)}
	Roots[`/syscall/S_IFSOCK`] = ConstRoot{Const: int64(syscall.S_IFSOCK)}
	Roots[`/syscall/S_IRUSR`] = ConstRoot{Const: int64(syscall.S_IRUSR)}
	Roots[`/syscall/S_ISGID`] = ConstRoot{Const: int64(syscall.S_ISGID)}
	Roots[`/syscall/S_ISUID`] = ConstRoot{Const: int64(syscall.S_ISUID)}
	Roots[`/syscall/S_ISVTX`] = ConstRoot{Const: int64(syscall.S_ISVTX)}
	Roots[`/syscall/S_IWUSR`] = ConstRoot{Const: int64(syscall.S_IWUSR)}
	Roots[`/syscall/S_IXUSR`] = ConstRoot{Const: int64(syscall.S_IXUSR)}
	Roots[`/syscall/TCP_NODELAY`] = ConstRoot{Const: int64(syscall.TCP_NODELAY)}
	Roots[`/text/scanner/Char`] = ConstRoot{Const: int64(text_scanner.Char)}
	Roots[`/text/scanner/Comment`] = ConstRoot{Const: int64(text_scanner.Comment)}
	Roots[`/text/scanner/EOF`] = ConstRoot{Const: int64(text_scanner.EOF)}
	Roots[`/text/scanner/Float`] = ConstRoot{Const: int64(text_scanner.Float)}
	Roots[`/text/scanner/GoTokens`] = ConstRoot{Const: int64(text_scanner.GoTokens)}
	Roots[`/text/scanner/GoWhitespace`] = ConstRoot{Const: int64(text_scanner.GoWhitespace)}
	Roots[`/text/scanner/Ident`] = ConstRoot{Const: int64(text_scanner.Ident)}
	Roots[`/text/scanner/Int`] = ConstRoot{Const: int64(text_scanner.Int)}
	Roots[`/text/scanner/RawString`] = ConstRoot{Const: int64(text_scanner.RawString)}
	Roots[`/text/scanner/ScanChars`] = ConstRoot{Const: int64(text_scanner.ScanChars)}
	Roots[`/text/scanner/ScanComments`] = ConstRoot{Const: int64(text_scanner.ScanComments)}
	Roots[`/text/scanner/ScanFloats`] = ConstRoot{Const: int64(text_scanner.ScanFloats)}
	Roots[`/text/scanner/ScanIdents`] = ConstRoot{Const: int64(text_scanner.ScanIdents)}
	Roots[`/text/scanner/ScanInts`] = ConstRoot{Const: int64(text_scanner.ScanInts)}
	Roots[`/text/scanner/ScanRawStrings`] = ConstRoot{Const: int64(text_scanner.ScanRawStrings)}
	Roots[`/text/scanner/ScanStrings`] = ConstRoot{Const: int64(text_scanner.ScanStrings)}
	Roots[`/text/scanner/SkipComments`] = ConstRoot{Const: int64(text_scanner.SkipComments)}
	Roots[`/text/scanner/String`] = ConstRoot{Const: int64(text_scanner.String)}
	Roots[`/text/tabwriter/AlignRight`] = ConstRoot{Const: text_tabwriter.AlignRight}
	Roots[`/text/tabwriter/Debug`] = ConstRoot{Const: text_tabwriter.Debug}
	Roots[`/text/tabwriter/DiscardEmptyColumns`] = ConstRoot{Const: text_tabwriter.DiscardEmptyColumns}
	Roots[`/text/tabwriter/Escape`] = ConstRoot{Const: text_tabwriter.Escape}
	Roots[`/text/tabwriter/FilterHTML`] = ConstRoot{Const: text_tabwriter.FilterHTML}
	Roots[`/text/tabwriter/StripEscape`] = ConstRoot{Const: text_tabwriter.StripEscape}
	Roots[`/text/tabwriter/TabIndent`] = ConstRoot{Const: text_tabwriter.TabIndent}
	Roots[`/text/template/parse/NodeAction`] = ConstRoot{Const: text_template_parse.NodeAction}
	Roots[`/text/template/parse/NodeBool`] = ConstRoot{Const: text_template_parse.NodeBool}
	Roots[`/text/template/parse/NodeCommand`] = ConstRoot{Const: text_template_parse.NodeCommand}
	Roots[`/text/template/parse/NodeDot`] = ConstRoot{Const: text_template_parse.NodeDot}
	Roots[`/text/template/parse/NodeField`] = ConstRoot{Const: text_template_parse.NodeField}
	Roots[`/text/template/parse/NodeIdentifier`] = ConstRoot{Const: text_template_parse.NodeIdentifier}
	Roots[`/text/template/parse/NodeIf`] = ConstRoot{Const: text_template_parse.NodeIf}
	Roots[`/text/template/parse/NodeList`] = ConstRoot{Const: text_template_parse.NodeList}
	Roots[`/text/template/parse/NodeNumber`] = ConstRoot{Const: text_template_parse.NodeNumber}
	Roots[`/text/template/parse/NodePipe`] = ConstRoot{Const: text_template_parse.NodePipe}
	Roots[`/text/template/parse/NodeRange`] = ConstRoot{Const: text_template_parse.NodeRange}
	Roots[`/text/template/parse/NodeString`] = ConstRoot{Const: text_template_parse.NodeString}
	Roots[`/text/template/parse/NodeTemplate`] = ConstRoot{Const: text_template_parse.NodeTemplate}
	Roots[`/text/template/parse/NodeText`] = ConstRoot{Const: text_template_parse.NodeText}
	Roots[`/text/template/parse/NodeVariable`] = ConstRoot{Const: text_template_parse.NodeVariable}
	Roots[`/text/template/parse/NodeWith`] = ConstRoot{Const: text_template_parse.NodeWith}
	Roots[`/time/ANSIC`] = ConstRoot{Const: time.ANSIC}
	Roots[`/time/April`] = ConstRoot{Const: time.April}
	Roots[`/time/August`] = ConstRoot{Const: time.August}
	Roots[`/time/December`] = ConstRoot{Const: time.December}
	Roots[`/time/February`] = ConstRoot{Const: time.February}
	Roots[`/time/Friday`] = ConstRoot{Const: time.Friday}
	Roots[`/time/Hour`] = ConstRoot{Const: time.Hour}
	Roots[`/time/January`] = ConstRoot{Const: time.January}
	Roots[`/time/July`] = ConstRoot{Const: time.July}
	Roots[`/time/June`] = ConstRoot{Const: time.June}
	Roots[`/time/Kitchen`] = ConstRoot{Const: time.Kitchen}
	Roots[`/time/March`] = ConstRoot{Const: time.March}
	Roots[`/time/May`] = ConstRoot{Const: time.May}
	Roots[`/time/Microsecond`] = ConstRoot{Const: time.Microsecond}
	Roots[`/time/Millisecond`] = ConstRoot{Const: time.Millisecond}
	Roots[`/time/Minute`] = ConstRoot{Const: time.Minute}
	Roots[`/time/Monday`] = ConstRoot{Const: time.Monday}
	Roots[`/time/Nanosecond`] = ConstRoot{Const: time.Nanosecond}
	Roots[`/time/November`] = ConstRoot{Const: time.November}
	Roots[`/time/October`] = ConstRoot{Const: time.October}
	Roots[`/time/RFC1123`] = ConstRoot{Const: time.RFC1123}
	Roots[`/time/RFC1123Z`] = ConstRoot{Const: time.RFC1123Z}
	Roots[`/time/RFC3339`] = ConstRoot{Const: time.RFC3339}
	Roots[`/time/RFC3339Nano`] = ConstRoot{Const: time.RFC3339Nano}
	Roots[`/time/RFC822`] = ConstRoot{Const: time.RFC822}
	Roots[`/time/RFC822Z`] = ConstRoot{Const: time.RFC822Z}
	Roots[`/time/RFC850`] = ConstRoot{Const: time.RFC850}
	Roots[`/time/RubyDate`] = ConstRoot{Const: time.RubyDate}
	Roots[`/time/Saturday`] = ConstRoot{Const: time.Saturday}
	Roots[`/time/Second`] = ConstRoot{Const: time.Second}
	Roots[`/time/September`] = ConstRoot{Const: time.September}
	Roots[`/time/Stamp`] = ConstRoot{Const: time.Stamp}
	Roots[`/time/StampMicro`] = ConstRoot{Const: time.StampMicro}
	Roots[`/time/StampMilli`] = ConstRoot{Const: time.StampMilli}
	Roots[`/time/StampNano`] = ConstRoot{Const: time.StampNano}
	Roots[`/time/Sunday`] = ConstRoot{Const: time.Sunday}
	Roots[`/time/Thursday`] = ConstRoot{Const: time.Thursday}
	Roots[`/time/Tuesday`] = ConstRoot{Const: time.Tuesday}
	Roots[`/time/UnixDate`] = ConstRoot{Const: time.UnixDate}
	Roots[`/time/Wednesday`] = ConstRoot{Const: time.Wednesday}
	Roots[`/unicode/LowerCase`] = ConstRoot{Const: int64(unicode.LowerCase)}
	Roots[`/unicode/MaxASCII`] = ConstRoot{Const: unicode.MaxASCII}
	Roots[`/unicode/MaxCase`] = ConstRoot{Const: int64(unicode.MaxCase)}
	Roots[`/unicode/MaxLatin1`] = ConstRoot{Const: unicode.MaxLatin1}
	Roots[`/unicode/MaxRune`] = ConstRoot{Const: unicode.MaxRune}
	Roots[`/unicode/ReplacementChar`] = ConstRoot{Const: unicode.ReplacementChar}
	Roots[`/unicode/TitleCase`] = ConstRoot{Const: int64(unicode.TitleCase)}
	Roots[`/unicode/UpperCase`] = ConstRoot{Const: int64(unicode.UpperCase)}
	Roots[`/unicode/UpperLower`] = ConstRoot{Const: int64(unicode.UpperLower)}
	Roots[`/unicode/Version`] = ConstRoot{Const: unicode.Version}
	Roots[`/unicode/utf8/MaxRune`] = ConstRoot{Const: unicode_utf8.MaxRune}
	Roots[`/unicode/utf8/RuneError`] = ConstRoot{Const: unicode_utf8.RuneError}
	Roots[`/unicode/utf8/RuneSelf`] = ConstRoot{Const: int64(unicode_utf8.RuneSelf)}
	Roots[`/unicode/utf8/UTFMax`] = ConstRoot{Const: int64(unicode_utf8.UTFMax)}
}
