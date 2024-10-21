# certexplorer

`certexplorer` is a command line tool designed to easily navigate through certificates with a nicely formatted output. Its use cases range from quickly check a single certificate in standard input to the inspection of hundreds of certificates contained in multiple ca bundle files.

![certexplorer screenshot](./screenshot.png)

## Installation
`certexplorer` requires Go 1.22 or later.

``` shell
go install github.com/dcommisso/certexplorer@latest
```

## Usage

``` shell
# Get all the default fields for all the certificates in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem

# Get only the serial number and the validity of the certificates in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem -f serial,validity

# Get only the serial number and the validity of the certificates 119 and 3 in provided files
$ certexplorer certificate.pem cabundle-full-of-certificates.pem -f serial,validity -c 119,3

# Get the expiration date and the subject of a certificate from standard input
$ echo "-----BEGIN CERTIFICATE----- ..." | certexplorer -f notafter,subject

# Get some certificates from cabundle in plain output (useful for get the raw certificate without indentation)
$ certexplorer cabundle-full-of-certificates.pem -c 1,5 -o plain
```

Run `certexplorer --help` for the full documentation.

## Features
- Certificates can be read from multiple files or standard input.
- Flexible output: it's possible to choose which certificates and fields (and in what order) to get.
- By default the expiration date is colored red if the certificate is expired or yellow if it expires in less than a month.
- The `source` field helps to easily find out in which file each certificate is contained (especially useful during the analysis of multiple cabundle files).
- Alternative format output: `plain` output makes easy to copy-and-paste the certificate without indentation.
- Colors can be disabled with `--no-color` flag.
