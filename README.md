# Ethereum Passphrase Bruteforcer

Forgot your Ethereum wallet passphrase and need a simple tool to help bruteforce
your passphrase. Of course you do.

## Install

You need to [Go][1] toolchain to install.

```sh
$ go get github.com/vasuman/eth-bruteforce
```

## Usage

```sh
$ eth-bruteforce --password-file=<password-file> --keystore-dir=<keystore-dir>
```

The script will loop through the every key in the `keystore-dir` sequentially
trying passwords from the `password-file`.

The `password-file` is a plain-text file containing a line separated list of
possible passwords. You can use the `#` symbol as the first character in the
line to skip entries (and `\#` to escape). For example,

```
password1
#failed2
password3
```

[1]: https://golang.org/doc/install#install
