# User REST API (L2, Golang | Ruby)

Hi there **Mayank**,

Hope all is well!

Here's your take-home project from Raksul.

Please submit your code before **Friday, 04 Dec 19:37 JST**.

Best wishes,

// The Raksul HR team.

:heartpulse: :two_hearts: :revolving_hearts: :sparkling_heart:<br/>
:sparkles: :sparkles: :sparkles:<br/>
:joy_cat: :star: :star2: :beers: :star: :star2: :joy_cat:<br/>
<br/>

# Instructions

Create a small REST API to handle users.

Limits:

- This challenge should take **max 4-6 hours** to complete.
- Your submission _must_ be less than 1,000 LOC (irrespective of chosen language)!
- Do _NOT_ fork this repo.
- Do _NOT_ merge your PR to main.

# Requirements:

1. Create an HTTP server that listens on `:9001`.

1. Create a `POST /users` route that creates a user, i.e performs a signup.

   Example request:

   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '<JSON payload below>' http://localhost:9001/users
   ```

   ```js
   {
     "email": "ace@base.se", // required, email format
     "password": "open1234"  // required, min 8 chars
   }
   ```

   Fields _must be_ validated as per the comments above.

   Example response:

   ```js
   {
     "id": 1000, // must be a number!
     "email": "ace@base.se",
     "created_at": "2020-09-15T19:56:10+09:00", // must be ISO 8601 format, e.g '2014-01-01T23:28:56.782Z' (UTC) or '2014-02-01T09:28:56+09:00' (GMT+9)
     "updated_at": null
   }
   ```

1. Create a `POST /login` route that performs a user login (verifies a user's password) and returns an access token.

   The token _must be_ an MD5 hash of some random value, e.g. `418638191e2b6c1d5ae24b471062dc03`.

   The token _must be_ stored in the database _in the same table as the user_.

   The token _must not be_ a JWT!

   Example request:

   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '<JSON payload below>' http://localhost:9001/login
   ```

   ```js
   {
     "email": "ace@base.se", // required, email format
     "password": "open1234"  // required, min 8 chars
   }
   ```

   Fields _must be_ validated as per the comments above.

   Example response:

   ```js
   {
     "token": "418638191e2b6c1d5ae24b471062dc03"
   }
   ```

1. Create a `GET /secret` route that returns a secret string to a logged in user.

   The token _must be_ passed as this HTTP header: `Authorization: Bearer <token>`

   The secret string _must be_ this value: `All your base are belong to us`

   Example request:

   ```bash
   curl -X GET \
     -H 'Content-Type: application/json' \
     -H 'Authorization: Bearer <token>' \
     http://localhost:9001/secret`
   ```

   Example response:

   ```js
   {
     "user_id": 1000, // User id of token owner!
     "secret": "All your base are belong to us"
   }
   ```

   Example error response:

   Calling this endpoint without a token (or a bad token) _must_ return the following error:

   ```js
   {
     "error":  "token invalid"
   }
   ```

1. Create a `PATCH /users/{id}` route updates user with id `{id}`.

   A user _must be_ able to update one or more fields in a single call, e.g. update only `email` or `password` or both.

   Example request:

   ```bash
   curl -X PATCH \
     -H 'Content-Type: application/json' \
     -d '<JSON payload below>' \
     http://localhost:9001/users/1000`
   ```

   ```js
   {
     "email": "base@base.se", // optional, email format
     "password": "close1234"  // optional, min 8 chars
   }
   ```

   Fields are _optional_ but if present _must be_ validated as per the comments above.

   Example response:

   ```js
   {
     "id": 1000,
     "email": "base@base.se",
     "created_at": "2020-09-15T19:56:10+09:00",
     "updated_at": "2020-09-15T20:00:30+09:00" // Set with each update!
   }
   ```

1. Your server _must_ accept the following ENV vars:

   ```bash
   HOST=:9001        # bind to port 9001 on all IPs.
   DB_HOST=localhost # connect to postgres on this host
   DB_USER=postgres  # postgres user
   DB_PASS=example   # postgres password
   DB_NAME=userapi   # database schema name
   ```

1. Validation errors _must be_ returned as with follows:

   Example error response:

   ```js
   {
     "error": "validation error: password" // e.g. when password isn't at least 8 chars long
   }
   ```

1. We've provided a `test-suite.sh` to _help you ensure_ that your API functions as expected.

   **Important Notes:**

   - Beware; Not all edge cases are covered by the given test suite. We _will_ run our own, more extensive test suite against your program!
   - You need a `bash` shell with `jq` installed run `test-suite.sh`.

1. Store users and tokens in **PostgreSQL**.

   Use the provided `docker-compose.yml` to start a local Postgres v11 server. Your version of Postgres _must be_ 11.x.

   Create a database named `userapi`.

   Create a table as follows:

   ```sql
   CREATE TABLE users ( ... );
   ```

   It's up to you to decide what data types to use.

   Remember to _create indexes_ on the columns that need them.

1. Create a `Dockerfile` that builds your server.

   We've provided an example `Dockerfile` for you to draw inspiration from.

   Your build _must be_ reproducable by us and you _must_ build / compile your server 100% inside Docker including fetching dependencies.

   **We will build and run your API server as follows:**

   1. Uncomment the lines in `docker-compose.yml`
   1. Run `docker-compose build`
   1. Run `docker-compose up`
   1. Run `test-suite.sh` against your server.
   1. Run our internal automated test suite against your server.

1. You can use _any_ of the following languages to implement your API server:

   - Golang
   - Ruby

# Bonus Points! (If you want to impress us and have time to spare!)

You will score extra points for the following things:

1. `PATCH /users/{id}` _must be_ require a token (i.e. same as the `GET /secret` route) and allow users to _update only their own accounts_.

   If a user tries to update another user, your server _must_ return the following error:

   ```js
   {
     "error":  "access denied"
   }
   ```

1. `POST /login` returns a token _that expires after 3 seconds_.

   Calling an endpoint with an expired token _must_ return the following error:

   ```js
   {
     "error":  "token expired"
   }
   ```

1. Write at least _one good e2e test_!

# Important Notes

- You _must_ drop and recreate your entire database _every time_ you start your API server.

- The password must be _stored securely_ following industry best practices.

- Avoid using an ORM library/framework if possible.

- You're _not required_ to provide tests, but you'll definitely score extra points if you do.

# How To Submit Your Code

1. Create a feature branch named `challenge` and submit your code as a PR.

   Once you're ready, you _must request a PR review_ from the user `raksul-code-review`.

   **Important:**

   - Do NOT fork this repo.
   - Do NOT merge your PR to main.
   - Do NOT include vendor folders.
   - Keep your own submission below 1,000 LOC total!

# Grading & What To Expect

1. We will grade you on code clarity, naming, code structure, project size and, most importantly, _correctness_.

   Our own internal reference solution implements everything in _~900 LOC of Go code_, including:

   - An extensive e2e test suite
   - Proper comments
   - Good formatting and spacing
   - Code layering, i.e. decoupling using interfaces, etc.

   Try to keep your submission as small and as compact as possible and remember: _KISS_ and _YAGNI_.

   **IMPORTANT: For this challange you will only be informed if you pass or not. We will NOT provide any additional feedback, so please do not ask for it!**

   **Good Luck >:)**
