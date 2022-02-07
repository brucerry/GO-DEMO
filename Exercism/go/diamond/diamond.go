package diamond

import (
    "errors"
    "bytes"
)

func Gen(char byte) (output string, e error) {
	if char < 'A' || char > 'Z' {
        return "", errors.New("")
    }
    row := (char - 'A' + 1) + (char - 'A')
    col := row
    buf := make([][]byte, row)
    for i := range buf {
        buf[i] = make([]byte, col)
    }
    c := byte('A')
    for i,j,k,l:=byte(0),byte(col-1),col/2,col/2; i<row/2+1; i,j,k,l,c=i+1,j-1,k-1,l+1,c+1 {
        init_line := bytes.Repeat([]byte{' '}, int(col))
        buf[i], buf[j] = init_line, init_line
        buf[i][k], buf[i][l], buf[j][k], buf[j][l] = c, c, c, c
    }
	output = string(bytes.Join(buf, []byte{'\n'}))
    output += "\n"
    return output, nil
} 
