import { split, map, lowerCase, join, slice, findIndex } from 'lodash';

export const array_equals = (a1, a2) => {
  let i = a1.length;
  if (i !== a2.length) return false;
  while (i--) {
    if (a1[i] !== a2[i]) return false;
  }
  return true;
}

export const round = (number, decimals = 2) => {
  return number.toFixed(decimals);
}

export const format_name = (name, forbidden = ['inc', 'corporation']) => {
  const words = split(name, ' ');
  const lower_words = map(words, (word) => lowerCase(word));

  for (let word of lower_words) {
    if (forbidden.includes(word)) {
      return join(slice(words, 0, findIndex(word, null, 1) + 1), ' ');
    }
  }
  return name;
}
