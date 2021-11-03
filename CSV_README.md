_This is one of the steps in the Truss interview process. If you've
stumbled upon this repository and are interested in a career with
Truss, [check out our jobs page](https://truss.works/jobs)._

# The problem: CSV normalization

## Introduction

Approach this solution the way you would a real world problem. Feel free
to use libraries where it makes sense to. Be prepared to explain your
thought process around your submision.

Please complete as much of the exercise you can within 3-4 hours.
Do not spend more than 4 hours on this exercise.
We welcome a working but incomplete solution.

If you have any questions, please contact [hiring@truss.works](mailto:hiring@truss.works); we're
happy to help.

We will be excited to review and discuss your code over video chat no matter the level of polish of your submission.

## What we want to learn about you from this challenge

This question is intentionally open-ended to scale to your level of comfort with the problem or technologies.
What you focus on and how you structure your code will tell us a lot about you and how you use technology.

We really want to know how you approach a problem and how your solution addresses it.
We want to know how you handle technical choices and communicate your reasoning behind your responses to them.

We know not all of your technical depth may be exercised by this problem alone.

With the above in mind, here are some tips:

* Ask us questions. We will try our best to answer your questions quickly over email.
* Use a language you are comfortable in. We’ll figure out how to run it one way or another.
* Use the internet liberally. We want this to be a reflection of what you can do day to day. It would be ridiculous if we asked you to code at work without the internet.
* Focus on getting something to share with us. This is the starting point for our technical conversation.
* Take notes on things you would like to follow up on if you had more time. We will be excited to discuss these points during our synchronous discussion.
* Taking the time to include helper scripts or instructions on how to get your project set up will be extremely helpful. Please count any work around this towards your limit of 4 hours.

If you have any time constraints for completing this assignment please reach out to the hiring team.
We will attempt to accommodate your reasonable requests.

## How to submit your response

Please submit your solution by emailing [hiring@truss.works](mailto:hiring@truss.works) with a link to a public git repository
(Github is fine) that contains your code and a README.md that tells us
how to build and run it. Your code will be run on either macOS 11.2+
or Ubuntu 20.04 LTS, your choice. Note, if your development machine is
running Windows, we suggest testing your response in a Docker image based on
Ubuntu or Alpine linux.

## Full Description

Please write a tool that reads a CSV formatted file on `stdin` and
emits a normalized CSV formatted file on `stdout`. For example, if
your program was named `normalizer` we would test your code on the
command line like this:

```sh
./normalizer < sample.csv > output.csv
```

Normalized, in this case, means:

* The entire CSV is in the UTF-8 character set.
* The `Timestamp` column should be formatted in RFC3339 format.
* The `Timestamp` column should be assumed to be in US/Pacific time;
  please convert it to US/Eastern.
* All `ZIP` codes should be formatted as 5 digits. If there are less
  than 5 digits, assume 0 as the prefix.
* The `FullName` column should be converted to uppercase. There will be
  non-English names.
* The `Address` column should be passed through as is, except for
  Unicode validation. Please note there are commas in the Address
  field; your CSV parsing will need to take that into account. Commas
  will only be present inside a quoted string.
* The `FooDuration` and `BarDuration` columns are in HH:MM:SS.MS
  format (where MS is milliseconds); please convert them to the
  total number of seconds.
* The `TotalDuration` column is filled with garbage data. For each
  row, please replace the value of `TotalDuration` with the sum of
  `FooDuration` and `BarDuration`.
* The `Notes` column is free form text input by end-users; please do
  not perform any transformations on this column. If there are invalid
  UTF-8 characters, please replace them with the Unicode Replacement
  Character.

Safe Assumptions:

* The input document is in UTF-8, although some characters may be incorrectly encoded.
* Invalid characters can be replaced with the Unicode Replacement Character. If that replacement makes data invalid (for example, because it turns a date field into something unparseable), print a warning to `stderr` and drop the row from your output.
* Times that are missing timezone information are in `US/Pacific`.
* The sample data we provide contains all date and time format variants you will need to handle.
* Any type of line endings are permissible in the output.
