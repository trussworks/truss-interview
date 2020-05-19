_This is one of the steps in the Truss interview process. If you've
stumbled upon this repository and are interested in a career with
Truss, [check out our jobs page](https://truss.works/jobs)._

# Truss Software Engineering Interview

## Introduction and expectations

Hi there! Please complete the problem described below to the best of
your ability, using the tools you're most comfortable with. Assume
you're sending your submission in for code review from peers;
we'll be talking about your submission in your interview in that
context.

We expect this to take less than 4 hours of actual coding time. Please
submit a working but incomplete solution instead of spending more time
on it. We're also aware that getting after-hours coding time can be
challenging; we'd like a submission within a week and if you need more
time please let us know.

Approach this solution the way you would a real world problem. Use
libraries where it makes sense to, and be prepared to explain your
thought process if you do.

If you have any questions, please contact hiring@truss.works; we're
happy to help if you're not sure what we're asking for or if you have
questions.

## How to submit your response

Please send hiring@truss.works a link to a public git repository
(Github is fine) that contains your code and a README.md that tells us
how to build and run it. Your code will be run on either macOS 10.15
or Ubuntu 16.04 LTS, your choice. If you choose the browser solution,
your code will be run in Firefox, Chrome, or Safari.

## The problem: CSV normalization

You can decide between writing your solution in code that will be run
on the command line or in a web browser. Pick one of the following methods
and follow the instructions below:

### Command Line (If you want us to run your code from the command line):

Please write a tool that reads a CSV formatted file on `stdin` and
emits a normalized CSV formatted file on `stdout`. For example, if
your program was named `normalizer` we would test your code on the
command line like this:

```sh
./normalizer < sample.csv > output.csv
```

### Web Browser (If you want us to run your code from a browser):

Please write a static webpage that normalizes and displays the data contained in
`sample.csv`. We have provided a function that will return the contents of this
file as plain text, which you can either include directly in the browser, or
import to your compiled JavaScript. If you would like to import it, you can
modify `getSampleData.js` to export the `getSampleData` function. For example:

```
// In-browser
<script src="getSampleData.js"></script>
<script>
  const data = getSampleData()
</script>
```

```
// As a JS module
import getSampleData from './getSampleData.js'
```

You should write JavaScript (or code that compiles to JavaScript) that takes the sample data and formats it accordingto the normalization rules listed below. Don't worry about making
it look too fancy; you can display and style the data however you wish,
but we will be looking for valid markup.

### Normalization rules:

Normalized, in this case, means:

- The entire data set is in the UTF-8 character set.
- The `Timestamp` column should be formatted in RFC3339 format.
- The `Timestamp` column should be assumed to be in US/Pacific time;
  please convert it to US/Eastern.
- All `ZIP` codes should be formatted as 5 digits. If there are less
  than 5 digits, assume 0 as the prefix.
- The `FullName` column should be converted to uppercase. There will be
  non-English names.
- The `Address` column should be passed through as is, except for
  Unicode validation. Please note there are commas in the Address
  field; your CSV parsing will need to take that into account. Commas
  will only be present inside a quoted string.
- The `FooDuration` and `BarDuration` columns are in HH:MM:SS.MS
  format (where MS is milliseconds); please convert them to the
  total number of seconds expressed in floating point format.
  You should not round the result.
- The `TotalDuration` column is filled with garbage data. For each
  row, please replace the value of `TotalDuration` with the sum of
  `FooDuration` and `BarDuration`.
- The `Notes` column is free form text input by end-users; please do
  not perform any transformations on this column. If there are invalid
  UTF-8 characters, please replace them with the Unicode Replacement
  Character.

You can assume that the input document is in UTF-8 and that any times
that are missing timezone information are in US/Pacific. If a
character is invalid, please replace it with the Unicode Replacement
Character. If that replacement makes data invalid (for example,
because it turns a date field into something unparseable), print a
warning to `stderr` and drop the row from your output (command line solution) OR
display an error message for that row instead (browser solution).

You can assume that the sample data we provide will contain all date
and time format variants you will need to handle.
