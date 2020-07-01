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

We will review your solution in a recent version of Firefox, Chrome, or Safari.

If you have any questions, please contact hiring@truss.works; we're
happy to help if you're not sure what we're asking for or if you have
questions.

## How to submit your response

Please submit your solution by emailing hiring@truss.works a link to **one** of the following:

- A public git repository (Github is fine) that contains your code and a `README.md` that tells us how to build and run it.
- Your solution running in an online IDE such as [CodePen](https://codepen.io) or [Glitch](https://glitch.com).

## The problem: Display data from an API for human consumption

Please write a static webpage that loads data from `https://swapi.dev/api/planets/` and displays that data
in an HTML `table`.

Name the entrypoint for your application `index.html`. 

You may use additional files to organize your code. You may also use frameworks and build tools as long as clear, simple instructions are provided for how to run. Please ensure your solution can be easily viewed in the browser. Be ready to answer questions about why you chose a given approach.


Don't worry about making the table look fancy; you can style it however you
would like to outside the requirements below. We will look at the markup you
generate and any code you write to load or display the data.

## What to display

For each planet in the dataset, please display:

- The planet's name
    - The name should be a link that, when clicked, opens the planet's API URL in a new window
- The planet's climate
- How many residents the planet has
- The terrains found on the planet
- The population
- The surface area covered by water
    - Assume that all planets are perfect spheres.
    - The radius of a sphere is half its diameter.
    - The value of `surface_water` from the API is a percentage, so a value of `50` means the planet is 50% covered in water.

Additionally, please satisfy these requirements:

- Show a loading message while loading the data, and hide this message once the data is displayed.
- Display an error message if the data load fails for some reason. We may test your code against an invalid URL to demonstrate this.
- Sort the planets alphabetically
- Any column that is unknown should be displayed as `?`.
- Format all numbers larger than 999 with spaces to group digits into groups of three.
  For example, ten thousand should be displayed as `10 000`.
- Cells in the table should be separated by 1 pixel gray lines.

## Stretch goals (optional)

If you complete the above and still have additional time, you can choose to complete one or more of the following extra requirements.

**This is absolutely not required for submitting the work sample and should not be worked on beyond the four hour time limit.**

- The API only returns the first page of data by default. Add a "load more" button that
  loads additional pages of data when clicked. Only show this button if there is more data to be loaded.
- Sort the table by a column's values when that column's header is clicked.
- Spruce up your page with some extra styling & design elements!
- Make sure your page is usable on mobile devices.
- Include some tests.
