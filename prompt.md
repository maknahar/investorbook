# InvestorBook

Welcome to InvestorBook, Inc! This is your first week at InvestorBook and you
are excited to build on top of the data. The dataset contains three tables:
Investors, Companies, and Investments.

## Your Task

You met with the CEO and promised to deliver three critical API endpoints in
four hours:

- GET /investors/:investorId/connections
  - Returns a list of 1st-degree connections for this investor.
- GET /investors/:investorId/mutual/:otherInvestorId
  - Returns a list of mutual connections between two investors.
- GET /investors/:investorId/search?q={searchString}
  - Returns all investors where the investor name is an exact, case-insensitive
    substring match on searchString
  - Include the degree of connection to investorId for each investor

## Definitions

- Connection - Two investors are connected if they have both invested in the
  same company.
- Degree of connection - This is a measure of connection between two investors.
  - 1st degree connection - Investors who have invested in the same company.
  - 2nd degree connection - For a given investor, 2nd degree connections are the
    1st degree connections of their 1st degree connections.
  - 3rd-degree connections - For a given investor, 3rd degree connections are
    the 1st-degree connections of their 2nd-degree connections. This feature is
    identical to LinkedIn connections:
    [reference](https://www.linkedin.com/help/linkedin/answer/110/your-network-and-degrees-of-connection?lang=en)
- Mutual connection - A given nvestor A is a mutual connection of investor B and
  investor C if investor A is a 1st degree connection of both investor B and
  investor C.

## Submission

Fork this repository and send us a link to your fork after pushing your changes.

## Evaluation Criteria

- Clean, modular, easy-to-maintain code
- Test coverage
  - Bonus points for tests that are short, fast, and reliable but offer an
    unexpectedly high ROI in terms of coverage.
- Performance - Aim for a response time of less than 100ms for each endpoint.
