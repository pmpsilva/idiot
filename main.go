package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	flagUUID      = flag.Bool("uuid", false, "Generate a UUID v4")
	flagULID      = flag.Bool("ulid", false, "Generate a ULID")
	flagPrefix    = flag.String("prefix", "", "Prefix to prepend when generating ULIDs")
	flagLuhn      = flag.String("luhn", "", "Calculate Luhn check digit and append it to the provided number (digits only)")
	flagCD        = flag.String("cd", "", "Alias for -luhn")
	flagValidate  = flag.String("validate", "", "Validate provided number (digits only) using Luhn algorithm")
	flagCDV       = flag.String("cdv", "", "Alias for -validate")
	flagPass      = flag.Bool("pass", false, "Generate a random password")
	flagPassShort = flag.Bool("p", false, "Alias for -pass")
	flagPassLen   = flag.Int("l", 12, "Password length")
	flagPassChars = flag.Bool("c", true, "Include letters in password generation")
	flagPassDigits = flag.Bool("d", true, "Include digits in password generation")
	flagPassSpecial = flag.Bool("s", true, "Include special characters in password generation")
	flagHelp      = flag.Bool("h", false, "Show help")
	flagHelpLong  = flag.Bool("help", false, "Show help")
	errNoAction   = errors.New("no action specified; use -h for help")
	errConflicted = errors.New("multiple actions requested; choose one")
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if *flagHelp || *flagHelpLong {
		usage()
		return
	}

	actions := 0
	if *flagUUID {
		actions++
	}
	if *flagULID {
		actions++
	}
	if *flagLuhn != "" {
		actions++
	}
	if *flagCD != "" {
		actions++
	}
	if *flagValidate != "" {
		actions++
	}
	if *flagCDV != "" {
		actions++
	}
	if *flagPass || *flagPassShort {
		actions++
	}

	if actions == 0 {
		exitWith(errNoAction)
	}
	if actions > 1 {
		exitWith(errConflicted)
	}

	switch {
	case *flagUUID:
		fmt.Println(generateUUID())
	case *flagULID:
		fmt.Println(generateULID(*flagPrefix))
	case *flagLuhn != "" || *flagCD != "":
		number := firstNonEmpty(*flagLuhn, *flagCD)
		if !isDigits(number) {
			exitWith(errors.New("luhn input must be digits only"))
		}
		checkDigit := calculateLuhnCheckDigit(number)
		fmt.Println(number + strconv.Itoa(checkDigit))
	case *flagValidate != "" || *flagCDV != "":
		number := firstNonEmpty(*flagValidate, *flagCDV)
		if !isDigits(number) {
			exitWith(errors.New("validate input must be digits only"))
		}
		if validateLuhn(number) {
			cd := number[len(number)-1] - '0'
			fmt.Printf("valid (check digit: %d)\n", cd)
		} else {
			fmt.Println("invalid")
		}
	case *flagPass || *flagPassShort:
		password, err := generatePassword(*flagPassLen, *flagPassChars, *flagPassDigits, *flagPassSpecial)
		if err != nil {
			exitWith(err)
		}
		fmt.Println(password)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `idiot - simple ID and Luhn utility

Usage:
  idiot -uuid                      Generate a UUID v4
  idiot -ulid [-prefix PREFIX]     Generate a ULID, optionally prefixed
  idiot -luhn NUMBER               Append Luhn check digit to NUMBER (digits only)
  idiot -cd NUMBER                 Alias for -luhn
  idiot -validate NUMBER           Validate NUMBER using Luhn
  idiot -cdv NUMBER                Alias for -validate
  idiot -pass | -p [-l LEN -c=true -d=true -s=true]
                                    Generate random password (default length 12, include letters, digits, and special characters)
  idiot -h | --help                Show this help
`)
}

func exitWith(err error) {
	log.Println(err)
	os.Exit(1)
}
