# Description

A simple web application in Go that converts user input text into ASCII art using different banner styles.


## Authors

- Ilyass Aboudou
- Ahmed Talbi
- Basma Toumanar

## Usage: how to run

1. Clone this repository:
git clone https://learn.zone01oujda.ma/git/iaboudou/ascii-art-web.git

2. Run the server:
go run main.go

3. Open your browser and navigate to:
http://localhost:8080

4. usage:
    - Enter your text in the textarea.
    - Select a banner style from the dropdown.
    - Click Submit.
    - View the ASCII art result on the new page.
    - Click Back to return to the form and try again.


## Implementation details: algorithm

1. Input:
    - The user types a message (e.g. "hello").
    - The selected banner style (e.g. "standard") determines which font file is used.
    - Characters must be in the printable ASCII range or a newline (\n).
    - Invalid characters result in an error returned to the frontend.
2. ascii-art app:
    - the ascii-art app open the selected file.
    - if the file doesn't exist, it is automatically downloaded it from the zone01oujda.ma API and made read-only
    - The font file is split into chunks, then stored into slice (from ASCII 32 to 126).
    - The input string is processed line by line.
    - all lines are combined into a final output string
3. server handling :
    - Extracts input and banner.
    - Calls AsciiArt (input, banner).
    - On success: stores result in a global Result variable and redirects to /artstyle. if there is an error re-renders the form and displays the error
    - Displays the stored ASCII result in a <pre>{{.Output}}</pre> block
    - Includes a Back button to return to /.