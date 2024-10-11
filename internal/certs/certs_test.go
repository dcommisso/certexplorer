package certs

import (
	"fmt"
	"testing"
)

func getSampleCert(name string) []byte {
	sampleCerts := map[string][]byte{
		"single certificate": []byte(`-----BEGIN CERTIFICATE-----
MIIH0zCCBbugAwIBAgIIXsO3pkN/pOAwDQYJKoZIhvcNAQEFBQAwQjESMBAGA1UE
AwwJQUNDVlJBSVoxMRAwDgYDVQQLDAdQS0lBQ0NWMQ0wCwYDVQQKDARBQ0NWMQsw
CQYDVQQGEwJFUzAeFw0xMTA1MDUwOTM3MzdaFw0zMDEyMzEwOTM3MzdaMEIxEjAQ
BgNVBAMMCUFDQ1ZSQUlaMTEQMA4GA1UECwwHUEtJQUNDVjENMAsGA1UECgwEQUND
VjELMAkGA1UEBhMCRVMwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQCb
qau/YUqXry+XZpp0X9DZlv3P4uRm7x8fRzPCRKPfmt4ftVTdFXxpNRFvu8gMjmoY
HtiP2Ra8EEg2XPBjs5BaXCQ316PWywlxufEBcoSwfdtNgM3802/J+Nq2DoLSRYWo
G2ioPej0RGy9ocLLA76MPhMAhN9KSMDjIgro6TenGEyxCQ0jVn8ETdkXhBilyNpA
lHPrzg5XPAOBOp0KoVdDaaxXbXmQeOW1tDvYvEyNKKGno6e6Ak4l0Squ7a4DIrhr
IA8wKFSVf+DuzgpmndFALW4ir50awQUZ0m/A8p/4e7MCQvtQqR0tkw8jq8bBD5L/
0KIV9VMJcRz/RROE5iZe+OCIHAr8Fraocwa48GOEAqDGWuzndN9wrqODJerWx5eH
k6fGioozl2A3ED6XPm4pFdahD9GILBKfb6qkxkLrQaLjlUPTAYVtjrs78yM2x/47
4KElB0iryYl0/wiPgL/AlmXz7uxLaL2diMMxs0Dx6M/2OLuc5NF/1OVYm3z61PMO
m3WR5LpSLhl+0fXNWhn8ugb2+1KoS5kE3fj5tItQo05iifCHJPqDQsGH+tUtKSpa
cXpkatcnYGMN285J9Y0fkIkyF/hzQ7jSWpOGYdbhdQrqeWZ2iE9x6wQl1gpaepPl
uUsXQA+xtrn13k/c4LOsOxFwYIRKQ26ZIMApcQrAZQIDAQABo4ICyzCCAscwfQYI
KwYBBQUHAQEEcTBvMEwGCCsGAQUFBzAChkBodHRwOi8vd3d3LmFjY3YuZXMvZmls
ZWFkbWluL0FyY2hpdm9zL2NlcnRpZmljYWRvcy9yYWl6YWNjdjEuY3J0MB8GCCsG
AQUFBzABhhNodHRwOi8vb2NzcC5hY2N2LmVzMB0GA1UdDgQWBBTSh7Tj3zcnk1X2
VuqB5TbMjB4/vTAPBgNVHRMBAf8EBTADAQH/MB8GA1UdIwQYMBaAFNKHtOPfNyeT
VfZW6oHlNsyMHj+9MIIBcwYDVR0gBIIBajCCAWYwggFiBgRVHSAAMIIBWDCCASIG
CCsGAQUFBwICMIIBFB6CARAAQQB1AHQAbwByAGkAZABhAGQAIABkAGUAIABDAGUA
cgB0AGkAZgBpAGMAYQBjAGkA8wBuACAAUgBhAO0AegAgAGQAZQAgAGwAYQAgAEEA
QwBDAFYAIAAoAEEAZwBlAG4AYwBpAGEAIABkAGUAIABUAGUAYwBuAG8AbABvAGcA
7QBhACAAeQAgAEMAZQByAHQAaQBmAGkAYwBhAGMAaQDzAG4AIABFAGwAZQBjAHQA
cgDzAG4AaQBjAGEALAAgAEMASQBGACAAUQA0ADYAMAAxADEANQA2AEUAKQAuACAA
QwBQAFMAIABlAG4AIABoAHQAdABwADoALwAvAHcAdwB3AC4AYQBjAGMAdgAuAGUA
czAwBggrBgEFBQcCARYkaHR0cDovL3d3dy5hY2N2LmVzL2xlZ2lzbGFjaW9uX2Mu
aHRtMFUGA1UdHwROMEwwSqBIoEaGRGh0dHA6Ly93d3cuYWNjdi5lcy9maWxlYWRt
aW4vQXJjaGl2b3MvY2VydGlmaWNhZG9zL3JhaXphY2N2MV9kZXIuY3JsMA4GA1Ud
DwEB/wQEAwIBBjAXBgNVHREEEDAOgQxhY2N2QGFjY3YuZXMwDQYJKoZIhvcNAQEF
BQADggIBAJcxAp/n/UNnSEQU5CmH7UwoZtCPNdpNYbdKl02125DgBS4OxnnQ8pdp
D70ER9m+27Up2pvZrqmZ1dM8MJP1jaGo/AaNRPTKFpV8M9xii6g3+CfYCS0b78gU
JyCpZET/LtZ1qmxNYEAZSUNUY9rizLpm5U9EelvZaoErQNV/+QEnWCzI7UiRfD+m
AM/EKXMRNt6GGT6d7hmKG9Ww7Y49nCrADdg9ZuM8Db3VlFzi4qc1GwQA9j9ajepD
vV+JHanBsMyZ4k0ACtrJJ1vnE5Bc5PUzolVt3OAJTS+xJlsndQAJxGJ3KQhfnlms
tn6tn1QwIgPBHnFk/vk4CpYY3QIUrCPLBhwepH2NDd4nQeit2hW3sCPdK6jT2iWH
7ehVRE2I9DZ+hJp4rPcOVkkO1jMl1oRQQmwgEh0q1b688nCBpHBgvgW1m54ERL5h
I6zppSSMEYCUWqKiuUnSwdzRp+0xESyeGabu4VXhwOrPDYTkF7eifKXeVSUG7szA
h1xA2syVP1XgNce4hL60Xc16gwFy7ofmXx2utYXGJt/mwZrpHgJHnyqobalbz+xF
d3+YJ5oyXSrjhO7FmGYvliAd3djDJ9ew+f7Zfc3Qn48LFFhRny+Lwzgt3uiP1o2H
pPVWQxaZLPSkVrQ0uGE3ycJYgBugl6H8WY3pEfbRD0tVNEYqi4Y7
-----END CERTIFICATE-----`),

		"multiple certificates with comments": []byte(`# AC RAIZ FNMT-RCM
-----BEGIN CERTIFICATE-----
MIIFgzCCA2ugAwIBAgIPXZONMGc2yAYdGsdUhGkHMA0GCSqGSIb3DQEBCwUAMDsx
CzAJBgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJ
WiBGTk1ULVJDTTAeFw0wODEwMjkxNTU5NTZaFw0zMDAxMDEwMDAwMDBaMDsxCzAJ
BgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJWiBG
Tk1ULVJDTTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALpxgHpMhm5/
yBNtwMZ9HACXjywMI7sQmkCpGreHiPibVmr75nuOi5KOpyVdWRHbNi63URcfqQgf
BBckWKo3Shjf5TnUV/3XwSyRAZHiItQDwFj8d0fsjz50Q7qsNI1NOHZnjrDIbzAz
WHFctPVrbtQBULgTfmxKo0nRIBnuvMApGGWn3v7v3QqQIecaZ5JCEJhfTzC8PhxF
tBDXaEAUwED653cXeuYLj2VbPNmaUtu1vZ5Gzz3rkQUCwJaydkxNEJY7kvqcfw+Z
374jNUUeAlz+taibmSXaXvMiwzn15Cou08YfxGyqxRxqAQVKL9LFwag0Jl1mpdIC
IfkYtwb1TplvqKtMUejPUBjFd8g5CSxJkjKZqLsXF3mwWsXmo8RZZUc1g16p6DUL
mbvkzSDGm0oGObVo/CK67lWMK07q87Hj/LaZmtVC+nFNCM+HHmpxffnTtOmlcYF7
wk5HlqX2doWjKI/pgG6BU6VtX7hI+cL5NqYuSf+4lsKMB7ObiFj86xsc3i1w4peS
MKGJ47xVqCfWS+2QrYv6YyVZLag13cqXM7zlzced0ezvXg5KkAYmY6252TUtB7p2
ZSysV4999AeU14ECll2jB0nVetBX+RvnU0Z1qrB5QstocQjpYL05ac70r8NWQMet
UqIJ5G+GR4of6ygnXYMgrwTJbFaai0b1AgMBAAGjgYMwgYAwDwYDVR0TAQH/BAUw
AwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFPd9xf3E6Jobd2Sn9R2gzL+H
YJptMD4GA1UdIAQ3MDUwMwYEVR0gADArMCkGCCsGAQUFBwIBFh1odHRwOi8vd3d3
LmNlcnQuZm5tdC5lcy9kcGNzLzANBgkqhkiG9w0BAQsFAAOCAgEAB5BK3/MjTvDD
nFFlm5wioooMhfNzKWtN/gHiqQxjAb8EZ6WdmF/9ARP67Jpi6Yb+tmLSbkyU+8B1
RXxlDPiyN8+sD8+Nb/kZ94/sHvJwnvDKuO+3/3Y3dlv2bojzr2IyIpMNOmqOFGYM
LVN0V2Ue1bLdI4E7pWYjJ2cJj+F3qkPNZVEI7VFY/uY5+ctHhKQV8Xa7pO6kO8Rf
77IzlhEYt8llvhjho6Tc+hj507wTmzl6NLrTQfv6MooqtyuGC2mDOL7Nii4LcK2N
JpLuHvUBKwrZ1pebbuCoGRw6IYsMHkCtA+fdZn71uSANA+iW+YJF1DngoABd15jm
fZ5nc8OaKveri6E6FO80vFIOiZiaBECEHX5FaZNXzuvO+FB8TxxuBEOb+dY7Ixjp
6o7RTUaN8Tvkasq6+yO3m/qZASlaWFot4/nUbQ4mrcFuNLwy+AwF+mWj2zs3gyLp
1txyM/1d8iC9djwj2ij3+RvrWWTV3F9yfiD8zYm1kGdNYno/Tq0dwzn+evQoFt9B
9kiABdcPUXmsEKvU7ANm5mqwujGSQkBqvjrTcuFqN1W8rB2Vt2lh8kORdOag0wok
RqEIr9baRRmW1FMdW4R58MD3R++Lj8UGrp1MYp3/RgT408m2ECVAdf4WqslKYIYv
uu8wd+RU4riEmViAqhOLUTpPSPaLtrM=
-----END CERTIFICATE-----

# AC RAIZ FNMT-RCM SERVIDORES SEGUROS
-----BEGIN CERTIFICATE-----
MIICbjCCAfOgAwIBAgIQYvYybOXE42hcG2LdnC6dlTAKBggqhkjOPQQDAzB4MQsw
CQYDVQQGEwJFUzERMA8GA1UECgwIRk5NVC1SQ00xDjAMBgNVBAsMBUNlcmVzMRgw
FgYDVQRhDA9WQVRFUy1RMjgyNjAwNEoxLDAqBgNVBAMMI0FDIFJBSVogRk5NVC1S
Q00gU0VSVklET1JFUyBTRUdVUk9TMB4XDTE4MTIyMDA5MzczM1oXDTQzMTIyMDA5
MzczM1oweDELMAkGA1UEBhMCRVMxETAPBgNVBAoMCEZOTVQtUkNNMQ4wDAYDVQQL
DAVDZXJlczEYMBYGA1UEYQwPVkFURVMtUTI4MjYwMDRKMSwwKgYDVQQDDCNBQyBS
QUlaIEZOTVQtUkNNIFNFUlZJRE9SRVMgU0VHVVJPUzB2MBAGByqGSM49AgEGBSuB
BAAiA2IABPa6V1PIyqvfNkpSIeSX0oNnnvBlUdBeh8dHsVnyV0ebAAKTRBdp20LH
sbI6GA60XYyzZl2hNPk2LEnb80b8s0RpRBNm/dfF/a82Tc4DTQdxz69qBdKiQ1oK
Um8BA06Oi6NCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYD
VR0OBBYEFAG5L++/EYZg8k/QQW6rcx/n0m5JMAoGCCqGSM49BAMDA2kAMGYCMQCu
SuMrQMN0EfKVrRYj3k4MGuZdpSRea0R7/DjiT8ucRRcRTBQnJlU5dUoDzBOQn5IC
MQD6SmxgiHPz7riYYqnOK8LZiqZwMR2vsJRM60/G49HzYqc8/5MuB1xJAWdpEgJy
v+c=
-----END CERTIFICATE-----

# ANF Secure Server Root CA
-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----


-----BEGIN CERTIFICATE-----
MIIFQjCCAyqgAwIBAgIUOSjtdkVvhNB3mssMD+L004fls2QwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wHhcNMjQxMDA3MTU0NDEyWhcNMzQx
MDA1MTU0NDEyWjAWMRQwEgYDVQQDDAtleGFtcGxlLmNvbTCCAiIwDQYJKoZIhvcN
AQEBBQADggIPADCCAgoCggIBANuEMkqyS0UrlkAMKMzPsmvdwp0fefH564JXXOZ5
v5TnkubqI7ijKBXwSvPCTCuebZFYqIpqN58dSyeMewkEjngAH99+NAuIlt5nyAg5
+BxZiRNgugdnc6kAcBH+2C/3T4P0HSd0NPGVQOC4GxMAy6Shz5LiztunPkWPUxWy
8OZYO86xilW4L6Dv6JuPCsX7vzv90a31rtfghPlFUz4frcfGpdmTUeXv38/aVZRz
ToGgMvQUIaVvfxfmTqrWLCTy/4sTvIY8mVgCaGw5wNDv2oUiZ7C18w0orWeY0vc4
b+F+ma2Mwa04pj0KluwJFx85qH3lEk50yfnLYWUr3ZYwu2TOTqpDr61aR+NZqbw+
QZWrUImipTzbtVqezmzCLQqXI3zPsmhXb6+k/Ykj1xj9deE7vYAIc8MBVxhw1pYQ
fxRgx+dhyZpqYDW8fCUF/5vXtABm9we+/edL7vFKfBuSipeZeAie+a+SO6aMLCIZ
rMnH1PdJjzKTbCuWGA6K75y6X3vh2qenmh0/qVJljKktWfWq4Gp5u6EjcZ2mljzh
kDMMrRikWq/ThLOFLHC1wjy3IQ7SLGfnM9qharpHTVinwU1g97RGdzFXiB+y69O1
MlIlIXWeNJsU94pf1VidKz1hLJynw5cx7ifU+eSvbRU0ZLIHjuX3G4gVlwkUKPNN
rdgXAgMBAAGjgYcwgYQwHQYDVR0OBBYEFBKXOJluZKJ+yy9XfVvmEBf3KspVMB8G
A1UdIwQYMBaAFBKXOJluZKJ+yy9XfVvmEBf3KspVMA8GA1UdEwEB/wQFMAMBAf8w
MQYDVR0RBCowKIILZXhhbXBsZS5jb22CDSouZXhhbXBsZS5jb22HBAoAAAGHBH8A
AAEwDQYJKoZIhvcNAQELBQADggIBAEFfjgkurGI/ouVaJgCJYXmf+mJtexN6JYAB
XuCcobhUM+t5bfYt9DbZhaNC3pvtT9OijujajnXmd20QJLgXbWJe3qMbiDYroXZ/
ry8JHX9Nlp3wF3V/iGUQw4zLnFRalAeSyEAwg+nsoTwA6vkCompNFqzozh2ViSo+
ucWkb71Ky7Fl3HmKxp9ohG/0REwLTMYYmlPaGvk4o3oWAH3jm6g0fuFiw3mC1Gvr
nFiMva0JJ3LvaOZfoe2U2dSCFEleNqVMEvLNRF4Sd6SFR7/IhwjzvyxSw0DRSmmt
FZPakLw1Dm2eibT0rhEhn+7fLVHbaTSgOg8diXKqI+kqxbReYbrGkfa4lIfrvAMi
+ax9OQ54KIdQU3uqugyyTqqN1WERjvuv0uhWSF1sEk6AdpMWR8ymircQ72yEJmjI
ycRC2okZKkhYwFHb4nYwhQJOtzkmIUc90xmXZK4EvC9SYbHDu/8RBFRqkF4BnE8y
4eoywGFMRIJ4GdAzaI7rpQmtSvbZsiismBkTGU/IibM6udnnp5xD4R3HDDSdrkMH
MpupMy3sYHTJ+pi4OdpvYBfQHy1Cq0RxRc5SCSb0Dn6IggUVgyG0QpT4SuIM5qt8
3cp1JPJuZpbks3EkjqeyzKgfWsJWWi/9q68zbubW5oOgHk2548BO9x2njvEVfrTg
rDo5uzq8
-----END CERTIFICATE-----

`),
		"3 certificates and private keys": []byte(`# AC RAIZ FNMT-RCM
-----BEGIN CERTIFICATE-----
MIIFgzCCA2ugAwIBAgIPXZONMGc2yAYdGsdUhGkHMA0GCSqGSIb3DQEBCwUAMDsx
CzAJBgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJ
WiBGTk1ULVJDTTAeFw0wODEwMjkxNTU5NTZaFw0zMDAxMDEwMDAwMDBaMDsxCzAJ
BgNVBAYTAkVTMREwDwYDVQQKDAhGTk1ULVJDTTEZMBcGA1UECwwQQUMgUkFJWiBG
Tk1ULVJDTTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBALpxgHpMhm5/
yBNtwMZ9HACXjywMI7sQmkCpGreHiPibVmr75nuOi5KOpyVdWRHbNi63URcfqQgf
BBckWKo3Shjf5TnUV/3XwSyRAZHiItQDwFj8d0fsjz50Q7qsNI1NOHZnjrDIbzAz
WHFctPVrbtQBULgTfmxKo0nRIBnuvMApGGWn3v7v3QqQIecaZ5JCEJhfTzC8PhxF
tBDXaEAUwED653cXeuYLj2VbPNmaUtu1vZ5Gzz3rkQUCwJaydkxNEJY7kvqcfw+Z
374jNUUeAlz+taibmSXaXvMiwzn15Cou08YfxGyqxRxqAQVKL9LFwag0Jl1mpdIC
IfkYtwb1TplvqKtMUejPUBjFd8g5CSxJkjKZqLsXF3mwWsXmo8RZZUc1g16p6DUL
mbvkzSDGm0oGObVo/CK67lWMK07q87Hj/LaZmtVC+nFNCM+HHmpxffnTtOmlcYF7
wk5HlqX2doWjKI/pgG6BU6VtX7hI+cL5NqYuSf+4lsKMB7ObiFj86xsc3i1w4peS
MKGJ47xVqCfWS+2QrYv6YyVZLag13cqXM7zlzced0ezvXg5KkAYmY6252TUtB7p2
ZSysV4999AeU14ECll2jB0nVetBX+RvnU0Z1qrB5QstocQjpYL05ac70r8NWQMet
UqIJ5G+GR4of6ygnXYMgrwTJbFaai0b1AgMBAAGjgYMwgYAwDwYDVR0TAQH/BAUw
AwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYDVR0OBBYEFPd9xf3E6Jobd2Sn9R2gzL+H
YJptMD4GA1UdIAQ3MDUwMwYEVR0gADArMCkGCCsGAQUFBwIBFh1odHRwOi8vd3d3
LmNlcnQuZm5tdC5lcy9kcGNzLzANBgkqhkiG9w0BAQsFAAOCAgEAB5BK3/MjTvDD
nFFlm5wioooMhfNzKWtN/gHiqQxjAb8EZ6WdmF/9ARP67Jpi6Yb+tmLSbkyU+8B1
RXxlDPiyN8+sD8+Nb/kZ94/sHvJwnvDKuO+3/3Y3dlv2bojzr2IyIpMNOmqOFGYM
LVN0V2Ue1bLdI4E7pWYjJ2cJj+F3qkPNZVEI7VFY/uY5+ctHhKQV8Xa7pO6kO8Rf
77IzlhEYt8llvhjho6Tc+hj507wTmzl6NLrTQfv6MooqtyuGC2mDOL7Nii4LcK2N
JpLuHvUBKwrZ1pebbuCoGRw6IYsMHkCtA+fdZn71uSANA+iW+YJF1DngoABd15jm
fZ5nc8OaKveri6E6FO80vFIOiZiaBECEHX5FaZNXzuvO+FB8TxxuBEOb+dY7Ixjp
6o7RTUaN8Tvkasq6+yO3m/qZASlaWFot4/nUbQ4mrcFuNLwy+AwF+mWj2zs3gyLp
1txyM/1d8iC9djwj2ij3+RvrWWTV3F9yfiD8zYm1kGdNYno/Tq0dwzn+evQoFt9B
9kiABdcPUXmsEKvU7ANm5mqwujGSQkBqvjrTcuFqN1W8rB2Vt2lh8kORdOag0wok
RqEIr9baRRmW1FMdW4R58MD3R++Lj8UGrp1MYp3/RgT408m2ECVAdf4WqslKYIYv
uu8wd+RU4riEmViAqhOLUTpPSPaLtrM=
-----END CERTIFICATE-----

# AC RAIZ FNMT-RCM SERVIDORES SEGUROS
-----BEGIN CERTIFICATE-----
MIICbjCCAfOgAwIBAgIQYvYybOXE42hcG2LdnC6dlTAKBggqhkjOPQQDAzB4MQsw
CQYDVQQGEwJFUzERMA8GA1UECgwIRk5NVC1SQ00xDjAMBgNVBAsMBUNlcmVzMRgw
FgYDVQRhDA9WQVRFUy1RMjgyNjAwNEoxLDAqBgNVBAMMI0FDIFJBSVogRk5NVC1S
Q00gU0VSVklET1JFUyBTRUdVUk9TMB4XDTE4MTIyMDA5MzczM1oXDTQzMTIyMDA5
MzczM1oweDELMAkGA1UEBhMCRVMxETAPBgNVBAoMCEZOTVQtUkNNMQ4wDAYDVQQL
DAVDZXJlczEYMBYGA1UEYQwPVkFURVMtUTI4MjYwMDRKMSwwKgYDVQQDDCNBQyBS
QUlaIEZOTVQtUkNNIFNFUlZJRE9SRVMgU0VHVVJPUzB2MBAGByqGSM49AgEGBSuB
BAAiA2IABPa6V1PIyqvfNkpSIeSX0oNnnvBlUdBeh8dHsVnyV0ebAAKTRBdp20LH
sbI6GA60XYyzZl2hNPk2LEnb80b8s0RpRBNm/dfF/a82Tc4DTQdxz69qBdKiQ1oK
Um8BA06Oi6NCMEAwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8EBAMCAQYwHQYD
VR0OBBYEFAG5L++/EYZg8k/QQW6rcx/n0m5JMAoGCCqGSM49BAMDA2kAMGYCMQCu
SuMrQMN0EfKVrRYj3k4MGuZdpSRea0R7/DjiT8ucRRcRTBQnJlU5dUoDzBOQn5IC
MQD6SmxgiHPz7riYYqnOK8LZiqZwMR2vsJRM60/G49HzYqc8/5MuB1xJAWdpEgJy
v+c=
-----END CERTIFICATE-----

# ANF Secure Server Root CA
-----BEGIN CERTIFICATE-----
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
-----END CERTIFICATE-----
-----BEGIN PRIVATE KEY-----
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
qA/T5R0K3G5f2YtkxuS29MeFywEK7sI=
-----END PRIVATE KEY-----
`),
	}
	return sampleCerts[name]
}

