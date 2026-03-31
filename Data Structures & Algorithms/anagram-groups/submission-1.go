func stringToKey(s string) [26]int {
    var key [26]int
    runes := []rune(s)
    for _, r := range runes {
        key[r - 'a'] += 1
    }
    return key
}

func groupAnagrams(strs []string) [][]string {
    // character frequency
    anagramMap := make(map[[26]int][]string)

    for _, str := range strs {
        key := stringToKey(str)
        anagramMap[key] = append(anagramMap[key], str)
    }

    var out [][]string
    for _,v := range anagramMap {
        out = append(out, v)
    }
    return out
}