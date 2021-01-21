function findFirstStringInBracket(string $string): string {
    return preg_match('~\((.*?)\)~', $string, $match,) ? $match[1] : '';
}