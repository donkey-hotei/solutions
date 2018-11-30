#!/usr/bin/python


def lcp(ls):
    if len(ls) < 1:
        return ""
    elif len(ls) == 1:
        return ls[0]

    min_len = len(min(ls, key=len))
    result = ""

    for i in range(min_len):
        curr_char = None
        for s in ls:
            if curr_char is None:
                curr_char = s[i]
            elif curr_char != s[i]:
                return result

        result += curr_char

    return result


def main():
    ls = ["flower", "flow", ""]
    print("{}".format(lcp(ls)))


if __name__ == "__main__":
    main()
