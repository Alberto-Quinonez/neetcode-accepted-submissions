class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = defaultdict(set)
        cols = defaultdict(set)        
        boxes = defaultdict(set)       
        for i,row in enumerate(board):
            for j,col in enumerate(board[i]):
                if col == ".":
                    continue
                #check if new value is in rows set already
                if col in rows[i]:
                    return False
                rows[i].add(col)

                if col in cols[j]:
                    return False
                cols[j].add(col)
                
                #sub box corrds
                rowCoord = i // 3
                colCoord = j // 3

                #check if new value is in subbox set already
                if col in boxes[(rowCoord,colCoord)]:
                    return False
                boxes[(rowCoord,colCoord)].add(col)
        return True        
