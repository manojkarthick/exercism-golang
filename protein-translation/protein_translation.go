package protein

import (
	"unicode/utf8"

	"github.com/pkg/errors"
)

const CODON_SIZE = 3

var codonProteinTranslation = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var (
	ErrStop        = errors.New("STOP protein encountered.")
	ErrInvalidBase = errors.New("Invalid base given.")
)

func getChunks(input string, size int) []string {
	runes := []rune(input)
	var chunks []string
	for i := 0; i < utf8.RuneCountInString(input); i += 3 {
		chunks = append(chunks, string(runes[i:i+size]))
	}
	return chunks
}

func FromCodon(codon string) (string, error) {
	protein, ok := codonProteinTranslation[codon]
	if !ok {
		return "", ErrInvalidBase

	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for _, chunk := range getChunks(rna, CODON_SIZE) {
		protein, err := FromCodon(chunk)
		if err != nil {
			if err == ErrStop {
				break
			}
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
