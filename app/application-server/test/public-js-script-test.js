var assert = chai.assert;

/**
  Before running this, the socket code in script.js
  needs to be commented out. I know this is terrible way
  test but in order to get a test case up fast thats what
  I did. Should defenitly be changed in the future.
*/

describe('Testing public/js/scritp.js', function() {

  describe('connect with script', function() {
    it('should be connected', function() {
      assert.equal('connected', TestConnection());
    });
  });

  describe('sort', function() {
    describe('one digit urgency', function() {
      var unsorted_list = [
        {name: 'name1', urgency: 0.43},
        {name: 'name2', urgency: 2.11},
        {name: 'name3', urgency: 1.78}
      ];
      var sorted_list = [
        {name: 'name2', urgency: 2.11},
        {name: 'name3', urgency: 1.78},
        {name: 'name1', urgency: 0.43}
      ];
      it('should be in order', function() {
        assert.deepEqual(sorted_list, sort(unsorted_list));
      });
    });

    describe('multi-digit urgency', function() {
      var unsorted_list = [
        {name: 'name1', urgency: 65.46},
        {name: 'name2', urgency: 6.19},
        {name: 'name3', urgency: 4.51},
        {name: 'name4', urgency: 2.60},
        {name: 'name5', urgency: 13.01},
        {name: 'name6', urgency: 123.42}
      ];
      var sorted_list = [
        {name: 'name6', urgency: 123.42},
        {name: 'name1', urgency: 65.46},
        {name: 'name5', urgency: 13.01},
        {name: 'name2', urgency: 6.19},
        {name: 'name3', urgency: 4.51},
        {name: 'name4', urgency: 2.60}
      ];
      it('sorting should have hanlded 2 digit urgency', function() {
        assert.deepEqual(sorted_list, sort(unsorted_list));
      });
    });

  });
});
