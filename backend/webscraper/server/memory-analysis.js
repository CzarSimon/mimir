const memwatch = require('memwatch-next');

memwatch.on('leak', info => {
  console.log('A leak has occured');
  console.log(info);
});

memwatch.on('stats', stats => {
  console.log('On stats:');
  console.log(stats);
});
