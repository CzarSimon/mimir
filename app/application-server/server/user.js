const r = require("rethinkdb");
const _ = require('lodash');
const database = require('./database');
const { nowUTCString, isEmpty } = require('./helper-methods');

const USER_TABLE = "users";

/**
* newUser() Creates, stores and sends back a new user
* Takes request and response objects and database connection as paramters
*/
const newUser = (req, res, conn) => {
  user = createNewUser();
  storeNewUser(user, conn, (err, userWithId) => {
    if (!err) {
      res.status(200).send(userWithId);
    } else {
      res.status(500).send("Unable to create user")
    }
  });
}

/**
* storeNewUser() Stores a new user and returns the genrated user id
* Takes user, database connection and a callback as parameters
*/
const storeNewUser = (user, conn, callback) => {
  r.table(USER_TABLE).insert(user).run(conn, (err, res) => {
    // Add id to user if there was no error, set to empty object if the was one
    userWithId = (err) ? {} : Object.assign({}, user, {id: parseGeneratedKey(res)})
    callback(err, userWithId);
  })
}

// createNewUser() Creates a new user shell
const createNewUser = (email = "") => ({
  email,
  joinDate: nowUTCString(),
  tickers: _initalTickers(),
  searchHistory: [],
  sessions: [ nowUTCString() ]
})

// _initalTickers() Returns an inital set of ticker for a new user
const _initalTickers = () => (['AAPL', 'FB', 'TSLA', 'TSLA', 'AMZN'])

// parseGeneratedKey() Returns the first gnereated key for a new user creation
const parseGeneratedKey = dbRes => _.head(dbRes.generated_keys)


/**
* getUser() Gets a user from the database based on a supplied user id
* Takes request and response objects and database connection as paramters
*/
const getUser = (req, res, conn) => {
  const userId = req.query.id;
  if (isEmpty(userId)) {
    res.status(400).send("No user id supplied");
    return;
  }
  getUserFromDB(userId, conn, (err, user) => {
    if (!err) {
      res.status(200).send(user);
    } else {
      console.log(err);
      res.status(500).send("Unable to fetch user");
    }
  });
}

/**
* getUserFromDB() Retrives a user from the database
* Takes user id, database connection and a callback as parameters
*/
const getUserFromDB = (userId, conn, callback) => {
  r.table(USER_TABLE).get(userId).run(conn, (err, res) => {
    const user = (!err) ? res : {}
    const error = (res == null) ? `unable to find user ${userId}` : err
    callback(error, user);
  })
}


/**
* recordSession() Stores the start time of a user session
* Takes a request and response objects and database connection as paramters
*/
const recordSession = (req, res, conn) => {
  const userId = req.body.id;
  if (isEmpty(userId)) {
    res.status(400).send("No user id supplied");
    return;
  }
  prependFieldInDB(userId, nowUTCString(), "sessions", conn, err => {
    if (!err) {
      res.sendStatus(200);
    } else {
      console.log(err);
      res.status(500).send('unable to record session')
    }
  });
}


/**
* saveSearch() Stores a search query that a user has executed
* Takes a request and response objects and database connection as paramters
*/
const saveSearch = (req, res, conn) => {
  const { id: userId, query } = req.body;
  if (isEmpty(userId) || isEmpty(query)) {
    res.status(400).send("No user id or query supplied");
    return;
  }
  prependFieldInDB(userId, query, "searchHistory", conn, err => {
    if (!err) {
      res.sendStatus(200);
    } else {
      console.log(err);
      res.status(500).send('unable to save search query')
    }
  });
}


/**
* addTicker() Stores new ticker that the supplied user has chosen to follow
* Takes a request and response objects and database connection as paramters
*/
const addTicker = (req, res, conn) => {
  const { id: userId, ticker } = req.body;
  if (isEmpty(userId) || isEmpty(ticker)) {
    res.status(400).send('No user id or ticker supplied');
    return;
  }
  getTickers(userId, conn, (err, tickers) => {
    if (!_.includes(tickers, ticker)) {
      prependFieldInDB(userId, ticker, 'tickers', conn, err => {
        if (!err) {
          res.sendStatus(200);
        } else {
          console.log(err);
          res.status(500).send('unable to add ticker')
        }
      });
    } else {
      res.sendStatus(200);
    }
  });
}

/**
* deleteTicker() Deleates a ticker that the supplied user has chosen to no longer track
* Takes request and response objects and database connection as paramters
*/
const deleteTicker = (req, res, conn) => {
  const { id: userId, ticker } = req.body;
  if (isEmpty(userId) || isEmpty(ticker)) {
    res.status(400).send('No user id or ticker supplied');
    return;
  }
  removeTickerFromDB(userId, ticker, conn, err => {
    if (!err) {
      res.sendStatus(200);
    } else {
      console.log(err);
      res.status(500).send('unable to delete ticker')
    }
  });
}


/**
* removeTickerFromDB() Removes a ticker from a user in the database
* Takes user id, ticker, a database connection and a callback as parameters
*/
const removeTickerFromDB = (userId, ticker, conn, callback) =>  {
  r.table(USER_TABLE).get(userId).update({
    "tickers": r.row("tickers").difference([ticker])
  })
  .run(conn, (err, res) => {
    callback(err);
  })
}

/**
* getTickers() Gets all current tickers of the user with the supplied user id
* Takes user id, database connection and a callback as parameters
*/
const getTickers = (userId, conn, callback) => {
  r.table(USER_TABLE).get(userId)('tickers').run(conn, (err, res) => {
    callback(err, res);
  })
}

/**
* prependFieldInDB() Prepends an field with a supplied value
* for a given user in the database
* Takes the following parameters:
* userId: used to identify which user to update
* value: the value to prepend
* field: the field in which to prepend the value
* conn: a datbase connection
* callback: a callback function
*/
const prependFieldInDB = (userId, value, field, conn, callback) => {
  r.table(USER_TABLE)
   .get(userId)
   .update({ [field]: r.row(field).prepend(value) })
   .run(conn, (err, res) => {
     callback(err)
  })
}


// Publicly exposed functions
module.exports = {
  newUser,
  getUser,
  recordSession,
  saveSearch,
  addTicker,
  deleteTicker
}