func TestNewCertstore(t *testing.T) {
	cases := map[string]struct {
		inputCerts       []byte
		inputSource      string
		expectedElements int
		expectedSource   string
	}{
		"single certificate": {
			inputCerts:       getSampleCert("single certificate"),
			inputSource:      "test",
			expectedElements: 1,
			expectedSource:   "test",
		},
		"multiple certificates with comments": {
			inputCerts:       getSampleCert("multiple certificates with comments"),
			inputSource:      "ca.crt",
			expectedElements: 4,
			expectedSource:   "ca.crt",
		},
		"empty input": {
			inputCerts:       []byte{},
			expectedElements: 0,
		},
		"3 certificate and 1 private key": {
			inputCerts:       getSampleCert("3 certificates and private keys"),
			inputSource:      "certificates/anchors/ca.crt",
			expectedElements: 3,
			expectedSource:   "certificates/anchors/ca.crt",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			cs := NewCertstore()
			cs.Load(tc.inputCerts, tc.inputSource)
			gotQuantity := len(cs.certs)
			if gotQuantity != tc.expectedElements {
				t.Errorf("expected number of elements: %v - got: %v", tc.expectedElements, gotQuantity)
			}
			for n, cert := range cs.certs {
				if cert.Source != tc.expectedSource {
					t.Errorf("expected Source: \"%s\" on certificate #%d - got: %v", tc.expectedSource, n, cert.Source)
				}
			}
		})
	}
}

