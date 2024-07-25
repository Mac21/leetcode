def lengthOfLongestSubstring(s: str) -> int:
    length = len(s)

    if length < 2:
        return length

    i = 0
    j = 1

    # Mapping of sub string -> (ending index of the key substr, length of the substring)
    substr_index_len_map = {
    }

    substring = ""
    while j < length:
        if s[i] not in substring:
            substring += s[i]
            i = j

        if s[j] not in substring:
            substring += s[j]
            j += 1

        if i < length and j < length:
            printDebug(i, j, s[i], s[j], substring, prefix="s[i] and s[j]")

        if j >= length or s[j] in substring:
            if substring not in substr_index_len_map:
                substr_index_len_map[substring] = (j-1, len(substring))
            j = i+1
            substring = ""

    maxLen = 0
    for tup in substr_index_len_map.values():
        maxLen = max(tup[1], maxLen)

    return maxLen


def fastest_impl(s: str) -> int:
    last_used_index = {}
    max_length = start = 0
    for i, c in enumerate(s):
        if c in last_used_index and start <= last_used_index[c]:
            start = last_used_index[c] + 1
        else:
            max_length = max(max_length, i - start + 1)

        last_used_index[c] = i

    return max_length


def printDebug(i, j, si, sj, substring, prefix=None):
    if prefix:
        print(f'\t{prefix}')
    print(f'\ti: {i}')
    print(f'\tj: {j}')
    print(f'\ts[i]: {si}')
    print(f'\ts[j]: {sj}')
    print(f'\tsubstr: {substring}')


def testImpl(s: str, answer: int) -> bool:
    print(f'Running test for {s}:')
    result = fastest_impl(s)
    fstr = 'Test of {s} {passfail}: got {result} expected {answer}\n'
    sameAnswers = result == answer
    if sameAnswers:
        print(fstr.format(s=s, passfail="passed", result=result, answer=answer))
    else:
        print(fstr.format(s=s, passfail="failed", result=result, answer=answer))


def main():
    testImpl("abcabcbb", 3)  # answer is 'abc'
    testImpl("bbbbb", 1)  # answer is 'b'
    testImpl("pwwkew", 3)  # answer is 'wke'
    testImpl(" ", 1)  # answer is ' '


if __name__ == '__main__':
    main()
