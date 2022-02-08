package protein

import "errors"

var (
	ErrStop        = errors.New("STOP protein encountered.")
	ErrInvalidBase = errors.New("Invalid base given.")
)

func FromCodon(codon string) (string, error) {
    switch codon {
        case "AUG": return "Methionine", nil
    	case "UUU", "UUC": return "Phenylalanine", nil
    	case "UUA", "UUG": return "Leucine", nil
    	case "UCU", "UCC", "UCA", "UCG": return "Serine", nil
    	case "UAU", "UAC": return "Tyrosine", nil
    	case "UGU", "UGC": return "Cysteine", nil
    	case "UGG": return "Tryptophan", nil
    	case "UAA", "UAG", "UGA": return "", ErrStop
    	default: return "", ErrInvalidBase
    }
} 

func FromRNA(rna string) ([]string, error) {
    collection := []string{}
    var codon string
    for i:=0; i<=len(rna); i++ {
        if i%3 == 0 && i!=0 {
            name, err := FromCodon(codon)
            if err != nil {
                if err == ErrStop {
                    return collection, nil
                } else {
                	return collection, err
                }
            }
            collection = append(collection, name)
			if i < len(rna) {
				codon = string(rna[i])
			}
        } else {
        	codon += string(rna[i])
        }
    }
	return collection, nil
}