func TestCertstoreMultipleLoads(t *testing.T) {
	cs := NewCertstore()
	cs.Load(getSampleCert("single certificate"), "single")
	cs.Load(getSampleCert("3 certificates and private keys"), "3certsAndPrivateKey")
	if len(cs.certs) != 4 {
		t.Errorf("expected 4 elements - got: %d", len(cs.certs))
	}
	if cs.certs[0].Source != "single" {
		t.Errorf("expected Source of cert #0: single - got: %s", cs.certs[0].Source)
	}
	for i := 1; i < 4; i++ {
		if cs.certs[i].Source != "3certsAndPrivateKey" {
			t.Errorf("expected Source of cert #%d: 3certsAndPrivateKey - got: %s", i, cs.certs[i].Source)
		}
	}
}

func TestGetCertAttributes(t *testing.T) {
	cs := NewCertstore()
	cs.Load(getSampleCert("multiple certificates with comments"), "test")
	cases := map[int]struct {
		expectedSubject      string
		expectedIssuer       string
		expectedSerialNumber string
		//		sans                    SANs
		expectedNotBefore string
		expectedNotAfter  string
		expectedSKID      string
		expectedAKID      string
	}{
		0: {
			expectedSubject:      "OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
			expectedIssuer:       "OU=AC RAIZ FNMT-RCM,O=FNMT-RCM,C=ES",
			expectedSerialNumber: "5d:93:8d:30:67:36:c8:06:1d:1a:c7:54:84:69:07",
			expectedNotBefore:    "2008-10-29 15:59:56 +0000 UTC",
			expectedNotAfter:     "2030-01-01 00:00:00 +0000 UTC",
			expectedSKID:         "F7:7D:C5:FD:C4:E8:9A:1B:77:64:A7:F5:1D:A0:CC:BF:87:60:9A:6D",
		},
		1: {
			expectedSubject:      "CN=AC RAIZ FNMT-RCM SERVIDORES SEGUROS,OU=Ceres,O=FNMT-RCM,C=ES,2.5.4.97=#130f56415445532d51323832363030344a",
			expectedIssuer:       "CN=AC RAIZ FNMT-RCM SERVIDORES SEGUROS,OU=Ceres,O=FNMT-RCM,C=ES,2.5.4.97=#130f56415445532d51323832363030344a",
			expectedSerialNumber: "62:f6:32:6c:e5:c4:e3:68:5c:1b:62:dd:9c:2e:9d:95",
			expectedNotBefore:    "2018-12-20 09:37:33 +0000 UTC",
			expectedNotAfter:     "2043-12-20 09:37:33 +0000 UTC",
			expectedSKID:         "01:B9:2F:EF:BF:11:86:60:F2:4F:D0:41:6E:AB:73:1F:E7:D2:6E:49",
		},
		2: {
			expectedSubject:      "SERIALNUMBER=G63287510,CN=ANF Secure Server Root CA,OU=ANF CA Raiz,O=ANF Autoridad de Certificacion,C=ES",
			expectedIssuer:       "SERIALNUMBER=G63287510,CN=ANF Secure Server Root CA,OU=ANF CA Raiz,O=ANF Autoridad de Certificacion,C=ES",
			expectedSerialNumber: "0d:d3:e3:bc:6c:f9:6b:b1",
			expectedNotBefore:    "2019-09-04 10:00:38 +0000 UTC",
			expectedNotAfter:     "2039-08-30 10:00:38 +0000 UTC",
			expectedSKID:         "9C:5F:D0:6C:63:A3:5F:93:CA:93:98:08:AD:8C:87:A5:2C:5C:C1:37",
			expectedAKID:         "9C:5F:D0:6C:63:A3:5F:93:CA:93:98:08:AD:8C:87:A5:2C:5C:C1:37",
		},
		3: {
			expectedSubject:      "CN=example.com",
			expectedIssuer:       "CN=example.com",
			expectedSerialNumber: "39:28:ed:76:45:6f:84:d0:77:9a:cb:0c:0f:e2:f4:d3:87:e5:b3:64",
			expectedNotBefore:    "2024-10-07 15:44:12 +0000 UTC",
			expectedNotAfter:     "2034-10-05 15:44:12 +0000 UTC",
			expectedSKID:         "12:97:38:99:6E:64:A2:7E:CB:2F:57:7D:5B:E6:10:17:F7:2A:CA:55",
			expectedAKID:         "12:97:38:99:6E:64:A2:7E:CB:2F:57:7D:5B:E6:10:17:F7:2A:CA:55",
			//			sans: SANs{
			//				DNS: []string{"example.com", "*.example.com"},
			//				IP:  []net.IP{net.IPv4(10, 0, 0, 1), net.IPv4(127, 0, 0, 1)},
		},
	}

	for n, tc := range cases {
		t.Run("test "+fmt.Sprint(n), func(t *testing.T) {
			gotSerialNumber := cs.certs[n].GetSerialNumber()
			if gotSerialNumber != tc.expectedSerialNumber {
				t.Errorf("expected: %v - got: %v\n", tc.expectedSerialNumber, gotSerialNumber)
			}

			gotIssuer := cs.certs[n].GetIssuer()
			if gotIssuer != tc.expectedIssuer {
				t.Errorf("expected: %v - got: %v\n", tc.expectedIssuer, gotIssuer)
			}

			gotSubject := cs.certs[n].GetSubject()
			if gotSubject != tc.expectedSubject {
				t.Errorf("expected: %v - got: %v\n", tc.expectedSubject, gotSubject)
			}

			gotNotBefore := cs.certs[n].GetNotBefore()
			if gotNotBefore != tc.expectedNotBefore {
				t.Errorf("expected NotBefore: %v - got: %v\n", tc.expectedNotBefore, gotNotBefore)
			}

			gotNotAfter := cs.certs[n].GetNotAfter()
			if gotNotAfter != tc.expectedNotAfter {
				t.Errorf("expected NotAfter: %v - got: %v\n", tc.expectedNotAfter, gotNotAfter)
			}

			gotSKID := cs.certs[n].GetSKID()
			if gotSKID != tc.expectedSKID {
				t.Errorf("expected SKID: %v - got: %v\n", tc.expectedSKID, gotSKID)
			}

			gotAKID := cs.certs[n].GetAKID()
			if gotAKID != tc.expectedAKID {
				t.Errorf("expected AKID: %v - got: %v\n", tc.expectedAKID, gotAKID)
			}
		})
	}
}
