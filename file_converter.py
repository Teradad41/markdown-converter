import sys


import markdown


def main():
    if len(sys.argv) != 4:
        print('Usage: python script.py markdown input_file output_file')
        sys.exit(1)

    command, input_file, output_file = sys.argv[1], sys.argv[2], sys.argv[3]

    if command != 'markdown':
        print('unsupported command')
        sys.exit(1)

    with open(input_file, 'r') as file:
        md_content = file.read()

    md = markdown.Markdown(extensions=['tables'])
    html = md.convert(md_content)

    with open(output_file, 'w') as file:
        file.write(html)


if __name__ == '__main__':
    main()