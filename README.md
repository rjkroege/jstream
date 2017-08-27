` jstream` is a simple tool that creates a JSON version of a
relational join
 over the input files

 An example will clarify. Given files letters and numbers (i.e. the
 ones checked into the repository) Then:

```
./jstream Numbers numbers Letters letters
```

will create an output JSON file that looks like this:

```
[
  {
    "Letters": "a",
    "Numbers": "1"
  },
  {
    "Letters": "b",
    "Numbers": "2"
  },
  {
    "Letters": "c",
    "Numbers": "3"
  }
]
```

Note how we end up with an array of objects, each object has the keys
`Letters` and `Numbers` that correspond to the command line argument
while the values of the keys in the array correspond to the lines read
from the input files.
