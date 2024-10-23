```markdown
# ASCII Art Generator

A Go-based tool for converting plain text into ASCII art using standard banner file with a specific graphical template representation.

## Features

- Handles multi-line inputs
- Converts numbers, letters, special characters and symbols

## How to Use

1. Clone the repository:
   ```
   git clone https://01.gritlab.ax/git/obalandi/ascii-art.git
   cd ascii-art
   ```

2. Build the application:
   ```
   go build
   ```

3. Run the application:
   ```
   ascii-art 'Your text here'
   ```

3.1. Running the without building:
   ```
   go run ascii-art 'Your text here'
   ```

4. Run the tests (only works with build application):
   ```
   go test
   ```


## Example Output

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```
