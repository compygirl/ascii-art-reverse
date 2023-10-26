## ascii-art-reverse by bbilisbe and ayessenb

### Objectives

Ascii-art-reverse is a program which consists in receiving a text file containing a graphic representation of a random string as an argument and  converting the graphic representation into a text.

The argument will be a flag, --reverse=<fileName>, in which --reverse is the flag and <fileName> is the file name. The program must then print this string in normal text.


- This project should handle an input with numbers, letters, spaces, special characters and `\n`.
- Usage: go run . [OPTION]
- Ascii-art-output project implemented, the program accept the correctly formatted [OPTION] [STRING] [BANNER].
- Ascii-art-fs project implemented, the program accept the correctly formatted [STRING] [BANNER].
- The program must still able to run with a single [STRING] argument.


### Banner Format

* shadow
* standard
* thinkertoy


### Usage

```console
student$ cat file.txt
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
student$ go run . --reverse=file.txt
hello
$

```

! Note that the test file has to have one extra new line as shown in the example.

### Allowed packages

- Only the [standard Go](https://golang.org/pkg/) packages are allowed
