package main

var stages []string = []string{`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========
`, `
  +---+
  |   |
  O   |
      |
      |
      |
=========
`, `
  +---+
  |   |
      |
      |
      |
      |
=========
`}

var logo string =`
 _                                             
| |                                            
| |__   __ _ _ __   __ _ _ __ ___   __ _ _ __  
| '_ \ / _| | '_ \ / _| | '_ \ _ \ / _| | '_ \ 
| | | | (_| | | | | (_| | | | | | | (_| | | | |
|_| |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|
                    __/ |                      
                   |___/    
`