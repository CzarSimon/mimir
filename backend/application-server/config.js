'use strict';

module.exports = {
  rethinkdb: {
    host: process.env.DB_HOST || 'localhost',
    port: 28015,
    db: 'mimir_app_server'
  },
  express: {
    port: 3000
  }
}
