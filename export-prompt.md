# InvestorBook

Welcome to InvestorBook, Inc! This is your first week at InvestorBook and you
are excited to build on top of the data. The dataset contains three tables:
Investors, Companies, and Investments.

## Your Task

Build a simple web server that returns an Excel export for a given company's
data.

- GET /companies/:companyId/export
  - Header(top of each page)
    - Row 1, Column A is `{Name of Company} Export`.
    - Row 2, Column B is `Generated on {date}`.
    - Row 3 is blank.
    - Row 4 has the column headers.
    - The background color of the first two header rows should be RGB(92, 128,
      185).
    - The text color in the first two header rows should be white.
    - ![Example](header-screenshot.png "Example")
  - Sheet 1 - Investors
    - In this sheet, list all of the investors that have invested in this
      company.
    - Column A is the name of the investor.
    - Column B is the total amount invested.
  - Sheet 2 - Related Companies
    - List the companies that are connected to this company through at least one
      shared investor.
    - Column A is the name of the related company.
    - Column B is the list of names of the shared investors

README.md has instructions for connecting to the database.

## Submission

Fork this repository and send us a link to your fork after pushing your changes.

## Evaluation Criteria

- Clean, modular, easy-to-maintain code
- Test coverage
  - Bonus points for tests that are short, fast, and reliable but offer an
    unexpectedly high ROI in terms of coverage.
- Performance - Aim for a response time of less than 100ms.
