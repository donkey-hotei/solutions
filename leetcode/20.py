#!/usr/bin/python
def matching(s, e):
    if s == "}":
        return e == "{"
    elif s == ")":
        return e == "("
    elif s == "]":
        return e == "["


def is_valid(s):
    stack = []

    for c in s:
        if c in ["{", "(", "["]:
            stack.append(c)
        elif c in ["}", ")", "]"]:
            if len(stack) == 0:
                return False
            last = stack.pop()
            if not matching(c, last):
                return False

    return len(stack) == 0


def main():
    assert is_valid("()") is True
    assert is_valid("{}[]()") is True
    assert is_valid("(]") is False
    assert is_valid("([)]") is False
    assert is_valid("{[]}") is True
    assert is_valid("]") is False

    print "tests pass."


if __name__ == "__main__":
    main()
