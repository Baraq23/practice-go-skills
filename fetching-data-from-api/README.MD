# learning to fetch data from api


### links

ARTIST API LINK : https://groupietrackers.herokuapp.com/api/artists

### Displaying json data in the terminal

To use curl to retrieve the structure of a JSON response from an API or a URL, you can follow this approach. The curl command fetches the JSON data, and you can combine it with jq to pretty-print the structure. Here's the basic command:

```bash
curl -s https://api.example.com/data | jq .
```
- -s: Silent mode, prevents curl from showing progress output.
- jq .: Pretty-prints the JSON data, allowing you to view its structure.

If you don't have jq installed, you can simply use curl to get the raw JSON response like this:

```bash
curl -s https://api.example.com/data
```

However, without jq, the output will not be formatted, so it might be harder to read.




# Some notes on styling (CSS)

### 