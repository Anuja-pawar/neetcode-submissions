func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	sqr := make([]map[byte]bool, 9)

	for i := range(9){
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		sqr[i] = make(map[byte]bool)
	}

	for r := range(9){
		for c := range(9){
			if string(board[r][c]) == "."{
				continue
			}else{
				val := board[r][c]
				sqrInd := (r/3) * 3 + c/3
				if rows[r][val] || cols[c][val] || sqr[sqrInd][val]{
					return false
				}
				rows[r][val] = true
				cols[c][val] = true
				sqr[sqrInd][val] = true
			}
		}	
	}
	return true
}
