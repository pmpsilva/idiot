# idiot

A small command-line utility for generating and validating identifiers commonly needed during development and testing.

## Install

```sh
go install github.com/pmpsilva/idiot@latest
```

Or build from source:

```sh
git clone https://github.com/pmpsilva/idiot
cd idiot
go build -o idiot .
```

## Usage

Only one action can be specified per invocation.

```
idiot <flag> [options]
```

---

### UUID v4

Generate a random UUID v4.

```sh
idiot -uuid
# e.g. b7315e5d-d669-4bf8-8076-72ba139824f4
```

---

### ULID

Generate a [ULID](https://github.com/ulid/spec) (Universally Unique Lexicographically Sortable Identifier).

```sh
idiot -ulid
# e.g. 01J5Z8K3QG1V2W3X4Y5Z6A7B8C

idiot -ulid -prefix order_
# e.g. order_01J5Z8K3QG1V2W3X4Y5Z6A7B8C
```

| Flag | Description |
|------|-------------|
| `-ulid` | Generate a ULID |
| `-prefix PREFIX` | Optional string to prepend to the generated ULID |

---

### Luhn — append check digit

Append a [Luhn](https://en.wikipedia.org/wiki/Luhn_algorithm) check digit to a digit string.

```sh
idiot -luhn 12345
# 123455

idiot -cd 12345
# 123455  (alias)
```

| Flag | Description |
|------|-------------|
| `-luhn NUMBER` | Append Luhn check digit to NUMBER (digits only) |
| `-cd NUMBER` | Alias for `-luhn` |

---

### Luhn — validate

Validate a number using the Luhn algorithm.

```sh
idiot -validate 123455
# valid (check digit: 5)

idiot -validate 123456
# invalid

idiot -cdv 123455
# valid (check digit: 5)  (alias)
```

| Flag | Description |
|------|-------------|
| `-validate NUMBER` | Validate NUMBER using Luhn (digits only) |
| `-cdv NUMBER` | Alias for `-validate` |

---

### EID — validate

Validate a telco EID (eUICC Identifier) per [GSMA SGP.02](https://www.gsma.com/esim/resources/sgp-02/).

An EID is exactly 32 decimal digits. The last two are check digits computed as `98 - (n mod 97)` where `n` is the 32-digit value with the check digits replaced by `00`.

```sh
idiot -eid 89049032005008882600331335531462
# valid (check digits: 62)

idiot -eid 89049032005008882600331335531400
# invalid
```

| Flag | Description |
|------|-------------|
| `-eid EID` | Validate a 32-digit EID using GSMA SGP.02 check digits |

---

### EID — generate

Generate a random EID with syntactically valid check digits.

> **For testing/development only.** Generated EIDs are not real, provisionable EIDs and will not resolve on any SM-DP+ or SM-DS.

```sh
idiot -neweid
# e.g. 89458820178882829923163983041645

idiot -neweid -prefix 8944
# e.g. 89444575119504363851885189849583
```

| Flag | Description |
|------|-------------|
| `-neweid` | Generate a random test EID |
| `-prefix DIGITS` | Optional leading digits (up to 30, digits only). Defaults to `89` (telecom MII) |

---

### Password

Generate a cryptographically random password.

```sh
idiot -pass
# e.g. aB3!xZ9@kL2#

idiot -p
# alias for -pass

idiot -pass -l 20
# 20-character password

idiot -pass -l 10 -s=false
# 10-character password, no special characters

idiot -pass -l 8 -c=false
# digits and special characters only
```

| Flag | Default | Description |
|------|---------|-------------|
| `-pass` | — | Generate a random password |
| `-p` | — | Alias for `-pass` |
| `-l LEN` | `12` | Password length |
| `-c` | `true` | Include letters (a-z, A-Z) |
| `-d` | `true` | Include digits (0-9) |
| `-s` | `true` | Include special characters (`!@#$%^&*…`) |

---

### Help

```sh
idiot -h
idiot -help
```
