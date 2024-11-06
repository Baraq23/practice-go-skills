## ASCII Art Color Program in Go!
![Go-Logo_Blue-removebg-preview(2)](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)
### Overview

This project is an ASCII art program implemented in Go. It takes a string as input with color using color flag and outputs a colored graphic representation of the string using ASCII characters.
### Content
- banners
    - shadow.txt
    - standard.txt
    - thinkertoy.txt
- functions/
    - formatarguments.go
    - getansi.go
    - printart.go
    - readbanner.go
    - substringpositions.go   
    - validatebannerfile.go
    - validate.go
    - worddistributors.go
- testFiles
    - functoins_test.go
- go.mod
- LICENCE
- main.go
- README.md


### Features

- This project is written in go lang.
- Handles input with ascii values between ascii values 32 to 126, and *`does not`* include special characters such as `newline` and `tab`.
- Uses graphical templates such as standard banner file, shadow banner file and also thinkertoy banner file.
- Uses `--color` flag with which users can define the color they want to see the graphics on.
    - the color flag takes upto two colors defined.
    - the color format can be in `rgb` format or `color name`.
    - there is an option of coloring only part of the string by providing substring.
- Includes test files for testing.

### Usage

To run the program,follow these steps:

* clone the program into your computer

```bash
https://learn.zone01kisumu.ke/git/egathang/ascii-art-color.git
```

* Open the cloned folder with your prefered text editor among VScode and Vim.

* In the terminal from your text editor type the following:
```bash
go run . -color="green" "type string here" | cat -e
```

* Hit Enter. The corresponding word art will be printed in the console.

## Output

The following shows input and results when the program is run:

```bash
:~ go run . --color="green" "Hello" | cat -e
```
<pre><span style="color: green">
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
</span></pre>

```bash
go run . --color="orange,blue" "hello" "hello world"
```

<pre>
<span style="color:orange"> _              _   _           </span>  <span style="color:blue">                              _       _ </span>
<span style="color:orange">| |            | | | |          </span>  <span style="color:blue">                             | |     | |</span> 
<span style="color:orange">| |__     ___  | | | |   ___    </span>  <span style="color:blue">   __      __   ___    _ __  | |   __| |</span> 
<span style="color:orange">|  _ \   / _ \ | | | |  / _ \   </span>  <span style="color:blue">   \ \ /\ / /  / _ \  | '__| | |  / _` |</span> 
<span style="color:orange">| | | | |  __/ | | | | | (_) |  </span>  <span style="color:blue">    \ V  V /  | (_) | | |    | | | (_| |</span> 
<span style="color:orange">|_| |_|  \___| |_| |_|  \___/   </span>  <span style="color:blue">     \_/\_/    \___/  |_|    |_|  \__,_|</span> 
<span style="color:orange">                                </span>  <span style="color:blue">                                        </span> 
</pre>      



### Packages

Only standard Go packages are used in this project.

### How to use the program
- all input must be written in this format:
```bash
:~ go run . --color=<color> <substring to be colored> "something"
```

_Example_
- `For common colors`:
```bash
go run . --color="grey" "hello" "hello world"
```
- `For rgb format`:
```bash
go run . --color="rgb(200, 90, 90)" "hello" "hello world"
```

- The following banner formats are included:
  - `shadow`    -sh
  - `standard`  -st
  - `thinkertoy`-th
 - To use any type the banners above:
    specify its corresponding flag as follows:
    -sh, -st and -th respectively.

 _Example_
```bash
go run . --color="teal,yellow" "hello" "hello world" -sh
``` 
<pre>
<span style="color:teal">                                  </span>   <span style="color:yellow">                                                     </span>
<span style="color:teal">_|                _| _|           </span>   <span style="color:yellow">                                         _|       _| </span>
<span style="color:teal">_|_|_|     _|_|   _| _|   _|_|    </span>   <span style="color:yellow">    _|      _|      _|   _|_|   _|  _|_| _|   _|_|_| </span>
<span style="color:teal">_|    _| _|_|_|_| _| _| _|    _|  </span>   <span style="color:yellow">    _|      _|      _| _|    _| _|_|     _| _|    _| </span>
<span style="color:teal">_|    _| _|       _| _| _|    _|  </span>   <span style="color:yellow">      _|  _|  _|  _|   _|    _| _|       _| _|    _| </span>
<span style="color:teal">_|    _|   _|_|_| _| _|   _|_|    </span>   <span style="color:yellow">        _|      _|       _|_|   _|       _|   _|_|_| </span>
<span style="color:teal">                                  </span>   <span style="color:yellow">                                                     </span>
</pre>

- Standard banner template will be automatically selected if user has not entered any banner template flag.

### Error handling
- The program notifies if the template bannerfile is corrupted or empty. In this case the user should download a new banner template from the following links;
    - <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt">shadow</a>
    - <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt">standard</a>
    - <a href="https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt">thinkertoy</a> 

  
### Collaborators


<table>
<tr>
<td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/elijah-gathanga-88b601262/ >
            <img src=https://learn.zone01kisumu.ke/git/avatars/728b6eaf4ffc61a682b2f84a2b0d6d3a?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Elijah/>
            <br />
            <sub style="font-size:14px"><b><i>Elijah Gathanga</i></b></sub>
        </a>
  </td> 
  <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/barrack-kope-064a43244/ >
            <img src=https://learn.zone01kisumu.ke/git/avatars/69ae4e7472c685f60beaf04edb53b624?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Barrack/>
            <br />
            <sub style="font-size:14px"><b><i>Barrack Otieno</i></b></sub>
        </a>
  </td> 
   
</tr>

</table>


# Goals
This project notions include:
- The Go file system(fs) API
- Color converters
- Data manipulation
- Terminal display


## HOW TO CONTRIBUTE ? 👷 

**1.** Clone [this](https://learn.zone01kisumu.ke/git/egathang/ascii-art-color.git) repository.

**2.** Clone the forked repository.

```terminal
git clone https://learn.zone01kisumu.ke/git/egathang/ascii-art-color.git
```

**3.** Navigate to the project directory.

```terminal
cd ascii-art-color
```

**4.**  MAKE A NEW FOLDER WITH YOUR PROJECT NAME INSIDE 
<br>

**5.**  Also Add a README file in your project folder which consists of Description/screenshots about your project !          
 
<br>

**6.** Create a new branch.

```terminal
git checkout -b <your_branch_name>
```

**7.** Add & Commit your changes.

```terminal
  git add .
  git commit -m "<your_commit_message>"
```

  Push your local branch to the remote repository.

```terminal
git push -u origin <your_branch_name>
```

**8.** Create a Pull Request!

**Review:** At this stage, we will review PR and merge your code into our codebase after approval by our team.

### Issues and Contributions

If you encounter any issues or have suggestions for improvements, feel free to submit an issue or propose a change.

## _____________________

Feel free to explore the codebase, run the program with different inputs, and contribute to enhancing the project.