package main

import (
	"testing"
)

func TestConvertToAsciiArt(t *testing.T) {
	font := "standard.txt"
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty",
			input:    "",
			expected: ``,
		},
		{
			name:  "newline",
			input: "\\n",
			expected: `
`,
		},
		{
			name:  "Hello\\n",
			input: "Hello\\n",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

`,
		},
		{
			name:  "HeLlO",
			input: "HeLlO",
			expected: ` _    _          _        _    ____   
| |  | |        | |      | |  / __ \  
| |__| |   ___  | |      | | | |  | | 
|  __  |  / _ \ | |      | | | |  | | 
| |  | | |  __/ | |____  | | | |__| | 
|_|  |_|  \___| |______| |_|  \____/  
                                      
                                      
`,
		},
		{
			name:  "Hello There",
			input: "Hello There",
			expected: ` _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 
                                                                             
                                                                             
`,
		},
		{
			name:  "1Hello 2There",
			input: "1Hello 2There",
			expected: `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`,
		},
		{
			name:  "{Hello There}",
			input: "{Hello There}",
			expected: `   __  _    _          _   _                 _______   _                           __    
  / / | |  | |        | | | |               |__   __| | |                          \ \   
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ 
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  
  \_\                                                                              /_/   
                                                                                         
`,
		},
		{
			name:  "Hello\\nThere",
			input: "Hello\\nThere",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
		},
		{
			name:  "Hello\\n\\nThere",
			input: "Hello\\n\\nThere",
			expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
		},
		{
			name:  "\\n\\n\\n",
			input: "\\n\\n\\n",
			expected: `


`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := convertToAsciiArt(testCase.input, font)
			if got != testCase.expected {
				t.Errorf("convertToAsciiArt(%s) = %v, expected %v", testCase.input, got, testCase.expected)
			}
		})
	}
}
