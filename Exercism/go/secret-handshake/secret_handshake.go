package secret

import "math"

var handshake = [4]string{
    "wink",
    "double blink",
    "close your eyes",
    "jump",
}

func Handshake(num uint) (output []string) {
  for i:=1; i<=8; i=i<<1 {
      if num & uint(i) > 0 {
          output = append(output, handshake[int(math.Log2(float64(i)))])
      }
  }
  if num & 16 > 0 {
    for i,j:=0,len(output)-1; i<j; i,j=i+1,j-1 {
      output[i], output[j] = output[j], output[i]
    }
  }
  return
}