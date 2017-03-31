'use strict'
import { KG_KEY } from '../credentials/config'

export const kgSearch = name => {
  const api = "https://kgsearch.googleapis.com/v1/entities:search";
  const queryURL = api + '?query=' + name + '&key=' + KG_KEY + '&limit=1&indent=True';
  return (
    fetch(queryURL)
    .then((res) => res.json())
    .then((res) => res.itemListElement[0].result.detailedDescription.articleBody)
  );
}
