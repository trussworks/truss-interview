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
how to build and run it. Your code will be run in Firefox, Chrome, or Safari.

## The problem: Display data from an API for human consumption

Please write a static webpage that loads data from `https://swapi.dev/api/planets/` and displays that data
in an HTML `table`.

Name the entrypoint for your application `index.html`. To test your solution, we will open `index.html` in a browser.

You may use additional files to organize your code, but please include submit
a built version of the project that doesn't require the use of any build tools
if you do. Accordingly, you may use any build tool you'd like as long as you
include files that are ready for viewing directly in a browser.

Don't worry about making the table look fancy; you can style it however you
would like to outside the requirements below. We will look at the markup you
generate and any code you write to load or display the data.

## What to display

For each planet in the dataset, please display:

- The planet's name
- The planet's climate
- How many residents the planet has
- The terrains found on the planet. Please use an unordered list (`ul`) to structure these and wrap each terrain in a list item (`li`).
- The population
- The surface area covered by water
    - Assume that all planets are perfect spheres.
    - The radius of a sphere is half its diameter.
    - The `surface_water` value in the data is a percent, so `50` means the planet is 50% covered in water.

Additionally, please satisfy these requirements:

- Show a loading message while loading the data, and hide this message once the data is displayed.
- Sort the planets by name, showing earlier letters first.
- Any column that is unknown should be displayed as `?`.
- Format all numbers larger than 999 with spaces to group digits into groups of three.
  For example, ten thousand should be displayed as `10 000`.
- Cells in the table should be separated by 1 pixel gray lines.

## Stretch goals (optional)

If you complete the above and still have additional time, you can choose to complete one or more of the following extra requirements.

**This is absolutely not required for submitting the work sample.**

- The API only returns the first page of data by default. Add a "load more" button that
  loads additional pages of data when clicked. Only show this button if there is more data to be loaded.
- Sort the table by a column's values when that column's header is clicked.