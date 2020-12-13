import re

PASSWORD_REGEX = re.compile(
    r"(?P<minimum>\d+)-(?P<maximum>\d+) (?P<letter>[a-z]): (?P<password>[a-z]+)"
)


def main():
    data = open("input.txt").read()

    for line in data.split("\n"):
        m = PASSWORD_REGEX.search(line)

        if m:
            print(m)
        else:
            print('not found')
    

if __name__ == "__main__":
    main()
