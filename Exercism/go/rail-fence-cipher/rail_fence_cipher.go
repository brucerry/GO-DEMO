package railfence

func Encode(message string, rails int) string {
    row, col := rails, len(message)
    arr := make([][]rune, row)
    for i := range arr {
        arr[i] = make([]rune, col)
    }
	left_to_right_index := 0
    row_index := 0
    move_downward := true
	for left_to_right_index != len(message) {
        arr[row_index][left_to_right_index] = rune(message[left_to_right_index])
        left_to_right_index++
        if move_downward {
            row_index++
        } else {
        	row_index--
        }
    	if row_index == row - 1 {
            move_downward = false
        }
    	if row_index == 0 {
            move_downward = true
        }
    }
	output := ""
	for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if arr[i][j] != 0 {
                output += string(arr[i][j])
            }
        }
    }
	return output
}

func Decode(message string, rails int) string {
    row, col := rails, len(message)
    arr := make([][]rune, row)
    for i := range arr {
        arr[i] = make([]rune, col)
    }
	left_to_right_index := 0
    row_index := 0
    move_downward := true
	for left_to_right_index != len(message) {
        arr[row_index][left_to_right_index] = '?'
        left_to_right_index++
        if move_downward {
            row_index++
        } else {
        	row_index--
        }
    	if row_index == row - 1 {
            move_downward = false
        }
    	if row_index == 0 {
            move_downward = true
        }
    }
	left_to_right_index = 0
	for i:=0; i<row; i++ {
        for j:=0; j<col; j++ {
            if arr[i][j] == '?' {
                arr[i][j] = rune(message[left_to_right_index])
                left_to_right_index++
            }
        }
    }
	output := ""
	left_to_right_index = 0
    row_index = 0
    move_downward = true
	for left_to_right_index != len(message) {
        output += string(arr[row_index][left_to_right_index])
        left_to_right_index++
        if move_downward {
            row_index++
        } else {
        	row_index--
        }
    	if row_index == row - 1 {
            move_downward = false
        }
    	if row_index == 0 {
            move_downward = true
        }
    }
	return output
}
