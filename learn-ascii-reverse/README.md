## ASCII Art Reverse Program in Go!
![Go-Logo_Blue-removebg-preview(2)](https://github.com/makebelief/makebelief/assets/166484145/ad53f422-f338-4dd7-9ef1-ab772aa1fbb5)

### Overview

This project is an ASCII art program implemented in Go. It take a file with graphical representation of ascii characters then converts them into a normal text.

### Content
- ascii-fs-func
    - formatspecialcharacters.go
    - printart.go
    - start.go
    - validatestr.go
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
- Uses graphical templates such as standard banner file, shadow banner file and also thinkertoy banner file.
- The program can still display asciiart to the console when given a string,and with banner template given using the `--banner` flag.
    - options for banner templates include: `standard`, `thinkertoy` and `shadow`.
- Uses `--reverse` flag with which users can define write the fileName containing the art to be converted to normal text.
- The program can reverse any file written in any of the 3 banner styles.
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
go run . -reverse="testSt.txt"
```

* Hit Enter. The corresponding text will be printed in the console.

## Output

The following shows input and results when the program is run:

```bash
:~ go run . -reverse="testSt.txt"
```

_Output_

    Nokia is a tech giant.



### Packages

Only standard Go packages are used in this project.

### How to use the program
- all input must be written in this format:
```bash
:~ go run . -reverse=<fileName.txt
```


### Error handling
- The program notifies if the template bannerfile is corrupted or empty. In this case the user should download a new banner template from the following links and sve them in the banners folder;
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
- Data manipulation
- Terminal display


## HOW TO CONTRIBUTE ? 👷 

**1.** Clone [this](https://learn.zone01kisumu.ke/git/egathang/ascii-art-reverse.git) repository.

**2.** Clone the forked repository.

```terminal
git clone https://learn.zone01kisumu.ke/git/egathang/ascii-art-reverse.git
```

**3.** Navigate to the project directory.

```terminal
cd ascii-art-reverse
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